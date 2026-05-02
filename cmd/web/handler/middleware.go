package handler

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/enmus"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				request, _ := httputil.DumpRequest(c.Request, false)
				fmt.Printf("[Recovery] panic recovered: %v request: %s ## stack: %s\n", err, string(request), string(stack(3)))
				global.G_LOG.Errorf("[Recovery] panic recovered: %v request: %s ## stack: %s", err, string(request), string(stack(3)))
				response.FailJSON(c, ecode.ServerErr)
				c.Abort()
				return
			}
		}()
		value, err := c.Cookie(enmus.LOGIN_COOKIE)
		if err != nil || len(value) == 0 {
			SetSessionCookie(c)
		} else {
			sessionIds := strings.Split(value, "_")
			if len(sessionIds) != 2 {
				SetSessionCookie(c)
			} else {
				// 从环境变量/配置读取密钥
				sessionAuthToken := tool.GetGlobalSecrets().SessionAuthToken
				signBase64 := base64.StdEncoding.EncodeToString([]byte(sessionIds[1] + sessionAuthToken))
				signStr := signBase64[:5]
				if signStr != sessionIds[0] {
					SetSessionCookie(c)
				}
			}
		}
		c.Next()
	}
}

func SetSessionCookie(c *gin.Context) string {
	sessionId := uuid.NewString()
	// 从环境变量/配置读取密钥
	sessionAuthToken := tool.GetGlobalSecrets().SessionAuthToken
	signBase64 := base64.StdEncoding.EncodeToString([]byte(sessionId + sessionAuthToken))
	newSessionId := signBase64[:5] + "_" + sessionId
	// Local dev often runs on http. Use secure cookie only when request is https.
	secureCookie := c.Request.TLS != nil || strings.EqualFold(c.GetHeader("X-Forwarded-Proto"), "https")
	c.SetCookie(enmus.LOGIN_COOKIE, newSessionId, 0, "/", "", secureCookie, true)
	return newSessionId
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOriginList := tool.GetAllowedCORSOrigins()
		if len(allowedOriginList) == 0 {
			allowedOriginList = []string{
				"http://localhost:3000",
				"http://localhost:5173",
				"http://localhost:8080",
				"http://127.0.0.1:5173",
				"http://127.0.0.1:8080",
			}
		}

		allowedOrigins := make(map[string]bool, len(allowedOriginList))
		for _, o := range allowedOriginList {
			allowedOrigins[o] = true
		}

		origin := c.Request.Header.Get("Origin")
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-Token")
			c.Header("Access-Control-Max-Age", "3600")
		}

		// 添加安全响应头防护 XSS、点击劫持等
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		_, _ = fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

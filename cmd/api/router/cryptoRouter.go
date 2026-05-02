package router

import (
	"bootpkg/cmd/api/handler"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/enmus"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/utils"
)

func getCryptoAuthToken() string {
	if token := strings.TrimSpace(global.CONFIG.General.CryptoAuthToken); token != "" {
		return token
	}
	return strings.TrimSpace(global.CONFIG.General.ApiSHA256Salt)
}

func NewCryptoRouter() *gin.Engine {
	if global.CONFIG.General.ENV == enmus.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	baseRouters := gin.New()
	baseRouters.NoRoute(func(c *gin.Context) {
		c.Status(200)
	})
	baseRouters.Static("/api/upload", "./upload") //静态文件
	baseRouters.Any("/crypto/:code", func(c *gin.Context) {
		authToken := getCryptoAuthToken()
		if authToken == "" {
			c.Status(619)
			c.Abort()
			return
		}

		timestampStr := c.GetHeader("GP-TM") //时间戳
		timestamp := tool.Int(timestampStr)
		timestampNow := time.Now().Unix()
		subUnix := timestampNow - timestamp
		if subUnix > 36000 || subUnix < -36000 {
			c.Status(619)
			c.Abort()
			return
		}

		payloadText := c.GetHeader("GP-TPM") //加密header头
		if len(payloadText) < 16 {
			c.Status(619)
			c.Abort()
			return
		}
		//验证URL链接
		code := c.Param("code")
		sign := strings.ToLower(tool.MD5([]byte(payloadText + timestampStr + authToken)))
		if sign[:16] != code {
			c.Status(619)
			c.Abort()
			return
		}

		aesKey := tool.MD5([]byte(timestampStr + strings.ToUpper(c.Request.Method) + authToken))[:16]

		dataByte, _ := base64.StdEncoding.DecodeString(payloadText)
		payloadTextByte, err := tool.DecryptAES(dataByte, []byte(aesKey))
		if err != nil {
			c.Status(619)
			c.Abort()
			return
		}
		headerJSON := string(payloadTextByte)

		var headers map[string]string
		utils.Unmarshal(headerJSON, &headers)
		for k, v := range headers {
			c.Request.Header.Set(k, v)
		}
		c.Request.Header.Set("DataAesKey", aesKey)
		//url处理
		urlStr := headers["url"]
		c.Request.RequestURI = urlStr
		index := strings.Index(urlStr, "?")
		if index > -1 {
			c.Request.URL.Path = urlStr[:index]
		} else {
			c.Request.URL.Path = urlStr
		}

		if c.Request.Method == "GET" {
			err := parseQuery(c, aesKey)
			if err != nil {
				response.FailErrJSON(c, 519, "处理失败")
				c.Abort()
				return
			}
		}
		routers.HandleContext(c)
		c.Abort()
	}).Use(handler.Recovery(), handler.Cors())
	return baseRouters
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	return w.body.WriteString(s)
}

func decryptString(encryptString []byte, aesKey string) (string, error) {
	if len(encryptString) < 1 {
		return "", nil
	}
	plaintext, err := tool.DecryptAES(encryptString, []byte(aesKey))
	if err != nil {
		return "", err
	}
	if len(plaintext) < 3 {
		//plaintext 应该是json 串 {}
		return "", nil
	}
	return string(plaintext), nil
}

func getStr(v interface{}) string {
	val := ""
	switch v := v.(type) {
	case float64:
		val = strconv.FormatFloat(v, 'f', -1, 64)
	default:
		val = fmt.Sprintf("%v", v)
	}
	return val
}

// get解密
// 处理get url的解密
func parseQuery(c *gin.Context, aesKey string) error {
	encryptString := c.DefaultQuery("x", "")

	if len(encryptString) < 1 {
		return nil
	}
	encryptByte, err := base64.StdEncoding.DecodeString(encryptString)
	if err != nil {
		return errors.New("base64解密失败")
	}

	queryData, err := decryptString(encryptByte, aesKey)
	if err != nil {
		return err
	}
	//var args []string
	//for k, v := range queryData {
	//	val := getStr(v)
	//	args = append(args, fmt.Sprintf("%s=%s", k, url.QueryEscape(val)))
	//}
	c.Request.URL.RawQuery = queryData
	return nil
}

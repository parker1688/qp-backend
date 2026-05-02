package handler

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ipWindowCounter struct {
	windowStart int64
	count       int
}

var (
	rateLimitMu   sync.Mutex
	ipRequestStat = map[string]*ipWindowCounter{}
	rateLimitOnce sync.Once
)

func startRateLimitGC() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		nowSec := time.Now().Unix()
		rateLimitMu.Lock()
		for ip, counter := range ipRequestStat {
			if nowSec-counter.windowStart > 120 {
				delete(ipRequestStat, ip)
			}
		}
		rateLimitMu.Unlock()
	}
}

// IPRateLimit limits requests per IP in a fixed 1-second window.
func IPRateLimit(maxPerSecond int) gin.HandlerFunc {
	return RequestRateLimit(maxPerSecond, map[string]int{})
}

// RequestRateLimit supports per-path rate limit overrides.
// key strategy: Token/IP + URL.Path to avoid one endpoint starving others.
func RequestRateLimit(defaultPerSecond int, pathSpecificLimit map[string]int) gin.HandlerFunc {
	if defaultPerSecond <= 0 {
		defaultPerSecond = 200
	}

	skipPathPrefix := []string{
		"/api/upload",
		"/favicon.ico",
		"/health",
	}

	getLimitByPath := func(path string) int {
		for prefix, limit := range pathSpecificLimit {
			if strings.HasPrefix(path, prefix) {
				if limit > 0 {
					return limit
				}
				break
			}
		}
		return defaultPerSecond
	}

	return func(c *gin.Context) {
		rateLimitOnce.Do(func() {
			go startRateLimitGC()
		})

		nowSec := time.Now().Unix()
		ip := c.ClientIP()
		token := strings.TrimSpace(c.GetHeader("Token"))
		path := c.Request.URL.Path
		for _, prefix := range skipPathPrefix {
			if strings.HasPrefix(path, prefix) {
				c.Next()
				return
			}
		}
		limit := getLimitByPath(path)

		identity := ip
		if token != "" {
			identity = token
		}
		key := identity + "|" + path

		rateLimitMu.Lock()
		counter, ok := ipRequestStat[key]
		if !ok {
			counter = &ipWindowCounter{windowStart: nowSec, count: 0}
			ipRequestStat[key] = counter
		}

		if counter.windowStart != nowSec {
			counter.windowStart = nowSec
			counter.count = 0
		}

		counter.count++
		exceeded := counter.count > limit
		rateLimitMu.Unlock()

		if exceeded {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": 429,
				"msg":  "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

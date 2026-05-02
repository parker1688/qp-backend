package handler

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/pkg/core/modules/enmus"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

/**
TODO: 根据URL 缓存输出的数据-减少数据计算
*/

var _cacheUrl = cache.New(60*time.Second, 120*time.Second)

type CustomResponseCacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseCacheWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseCacheWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func UrlLocalCacheJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER) //头部
		domain := c.GetHeader(enmus.Domain_HEADER)          //域名
		merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)  //商户Code
		path := domain + clientType + merchantCode + c.Request.URL.String()
		if val, ok := _cacheUrl.Get(path); ok {
			v := val.(string)
			c.Header("Cache-URL", "HIT")
			c.Data(200, "application/json", []byte(v))
			c.Abort()
			return
		}
		blw := &CustomResponseCacheWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		_cacheUrl.SetDefault(path, blw.body.String())
	}
}

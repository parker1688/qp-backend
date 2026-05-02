package handler

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/pkg/core/modules/enmus"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

/**
TODO: 根据URL 缓存输出的数据-减少数据计算
*/

func UrlLocalCacheJsonTow() gin.HandlerFunc {
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
		_cacheUrl.Set(path, blw.body.String(), 10*time.Second)
	}
}

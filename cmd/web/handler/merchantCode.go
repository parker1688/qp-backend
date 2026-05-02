package handler

import (
	"bootpkg/common/response"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"io"
	"net/url"
	"strings"
)

func MerchantCodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqContentType := c.Request.Header.Get("Content-Type")
		isJsonRequest := strings.Contains(reqContentType, "application/json")
		isFormUrl := strings.Contains(reqContentType, "application/x-www-form-urlencoded")
		selectMerchantCode := c.Request.Header.Get("MerchantCode")
		if c.Request.Method == "GET" {
			values, err := url.ParseQuery(c.Request.URL.RawQuery)
			if err != nil {
				response.FailErrJSON(c, 899, "处理失败")
				c.Abort()
				return
			}
			if selectMerchantCode != "" {
				values.Set("merchant_code", selectMerchantCode)
			}
			queryData := values.Encode()
			//global.G_LOG.Infof("url: %s queryData: %s", c.Request.RequestURI, queryData)
			c.Request.URL.RawQuery = queryData
		} else if isJsonRequest {
			payload, err := c.GetRawData()
			if err != nil {
				response.FailErrJSON(c, 899, "处理失败")
				c.Abort()
				return
			}
			if selectMerchantCode != "" {
				payload, err = sjson.SetBytes(payload, "merchant_code", selectMerchantCode)
				if err != nil {
					response.FailErrJSON(c, 899, "处理失败")
					c.Abort()
					return
				}
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
		} else if isFormUrl {
			payload, err := c.GetRawData()
			if err != nil {
				response.FailErrJSON(c, 899, "处理失败")
				c.Abort()
				return
			}
			values, err := url.ParseQuery(string(payload))
			if err != nil {
				response.FailErrJSON(c, 899, "处理失败")
				c.Abort()
				return
			}
			if selectMerchantCode != "" {
				values.Set("merchant_code", selectMerchantCode)
			}
			formData := values.Encode()
			payload = []byte(formData)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
		}
		c.Next()
	}
}

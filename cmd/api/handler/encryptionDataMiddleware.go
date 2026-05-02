package handler

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

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

var notEncryption = map[string]struct{}{}

// 请求&加密返回数据
func EncryptionDataMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		gpReqToken := c.GetHeader("DataAesKey")
		if len(gpReqToken) == 0 {
			c.Next()
			return
		}
		reqContentType := c.Request.Header.Get("Content-Type")
		isJsonRequest := strings.Contains(reqContentType, "application/json")
		isFormUrl := strings.Contains(reqContentType, "application/x-www-form-urlencoded")

		aesKey := gpReqToken[:16]
		if isJsonRequest && len(aesKey) > 0 {
			err := parseJson(c, aesKey)
			if err != nil {
				c.Status(618)
				c.Abort()
				return
			}
		} else if isFormUrl && len(aesKey) > 0 {
			err := parseForm(c, aesKey, 0)
			if err != nil {
				c.Status(618)
				c.Abort()
				return
			}
		}
		rawData, err := c.GetRawData()
		if err == nil {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))
		}
		oldWriter := c.Writer
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		contentType := blw.ResponseWriter.Header().Get("Content-Type")
		responseByte := blw.body.Bytes()

		c.Writer = oldWriter
		if strings.Contains(contentType, "application/json") && len(aesKey) >= 16 {
			cacheHeader := blw.ResponseWriter.Header().Get("Cache-URL")
			if cacheHeader != "HIT" && len(responseByte) < 1024*5 {
				urlPath := c.FullPath()
				global.G_LOG.Infof("urlPath:%v request-id: %v token: %v request: %v  respone:%v", c.Request.Host+urlPath, c.GetHeader("Gp-Request-Id"), c.GetHeader("Token"), string(rawData), string(responseByte))
			}
			encryptByte, err := tool.EncryptAES(responseByte, []byte(aesKey))
			if err != nil {
				response.FailErrJSON(c, 61, "加密失败,请重试")
				return
			}

			str := base64.StdEncoding.EncodeToString(encryptByte)
			str = `{"x":"` + str + `","d":true}`
			c.Writer.WriteString(str)
			return
		}
		_, _ = c.Writer.Write(responseByte)
	}
}

func parseJson(c *gin.Context, aesKey string) error {
	//读取数据 body处理
	payload, err := c.GetRawData()
	if err != nil {
		return err
	}

	if len(payload) > 20 {
		payloadText := gjson.Get(string(payload), "x").String()
		if len(payloadText) == 0 {
			return errors.New("数据错误")
		}
		if len(payloadText) > 0 {
			payloadTextBase64Byte, err := base64.StdEncoding.DecodeString(payloadText)
			if err != nil {
				return errors.New("base64解密失败")
			}
			payloadText = string(payloadTextBase64Byte)
			payloadTextByte, err := tool.DecryptAES([]byte(payloadText), []byte(aesKey))
			if err != nil {
				return errors.New("非法数据")
			}
			payload = payloadTextByte
		}
	}
	if len(payload) == 0 {
		payload = []byte("{}")
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(payload))

	return nil
}

func decryptString(encryptString []byte, aesKey string) (map[string]interface{}, error) {
	formData := make(map[string]interface{}, 0)
	if len(encryptString) < 1 {
		return formData, nil
	}

	plaintext, err := tool.DecryptAES(encryptString, []byte(aesKey))
	if err != nil {
		return formData, err
	}

	if len(plaintext) < 3 {
		//plaintext 应该是json 串 {}
		return formData, nil
	}

	err = json.Unmarshal(plaintext, &formData)
	if err != nil {
		return formData, err
	}

	return formData, nil
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

func parseForm(c *gin.Context, aesKey string, randIndex int64) error {
	//读取数据 body处理
	payload, err := c.GetRawData()
	if err != nil {
		return err
	}
	if len(payload) < 1 {
		return errors.New("数据错误")
	}

	encryptByte, err := base64.StdEncoding.DecodeString(string(payload)[randIndex:])
	if err != nil {
		return errors.New("base64解密失败")
	}
	payload = encryptByte

	if len(payload) > 1 {

		values, err := url.ParseQuery(string(payload))
		if err != nil {
			return err
		}
		payloadText := values.Get("x")
		if len(payloadText) > 0 {
			mapData, err := decryptString([]byte(payloadText), aesKey)
			if err != nil {
				return err
			}
			for k, v := range mapData {
				values.Add(k, getStr(v))
			}
			formData := values.Encode()
			payload = []byte(formData)
		}
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(payload))

	return nil
}

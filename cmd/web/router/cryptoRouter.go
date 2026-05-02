package router

import (
	"bootpkg/cmd/web/handler"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/utils"
)

func NewCryptoRouter() *gin.Engine {
	baseRouters := gin.Default()
	baseRouters.NoRoute(func(c *gin.Context) {
		c.Status(404)
	})

	baseRouters.Static("/api/upload", "./upload") //静态文件
	baseRouters.Any("/api/sysmg/:code", func(c *gin.Context) {

		payloadText := c.GetHeader("XX-Req-MOI")
		if len(payloadText) < 16 {
			c.Status(619)
			c.Abort()
			return
		}
		code := c.Param("code")

		timestampStr := c.GetHeader("XX-Req-TM")
		timestamp := tool.Int(timestampStr)
		tpm := c.GetHeader("XX-Req-TPM")
		payloadTextSubIndex := timestamp % 9
		if payloadTextSubIndex < 5 {
			payloadTextSubIndex = 5
		}

		sign := strings.ToUpper(tool.MD5([]byte(payloadText + auth_token))[:8+payloadTextSubIndex])
		if len(tpm) == 0 && code != sign {
			c.Status(619)
			c.Abort()
			return
		}

		randStr := payloadText[:payloadTextSubIndex]
		signStr := tool.MD5([]byte(auth_token + timestampStr + strings.ToUpper(c.Request.Method) + randStr))
		aesKey := signStr[:16]
		dataAesKey := payloadText[:16]
		payloadTextBase64Byte, err := base64.StdEncoding.DecodeString(payloadText[payloadTextSubIndex:])
		if err != nil {
			c.Status(619)
			c.Abort()
			return
		}
		payloadText = string(payloadTextBase64Byte)
		payloadTextByte, err := tool.DecryptAES([]byte(payloadText), []byte(aesKey))
		if err != nil {
			c.Status(619)
			c.Abort()
			return
		}
		headerJSON := string(payloadTextByte)
		var headers map[string]string
		utils.Unmarshal(headerJSON, &headers)

		timestamp = tool.Int(headers["t"]) //时间判断
		remaining := time.Now().Unix() - timestamp
		if len(tpm) > 0 && remaining > 7200 { //图片上传界面
			c.Status(618)
			c.Abort()
			return
		}
		if len(tpm) == 0 && (remaining > 7200 || remaining < -7200) { //前后相差不能120分钟
			c.Status(618)
			c.Abort()
			return
		}
		c.Request.Header.Set("Token", headers["Token"])
		c.Request.Header.Set("E-DataAesKey", dataAesKey)
		c.Request.Header.Set("E-RandDataLen", tool.String(payloadTextSubIndex))
		c.Request.Header.Set("MerchantCode", headers["MC"])

		urlStr := headers["url"]
		c.Request.RequestURI = urlStr
		index := strings.Index(urlStr, "?")
		if index > -1 {
			c.Request.URL.Path = urlStr[:index]
		} else {
			c.Request.URL.Path = urlStr
		}
		if c.Request.Method == "GET" && len(dataAesKey) > 0 {
			err := parseQuery(c, dataAesKey, payloadTextSubIndex)
			if err != nil {
				response.FailErrJSON(c, 519, "处理失败")
				c.Abort()
				return
			}
		}

		routers.HandleContext(c)
		c.Abort()
	}).Use(handler.Recovery())
	routersFun = nil
	return baseRouters
}

// get解密
// 处理get url的解密
func parseQuery(c *gin.Context, aesKey string, randIndex int64) error {
	encryptString := c.DefaultQuery("x", "")

	if len(encryptString) < 1 {
		return nil
	}
	encryptByte, err := base64.StdEncoding.DecodeString(encryptString[randIndex:])
	if err != nil {
		return errors.New("base64解密失败")
	}
	encryptString = string(encryptByte)

	queryData, err := decryptString(encryptByte, aesKey)
	if err != nil {
		return err
	}

	var args []string
	//var logs []string
	for k, v := range queryData {
		val := getStr(v)
		args = append(args, fmt.Sprintf("%s=%s", k, url.QueryEscape(val)))
		//logs = append(logs, fmt.Sprintf("%s=%s", k, val))
	}

	c.Request.URL.RawQuery = strings.Join(args, "&")
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

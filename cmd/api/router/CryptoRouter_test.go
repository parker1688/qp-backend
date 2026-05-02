package router

import (
	"bootpkg/common/tool"
	"encoding/base64"
	"strings"
	"testing"
	"time"
)

const testAuthToken = "test-crypto-auth-token"

func TestCryptoRouter(t *testing.T) {
	//header加密计算
	timestamp := time.Now().Unix()
	aesKey := tool.MD5([]byte(tool.String(timestamp) + strings.ToUpper("post") + testAuthToken))[:16]
	headers := map[string]string{
		"url":   "/api/login",
		"Token": "",
	}
	B, _ := tool.EncryptAES([]byte(tool.String(headers)), []byte(aesKey))
	payloadText := base64.StdEncoding.EncodeToString(B)
	//code计算
	code := strings.ToLower(tool.MD5([]byte(payloadText + tool.String(timestamp) + testAuthToken))[:16])

	body := `{
	"username":"test",
	"password":"test"
	}`

	dataB, _ := tool.EncryptAES([]byte(body), []byte(aesKey))
	BodyText := base64.StdEncoding.EncodeToString(dataB)

	t.Log("GP-TM", timestamp)
	t.Log("URL:", code)
	t.Log("GP-TPM:", payloadText)
	t.Log("body:", BodyText)
	t.Log("aes-Key:", aesKey)
}

func TestCryptoRouterCallBack(t *testing.T) {
	callbackText := `kIYLGhK/HHQuMk2aD1L9WxoVG4ycrWM4aI2ODO/zzaHfgYH6EwVX7FbUw1bownsg8aRNSDtQcT8wCuu6Iqg5f5Q8Xx82Yp7K8mlsAN+OKN15j8DCax6xowq872FWRB6imlhI1kWwbTQ0565inQ+pK6FeId/wGpkhAryPxWxd6y0gyQC5xIKjQY263golM5DAukYz/Y56Sb4WhiYBd4OfbqxMxQZRvE7rxxaislv1GD1ejunZzxNfJK2H46PFTqsroeDTa6wDG0n7h7bR5izq9JV0p74vUNjTLAYkZDfWO52oEfP/UHQassWoFMberQe1U9ARuhZ2PP8WoG1kilPcZp9ntV77A0iiRQofJSJDd2Unzq+60ctgOYL3iQ4wnr0o`
	callbackBodyByte, _ := base64.StdEncoding.DecodeString(callbackText)
	callbackB, _ := tool.DecryptAES(callbackBodyByte, []byte("d862aafd06b0a721"))
	t.Log("callbackB", string(callbackB))
}

func TestCryptoRouterUrl(t *testing.T) {
	header := "xB6AayslYMor2HqmGkLmlw2LxV0yDVme7ZP40V6kFRnMGnt6UatndMzA75w3swNnWURl3IfxfTuUYVNn0xPS/mn8uK+uHlQOl21xpDarAjY="
	tm := "1692274866"
	m := tool.MD5([]byte(header + tm + testAuthToken))
	t.Log(m)
}

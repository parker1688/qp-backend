package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func HashSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	//fmt.Println(base64.StdEncoding.EncodeToString(sum))
	//由于是十六进制表示，因此需要转换
	s := hex.EncodeToString(sum)
	return s
}

func MakeSign(merchantCode, timestamp int64, md5key, data string) string {
	str := fmt.Sprintf("%v%d%s%s",
		merchantCode,
		timestamp,
		data,
		md5key,
	)
	return MD5([]byte(str))
}

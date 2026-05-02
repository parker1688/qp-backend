package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HMACSignatureSha256(secret2, message string) string {
	hash := hmac.New(sha256.New, []byte(secret2))
	hash.Write([]byte(message))
	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))
	// to base64
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Encryption(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

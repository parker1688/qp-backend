package tool

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) ([]byte, error) {
	n := len(src)
	unPadNum := int(src[n-1])
	if len(src) < (n - unPadNum) {
		return nil, errors.New("data PKCS7 Slice err range out")
	}
	return src[:(n - unPadNum)], nil
}

// 加密
//func EncryptAES(src []byte, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	src = padding(src, block.BlockSize())
//	blockMode := cipher.NewCBCEncrypter(block, key)
//	blockMode.CryptBlocks(src, src)
//	return src, nil
//}
//
//// 解密
//func DecryptAES(src []byte, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	blockMode := cipher.NewCBCDecrypter(block, key)
//	blockMode.CryptBlocks(src, src)
//	src,err = unpadding(src)
//	return src, err
//}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return []byte("")
	}
	if origData == nil {
		return []byte("")
	}

	unpadding := int(origData[length-1])
	//fmt.Printf("origData: %v length: %v unpadding: %d capLen: %v\n", string(origData), length, unpadding, cap(origData))
	if length < unpadding {
		return []byte("")
	}

	return origData[:(length - unpadding)]
}

func DecryptAES(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		if len(decrypted) < be || len(data) < be {
			return nil, errors.New("data error")
		}
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted), nil
}

func EncryptAES(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted, nil
}

// EncryptAESPrefixRandKey
//
//	@Description: 加密数据随机KEY
//	@param data 需要加密的数据
//	@return string 加密Key + base64数据
//	@return string 加密Key
func EncryptAESPrefixRandKey(data string) (string, string) {
	aesKey := strings.ToUpper(RandString(16))
	B, err := EncryptAES([]byte(data), []byte(aesKey))
	if err != nil {
		return "", ""
	}
	return aesKey + base64.StdEncoding.EncodeToString(B), aesKey
}

// DecryptAESPrefixRandKey
//
//	@Description: 加密数据随机KEY
//	@param data 加密后的数据
//	@return string 解密后数据
func DecryptAESPrefixRandKey(data string) string {
	if len(data) < 16 {
		return ""
	}
	secret := data[16:]
	secretByte, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return ""
	}
	aesKey := data[:16]
	B, err := DecryptAES(secretByte, []byte(aesKey))
	if err != nil {
		return ""
	}
	return string(B)
}

// EncryptAESPrefixRandKeySalt
//
//	@Description: Aes 128加密数据
//	@param data 需要加密的数据
//	@param aesKey 128位加密Key, 长度不够补全16位
//	@return string 加密后的字符串
//	@return string 加密Key
func EncryptAESPrefixRandKeySalt(data, aesKey string) (string, string) {
	randKey := strings.ToUpper(RandString(16))
	if len(aesKey) < 16 {
		aesKey = aesKey + randKey[:16-len(aesKey)]
	} else {
		aesKey = aesKey[:16]
	}
	B, err := EncryptAES([]byte(data), []byte(aesKey))
	if err != nil {
		return "", ""
	}

	return randKey + base64.StdEncoding.EncodeToString(B), aesKey
}

// EncryptAESBase64EnCode
//
//	@Description: Aes 128加密数据
//	@param data 需要加密的数据
//	@param aesKey 128位加密Key
//	@return string 加密后的字符串
//	@return string 加密Key
func EncryptAESBase64EnCode(data, aesKey string) (string, string) {
	if data == "" || aesKey == "" {
		return "", ""
	}

	b, err := EncryptAES([]byte(data), []byte(aesKey))
	if err != nil {
		return "", ""
	}

	return base64.StdEncoding.EncodeToString(b), aesKey
}

// DecryptAESPrefixRandKeySalt
//
//	@Description: 解密AES 128加密
//	@param data 加密数据
//	@param aesKey 128位加密Key
//	@return string 解密后字符串
func DecryptAESPrefixRandKeySalt(data, aesKey string) string {
	if len(data) < 17 {
		return ""
	}
	//secret := data[16:]
	secretByte, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return ""
	}
	if len(aesKey) < 17 {
		aesKey = aesKey + data[:16-len(aesKey)]
	} else {
		aesKey = aesKey[:16]
	}
	B, err := DecryptAES(secretByte, []byte(aesKey))
	if err != nil {
		fmt.Println("DecryptAES err:", err.Error())
		return ""
	}
	return string(B)
}

// DecryptAESDecodeStr
//
//	@Description: 解密AES 128加密
//	@param data 加密数据
//	@param aesKey 128位加密Key
//	@return string 解密后字符串
func DecryptAESDecodeStr(data, aesKey string) string {
	if data == "" || aesKey == "" {
		return ""
	}

	secretByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}

	B, err := DecryptAES(secretByte, []byte(aesKey))
	if err != nil {
		return ""
	}
	return string(B)
}

// DecryptAESPrefixRandKeySaltDefault
//
//	@Description: 解密AES 128加密 (解密失败默认返回data)
//	@param data 加密数据
//	@param aesKey 128位加密Key
//	@return string 解密后字符串
func DecryptAESPrefixRandKeySaltDefault(data, aesKey string) string {
	defaultData := DecryptAESPrefixRandKeySalt(data, aesKey)
	if len(defaultData) == 0 {
		return data
	}
	return defaultData
}

func AesEcbPk7EncryptBase64(data, key string) (string, error) {
	if data == "" || key == "" {
		return "", errors.New("data or key is empty")
	}

	encryptData, err := AesECBPk7Encrypt([]byte(data), []byte(key))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptData), nil
}

func AesEcbPk7DecryptBase64(data, key string) (string, error) {
	if data == "" || key == "" {
		return "", errors.New("data or key is empty")
	}

	secretByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	b, err := AesECBPk7Decrypt(secretByte, []byte(key))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// aes ecb 模式加密
func AesECBPk7Encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCS7Padding(data, block.BlockSize())

	encrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(encrypted[bs:be], data[bs:be])
	}

	return encrypted, nil
}

// aes ecb 模式解密
func AesECBPk7Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		if len(decrypted) < be || len(data) < be {
			return nil, errors.New("data error")
		}
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted), nil
}

// EncryptAESPrefixAesKey
//
//	@Description: Aes 128加密数据
//	@param data 需要加密的数据
//	@param aesKey 128位加密Key, 长度不够补全16位
//	@return string 加密后的字符串
//	@return string 加密Key
func EncryptAESPrefixAesKey(data, aesKey string) (string, string) {
	if len(aesKey) < 16 {
		return "", ""
	} else {
		aesKey = aesKey[:16]
	}
	B, err := EncryptAES([]byte(data), []byte(aesKey))
	if err != nil {
		return "", ""
	}
	return base64.StdEncoding.EncodeToString(B), aesKey
}

// DecryptAESPrefixAesKey
//
//	@Description: Aes 128加密数据
//	@param data 需要加密的数据
//	@param aesKey 128位加密Key, 长度不够补全16位
//	@return string 加密后的字符串
//	@return string 加密Key
func DecryptAESPrefixAesKey(data, aesKey string) (string, string) {
	if len(aesKey) < 16 {
		return "", ""
	} else {
		aesKey = aesKey[:16]
	}
	secretByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", ""
	}
	B, err := DecryptAES(secretByte, []byte(aesKey))
	if err != nil {
		return "", ""
	}
	return string(B), ""
}

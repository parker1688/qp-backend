package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func AesEncrypt(plainText, key, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("invalid encryption key: %w", err)
	}

	// Apply zero padding
	plainText = zeroPadding(plainText, block.BlockSize())

	// Encrypt
	cipherText := make([]byte, len(plainText))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherText, plainText)

	// Encode cipherText to base64
	encryptedString := base64.RawURLEncoding.EncodeToString(cipherText)
	return encryptedString, nil
}

// zeroPadding pads plainText with zeros to a multiple of blockSize.
func zeroPadding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	paddingText := bytes.Repeat([]byte{0}, padding)
	return append(plainText, paddingText...)
}

func aesDecrypt(encryptedText string, key, iv []byte) (string, error) {
	// Decode base64-encoded encryptedText
	decodedData, err := base64.RawURLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("invalid encrypted data: %w", err)
	}

	// Decrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("invalid decryption key: %w", err)
	}

	plainText := make([]byte, len(decodedData))
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(plainText, decodedData)

	// Remove zero padding
	plainText = bytes.TrimRight(plainText, string([]byte{0}))
	return string(plainText), nil
}

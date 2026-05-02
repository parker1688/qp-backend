package crypt

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// Parse the private key from a PEM encoded string
func GetPrivateKey(pemEncodedKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemEncodedKey))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return rsaPrivateKey, nil
}

// Encrypt data using the RSA private key with appropriate padding
func PrivateEncrypt(data string, privateKey *rsa.PrivateKey) (string, error) {
	buffer := []byte(data)
	keySize := privateKey.Size()
	maxBlockSize := keySize - 11 // 11 bytes for PKCS#1 v1.5 padding

	var encryptedChunks [][]byte

	for offset := 0; offset < len(buffer); offset += maxBlockSize {
		end := offset + maxBlockSize
		if end > len(buffer) {
			end = len(buffer)
		}

		chunk := buffer[offset:end]
		encryptedChunk, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.Hash(0), chunk)
		if err != nil {
			return "", err
		}
		encryptedChunks = append(encryptedChunks, encryptedChunk)
	}

	encryptedBuffer := bytes.Join(encryptedChunks, nil)
	return base64.StdEncoding.EncodeToString(encryptedBuffer), nil
}

var (
	ErrWrongFormatKey          = errors.New("the key has a wrong format")
	ErrUnexpectedBytesPayload  = errors.New("unexpected bytes payload")
	ErrUnexpectedPayloadLength = errors.New("unexpected payload length")

	paddingBytesSequence = []byte{0xff, 0}
)

// ParsePublic parses *rsa.PublicKey from the raw bytes data.
func ParsePublic(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)

	parsed, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	puk, ok := parsed.(*rsa.PublicKey)
	if !ok {
		return nil, ErrWrongFormatKey
	}

	return puk, nil
}

// getPublicKey formats the key and wraps it with PEM headers.
func ForPublicKey(key string) string {
	result := insertStr(key, "\n", 64)
	return "-----BEGIN PUBLIC KEY-----\n" + result + "-----END PUBLIC KEY-----"
}

func insertStr(str, insertStr string, sn int) string {
	var newStr strings.Builder
	for i := 0; i < len(str); i += sn {
		end := i + sn
		if end > len(str) {
			end = len(str)
		}
		newStr.WriteString(str[i:end])
		newStr.WriteString(insertStr)
	}
	return newStr.String()
}

func SecretKeypublicDecrypt(payload []byte, pubKey *rsa.PublicKey) (string, error) {
	buffer := payload
	keySize := pubKey.Size()
	maxBlockSize := keySize

	var encryptedChunks [][]byte

	for offset := 0; offset < len(buffer); offset += maxBlockSize {
		end := offset + maxBlockSize
		if end > len(buffer) {
			end = len(buffer)
		}

		chunk := buffer[offset:end]
		encryptedChunk, err := PublicDecrypt(chunk, pubKey)
		if err != nil {
			return "", err
		}
		encryptedChunks = append(encryptedChunks, encryptedChunk)
	}

	encryptedBuffer := bytes.Join(encryptedChunks, nil)
	return string(encryptedBuffer), nil
}

// PublicDecrypt decrypts a payload with public key and then return decrypted bytes.
func PublicDecrypt(payload []byte, pubKey *rsa.PublicKey) ([]byte, error) {
	if len(payload) < 1 {
		return nil, ErrUnexpectedPayloadLength
	}

	m := new(big.Int).SetBytes(payload)
	e := big.NewInt(int64(pubKey.E))
	c := new(big.Int).Exp(m, e, pubKey.N)

	bbytes := bytes.Split(c.Bytes(), paddingBytesSequence)
	if len(bbytes) != 2 {
		return nil, ErrUnexpectedBytesPayload
	}

	return bbytes[1], nil
}

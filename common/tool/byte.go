package tool

import "crypto/rand"

// GenerateRandomBytes
//
//	@Description: 随机生成byte值
//	@param length
//	@return []byte
//	@return error
func GenerateRandomBytes(length int) ([]byte, error) {
	// 创建一个切片，用于存储随机字节
	randomBytes := make([]byte, length)

	// 使用 crypto/rand 生成指定长度的随机字节
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err // 如果生成失败，返回错误
	}

	return randomBytes, nil
}

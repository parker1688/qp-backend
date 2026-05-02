package tool

import (
	"math/rand"
	"strconv"
	"strings"
)

func Int(v interface{}) int64 {
	switch v := v.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return 0
		}
		return i
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

func Atoi(v string) int {
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return n
}

func ToFloat64Zero(str string) []float64 {
	strSlice := strings.Split(str, ",")
	// 初始化 int 数组
	var intSlice []float64
	// 遍历字符串切片并转换为 int
	for _, s := range strSlice {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}
		if num == 0 {
			continue
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func StringToFloat64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// 生成 [min,max) 范围的随机整数
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// 生成 [min,max] 范围的随机浮点数
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

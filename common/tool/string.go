package tool

import (
	"bootpkg/pkg/core/modules/vo"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Length returns number of characters
func Length(s string) int {
	return len([]rune(s))
}

// Before returns the string before the first occurrence of the substr string
func Before(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.Index(s, substr)
	if i != -1 {
		return s[:i]
	}
	return s
}

// BeforeLast returns the string before the last occurrence of the substr string
func BeforeLast(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.LastIndex(s, substr)
	if i != -1 {
		return s[:i]
	}
	return s
}

// After returns the string after the first occurrence of the substr string
func After(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.Index(s, substr)
	if i != -1 {
		i = i + len(substr)
		return s[i:]
	}
	return s
}

// AfterLast returns the string after the last occurrence of the substr string
func AfterLast(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.LastIndex(s, substr)
	if i != -1 {
		i = i + len(substr)
		return s[i:]
	}
	return s
}

func Index(s, substr string) int {
	return strings.Index(s, substr)
}

func RuneIndex(s, substr string) int {
	p := strings.Index(s, substr)
	if p == -1 || p == 0 {
		return p
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(s)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		pos++
		totalSize += size

		if totalSize == p {
			return pos
		}
	}
	return pos
}

func Contians(s, substr string) bool {
	return strings.Contains(s, substr)
}

func StartsWith(s, substr string) bool {
	if substr != "" && Substr(s, 0, len([]rune(substr))) == substr {
		return true
	}
	return false
}

func EndsWith(s, substr string) bool {
	if Substr(s, -len([]rune(substr)), len(s)) == substr {
		return true
	}
	return false
}

// Substr returns a string of length length from the start position
func Substr(s string, start int, strlength ...int) string {
	charlist := []rune(s)
	l := len(charlist)
	length := 0
	end := 0

	if len(strlength) == 0 {
		length = l
	} else {
		length = strlength[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}

	return string(charlist[start:end])
}

func SubByte(str string, length int) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

// Char returns a char slice
func Char(str string) []string {
	c := make([]string, 0)
	for _, v := range str {
		c = append(c, string(v))
	}
	return c
}

func Escape(s string) string {
	str := strconv.Quote(s)
	str = strings.Replace(str, "'", "\\'", -1)
	strlist := []rune(str)
	l := len(strlist)
	return Substr(str, 1, l-2)
}

func Ufirst(s string) string {
	r := []rune(s)
	if len(s) > 0 && unicode.IsLetter(r[0]) && unicode.IsLower(r[0]) {
		r[0] -= 32
		return string(r)
	}
	return s
}

// String returns a string of any type
func String(iface interface{}) string {
	switch val := iface.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(iface)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Slice:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", iface)
}

func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return ""
	}
	if start == 0 && end == length {
		return source
	}
	return string(r[start:end])
}

func MaskString(s string, start, end int) string {
	if len(s) <= start+end {
		return strings.Repeat("*", len(s))
	}
	return s[:start] + strings.Repeat("*", len(s)-start-end) + s[len(s)-end:]
}

// 用户名合法性校验函数
func IsValidUsername(username string) bool {
	// 允许字母、数字、下划线，禁止汉字和特殊字符
	//reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	reg := regexp.MustCompile(`(?i)^[a-z0-9]{6,15}$`)
	return reg.MatchString(username)
}

func IsValidUsername2(username string) bool {
	// 允许字母、数字、下划线，禁止汉字和特殊字符
	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return reg.MatchString(username)
}

// StringReplaceAll
//
//	@Description: 将字符串中多个符号替换为指定的字符串
//	@param input 原始字符串
//	@param replace 替换结果
//	@param replaceSrc 需要替换的字符
func StringReplaceAll(input, replace string, allSpace bool, replaceSrc ...string) string {
	for _, v := range replaceSrc {
		input = strings.ReplaceAll(input, v, replace)
	}
	if allSpace {
		input = strings.TrimSpace(input)
	}
	return input
}

// 敏感数据处理
func PrivateDataHandler(v *vo.FcPrivate) {
	emailLen := len(v.Email)
	if emailLen > 4 {
		v.Email = v.Email[0:3] + "***" + v.Email[emailLen-4:]
	}

	telLen := len(v.Tel)
	if telLen > 4 {
		v.Tel = v.Tel[0:3] + "****" + v.Tel[telLen-4:]
	}

	alipayLen := len(v.Alipay)
	if alipayLen > 4 {
		v.Alipay = v.Alipay[0:3] + "****" + v.Alipay[alipayLen-4:]
	}

	alipayRealnameLen := len(v.AlipayRealname)
	if alipayRealnameLen > 4 {
		v.AlipayRealname = v.AlipayRealname[0:3] + "*****" + v.AlipayRealname[alipayRealnameLen-4:]
	}

	if v.WalletPassword != "" {
		v.WalletPassword = "已设置"
	}

	// 银行卡号
	accountNumLen := len(v.AccountNumber)
	if accountNumLen > 4 {
		v.AccountNumber = v.AccountNumber[0:3] + "*****" + v.AccountNumber[accountNumLen-4:]
	}

	// 银行卡号持有人
	accountHolderLen := len(v.AccountHolder)
	if accountHolderLen > 1 {
		tmpStr := ""
		for i := 0; i < accountHolderLen-1; i++ {
			tmpStr += "*"
		}
		v.AccountHolder = v.AccountHolder[0:1] + tmpStr
	}

}

func StrToListForSql(vals string) []interface{} {
	s1 := strings.Split(vals, ",")
	args := make([]interface{}, len(s1))
	for i, v := range s1 {
		args[i] = v
	}
	return args
}

func PlaceholderFormat(template string, args ...interface{}) string {
	re := regexp.MustCompile(`\{(\d+)\}`)
	return re.ReplaceAllStringFunc(template, func(m string) string {
		index, _ := strconv.Atoi(m[1 : len(m)-1])
		if index >= 0 && index < len(args) {
			return fmt.Sprintf("%v", args[index])
		}
		return m
	})
}

// ReplaceTemplate 替换字符串中{}包裹的内容为指定值
func ReplaceTemplate(template string, value interface{}) string {
	// 使用正则匹配所有{}包裹的内容
	re := regexp.MustCompile(`\{[^{}]+\}`)
	return re.ReplaceAllStringFunc(template, func(match string) string {
		return fmt.Sprintf("%v", value)
	})
}

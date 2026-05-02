package tool

import "unicode"

func PwdFormatValid(pwd string) bool {
	// 检查长度
	if len(pwd) < 6 || len(pwd) > 12 {
		return false
	}

	var hasDigit, hasLetter bool
	for _, r := range pwd {
		switch {
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsLetter(r):
			hasLetter = true
		default:
			return false // 包含非字母数字字符
		}
	}

	return hasDigit && hasLetter
}

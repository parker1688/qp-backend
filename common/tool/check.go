package tool

import (
	"fmt"
	"regexp"
)

const (
	// 数字或字母
	RegexNumStr = "^[a-zA-Z0-9]*$"

	// 仅数字
	RegexNum = "[0-9]"

	// 仅字母
	RegexStr = "[a-zA-Z]"

	// 必须字母开头，包含数字或字母
	RegexStrStartNumAndStr = "^[a-zA-Z][a-zA-Z0-9]*$"

	// 必须字母开头，包含数字或字母下划线横线
	RegexStrNumUlHl = "^[a-zA-Z][a-zA-Z0-9_-]*$"

	// 邮箱
	RegexEmail = "^[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)*@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$"
)

func IsNumStr(str string) (bool, error) {
	matched, err := regexp.MatchString(RegexNumStr, str)
	if err != nil {
		return false, fmt.Errorf("match str=%v err=%v", str, err)
	}
	if !matched {
		return false, nil
	}

	return true, nil
}

func IsNum(str string) (bool, error) {
	matched, err := regexp.MatchString(RegexNum, str)
	if err != nil {
		return false, fmt.Errorf("match str=%v err=%v", str, err)
	}
	if !matched {
		return false, nil
	}

	return true, nil
}

// 字母开头, 包含数字字母
func StrStartNum(str string) (bool, error) {
	matched, err := regexp.MatchString(RegexStrStartNumAndStr, str)
	if err != nil {
		return false, fmt.Errorf("match str=%v err=%v", str, err)
	}
	if !matched {
		return false, nil
	}

	return true, nil
}

// 字母开头, 包含数字字母横线下划线
func StrStartNumUlHl(str string) (bool, error) {
	matched, err := regexp.MatchString(RegexStrNumUlHl, str)
	if err != nil {
		return false, fmt.Errorf("match str=%v err=%v", str, err)
	}
	if !matched {
		return false, nil
	}

	return true, nil
}

func IsEmail(str string) (bool, error) {
	matched, err := regexp.MatchString(RegexEmail, str)
	if err != nil {
		return false, fmt.Errorf("match str=%v err=%v", str, err)
	}
	if !matched {
		return false, nil
	}

	return true, nil
}

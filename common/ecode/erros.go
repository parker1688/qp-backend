package ecode

import "bootpkg/langs"

// NewErrorCode
//
//	@Description: 自定义错误Code
//	@param code
//	@param text
//	@return error
func NewErrorCode(langCode string, text string, replacements ...*langs.Replacements) error {
	return &ErrorLangString{langCode, text, replacements}
}

// errorString is a trivial implementation of error.
type ErrorLangString struct {
	langCode     string //多语言错误标识
	s            string
	replacements []*langs.Replacements
}

func (e *ErrorLangString) Error() string {
	return e.s
}

func (e *ErrorLangString) LangCode() string {
	return e.langCode
}

func (e *ErrorLangString) Replacements() []*langs.Replacements {
	return e.replacements
}

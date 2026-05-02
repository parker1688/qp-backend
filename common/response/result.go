package response

import (
	"bootpkg/common/ecode"
	"bootpkg/langs"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type MerchantResponse struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Status    int         `json:"status"`
	Result    interface{} `json:"result"`
}

type PageResult struct {
	Response
	Total    int64 `json:"total"`
	PageSize int   `json:"page_size"`
	PageNo   int   `json:"page_no"`
}

type ItemResult struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (r Response) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

func SuccessPageJSON(c *gin.Context, pageNo, pageSize int, total int64, data interface{}) {
	rsp := &PageResult{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	rsp.Success = true
	rsp.Code = ecode.OK.Code()
	rsp.Msg = ecode.OK.Message()
	rsp.Timestamp = time.Now().UnixNano() / 1000000
	rsp.Data = data
	c.JSON(http.StatusOK, rsp)
}

func SuccessJSON(c *gin.Context, data interface{}) {
	rsp := &Response{
		Success:   true,
		Code:      ecode.OK.Code(),
		Msg:       ecode.OK.Message(),
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      data,
	}
	c.JSON(http.StatusOK, rsp)
}

func SuccessCodeJSON(c *gin.Context, code ecode.Code, data interface{}) {
	rsp := &Response{
		Success:   true,
		Code:      code.Code(),
		Msg:       code.Message(),
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      data,
	}
	c.JSON(http.StatusOK, rsp)
}

func SuccessMsgJSON(c *gin.Context, data interface{}, msg string) {
	rsp := &Response{
		Code:      ecode.OK.Code(),
		Msg:       msg,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, rsp)
}

func SuccessJSONValue(c *gin.Context, val string) {
	c.Data(200, "application/json", []byte(val))
}

func FailJSON(c *gin.Context, code ecode.Code) {
	rsp := &Response{
		Success:   false,
		Code:      code.Code(),
		Msg:       code.Message(),
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      nil,
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrJSON(c *gin.Context, code ecode.Code, err string) {
	rsp := &Response{
		Success:   false,
		Code:      code.Code(),
		Msg:       err,
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      nil,
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrJSONCode(c *gin.Context, code int, err string) {
	rsp := &Response{
		Success:   false,
		Code:      code,
		Msg:       err,
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      nil,
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrDataJSONData(c *gin.Context, code int, err string, data interface{}) {
	rsp := &Response{
		Success:   false,
		Code:      code,
		Msg:       err,
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      data,
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrDataJSON(c *gin.Context, code ecode.Code, err string, data interface{}) {
	rsp := &Response{
		Success:   false,
		Code:      code.Code(),
		Msg:       err,
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      data,
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrHttpCodeJSON(c *gin.Context, httpCode int, err string, data interface{}) {
	rsp := &Response{
		Success:   false,
		Msg:       err,
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      data,
	}
	c.JSON(httpCode, rsp)
}

// ErrorLanguage
//
//	@Description: 自定义错误,返回多语言Code
//	@param c gin
//	@param err 错误信息
//	@param replacements
//	@return string
func ErrorLanguage(c *gin.Context, err error) string {
	if err != nil {
		e, ok := err.(*ecode.ErrorLangString)
		if ok {
			return langs.GetWithLocaleGin(c, e.LangCode(), e.Replacements()...)
		}
		return err.Error()
	}
	return ""
}

package response

import (
	"bootpkg/common/ecode"
	"bootpkg/common/tool"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ERROR_PARAMETER = ecode.AddCode(501, "参数解析错误")
	ERROR_SERVER    = ecode.AddCode(504, "服务器错误")

	Merchant_Not_Exsit     = ecode.AddCode(1, "商户不存在")
	Params_Is_Bad          = ecode.AddCode(2, "参数错误")
	Params_Not_Unmarshal   = ecode.AddCode(3, "参数解析错误")
	Address_not_exsit      = ecode.AddCode(4, "地址不存在")
	Address_type_Only_One  = ecode.AddCode(5, "该类型地址不能重复添加")
	Member_Have_Exsit      = ecode.AddCode(6, "用户已经存在")
	Payment_Strategy_Limit = ecode.AddCode(7, "支付风控限制")

	ERROR_DEFAULT = ecode.AddCode(500, "默认错误")
)

type Response struct {
	Timestamp int64       `json:"timestamp"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

type PageResult struct {
	Response
	Total    int64 `json:"total"`
	PageSize int   `json:"page_size"`
	PageNo   int   `json:"page_no"`
}

type PageQuery struct {
	PageSize int `json:"pageSize"  form:"pageSize" uri:"pageSize"`
	PageNo   int `json:"current"  form:"current" uri:"current"`
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
	rsp.Code = ecode.OK.Code()
	rsp.Msg = ecode.OK.Message()
	rsp.Data = data
	rsp.Timestamp = time.Now().Unix()
	c.JSON(http.StatusOK, rsp)
}

func SuccessJSON(c *gin.Context, data interface{}) {
	rsp := &Response{
		Code:      ecode.OK.Code(),
		Msg:       "",
		Data:      data,
		Timestamp: time.Now().Unix(),
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

func SuccessMsgJSONData(c *gin.Context, data string, msg string) {
	c.Data(200, "application/json", []byte(`{"timestamp": `+tool.String(time.Now().Unix())+`,"code":0,"msg":"`+msg+`","data":`+data+`}`))
}

func FailJSON(c *gin.Context, code ecode.Code) {
	rsp := &Response{
		Code:      code.Code(),
		Msg:       code.Message(),
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrJSON(c *gin.Context, code ecode.Code, err string) {
	rsp := &Response{
		Code:      code.Code(),
		Msg:       err,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrDataJSON(c *gin.Context, code ecode.Code, err string, data interface{}) {
	rsp := &Response{
		Code:      code.Code(),
		Msg:       err,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, rsp)
}

func FailErrHttpCodeJSON(c *gin.Context, httpCode int, err string, data interface{}) {
	rsp := &Response{
		Code:      httpCode,
		Msg:       err,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	c.JSON(httpCode, rsp)
}

func FailErrCodeJSON(c *gin.Context, code int, err string, data interface{}) {
	rsp := &Response{
		Code:      code,
		Msg:       err,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	c.JSON(200, rsp)
}

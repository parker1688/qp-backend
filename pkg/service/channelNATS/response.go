package channelNATS

import (
	"bootpkg/common/tool"
)

type Response struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

// SuccessJSON
//
//	@Description: 发送消息结构体
//	@param action 方法名称
//	@param data 数据结构体
func SuccessJSON(action string, data interface{}) []byte {
	rsp := &Response{
		Action: action,
		Data:   data,
	}
	b, _ := tool.JsonMarshal(rsp)
	return b
}

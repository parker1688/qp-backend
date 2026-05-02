package message

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"github.com/go-resty/resty/v2"
)

type IMessage interface {
	//
	// GetMessage
	//  @Description: 获取消息体
	//  @return string 消息体字符串
	//
	GetMessage() string

	//
	// SetTags
	//  @Description: 设置发送消息的标签
	//  @param tags 消息标签集合
	//  @param tagsType 0: 包含匹配  1: 完全匹配标签
	//
	SetTags(tags []string, tagsType int)

	//
	// SetMessage
	//  @Description: 设置消息
	//  @param s 消息类型
	//
	SetMessage(s IContentMessage)
}

// BaseMessage
// @Description: 基础消息结构体
type BaseMessage struct {
	Timestamp int64  `json:"timestamp"` //发送时间戳
	MsgType   int    `json:"msg_type"`
	MsgId     string `json:"msg_id"` //消息ID
}

// MatchMessage
// @Description: 基本消息
type WsMessage struct {
	Tags     []string `json:"tags"`      //用户标签集合
	TagsType int      `json:"tags_type"` // 0 包含匹配  1 全部匹配
	Content  string   `json:"content"`   //消息结构体
}

func (m *WsMessage) SetTags(tags []string, tagsType int) {
	m.Tags = tags
	m.TagsType = tagsType
}

func (m *WsMessage) GetMessage() string {
	return tool.String(m)
}

func (m *WsMessage) SetMessage(s IContentMessage) {
	s.SetMsgType()
	data, _ := tool.EncryptAESPrefixRandKey(tool.String(s))
	g, _ := tool.JsonMarshalString(data)
	m.Content = g
}

// IContentMessage
// @Description: 消息体
type IContentMessage interface {
	//
	// SetMsgType
	//  @Description: 设置消息类型
	//
	SetMsgType()
}

func SendWsMessage(msg IMessage) bool {
	body := msg.GetMessage()
	client := resty.New()
	_, err := client.R().SetBody(body).
		SetHeader("content-type", "application/json").
		Post(global.CONFIG.Chatroom.WebsocketAddr + "/api/ws/message/send")
	if err != nil {
		return false
	}
	return true
}

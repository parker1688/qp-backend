package message

const (
	MESSAGE_MACTH = iota
)

type MatchMessage struct {
	BaseMessage
	UserName string `json:"user_name"`
	Vip      int    `json:"vip"`
	Content  string `json:"content"`
}

// SetMsgType
//
//	@Description: 关联接口IContentMessage
//	@receiver *MatchMessage
func (m *MatchMessage) SetMsgType() {
	m.MsgType = MESSAGE_MACTH
}

package vo

type WsRequest struct {
	MatchId         string `json:"match_id"`           //比赛ID
	RedPacketRoomId string `json:"red_packet_room_id"` //红包房间ID
}

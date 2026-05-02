package channelData

import "bootpkg/common/global"

func SendHistoryBetRecordData(s string) error {
	err := safeSend(Kakfa_Topic_History_Bet_Record_Data, s)
	if err != nil {
		global.G_LOG.Errorf("SendHistoryBetRecordData err:%s", err.Error())
	}
	return err
}

func SendBetRecordData(s string) error {
	err := safeSend(Kakfa_Topic_Bet_Record_Data, s)
	if err != nil {
		global.G_LOG.Errorf("SendUserBetRecord err:%s", err.Error())
	}
	return err
}

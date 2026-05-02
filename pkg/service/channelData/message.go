package channelData

import "bootpkg/common/global"

func SendUserSiteMsgData(s string) error {
	err := safeSend(Kakfa_Topic_User_Site_Msg_Data, s)
	if err != nil {
		global.G_LOG.Errorf("SendUserSiteMsgData msg: %v err: %s", s, err.Error())
	}
	return err
}

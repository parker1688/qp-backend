package setnotify

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
)

/*
站点公告同步
*/
func SynUserSiteNotify(userinfo *dos.FcUserMaterial, clientType string, language string) {

	var row *dos.FcUserSiteNotify
	global.G_DB.Model(&dos.FcUserSiteNotify{}).Where("user_id=?", userinfo.UserId).Order("create_time DESC").Take(&row)
	var notifys []*dos.FcSiteNotify
	if row.Id == "" {
		global.G_DB.Model(&dos.FcSiteNotify{}).Where("merchant_code=? and language =? and notify_type=?", userinfo.MerchantCode, language, clientType).Find(&notifys)
	} else {
		global.G_DB.Model(&dos.FcSiteNotify{}).Where("merchant_code=? AND  create_time>? and language =? and notify_type=?", userinfo.MerchantCode, row.CreateTime, language, clientType).Find(&notifys)
	}

	for _, v := range notifys {
		modules.SaveFcUserSiteNotify(&dos.FcUserSiteNotify{
			UserId:       userinfo.UserId,
			NotifyId:     v.Id,
			Title:        v.Title,
			TitleImg:     v.TitleImg,
			Content:      v.Content,
			Language:     v.Language,
			Sort:         v.Sort,
			NotifyType:   v.NotifyType,
			CreateTime:   v.CreateTime,
			CreateBy:     v.CreateBy,
			UpdateBy:     v.UpdateBy,
			UpdateTime:   v.UpdateTime,
			Status:       1,
			ClassType:    v.ClassType,
			MerchantCode: v.MerchantCode,
		})
	}
}

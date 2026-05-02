// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 后台邮件系列方法
func SaveMail(vo *dos.Mail) error {
	return global.G_DB.Create(vo).Error
}

func FindPageMail(page, pageSize int, vo *dos.Mail) (ret []*dos.Mail, total int64) {
	query := global.G_DB.Model(&dos.Mail{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Type > -1 {
		query = query.Where("type = ?", vo.Type)
	} else {
		query = query.Where("type != ?", enmus.MailType_Manual)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.Mail
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func UpdateMail(vo *dos.Mail) error {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"merchant_code": vo.MerchantCode,
		"user_ids":      vo.UserIds,
		"type":          vo.Type,
		"title":         vo.Title,
		"content":       vo.Content,
		"is_popup":      vo.IsPopup,
		"is_keep":       vo.IsKeep,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"update_time":   automaticType.Now(),
		"status":        vo.Status,
	}).Error
}

func DelMail(mailId string) error {
	return global.G_DB.Model(&dos.Mail{}).Where("id = ?", mailId).
		Delete(&dos.Mail{}).Error
}

// 用户邮件系列方法
func SaveUserMailMulit(models []dos.FcUserMail) error {
	return global.G_DB.Create(&models).Error
}

func FindPageFcUserMail(page, pageSize int, vo *dos.FcUserMail, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcUserMail, total int64) {
	query := global.G_DB.Model(&dos.FcUserMail{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Type > -1 {
		if c != nil {
			query = query.Where("type = ?", vo.Type)
		} else {
			switch vo.Type {
			case 0: // 人工
				query = query.Where("type = ?", enmus.MailType_Manual)
			case 1: // 系统
				query = query.Where("type != ?", enmus.MailType_Manual)
			}
		}
	}

	if vo.IsPopup > -1 {
		query = query.Where("is_popup = ?", vo.IsPopup)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if vo.ReadStatus > -1 {
		query = query.Where("read_status = ?", vo.ReadStatus)
	}

	//if vo.DelStatus > -1 {
	if c == nil {
		query = query.Where("del_status = ?", enmus.MailDelStats_No)
	} else {
		query = query.Where("del_status != ?", enmus.MailDelStats_Destroy)
	}
	//}

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", vo.CreateTime)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", vo.CreateTime)
	}

	/*if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}*/

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserMail
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func DelUserMailByIds(ids []string) bool {
	return global.G_DB.Model(&dos.FcUserMail{}).Where("id in ?", ids).
		Delete(&dos.Mail{}).RowsAffected > 0
}

// SetUserMailDataByIds - 根据邮件id组更新用户邮件数据
// @param {[]string} ids
// @param {map[string]interface{}} sets
// @returns bool
func SetUserMailDataByIds(ids []string, sets map[string]interface{}) bool {
	return global.G_DB.Model(&dos.FcUserMail{}).Where("id in ?", ids).
		Updates(sets).RowsAffected > 0
}

func DelUserMail(vo *dos.FcUserMail) bool {
	return global.G_DB.Model(&dos.FcUserMail{}).Where("id = ? AND user_id = ?", vo.Id, vo.UserId).
		Delete(&dos.Mail{}).RowsAffected > 0
}

// DoUserMailAction - 处理用户邮件（人工邮件）
// @param {*dos.FcUserMaterial} u
// @returns
func DoUserMailAction(u *dos.FcUserMaterial) {
	startAt := time.Now().Add(-7 * 24 * time.Hour)
	registerAt := time.Time(u.CreateTime)
	if registerAt.After(startAt) {
		startAt = registerAt
	}
	startAtStr := startAt.Format(tool.TimeLayout)

	// 处理用户邮件过期
	err := global.G_DB.Model(&dos.FcUserMail{}).
		Where("user_id = ? AND create_time < ? AND del_status = ?", u.UserId, startAt, enmus.MailDelStats_No).
		Update("del_status", enmus.MailDelStats_Yes).Error
	if err != nil {
		global.G_LOG.Errorf("[DoUserMailAction] Update user mail expires failed: userId=%s, err=%s",
			u.UserId, err.Error())
	}

	// 处理人工邮件
	query := global.G_DB.Model(&dos.Mail{})
	query = query.Where("type = ?", enmus.MailType_Manual)
	query = query.Where("status = 1")
	query = query.Where("merchant_code = '' OR user_ids = ''")
	query = query.Where("create_time >= ?", startAtStr)
	data := []dos.Mail{}
	err = query.Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[DoUserMailAction] Find admin mails failed: userId=%s, err=%s",
			u.UserId, err.Error())
		return
	}

	// 需排除的邮件id组
	excludeIds := []string{}
	for _, v := range data {
		excludeIds = append(excludeIds, v.Id)
	}

	// 排除已获取的邮件
	excludeMails := []dos.FcUserMail{}
	err = global.G_DB.Model(&dos.FcUserMail{}).Select("msg_id").
		Where("user_id = ? AND msg_id in ?", u.UserId, excludeIds).Find(&excludeMails).Error
	if err != nil {
		global.G_LOG.Errorf("[DoUserMailAction] Find user mails failed: userId=%s, err=%s",
			u.UserId, err.Error())
		return
	}

	excludeIdMap := map[string]int{}
	for _, v := range excludeMails {
		excludeIdMap[v.MsgId] = 1 // 排除后台邮件
	}

	mails := []dos.FcUserMail{}
	for _, v := range data {
		if v.MerchantCode != "" &&
			u.MerchantCode != v.MerchantCode { // 需要排除指定商户的邮件
			continue
		}

		if v.UserIds != "" &&
			!strings.Contains(v.UserIds, u.UserId) {
			// 一般不会出现这种情况，指定用户一般在后台发送就直接发送进用户邮箱了
			continue
		}

		merchantCode := v.MerchantCode
		if v.MerchantCode == "" {
			merchantCode = u.MerchantCode
		}

		if _, ok := excludeIdMap[v.Id]; !ok {
			mails = append(mails, dos.FcUserMail{
				MsgId:        v.Id,
				UserId:       u.UserId,
				MerchantCode: merchantCode,
				Type:         v.Type,
				Title:        v.Title,
				Content:      v.Content,
				IsPopup:      v.IsPopup,
				IsKeep:       v.IsKeep,
				CreateTime:   automaticType.Now(),
				ReadStatus:   enmus.MailStats_Unread,
				DelStatus:    enmus.MailDelStats_No,
			})
		}
	}

	if len(mails) == 0 {
		return
	}

	err = SaveUserMailMulit(mails)
	if err != nil {
		global.G_LOG.Errorf("[DoUserMailAction] save user mails failed: userId=%s, data=%+v, err=%s",
			u.UserId, mails, err.Error())
	}
}

// DoUserSystemEmail - 处理用户系统邮件
// @param {dos.FcUserMaterial} user
// @param {string} merchantCode
// @param {int} mailType
// @returns
func DoUserSystemEmail(user *dos.FcUserMaterial, mailType int) {
	data := []dos.Mail{}
	err := global.G_DB.Model(&dos.Mail{}).Where("type = ? AND status = 1", mailType).
		Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[DoUserSystemEmail] Find admin mail failed: userId=%s, merchantCode=%s, mailType=%d, err=%s",
			user.UserId, user.MerchantCode, mailType, err.Error())
		return
	}

	if len(data) == 0 {
		return
	}

	excludeIds := []string{}
	for _, v := range data {
		excludeIds = append(excludeIds, v.Id)
	}

	excludeMails := []dos.FcUserMail{}
	err = global.G_DB.Model(&dos.FcUserMail{}).
		Where("user_id = ? AND msg_id in ?", user.UserId, excludeIds).
		Find(&excludeMails).Error
	if err != nil {
		global.G_LOG.Errorf("[DoUserSystemEmail] Find user exclude mail failed: userId=%s, merchantCode=%s, mailType=%d, err=%s",
			user.UserId, user.MerchantCode, mailType, err.Error())
		return
	}

	excludeMap := map[string]int{}
	for _, v := range excludeMails {
		excludeMap[v.MsgId] = 1
	}

	mails := []dos.FcUserMail{}
	for _, v := range data {
		if v.MerchantCode != "" &&
			v.MerchantCode != user.MerchantCode {
			// 排除指定商户
			continue
		}

		if v.UserIds != "" &&
			!strings.Contains(v.UserIds, user.UserId) {
			// 排除指定用户
			continue
		}

		if _, ok := excludeMap[v.Id]; !ok {
			var newContent string
			switch mailType {
			case enmus.MailType_FirstLogin:
				newContent = tool.PlaceholderFormat(v.Content, user.UserName)
			}

			merchantCode := v.MerchantCode
			if v.MerchantCode == "" {
				merchantCode = user.MerchantCode
			}

			mails = append(mails, dos.FcUserMail{
				MsgId:        v.Id,
				UserId:       user.UserId,
				MerchantCode: merchantCode,
				Type:         v.Type,
				Title:        v.Title,
				Content:      newContent,
				IsPopup:      v.IsPopup,
				IsKeep:       v.IsKeep,
				CreateTime:   automaticType.Now(),
				ReadStatus:   enmus.MailStats_Unread,
				DelStatus:    enmus.MailDelStats_No,
			})
		}
	}

	if len(mails) > 0 {
		err = SaveUserMailMulit(mails)
		if err != nil {
			global.G_LOG.Errorf("[DoUserSystemEmail] Save user new mails failed: userId=%s, merchantCode=%s, mailType=%d, err=%s",
				user.UserId, user.MerchantCode, mailType, err.Error())
		}
	}
}

// 通用发送系统邮件入口
func SendSystemMail(userId, merchantCode string, mailType int, args ...interface{}) {
	//global.G_LOG.Infof("send sysmail -----------------------------1:%v, %v, %v, %v", userId, merchantCode, mailType, args[0])
	data := []dos.Mail{}
	err := global.G_DB.Model(&dos.Mail{}).Where("type = ? AND status = 1", mailType).
		Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[DoUserSystemEmail] Find admin mail failed: userId=%s, merchantCode=%s, mailType=%d, err=%s",
			userId, merchantCode, mailType, err.Error())
		return
	}
	//global.G_LOG.Infof("send sysmail -----------------------------2:%v, %v", data, len(data))

	if len(data) == 0 {
		return
	}

	mails := []dos.FcUserMail{}
	for _, v := range data {
		if v.MerchantCode != "" &&
			v.MerchantCode != merchantCode {
			// 排除指定商户
			continue
		}

		if v.UserIds != "" &&
			!strings.Contains(v.UserIds, userId) {
			// 排除指定用户
			continue
		}

		var newContent string
		switch mailType {
		case enmus.MailType_RechargeSuccess:
			newContent = tool.ReplaceTemplate(v.Content, args[0])
		case enmus.MailType_WithdrawSuccess:
			newContent = tool.ReplaceTemplate(v.Content, args[0])
		case enmus.MailType_WithdrawFail:
			newContent = tool.ReplaceTemplate(v.Content, args[0])
		}
		//global.G_LOG.Infof("send sysmail -----------------------------3:%v, %v", userId, newContent)

		merchantCode1 := v.MerchantCode
		if v.MerchantCode == "" {
			merchantCode1 = merchantCode
		}
		mails = append(mails, dos.FcUserMail{
			MsgId:        v.Id,
			UserId:       userId,
			MerchantCode: merchantCode1,
			Type:         v.Type,
			Title:        v.Title,
			Content:      newContent,
			IsPopup:      v.IsPopup,
			IsKeep:       v.IsKeep,
			CreateTime:   automaticType.Now(),
			ReadStatus:   enmus.MailStats_Unread,
			DelStatus:    enmus.MailDelStats_No,
		})
	}
	//global.G_LOG.Infof("send sysmail -----------------------------4:%v, %v", mails, len(mails))

	if len(mails) > 0 {
		err = SaveUserMailMulit(mails)
		if err != nil {
			global.G_LOG.Errorf("[DoUserSystemEmail] Save user new mails failed: userId=%s, merchantCode=%s, mailType=%d, err=%s",
				userId, merchantCode, mailType, err.Error())
		}
	}
}

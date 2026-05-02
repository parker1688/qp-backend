// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveDailyTask(vo *dos.DailyTask) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func SaveDailyTaskMulit(models []dos.DailyTask) error {
	return global.G_DB.Create(&models).Error
}

func FindPageDailyTask(pageQuery response.PageTimeQuery, vo *dos.DailyTask, c *gin.Context) (ret []*dos.DailyTaskEx, total int64) {
	response.NormalizePageTimeQuery(&pageQuery)
	DoDailyTaskStats() // 处理任务状态

	query := global.G_DB.Model(&dos.DailyTask{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.GroupId) > 0 {
		//global.G_LOG.Infof("dailyTask data ------------------------------1:%v", vo.GroupId)
		query = query.Where("groupid = ?", vo.GroupId)
	}

	if len(pageQuery.StartAt) > 0 {
		query = query.Where("start_at >= ?", pageQuery.StartAt)
	}

	if len(pageQuery.EndAt) > 0 {
		query = query.Where("end_at <= ?", pageQuery.EndAt)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	dataSlice := []*dos.DailyTaskEx{}
	query.Offset((pageQuery.PageNo - 1) * pageQuery.PageSize).
		Preload("Merchant").
		Order("sort").
		Limit(pageQuery.PageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyDailyTask(vo *dos.DailyTask, c *gin.Context) []*dos.DailyTaskEx {
	DoDailyTaskStats() // 处理任务状态

	var data []*dos.DailyTaskEx
	query := global.G_DB.Model(&dos.DailyTask{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.StartAt != nil && !vo.StartAt.Timer().IsZero() {
		query = query.Where("start_at >= ?", vo.StartAt)
	}

	if vo.EndAt != nil && !vo.EndAt.Timer().IsZero() {
		query = query.Where("end_at <= ?", vo.EndAt)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Preload("Merchant").Order("sort").Find(&data)
	return data
}

func FindByKeyDailyTaskFirst(vo *dos.DailyTask) *dos.DailyTask {
	var data *dos.DailyTask
	query := global.G_DB.Model(&dos.DailyTask{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	query.Take(&data)
	return data
}

func UpdateDailyTask(vo *dos.DailyTask) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"merchant_code":      vo.MerchantCode,
		"type":               vo.Type,
		"subtype":            vo.Subtype,
		"name":               vo.Name,
		"sort":               vo.Sort,
		"intro":              vo.Intro,
		"detail":             vo.Detail,
		"amount":             vo.Amount,
		"bonus_amount":       vo.BonusAmount,
		"game_type":          vo.GameType,
		"venue_code":         vo.VenueCode,
		"channel_code":       vo.ChannelCode,
		"start_at":           vo.StartAt,
		"end_at":             vo.EndAt,
		"update_time":        automaticType.Time(time.Now()),
		"update_by":          vo.UpdateBy,
		"include_game_codes": vo.IncludeGameCodes,
		"exclude_game_codes": vo.ExcludeGameCodes,
		"cycle":              vo.Cycle,
		"status":             vo.Status,
	}).Error == nil
}

func DeleteDailyTask(vo *dos.DailyTask) bool {
	return global.G_DB.Model(&dos.DailyTask{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// DoDailyTaskStats - 处理任务状态
func DoDailyTaskStats() {
	// 进行中任务处理是否结束
	err := global.G_DB.Model(&dos.DailyTask{}).Where("status = ? AND end_at < ?", enmus.DailyTaskStats_Normal, time.Now()).
		Update("status", enmus.DailyTaskStats_Over).Error
	if err != nil {
		global.G_LOG.Errorf("[DoDailyTaskStats] Update normal daily task status failed: %v", err.Error())
	}

	// 结束中任务处理是否进行
	err = global.G_DB.Model(&dos.DailyTask{}).Where("status = ? AND (end_at > ? OR end_at IS NULL) AND type != ?",
		enmus.DailyTaskStats_Over, time.Now(), enmus.DailyTaskType_Daily).
		Update("status", enmus.DailyTaskStats_Normal).Error
	if err != nil {
		global.G_LOG.Errorf("[DoDailyTaskStats] Update over daily task status failed: %v", err.Error())
	}
}

// CheckDailyTaskGameCodes - 验证游戏码是否存在
// @param {string} venueCode
// @param {string} gameType
// @param {string} gameCodes
// @returns string, error
func CheckDailyTaskGameCodes(venueCode, gameType, gameCodes string) (string, error) {
	if len(gameCodes) == 0 {
		return "", nil
	}

	query := global.G_DB.Model(&dos.FcVenueGame{})

	venueCodes := strings.Split(venueCode, ",")
	if len(venueCodes) > 0 {
		query = query.Where("venue_code in ?", venueCodes)
	}

	gameTypes := strings.Split(gameType, ",")
	if len(gameType) > 0 {
		query = query.Where("game_type in ?", gameTypes)
	}

	gameCodeLis := strings.Split(gameCodes, ",")

	data := []dos.FcVenueGame{}
	err := query.Select("game_code").
		Where("game_code in ?", gameCodeLis).
		Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckDailyTaskGameCodes] Find venue games failed: gameCodes=%s, err=%s",
			gameCodes, err.Error())
		return "", err
	}

	if len(data) == 0 {
		return gameCodes, nil
	}

	tmps := map[string]int{}
	for _, v := range data {
		tmps[v.GameCode] = 1
	}

	errGameCodes := []string{}
	for _, gameCode := range gameCodeLis {
		if _, ok := tmps[gameCode]; !ok {
			// 不存在的游戏码
			errGameCodes = append(errGameCodes, gameCode)
		}
	}

	if len(errGameCodes) > 0 {
		return strings.Join(errGameCodes, ","), nil
	}

	return "", nil
}

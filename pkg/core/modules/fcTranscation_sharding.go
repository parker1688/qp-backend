// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getTranscationDB() *gorm.DB {
	if global.G_DB_SHARDING != nil {
		return global.G_DB_SHARDING
	}
	return global.G_DB
}

func SaveFcTranscationSharding(vo *dos.FcTranscation) (bool, ecode.Code) {
	err := getTranscationDB().Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcTranscationSharding(page, pageSize int, vo *dos.FcTranscationSharding, pageQuery response.PageTimeQuery, c *gin.Context) (ret []*dos.FcTranscationSharding, total int64) {
	page, pageSize = response.NormalizePage(page, pageSize)
	query := getTranscationDB().Model(&dos.FcTranscationSharding{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		userId := tool.Int(vo.UserId)
		query = query.Where("user_id = ?", tool.String(userId))
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.FundingType > -1 {
		query = query.Where("funding_type = ?", vo.FundingType)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(pageQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageQuery.StartAt)
	}

	if len(pageQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageQuery.EndAt)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcTranscationSharding
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcTranscationSharding(vo *dos.FcTranscation, c *gin.Context) []*dos.FcTranscation {
	var data []*dos.FcTranscation
	query := getTranscationDB().Model(&dos.FcTranscation{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		userId := tool.Int(vo.UserId)
		query = query.Where("user_id = ?", tool.String(userId))
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.FundingType > 0 {
		query = query.Where("funding_type = ?", vo.FundingType)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
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

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcTranscationSharding(vo *dos.FcTranscation) bool {
	return getTranscationDB().Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":            vo.UserId,
		"user_name":          vo.UserName,
		"status":             vo.Status,
		"amount":             vo.Amount,
		"transcation_before": vo.TranscationBefore,
		"transcation_after":  vo.TranscationAfter,
		"remark":             vo.Remark,
		"funding_type":       vo.FundingType,
		"create_by":          vo.CreateBy,
		"update_by":          vo.UpdateBy,
		"merchant_code":      vo.MerchantCode,
	}).Error == nil
}

func DeleteFcTranscationSharding(vo *dos.FcTranscation) bool {
	return getTranscationDB().Model(&dos.FcTranscation{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func FindByKeyFcTranscationFirst(vo *dos.FcTranscation) *dos.FcTranscation {
	var data *dos.FcTranscation
	query := getTranscationDB().Model(&dos.FcTranscation{})

	if len(vo.UserId) > 0 {
		query = query.Where("id = ?", vo.UserId)
	}

	query.Take(&data)
	return data
}

// GetFcTranscationFundingSubType - 获取资金流水子类型
// @param {int} fundingType 资金类型（大类）
// @param {string} key 查询键
// @returns bool
func GetFcTranscationFundingSubType(fundingType int, key string) string {
	if _, ok := enmus.FundingSubTypeEnums[fundingType]; ok {
		if val, ok := enmus.FundingSubTypeEnums[fundingType][key]; ok {
			return val
		}
	}

	return ""
}

func getFcTranscationQuery(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, query *gorm.DB, c *gin.Context) (*gorm.DB, bool) {
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		userId := tool.Int(vo.UserId)
		query = query.Where("user_id = ?", tool.String(userId))
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.FundingType > 0 {
		query = query.Where("funding_type = ?", vo.FundingType)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(timeQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", timeQuery.StartAt)
	}

	if len(timeQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", timeQuery.EndAt)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return query, false
		}
	}

	return query, true
}

// AcumulateQuotaCovertStatisData - 统计额度转换数据
// @returns map[string]interface{}
func AcumulateQuotaCovertStatisData(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, c *gin.Context) map[string]interface{} {
	var depositTotalAmount float64
	depositItems := []response.ItemResult{}
	var withdrawTotalAmount float64
	withdrawItems := []response.ItemResult{}

	var ok bool
	venueMap := GetFcVenueNameMap([]string{})
	venueList := GetFcVenueList()
	inVenues := []string{}
	outVenues := []string{}
	for _, venue := range venueList {
		inVenues = append(inVenues, "转入"+venue.VenueCode)
		outVenues = append(outVenues, "转出"+venue.VenueCode)
	}

	query := global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	// 转入场馆统计
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ?", inVenues).Group("funding_subtype").Find(&depositItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulateQuotaCovertStatisData] sum venue deposit total failed: %v", err.Error())
		}
	}
	for k, v := range depositItems {
		depositItems[k].Value = tool.TruncateFloat(v.Value, 2)
		depositTotalAmount += depositItems[k].Value
		depositItems[k].Name = venueMap[strings.ReplaceAll(v.Name, "转入", "")] + "转入"
	}

	query = global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	// 转出场馆统计
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ?", outVenues).Group("funding_subtype").Find(&withdrawItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulateQuotaCovertStatisData] sum venue withdraw total failed: %v", err.Error())
		}
	}
	for k, v := range withdrawItems {
		withdrawItems[k].Value = tool.TruncateFloat(v.Value, 2)
		withdrawTotalAmount += withdrawItems[k].Value
		withdrawItems[k].Name = venueMap[strings.ReplaceAll(v.Name, "转出", "")] + "转出"
	}

	result := map[string]interface{}{}
	result["depositSumAmount"] = tool.TruncateFloat(depositTotalAmount, 2)
	result["depositItems"] = depositItems
	result["withdrawSumAmount"] = tool.TruncateFloat(withdrawTotalAmount, 2)
	result["withdrawItems"] = withdrawItems

	return result
}

// AcumulateOnlineDepositStatisData - 统计在线存款数据
// @returns map[string]interface{}
func AcumulateOnlineDepositStatisData(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, c *gin.Context) map[string]interface{} {
	var sumAmount float64
	depositItems := []response.ItemResult{}

	var ok bool

	categoryLis := enmus.ConstTransacationStatisDesposit
	query := global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ?", categoryLis).Group("funding_subtype").Find(&depositItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulateOnlineDepositStatisData] sum online deposit total failed: %v", err.Error())
		}
	}
	for k, v := range depositItems {
		depositItems[k].Value = tool.TruncateFloat(v.Value, 2)
		sumAmount += depositItems[k].Value
	}

	result := map[string]interface{}{}
	result["sumAmount"] = tool.TruncateFloat(sumAmount, 2)
	result["items"] = depositItems

	return result
}

// AcumulateManualDepositStatisData - 统计手动存款数据
// @returns map[string]interface{}
func AcumulateManualDepositStatisData(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, c *gin.Context) map[string]interface{} {
	var sumAmount float64
	depositItems := []response.ItemResult{}

	var ok bool

	categoryLis := enmus.ConstTransacationStatisManual
	query := global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ?", categoryLis).Group("funding_subtype").Find(&depositItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulateManualDepositStatisData] sum manual deposit total failed: %v", err.Error())
		}
	}
	for k, v := range depositItems {
		depositItems[k].Value = tool.TruncateFloat(v.Value, 2)
		sumAmount += depositItems[k].Value
	}

	result := map[string]interface{}{}
	result["sumAmount"] = tool.TruncateFloat(sumAmount, 2)
	result["items"] = depositItems

	return result
}

// AcumulateWithdrawStatisData - 统计提款数据
// @returns map[string]interface{}
func AcumulateWithdrawStatisData(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, c *gin.Context) map[string]interface{} {
	var sumAmount float64
	depositItems := []response.ItemResult{}

	var ok bool

	categoryLis := enmus.ConstTransacationStatisWithdraw
	query := global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ? AND status = 1 AND funding_type = 4",
			categoryLis).Group("funding_subtype").
			Find(&depositItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulateWithdrawStatisData] sum withdraw total failed: %v", err.Error())
		}
	}
	for k, v := range depositItems {
		depositItems[k].Value = tool.TruncateFloat(math.Abs(v.Value), 2)
		sumAmount += depositItems[k].Value
	}

	result := map[string]interface{}{}
	result["sumAmount"] = tool.TruncateFloat(sumAmount, 2)
	result["items"] = depositItems

	return result
}

// AcumulatePromotionStatisData - 统计优惠数据
// @returns map[string]interface{}
func AcumulatePromotionStatisData(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, c *gin.Context) map[string]interface{} {
	var sumAmount float64
	depositItems := []response.ItemResult{}

	var ok bool
	lis := enmus.ConstTransacationStatisPromotion

	categoryLis := GetFcWelfareTitles("彩金活动")
	for _, v := range categoryLis {
		lis = append(lis, v)
	}

	query := global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ?", lis).Group("funding_subtype").Find(&depositItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulatePromotionStatisData] sum promotion total failed: %v", err.Error())
		}
	}
	for k, v := range depositItems {
		depositItems[k].Value = tool.TruncateFloat(v.Value, 2)
		sumAmount += depositItems[k].Value
	}

	result := map[string]interface{}{}
	result["sumAmount"] = tool.TruncateFloat(sumAmount, 2)
	result["items"] = depositItems

	return result
}

// AcumulateRebateStatisData - 统计返水数据
// @returns map[string]interface{}
func AcumulateRebateStatisData(vo *dos.FcTranscationSharding, timeQuery response.PageTimeQuery, c *gin.Context) map[string]interface{} {
	var sumAmount float64
	rebateItems := []response.ItemResult{}

	var ok bool
	venueMap := GetFcVenueNameMap([]string{})
	venueList := GetFcVenueList()
	rebateVenues := []string{}
	for _, venue := range venueList {
		rebateVenues = append(rebateVenues, venue.VenueCode+"平台洗码")
	}

	query := global.G_DB.Model(&dos.FcTranscation{})
	query, ok = getFcTranscationQuery(vo, timeQuery, query, c)
	// 反水场馆统计
	if ok {
		if err := query.Select("funding_subtype as `name`, IFNULL(sum(amount), 0) as `value`").Where("funding_subtype in ?", rebateVenues).Group("funding_subtype").Find(&rebateItems).Error; err != nil {
			global.G_LOG.Errorf("[AcumulateRebateStatisData] sum venue rebate total failed: %v", err.Error())
		}
	}
	for k, v := range rebateItems {
		rebateItems[k].Value = tool.TruncateFloat(v.Value, 2)
		sumAmount += rebateItems[k].Value
		rebateItems[k].Name = venueMap[strings.ReplaceAll(v.Name, "平台洗码", "")] + "反水"
	}

	result := map[string]interface{}{}
	result["sumAmount"] = tool.TruncateFloat(sumAmount, 2)
	result["items"] = rebateItems
	return result
}

// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
)

func SaveFcClientLog(vo *dos.FcClientLog) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcClientLog(page, pageSize int, vo *dos.FcClientLog, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcClientLog, total int64) {
	query := global.G_DB.Model(&dos.FcClientLog{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(pageTimeQuery.StartAt) > 0 {
		query = query.Where("visit_time >= ?", pageTimeQuery.StartAt)
	}
	if len(pageTimeQuery.EndAt) > 0 {
		query = query.Where("visit_time <= ?", pageTimeQuery.EndAt)
	}

	if len(vo.IP) > 0 {
		query = query.Where("ip = ?", vo.IP)
	}

	if vo.AgentId > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}
	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcClientLog
	query.Order("visit_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

// 根据主键Update
func UpdateFcClientLog(vo *dos.FcClientLog) bool {
	res := global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"agent_id":      vo.AgentId,
		"merchant_code": vo.MerchantCode,
		"merchant_name": vo.MerchantName,
		"ip":            vo.IP,
		"device":        vo.Device,
		"address":       vo.Address,
		"download":      vo.Download,
		"customer":      vo.Customer,
		"visit_time":    vo.VisitTime,
	}).Error
	return res == nil
}

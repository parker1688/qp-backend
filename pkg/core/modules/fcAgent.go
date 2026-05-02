// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcAgent(vo *dos.FcAgent) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcAgent(page, pageSize int, vo *dos.FcAgent, c *gin.Context) (ret []*dos.FcAgent, total int64) {
	query := global.G_DB.Model(&dos.FcAgent{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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
	var dataSlice []*dos.FcAgent
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcAgent(vo *dos.FcAgent, c *gin.Context) []*dos.FcAgent {
	var data []*dos.FcAgent
	query := global.G_DB.Model(&dos.FcAgent{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	query.Find(&data)
	return data
}

func FindByKeyFcAgentFirst(vo *dos.FcAgent) *dos.FcAgent {
	var data *dos.FcAgent
	query := global.G_DB.Model(&dos.FcAgent{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
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

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcAgent(vo *dos.FcAgent) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"agent_name":    vo.AgentName,
		"invite_code":   vo.InviteCode,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"status":        vo.Status,
		"merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcAgent(vo *dos.FcAgent) bool {
	return global.G_DB.Model(&dos.FcAgent{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SaveFcAgentDomain(vo *dos.FcAgentDomain) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcAgentDomain(page, pageSize int, vo *dos.FcAgentDomain, c *gin.Context) (ret []*dos.FcAgentDomainExt, total int64) {
	query := global.G_DB.Model(&dos.FcAgentDomain{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Type > 0 {
		query = query.Where("`type` = ?", vo.Type)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
	}

	if len(vo.Domain) > 0 {
		query = query.Where("domain = ?", vo.Domain)
	}

	if len(vo.ShortLink) > 0 {
		query = query.Where("short_link = ?", vo.ShortLink)
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

	if len(vo.AgentId) > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.JumpLink) > 0 {
		query = query.Where("jump_link = ?", vo.JumpLink)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcAgentDomainExt
	query.Offset((page - 1) * pageSize).Limit(pageSize).Preload("Merchant").Preload("MerchantLink").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcAgentDomain(vo *dos.FcAgentDomain, c *gin.Context) []*dos.FcAgentDomain {
	var data []*dos.FcAgentDomain
	query := global.G_DB.Model(&dos.FcAgentDomain{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Type > 0 {
		query = query.Where("`type` = ?", vo.Type)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
	}

	if len(vo.Domain) > 0 {
		query = query.Where("domain = ?", vo.Domain)
	}

	if len(vo.ShortLink) > 0 {
		query = query.Where("short_link = ?", vo.ShortLink)
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

	if len(vo.AgentId) > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.JumpLink) > 0 {
		query = query.Where("jump_link = ?", vo.JumpLink)
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

func FindByKeyFcAgentDomainFirst(vo *dos.FcAgentDomain) *dos.FcAgentDomain {
	var data *dos.FcAgentDomain
	query := global.G_DB.Model(&dos.FcAgentDomain{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.Type > 0 {
		query = query.Where("`type` = ?", vo.Type)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
	}

	if len(vo.Domain) > 0 {
		query = query.Where("domain = ?", vo.Domain)
	}

	if len(vo.ShortLink) > 0 {
		query = query.Where("short_link = ?", vo.ShortLink)
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

	if len(vo.AgentId) > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.JumpLink) > 0 {
		query = query.Where("jump_link = ?", vo.JumpLink)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcAgentDomain(vo *dos.FcAgentDomain) bool {
	fmt.Printf("[UpdateFcAgentDomain] id=%s\n", vo.Id)
	tx := global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"agent_name":     vo.AgentName,
		"invite_code":    vo.InviteCode,
		"domain":         vo.Domain,
		"short_link":     vo.ShortLink,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"agent_id":       vo.AgentId,
		"merchant_code":  vo.MerchantCode,
		"status":         vo.Status,
		"jump_link":      vo.JumpLink,
		"type":           vo.Type,
		//"customer_link": vo.CustomerLink,
		"ios_link":       vo.IosLink,
		"ios_link2":      vo.IosLink2,
		"android_link":   vo.AndroidLink,
		"android_link2":  vo.AndroidLink2,
		"banner_img":     vo.BannerImg,
		"logo_img":       vo.LogoImg,
	})
	if tx.Error != nil {
		fmt.Printf("[UpdateFcAgentDomain] FAILED error=%v\n", tx.Error)
		return false
	}
	fmt.Printf("[UpdateFcAgentDomain] rows_affected=%d\n", tx.RowsAffected)
	return tx.RowsAffected > 0
}

func DeleteFcAgentDomain(vo *dos.FcAgentDomain) bool {
	return global.G_DB.Model(&dos.FcAgentDomain{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func GetAgentDomainMerchantCodeByHeader(c *gin.Context) string {
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		// 查找推广ID
		inviteCode := c.GetHeader(vo.MerchantID_KEY_G)
		if len(inviteCode) == 0 {
			// 查找域名地址
			domainUrl := c.GetHeader(vo.MerchantUrl_KEY_G)
			if len(domainUrl) > 0 {
				agentDomain := dos.FcAgentDomain{}
				global.G_DB.Model(&dos.FcAgentDomain{}).Select("merchant_code").
					Where("jump_link like ?", "%"+strings.Trim(domainUrl, "/")+"%").Find(&agentDomain)
				merchantCode = agentDomain.MerchantCode
			}
		} else {
			agentDomain := dos.FcAgentDomain{}
			global.G_DB.Model(&dos.FcAgentDomain{}).Select("merchant_code").
				Where("invite_code = ?", inviteCode).Find(&agentDomain)
			merchantCode = agentDomain.MerchantCode
		}
	}

	if len(merchantCode) == 0 {
		global.G_LOG.Errorf("[GetAgentDomainMerchantCode] Find merchant code failed merchantCode=%s, merchantId=%s, merchantUrl=%s",
			c.GetHeader(vo.MerchantCode_KEY_G),
			c.GetHeader(vo.MerchantID_KEY_G),
			c.GetHeader(vo.MerchantUrl_KEY_G))
	}

	return merchantCode
}

// 获取推广域名代理列表项
func GetFcAgentDomainAgentOptions(c *gin.Context, pageQuery response.PageQuery, vo *dos.FcAgentDomain) (ret []*dos.FcAgentDomainOptionResp, total int64) {
	data := []*dos.FcMerchant{}

	query := global.G_DB.Model(&dos.FcMerchant{})

	ok := true
	if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
		return ret, total
	}

	var count int64
	query.Count(&count)

	query.Select("agent_invite_code").Find(&data)

	customerLink := SyncFcAgentDomainCustomerLink(vo.MerchantCode, false)

	list := []*dos.FcAgentDomainOptionResp{}
	for _, v := range data {
		list = append(list, &dos.FcAgentDomainOptionResp{
			Id:           strconv.Itoa(v.AgentInviteCode),
			CustomerLink: customerLink,
		})
	}

	return list, count
}

// 获取推广域名推广列表项
func GetFcAgentDomainInviteOptions(c *gin.Context, pageQuery response.PageQuery, vo *dos.FcAgentDomain) (ret []*dos.FcAgentDomainOptionResp, total int64) {
	data := []*dos.FcAgent{}

	query := global.G_DB.Model(&dos.FcAgent{})

	ok := true
	if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
		return ret, total
	}

	var count int64
	query.Count(&count)

	query.Select("invite_code").Find(&data)

	agentDomain := FindByKeyFcAgentDomainFirst(&dos.FcAgentDomain{
		MerchantCode: vo.MerchantCode,
		Type:         enmus.AgentDomainType_Agent,
	})

	customerLink := SyncFcAgentDomainCustomerLink(vo.MerchantCode, false)

	list := []*dos.FcAgentDomainOptionResp{}
	for _, v := range data {
		list = append(list, &dos.FcAgentDomainOptionResp{
			Id:           strconv.Itoa(v.InviteCode),
			CustomerLink: customerLink,
			ShortLink:    agentDomain.ShortLink,
			JumpLink:     agentDomain.JumpLink,
			Domain:       agentDomain.Domain,
		})
	}

	return list, count
}

// 同步客服链接到推广域名中
func SyncFcAgentDomainCustomerLink(merchantCode string, isSave bool) string {
	if len(merchantCode) == 0 {
		global.G_LOG.Errorf("[SyncFcAgentDomainCustomerLink] merchantCode is empty")
		return ""
	}

	customerLink := dos.FcCustomerLink{}
	err := global.G_DB.Model(&dos.FcCustomerLink{}).Select("link").
		Where("merchant_code = ?", merchantCode).First(&customerLink).Error
	if err != nil {
		global.G_LOG.Errorf("[SyncFcAgentDomainCustomerLink] find customer link failed: %v", err.Error())
		return ""
	}

	if isSave {
		err = global.G_DB.Model(&dos.FcAgentDomain{}).Where("merchant_code = ?", merchantCode).
			Update("customer_link", customerLink.Link).Error
		if err != nil {
			global.G_LOG.Errorf("[SyncFcAgentDomainCustomerLink] update agent domain customer link failed: %v", err.Error())
			return customerLink.Link
		}
	}

	return customerLink.Link
}

package fcVipReport

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func VipReport(c *gin.Context) {
	jsonp := struct {
		vo.VipReportReq
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	merchantNames := jsonp.MerchantName
	merchantCodes := jsonp.MerchantCode
	if merchantCodes == "" {
		userInfo, ok := c.Get("UserInfo")
		if ok {
			merchantCodes = userInfo.(*dos.AdminUser).MerchantCodes
		}
	}
	//query.Where("merchant_code in (?)", merchantCodes2)
	dataList := []vo.VipReportData{}
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	merchantNames2 := tool.StrToListForSql(merchantNames)
	//global.G_LOG.Infof("merchants --------------------:%v, %v", merchantCodes, merchantCodes2)
	for id, mc := range merchantCodes2 {
		query1 := global.G_DB.Model(&dos.FcUserMaterial{})
		query2 := global.G_DB.Model(&dos.FcUserMaterial{})
		query1.Where("merchant_code = ? ", mc)
		query2.Where("merchant_code = ? ", mc)
		num1 := 0
		query1.Select("count(user_id) as num1").Scan(&num1)
		num2 := 0
		query2.Select("count(user_id) as num2").Where("level > 0 ").Scan(&num2)
		vipList := []int{}
		for lv := 0; lv <= 30; lv++ {
			query3 := global.G_DB.Model(&dos.FcUserMaterial{})
			query3.Where("merchant_code = ? ", mc)
			num := 0
			query3.Select("count(user_id) as num").Where("level = ? ", lv).Scan(&num)
			vipList = append(vipList, num)
		}
		name := merchantNames2[id].(string)
		rec := vo.VipReportData{MerchantName: name, UserCount: num1, VipCount: num2, VipList: vipList}
		dataList = append(dataList, rec)
	}
	//data := vo.VipReportResp{VipData: dataList}
	//response.SuccessJSON(c, data)
	total := len(dataList)
	response.SuccessPageJSON(c, 1, 100000, int64(total), dataList)
}

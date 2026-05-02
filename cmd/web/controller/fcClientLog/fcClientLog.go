package fcClientLog

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcClientLog/save
func SaveFcClinetLogControl(c *gin.Context) {
	var jsonp dos.FcClientLog
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

	//if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
	//	return
	//}

	jsonp.Id = tool.SnowflakeIdByKey("fcClientLog-id")
	_, err = modules.SaveFcClientLog(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, jsonp.Id)
}

// api: api/fcClientLog/findPage
func FindPageFcClinetLogControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcClientLog
	}{}

	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "") // 投注时间
	jsonp.EndAt = c.DefaultQuery("endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.IP = c.DefaultQuery("ip", "")
	jsonp.AgentId = tool.Atoi(c.DefaultQuery("agent_id", ""))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//global.G_LOG.Infof("client-log -----------------------1:%v", jsonp)
	data, total := modules.FindPageFcClientLog(jsonp.PageNo, jsonp.PageSize, &jsonp.FcClientLog, &jsonp.PageTimeQuery, c)
	//global.G_LOG.Infof("client-log -----------------------2:%v, %v", data, total)
	//for _, v := range data {
	//global.G_LOG.Infof("client-log -----------------------3:%v, %v", v.VisitTime, v)
	//}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcClientLog/update
func UpdateFcClientLogControl(c *gin.Context) {
	var jsonp dos.FcClientLog
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
	//global.G_LOG.Infof("update client-log ----------------1:%v, %v", jsonp.Id, jsonp)
	data := modules.UpdateFcClientLog(&jsonp)
	response.SuccessJSON(c, data)
}

func Statics(c *gin.Context) {
	count := 0
	count1 := 0
	count2 := 0
	count3 := 0
	query1 := global.G_DB.Model(&dos.FcClientLog{})
	query1.Where("").Select("count(1)").Scan(&count)
	query2 := global.G_DB.Model(&dos.FcClientLog{})
	query2.Where("device = 0").Select("count(1)").Scan(&count3)
	query3 := global.G_DB.Model(&dos.FcClientLog{})
	query3.Where("device = 1").Select("count(1)").Scan(&count1)
	query4 := global.G_DB.Model(&dos.FcClientLog{})
	query4.Where("device = 2").Select("count(1)").Scan(&count2)

	downcount := 0
	query5 := global.G_DB.Model(&dos.FcClientLog{})
	query5.Where("download = 1 ").Select("count(1)").Scan(&downcount)
	downcount1 := 0
	query6 := global.G_DB.Model(&dos.FcClientLog{})
	query6.Where("device = 1 and download = 1 ").Select("count(1)").Scan(&downcount1)
	downcount2 := 0
	query7 := global.G_DB.Model(&dos.FcClientLog{})
	query7.Where("device = 2 and download = 1").Select("count(1)").Scan(&downcount2)
	customercount := 0
	query8 := global.G_DB.Model(&dos.FcClientLog{})
	query8.Where("customer = 1 ").Select("count(1)").Scan(&customercount)

	data := dos.FcClientLogResp{TotalVisitCount: count, VisitCount1: count1, VisitCount2: count2, VisitCount3: count3, TotalDownloadCount: downcount,
		DownloadCount1: downcount1, DownloadCount2: downcount2, CustomerCount: customercount}
	response.SuccessJSON(c, data)
}

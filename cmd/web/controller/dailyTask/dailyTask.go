// The build tag makes sure the stub is not built in the final build.

package dailyTask

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/dailyTask/save
func SaveDailyTaskControl(c *gin.Context) {
	var jsonp struct {
		dos.DailyTask
		MultiList []dos.DailyTaskMultiField `json:"multilist"`
	}
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	// 包含游戏验证
	sErrGameCodes, err2 := modules.CheckDailyTaskGameCodes(
		jsonp.VenueCode,
		jsonp.GameType,
		jsonp.IncludeGameCodes,
	)
	if err2 != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err2.Error())
		return
	}

	if len(sErrGameCodes) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER,
			fmt.Sprintf("存在错误的包含游戏参数: %s", sErrGameCodes))
		return
	}

	// 屏蔽游戏验证
	sErrGameCodes, err2 = modules.CheckDailyTaskGameCodes(
		jsonp.VenueCode,
		jsonp.GameType,
		jsonp.ExcludeGameCodes,
	)
	if err2 != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err2.Error())
		return
	}

	if len(sErrGameCodes) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER,
			fmt.Sprintf("存在错误的屏蔽游戏参数: %s", sErrGameCodes))
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateTime = automaticType.Time(time.Now())
		jsonp.UpdateTime = jsonp.CreateTime
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}

	if len(jsonp.MultiList) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "额度列表不能为空")
		return
	}

	// 循环保存多条任务数据
	tasks := []dos.DailyTask{}
	groupId := tool.SnowflakeIdByKey("dailytask-groupid")
	for _, v := range jsonp.MultiList {
		jsonp.DailyTask.Sort = v.Sort
		jsonp.DailyTask.Amount = v.Amount
		jsonp.DailyTask.BonusAmount = v.BonusAmount
		jsonp.DailyTask.GroupId = groupId
		tasks = append(tasks, jsonp.DailyTask)
	}

	err = modules.SaveDailyTaskMulit(tasks)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	modules.AddSyncDailyTaskConfig(modules.SyncDailyTaskConfigIndexMark, jsonp.MerchantCode)

	response.SuccessJSON(c, true)
}

// api: api/dailyTask/findPage
func FindPageDailyTaskControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.DailyTask
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Name = c.DefaultQuery("name", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "0"))
	jsonp.GroupId = c.DefaultQuery("groupid", "")

	jsonp.PageTimeQuery.StartAt = c.DefaultQuery("startAt", "")
	jsonp.PageTimeQuery.EndAt = c.DefaultQuery("endAt", "")

	//global.G_LOG.Infof("dailytask--------------------------0:%v, %v", c.DefaultQuery("startAt", ""), c.DefaultQuery("endAt", ""))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//global.G_LOG.Infof("dailytask--------------------------1-1:%v", jsonp)

	//global.G_LOG.Infof("dailytask--------------------------1-2:%v, %v, %v, %v", jsonp.PageTimeQuery, jsonp.DailyTask, c.DefaultQuery("groupid", "0"), jsonp.GroupId)
	data, total := modules.FindPageDailyTask(jsonp.PageTimeQuery, &jsonp.DailyTask, c)
	//global.G_LOG.Infof("dailytask--------------------------2:%v, %v", data, total)

	list := []*dos.DailyTaskResp{}
	for _, v := range data {
		dailyTask := dos.DailyTaskResp{}
		tool.JsonMapper(v, &dailyTask)
		dailyTask.MerchantName = v.Merchant.MerchantName
		list = append(list, &dailyTask)
	}
	//global.G_LOG.Infof("dailytask--------------------------3:%v", list)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
}

// api: api/dailyTask/findByKey
func FindByKeyDailyTaskControl(c *gin.Context) {
	jsonp := dos.DailyTask{}
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Name = c.DefaultQuery("name", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "0"))
	jsonp.StartAt = tool.CovertAutomaticTypeTimeFromDateString(c.DefaultQuery("startAt", ""))
	jsonp.EndAt = tool.CovertAutomaticTypeTimeFromDateString(c.DefaultQuery("endAt", ""))

	data := modules.FindByKeyDailyTask(&jsonp, c)

	list := []*dos.DailyTaskResp{}
	for _, v := range data {
		dailyTask := dos.DailyTaskResp{}
		tool.JsonMapper(v, &dailyTask)
		dailyTask.MerchantName = v.Merchant.MerchantName
		list = append(list, &dailyTask)
	}

	response.SuccessJSON(c, list)
}

// api: api/dailyTask/update
func UpdateDailyTaskControl(c *gin.Context) {
	var jsonp struct {
		dos.DailyTask
		MultiList []dos.DailyTaskMultiField `json:"multilist"`
	}
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
	//global.G_LOG.Infof("dailytask update--------------------------9-1:%v", 1)
	merchant := modules.FindByKeyDailyTaskFirst(&jsonp.DailyTask)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	// 包含游戏验证
	sErrGameCodes, err2 := modules.CheckDailyTaskGameCodes(
		jsonp.VenueCode,
		jsonp.GameType,
		jsonp.IncludeGameCodes,
	)
	if err2 != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err2.Error())
		return
	}

	if len(sErrGameCodes) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER,
			fmt.Sprintf("存在错误的包含游戏参数: %s", sErrGameCodes))
		return
	}

	// 屏蔽游戏验证
	sErrGameCodes, err2 = modules.CheckDailyTaskGameCodes(
		jsonp.VenueCode,
		jsonp.GameType,
		jsonp.ExcludeGameCodes,
	)
	if err2 != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err2.Error())
		return
	}

	if len(sErrGameCodes) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER,
			fmt.Sprintf("存在错误的屏蔽游戏参数: %s", sErrGameCodes))
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateTime = automaticType.Time(time.Now())
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	// 第一条数据为保存其他为新增
	ret := modules.UpdateDailyTask(&jsonp.DailyTask)
	if ret {
		modules.AddSyncDailyTaskConfig(modules.SyncDailyTaskConfigMark, jsonp.Id)
	}
	//global.G_LOG.Infof("dailytask update--------------------------9-2:%v", 1)

	for _, v := range jsonp.MultiList {
		jsonp.DailyTask.Id = v.Id
		jsonp.DailyTask.Sort = v.Sort
		jsonp.DailyTask.Amount = v.Amount
		jsonp.DailyTask.BonusAmount = v.BonusAmount
		ret1 := modules.UpdateDailyTask(&jsonp.DailyTask)
		//global.G_LOG.Infof("dailytask update--------------------------9-3:%v", ret1)
		if ret1 {
			//global.G_LOG.Infof("dailytask update--------------------------9-4:%v", 1)
			modules.AddSyncDailyTaskConfig(modules.SyncDailyTaskConfigMark, v.Id)
		}
	}
	//global.G_LOG.Infof("dailytask update--------------------------9-3:%v", 1)

	response.SuccessJSON(c, ret)
}

// api: api/dailyTask/delete
func DeleteDailyTaskControl(c *gin.Context) {
	var jsonp dos.DailyTask
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

	merchant := modules.FindByKeyDailyTaskFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteDailyTask(&jsonp)
	response.SuccessJSON(c, data)
}

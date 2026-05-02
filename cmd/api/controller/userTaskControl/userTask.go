package userTaskControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/srv"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 用户任务列表
func UserTaskList(c *gin.Context) {
	var jsonp struct {
		Type     int `json:"type" form:"type"`
		Page     int `json:"current" form:"current"`
		PageSize int `json:"pageSize" form:"pageSize"`
	}

	err := c.ShouldBindQuery(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfo, err := srv.GetUserInfo(c)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户码为空")
		return
	}

	// 获取用户任务列表
	list, total := modules.GetUserTaskList(merchantCode, userInfo.UserId, jsonp.Type,
		jsonp.Page, jsonp.PageSize)

	response.SuccessPageJSON(c, jsonp.Page, jsonp.PageSize, total, list)
}

// 用户任务领取
func UserTaskReward(c *gin.Context) {
	var jsonp struct {
		Id string `json:"id"`
	}

	c.ShouldBindJSON(&jsonp)

	userInfo, err := srv.GetUserInfo(c)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户码为空")
		return
	}

	if modules.CheckVenueEntryRecordVal(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "场馆金额未转出暂时无法领取")
		return
	}

	userTaskMp := modules.GetUserTaskDataToMap(userInfo.UserId)
	if v, ok := userTaskMp[jsonp.Id]; ok {
		// 判断状态是否可领取
		switch v.Status {
		case enmus.UserTaskStats_None:
			response.FailErrJSON(c, response.ERROR_PARAMETER, "任务奖励不可领")
			return
		case enmus.UserTaskStats_Rewarded:
			response.FailErrJSON(c, response.ERROR_PARAMETER, "任务奖励已领取")
			return
		}

		// 获取奖励
		taskConfig := modules.GetDailyTaskConfig(jsonp.Id)
		if taskConfig == nil {
			response.FailErrJSON(c, response.ERROR_SERVER, "任务配置错误")
			return
		}

		err = global.G_DB.Transaction(func(tx *gorm.DB) error {
			err1 := userTransfer.UserAmountChange(tx, taskConfig.BonusAmount, userTransfer.TranDiscount, global.CONFIG.General.DefaultCurrency,
				taskConfig.Name, userInfo.UserId, "", "",
				modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, "daily_task"))
			if err1 != nil {
				global.G_LOG.Errorf("[UserTaskReward][Reward] userTransfer.UserAmountChange  Amount:%0.2f  userId:%s  err :%s",
					taskConfig.BonusAmount, userInfo.UserId, err1.Error())
				return err1
			}
			return err1
		})

		if err != nil {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}

		// 更改用户数据状态
		v.Status = enmus.UserTaskStats_Rewarded
		userTaskMp[jsonp.Id] = v
		modules.SetUserTaskDataByMap(userInfo.UserId, userTaskMp)

		response.SuccessMsgJSON(c, nil, "领取成功")
		return
	}

	response.FailErrJSON(c, response.ERROR_PARAMETER, "任务奖励领取失败")
}

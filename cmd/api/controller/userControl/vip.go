package userControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LevelList(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)

	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户不存在")
		return
	}
	userInfo = modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: userInfo.UserId,
	})
	data := modules.FindByKeyFcVip(&dos.FcVip{})
	newData := make([]*vo.FcVipListVO, 0, len(data))
	tool.JsonMapper(&data, &newData)

	weeklyGiftApply := true
	monthlyApply := true
	//周礼金判断是否可领取
	y, m, w := tool.GetDateWeek(time.Now().Format(time.DateOnly))
	week := fmt.Sprintf("%d-%d-%d", y, m, w)
	weekGift := modules.FindByKeyFcVipWeekGiftFirst(&dos.FcVipWeekGift{Week: week, UserId: userInfo.UserId})
	if len(weekGift.Id) > 0 {
		weeklyGiftApply = false
	}

	//月礼金判断是否可取
	month := time.Now().Format("2006-01")
	monthGift := modules.FindByKeyFcVipMonthGiftFirst(&dos.FcVipMonthGift{Month: month, UserId: userInfo.UserId})
	if len(monthGift.Id) > 0 {
		monthlyApply = false
	}

	for _, v := range newData {
		vipUpLog := modules.FindByKeyFcUserVipRecordFirst(&dos.FcUserVipRecord{Level: v.Level, UserId: userInfo.UserId})
		if len(vipUpLog.Id) > 0 && vipUpLog.IssueBonus == 0 && vipUpLog.Bonus > 0 {
			v.UpgradeGiftApply = true
		} else {
			v.UpgradeGiftApply = false
		}
		if v.WeeklyGift == 0 || v.Level != userInfo.Level {
			v.WeeklyGiftApply = false
		} else {
			v.WeeklyGiftApply = weeklyGiftApply
		}

		if v.MonthlyGift == 0 || v.Level != userInfo.Level {
			v.MonthlyApply = false
		} else {
			v.MonthlyApply = monthlyApply
		}
	}
	response.SuccessJSON(c, newData)
}

func LevelrebateList(c *gin.Context) {
	data := modules.FindByKeyFcVipRebate(&dos.FcVipRebate{})
	response.SuccessJSON(c, data)
}

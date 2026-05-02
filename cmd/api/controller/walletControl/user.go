package walletControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/service/venues"
	"context"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gin-gonic/gin"
	"time"
)

const UserWalletKey = "UserVenueWalletKey:%v:%v"

func GetUserWalletMoney(c *gin.Context) {
	var jsonp vo.UserWalletMoneyRequest
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	userName := userInfo.UserName
	userWallet := modules.FindByKeyFcUserWalletFirst(&dos.FcUserWallet{UserId: userInfo.UserId, Currency: jsonp.Currency})
	if len(userWallet.Id) == 0 {
		userWallet = &dos.FcUserWallet{
			UserId:        userInfo.UserId,
			UserName:      userName,
			Currency:      jsonp.Currency,
			TotalAmount:   0,
			AvaAmount:     0,
			FronzenAmount: 0,
			IsLock:        0,
			MerchantCode:  merchantCode,
		}
		modules.SaveFcUserWallet(userWallet)
	}
	var resp = new(vo.WalletMoneyResp)
	tool.Map(userWallet, resp)
	resp.TotalAmount = resp.AvaAmount

	var venueAmount float64
	if !jsonp.VenueBalance && userInfo.IsFree == false {
		value := global.G_REDIS.Get(context.Background(), fmt.Sprintf(UserWalletKey, userInfo.UserId, jsonp.Currency)).Val()
		if len(value) == 0 {
			venueAmount = venues.VenueBalancesAllAsync(userInfo, jsonp.Currency)
			global.G_REDIS.Set(context.Background(), fmt.Sprintf(UserWalletKey, userInfo.UserId, jsonp.Currency), tool.String(venueAmount), 6*time.Second)
		} else {
			venueAmount, _ = convertor.ToFloat(value)
		}
	}
	resp.TotalAmount += venueAmount
	response.SuccessJSON(c, resp)
}

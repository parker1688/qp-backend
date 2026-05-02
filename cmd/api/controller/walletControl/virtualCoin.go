package walletControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
)

func GetDepositVirtualCoin(c *gin.Context) {
	fvc := modules.FindByKeyFcVirtualCurrency(&dos.FcVirtualCurrency{
		Status: 1,
	})
	fx := modules.FindByKeyFcVirtualCurrencyFx(&dos.FcVirtualCurrencyFx{
		OptType: 1, //存款
	})
	data := make([]*vo.VirtualCoinResp, 0, len(fvc))
	for _, v := range fvc {
		f := &vo.VirtualCoinResp{
			CurrencyName:     v.CurrencyName,
			CurrencyNameImg:  v.CurrencyNameImg,
			CurrencyChain:    v.CurrencyChain,
			CurrencyProtocol: v.CurrencyProtocol,
		}
		for _, currencyFx := range fx {
			if v.CurrencyChain == currencyFx.CurrencyChain && v.CurrencyName == currencyFx.CurrencyName {
				f.FxAmount = currencyFx.FxAmount
				break
			}
		}
		data = append(data, f)
	}
	response.SuccessJSON(c, data)
}

func GetWithdrawVirtualCoin(c *gin.Context) {
	fvc := modules.FindByKeyFcVirtualCurrency(&dos.FcVirtualCurrency{
		Status: 1,
	})
	fx := modules.FindByKeyFcVirtualCurrencyFx(&dos.FcVirtualCurrencyFx{
		OptType: 2, //提款
	})
	data := make([]*vo.VirtualCoinResp, 0, len(fvc))
	for _, v := range fvc {
		f := &vo.VirtualCoinResp{
			CurrencyName:     v.CurrencyName,
			CurrencyNameImg:  v.CurrencyNameImg,
			CurrencyChain:    v.CurrencyChain,
			CurrencyProtocol: v.CurrencyProtocol,
		}
		for _, currencyFx := range fx {
			if v.CurrencyChain == currencyFx.CurrencyChain && v.CurrencyName == currencyFx.CurrencyName {
				f.FxAmount = currencyFx.FxAmount
				break
			}
		}
		data = append(data, f)
	}
	response.SuccessJSON(c, data)
}

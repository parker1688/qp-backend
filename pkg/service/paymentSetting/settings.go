package paymentSetting

import (
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var _cache = cache.New(2*time.Minute, 5*time.Second)

// GetPaymentSettings
//
//	@Description: 获取支付Key配置(数据库)
//	@return []*dos.FcPaymentSetting
func GetPaymentSettings(paymentCode, merchantCode string) []*dos.FcPaymentSetting {
	key := fmt.Sprintf("%s:%s", paymentCode, merchantCode)
	v, ok := _cache.Get(key)
	if ok {
		return v.([]*dos.FcPaymentSetting)
	}
	data := modules.FindByKeyFcPaymentSetting(&dos.FcPaymentSetting{
		BaseDos:      dos.BaseDos{},
		PaymentCode:  paymentCode,
		MerchantCode: merchantCode,
	})
	_cache.SetDefault(key, data)
	return data
}

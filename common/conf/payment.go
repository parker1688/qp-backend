package conf

// Payment
// @Description: 支付配置
type PaymentSetting struct {
	YinRunPay struct {
		AppID                     string `yaml:"AppId"`
		MerchantCode              int    `yaml:"MerchantCode"`
		Md5Key                    string `yaml:"Md5Key"`
		APIURL                    string `yaml:"ApiUrl"`
		NotifyURL                 string `yaml:"NotifyUrl"`
		BankNotifyURLOut          string `yaml:"BankNotifyURLOut"`
		VirtualNotifyURLOut       string `yaml:"VirtualNotifyURLOut"`
		AliPayVirtualNotifyURLOut string `yaml:"AliPayVirtualNotifyURLOut"`
		ReturnUrl                 string `yaml:"ReturnUrl"`
		Currency                  string `yaml:"Currency"`
		Subject                   string `yaml:"Subject"`
		Body                      string `yaml:"Body"`
	} `yaml:"YinRunPay"`
}

package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	url2 "net/url"
	"strconv"
	"time"
)

var (
	DBDJTransferProccesCode map[string]string = map[string]string{
		"processing": "处理中，请稍后重试",
	}

	DBDJCurrencyMap = map[string]string{
		"CNY":     "1",
		"USD":     "2",
		"HKD":     "3",
		"VND1000": "4",
		"PHP":     "11",
		"MYR":     "22",
		"INR":     "24",
		"THB":     "26",
		"BRL":     "29",
	}

	DBDJLangMap = map[string]string{
		"zh-cn": "cn",
		"zh-hk": "zh",
		"en":    "en",
		"vi":    "vi",
		"th":    "th",
		"ml":    "ml",
		"ni":    "ni",
	}
)

type VenueDBDJ struct {
	VenueConfig conf.DBDJ
}

func NewDBDJ(venueConfig *conf.DBDJ) IVenues {
	return &VenueDBDJ{
		VenueConfig: *venueConfig,
	}
}

func (v VenueDBDJ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	userType := 0              // 默认为正式用户
	if request.UserType == 2 { // 试玩
		userType = 1
	}

	currenType, _ := DBDJCurrencyMap[request.Currency]
	if currenType == "" {
		currenType = "1" // 默认为 cny
	}

	reqMap := map[string]string{}
	reqMap["username"] = request.UserName
	reqMap["password"] = request.Password
	reqMap["tester"] = strconv.Itoa(userType)
	reqMap["merchant"] = v.VenueConfig.Merchant
	reqMap["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	reqMap["currency_code"] = currenType
	reqMap["key"] = v.VenueConfig.Md5key

	signOrg := tool.MapSortKeyAZString(reqMap)
	sign := tool.MD5([]byte(signOrg))
	// 从参数中移除密钥
	delete(reqMap, "key")
	reqMap["sign"] = sign

	// 将参数转换为有效的 URL 查询字符串
	urlMap := url2.Values{}
	for k, v := range reqMap {
		urlMap.Add(k, v)
	}
	paramStr := urlMap.Encode()
	apiUrl := url + "/api/member/register" + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("DBDJ-CreateUser-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.DBDJCreateUserResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "true" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueDBDJ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url

	reqMap := map[string]string{}
	reqMap["username"] = request.UserName
	reqMap["merchant"] = v.VenueConfig.Merchant
	reqMap["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	reqMap["key"] = v.VenueConfig.Md5key

	signOrg := tool.MapSortKeyAZString(reqMap)
	sign := tool.MD5([]byte(signOrg))
	// 从参数中移除密钥
	delete(reqMap, "key")
	reqMap["sign"] = sign

	// 将参数转换为有效的 URL 查询字符串
	urlMap := url2.Values{}
	for k, v := range reqMap {
		urlMap.Add(k, v)
	}
	paramStr := urlMap.Encode()
	apiUrl := url + "/api/fund/getBalance" + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("DBDJ-GetUserBalance-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.DBDJGetUserBalanceResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "true" {
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		amount, err := strconv.ParseFloat(respData.Data, 64)
		if err != nil {
			venueResp.Code = GetUserBalance_FAIL_CODE
			venueResp.Msg = err.Error()
			return &venueResp
		}
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueDBDJ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueDBDJ) GetOrderNo() string {
	// 20-32位数字
	orderSn := tool.SnowflakeId() + tool.GetRandStr(2, 5, false)

	return orderSn
}

// 转入
func (v VenueDBDJ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	transferType := "1"

	currenType, _ := DBDJCurrencyMap[request.Currency]
	if currenType == "" {
		currenType = "1" // 默认为 cny
	}

	reqMap := map[string]string{}
	reqMap["username"] = request.UserName
	reqMap["merchant"] = v.VenueConfig.Merchant
	reqMap["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	reqMap["key"] = v.VenueConfig.Md5key
	reqMap["type"] = transferType
	reqMap["amount"] = strconv.FormatFloat(request.Amount, 'f', -1, 64)
	reqMap["merOrderId"] = request.OrderSn
	reqMap["currency_code"] = currenType

	signOrg := tool.MapSortKeyAZString(reqMap)
	sign := tool.MD5([]byte(signOrg))
	// 从参数中移除密钥
	delete(reqMap, "key")
	reqMap["sign"] = sign

	// 将参数转换为有效的 URL 查询字符串
	urlMap := url2.Values{}
	for k, v := range reqMap {
		urlMap.Add(k, v)
	}
	paramStr := urlMap.Encode()
	apiUrl := url + "/api/fund/transfer" + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("DBDJ-Deposit-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.DBDJTransferResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "true" {
		venueResp.Code = Deposit_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueDBDJ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	transferType := "2"
	currenType, _ := DBDJCurrencyMap[request.Currency]
	if currenType == "" {
		currenType = "1" // 默认为 cny
	}

	reqMap := map[string]string{}
	reqMap["username"] = request.UserName
	reqMap["merchant"] = v.VenueConfig.Merchant
	reqMap["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	reqMap["key"] = v.VenueConfig.Md5key
	reqMap["type"] = transferType
	reqMap["amount"] = strconv.FormatFloat(request.Amount, 'f', -1, 64)
	reqMap["merOrderId"] = request.OrderSn
	reqMap["currency_code"] = currenType

	signOrg := tool.MapSortKeyAZString(reqMap)
	sign := tool.MD5([]byte(signOrg))
	// 从参数中移除密钥
	delete(reqMap, "key")
	reqMap["sign"] = sign

	// 将参数转换为有效的 URL 查询字符串
	urlMap := url2.Values{}
	for k, v := range reqMap {
		urlMap.Add(k, v)
	}
	paramStr := urlMap.Encode()
	apiUrl := url + "/api/fund/transfer" + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("DBDJ-Withdraw-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.DBDJTransferResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "true" {
		venueResp.Code = Withdraw_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

func (v VenueDBDJ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	lang, _ := DBDJLangMap[request.Language]
	if lang == "" {
		lang = "cn" // 默认中文简体
	}

	reqMap := map[string]string{}
	reqMap["username"] = request.UserName
	reqMap["merchant"] = v.VenueConfig.Merchant
	reqMap["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	reqMap["key"] = v.VenueConfig.Md5key
	reqMap["password"] = request.Password
	reqMap["client_ip"] = "1270001"
	reqMap["lang"] = lang

	signOrg := tool.MapSortKeyAZString(reqMap)
	sign := tool.MD5([]byte(signOrg))
	// 从参数中移除密钥
	delete(reqMap, "key")
	reqMap["sign"] = sign

	// 将参数转换为有效的 URL 查询字符串
	urlMap := url2.Values{}
	for k, v := range reqMap {
		urlMap.Add(k, v)
	}
	paramStr := urlMap.Encode()
	apiUrl := url + "/api/v2/member/login" + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("DBDJ-LoginGame-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.DBDJLoginGameResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "true" {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.Data.H5
		if request.ClientType == "web" {
			venueResp.Data.GameUrl = respData.Data.PC
		}
		return &venueResp
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueDBDJ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url

	reqMap := map[string]string{}
	reqMap["merOrderId"] = request.OrderSn
	reqMap["merchant"] = v.VenueConfig.Merchant
	reqMap["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	reqMap["key"] = v.VenueConfig.Md5key

	signOrg := tool.MapSortKeyAZString(reqMap)
	sign := tool.MD5([]byte(signOrg))
	// 从参数中移除密钥
	delete(reqMap, "key")
	reqMap["sign"] = sign

	// 将参数转换为有效的 URL 查询字符串
	urlMap := url2.Values{}
	for k, v := range reqMap {
		urlMap.Add(k, v)
	}
	paramStr := urlMap.Encode()
	apiUrl := url + "/api/fund/transferQuery" + "?" + paramStr

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	respStr := resp.String()
	apiLog := fmt.Sprintf("DBDJ-TransferConfirm-param: %s apiUrl: %s response: %s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	var respData venuevo.DBDJTransferConfirmResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "true" {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = TransferConfirm_FAIL_CODE
	venueResp.Msg = respStr
	_, ok := DBDJTransferProccesCode[respData.Data]
	if ok {
		venueResp.Code = TransferConfirm_Processing_CODE
	}
	return &venueResp
}

func (v VenueDBDJ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueDBDJ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueDBDJ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

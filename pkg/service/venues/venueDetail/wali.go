package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm/utils"
	url2 "net/url"
	"strconv"
	"time"
)

type VenueWALI struct {
	VenueConfig conf.WALI
}

func NewWALI(venueConfig *conf.WALI) IVenues {
	return &VenueWALI{
		VenueConfig: *venueConfig,
	}
}

func (v VenueWALI) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	acount := v.VenueConfig.Acount
	signKey := v.VenueConfig.SignKey
	aesKey := v.VenueConfig.AesKey
	url := v.VenueConfig.Url
	timeStamp := utils.ToString(time.Now().Unix())
	params0 := "text=hello"
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("WALI CreateUser Aes err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := params1 + timeStamp + signKey
	k := tool.MD5([]byte(param))
	apiUrl := url + "/ping?a=" + acount + "&t=" + timeStamp + "&p=" + params + "&k=" + k

	global.G_LOG.Infof("WALI-CreateUser info param:%v, url:%v", param, apiUrl)
	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("WALI-CreateUser %s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("WALI-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof("WALI-CreateUser-Login info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.WALICreateResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("WALI-CreateUser%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Code == 0 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueWALI) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	acount := v.VenueConfig.Acount
	signKey := v.VenueConfig.SignKey
	aesKey := v.VenueConfig.AesKey
	url := v.VenueConfig.Url
	timeStamp := utils.ToString(time.Now().Unix())
	params0 := "uid=" + request.UserName
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("WALI Balance Aes err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := params1 + timeStamp + signKey
	k := tool.MD5([]byte(param))
	apiUrl := url + "/getBalance?a=" + acount + "&t=" + timeStamp + "&p=" + params + "&k=" + k

	global.G_LOG.Infof("WALI-Balance info param:%v, url:%v", param, apiUrl)
	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("WALI-Balance-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.WALIBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Msg)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	// PGDZ 金额单位分
	if respData.Code == 0 {
		amount, err1 := strconv.ParseFloat(respData.Data.Balance, 64)
		if err1 != nil {
			global.G_LOG.Errorf("%s respdata err: %v", apiLog, err1.Error())
			venueResp.Code = GetUserBalance_FAIL_CODE
			venueResp.Msg = err1.Error()
			return &venueResp
		}
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueWALI) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueWALI) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueWALI) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	acount := v.VenueConfig.Acount
	signKey := v.VenueConfig.SignKey
	aesKey := v.VenueConfig.AesKey
	url := v.VenueConfig.Url
	agentName := v.VenueConfig.AgentName
	timeStamp := utils.ToString(time.Now().Unix())
	orderId := agentName + "_" + utils.ToString(time.Now().UnixMilli()) + "_" + request.UserName
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	params0 := "orderId=" + orderId + "&uid=" + request.UserName + "&ccy=CNY&credit=" + amount
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("WALI Deposit Aes err: %v", err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := params1 + timeStamp + signKey
	k := tool.MD5([]byte(param))
	apiUrl := url + "/transferV3?a=" + acount + "&t=" + timeStamp + "&p=" + params + "&k=" + k

	global.G_LOG.Infof("WALI-Balance info param:%v, url:%v", param, apiUrl)
	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("%s PGDZ-Deposit-httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("WALI-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.WALITransferResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		amount1, err2 := strconv.ParseFloat(respData.Data.Balance, 64)
		if err2 != nil {
			global.G_LOG.Errorf("%s float err: %v", apiLog, err2.Error())
			venueResp.Code = Deposit_Processing_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Data.Amount = amount1
		return &venueResp
	}

	global.G_LOG.Infof(apiLog)
	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueWALI) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	acount := v.VenueConfig.Acount
	signKey := v.VenueConfig.SignKey
	aesKey := v.VenueConfig.AesKey
	url := v.VenueConfig.Url
	agentName := v.VenueConfig.AgentName
	timeStamp := utils.ToString(time.Now().Unix())
	orderId := agentName + "_" + utils.ToString(time.Now().UnixMilli()) + "_" + request.UserName
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(-request.Amount, 2))
	params0 := "orderId=" + orderId + "&uid=" + request.UserName + "&ccy=CNY&credit=" + amount
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("WALI Withdraw Aes err: %v", err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := params1 + timeStamp + signKey
	k := tool.MD5([]byte(param))
	apiUrl := url + "/transferV3?a=" + acount + "&t=" + timeStamp + "&p=" + params + "&k=" + k

	global.G_LOG.Infof("WALI-Withdraw info param:%v, url:%v", param, apiUrl)
	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("%s WALI-Withdraw-httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("WALI-withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.WALITransferResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		amount1, err2 := strconv.ParseFloat(respData.Data.Balance, 64)
		if err2 != nil {
			global.G_LOG.Errorf("%s float err: %v", apiLog, err2.Error())
			venueResp.Code = Withdraw_Processing_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Data.Amount = amount1
		return &venueResp
	}

	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenueWALI) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	acount := v.VenueConfig.Acount
	signKey := v.VenueConfig.SignKey
	aesKey := v.VenueConfig.AesKey
	url := v.VenueConfig.Url
	timeStamp := utils.ToString(time.Now().Unix())
	params0 := "uid=" + request.UserName + "&ip=" + request.IP + "&ccy=CNY&game=" + request.GameCode
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("WALI Login Aes err: %v", err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := params1 + timeStamp + signKey
	k := tool.MD5([]byte(param))
	apiUrl := url + "/enterGame?a=" + acount + "&t=" + timeStamp + "&p=" + params + "&k=" + k

	global.G_LOG.Infof("WALI-EnterGame info param:%v, url:%v", params0, apiUrl)
	client := resty.New()
	resp, err2 := client.R().Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("WALI-login error :%v", err2.Error())
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("WALI-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.WALIEnterGameResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code != 0 {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = respData.Data.GameUrl
	return &venueResp
}
func (v VenueWALI) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenueWALI) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueWALI) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueWALI) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

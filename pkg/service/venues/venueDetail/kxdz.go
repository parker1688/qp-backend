package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	url2 "net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm/utils"
)

type VenueKXDZ struct {
	VenueConfig conf.KXDZ
}

func NewKXDZ(venueConfig *conf.KXDZ) IVenues {
	return &VenueKXDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenueKXDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}

	venueResp.Code = CreateUser_SUCCESS_CODE
	return &venueResp
}

func (v VenueKXDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	agent := v.VenueConfig.Agent
	md5Key := v.VenueConfig.Md5Key
	aesKey := v.VenueConfig.DesKey
	url := v.VenueConfig.Url
	timeStamp := utils.ToString(time.Now().Unix())
	params0 := "s=1&account=" + request.UserName + "&currency=CNY"
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("KXDZ Balance Aes err: %v", err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := agent + timeStamp + md5Key
	k := tool.MD5([]byte(param))
	apiUrl := url + "?agent=" + agent + "&timestamp=" + timeStamp + "&param=" + params + "&key=" + k

	global.G_LOG.Infof("KXDZ-balance info param:%v, url:%v", params0, apiUrl)

	client := resty.New()
	resp, err2 := client.R().Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("KXDZ-balance error :%v", err2.Error())
	}
	respStr := resp.String()

	apiLog := fmt.Sprintf("KXDZ-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	respBytes := []byte(respStr)
	global.G_LOG.Infof(apiLog)

	var respData venuevo.KXDZBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.D.Code)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	// PGDZ 金额单位分
	if respData.D.Code == 0 {
		amount := float64(respData.D.Money)
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	global.G_LOG.Infof(apiLog)

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueKXDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueKXDZ) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueKXDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	agent := v.VenueConfig.Agent
	md5Key := v.VenueConfig.Md5Key
	aesKey := v.VenueConfig.DesKey
	url := v.VenueConfig.Url
	now := time.Now()
	timeform := now.Format("20060102150405000")
	orderId := agent + timeform + request.UserName
	timeStamp := utils.ToString(time.Now().Unix())
	amount := tool.String(request.Amount)
	params0 := "s=2&account=" + request.UserName + "&money=" + amount + "&orderid=" + orderId + "&currency=CNY"
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("KXDZ deposit Aes err: %v", err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := agent + timeStamp + md5Key
	k := tool.MD5([]byte(param))
	apiUrl := url + "?agent=" + agent + "&timestamp=" + timeStamp + "&param=" + params + "&key=" + k

	global.G_LOG.Infof("KXDZ- deposit param:%v, url:%v", params0, apiUrl)

	client := resty.New()
	resp, err2 := client.R().Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("KXDZ-deposit error :%v", err2.Error())
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("KXDZ-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.KXDZDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.D.Money
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueKXDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	agent := v.VenueConfig.Agent
	md5Key := v.VenueConfig.Md5Key
	aesKey := v.VenueConfig.DesKey
	now := time.Now()
	timeform := now.Format("20060102150405000")
	orderId := agent + timeform + request.UserName
	timeStamp := utils.ToString(time.Now().Unix())

	params0 := "s=3&account=" + request.UserName + "&money=" + amount + "&orderid=" + orderId + "&currency=CNY"
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("KXDZ wihtdraw Aes err: %v", err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := agent + timeStamp + md5Key
	k := tool.MD5([]byte(param))
	apiUrl := url + "?agent=" + agent + "&timestamp=" + timeStamp + "&param=" + params + "&key=" + k

	global.G_LOG.Infof("KXDZ- info param:%v, url:%v", params0, apiUrl)

	client := resty.New()
	resp, err2 := client.R().Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("KXDZ-withdraw error :%v", err2.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}
	respStr := resp.String()

	apiLog := fmt.Sprintf("KXDZ-Withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)
	respBytes := []byte(respStr)

	var respData venuevo.PGDZWithdrawResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.AfterBalance
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenueKXDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	agent := v.VenueConfig.Agent
	md5Key := v.VenueConfig.Md5Key
	aesKey := v.VenueConfig.DesKey
	url := v.VenueConfig.Url
	now := time.Now()
	timeform := now.Format("20060102150405000")
	orderId := agent + timeform + request.UserName
	timeStamp := utils.ToString(time.Now().Unix())
	params0 := "s=0&account=" + request.UserName + "&money=0&orderid=" + orderId + "&ip=127.0.0.1&lineCode=test11&KindID=" + request.GameCode + "&tiny=0&currency=CNY"
	params1, err := tool.AesEcbPk7EncryptBase64(params0, aesKey)
	params := url2.QueryEscape(params1)
	if err != nil {
		global.G_LOG.Errorf("KXDZ Login Aes err: %v", err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	param := agent + timeStamp + md5Key
	k := tool.MD5([]byte(param))
	apiUrl := url + "?agent=" + agent + "&timestamp=" + timeStamp + "&param=" + params + "&key=" + k

	global.G_LOG.Infof("KXDZ-EnterGame info param:%v, url:%v", params0, apiUrl)
	client := resty.New()
	resp, err2 := client.R().Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("KXDZ-login error :%v", err2.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}
	respStr := resp.String()

	apiLog := fmt.Sprintf("KXDZ-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.KXDZLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.D.Code != 0 {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = respData.D.Url
	return &venueResp
}
func (v VenueKXDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenueKXDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueKXDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueKXDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

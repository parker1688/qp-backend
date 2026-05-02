package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm/utils"
)

type VenueBBIN struct {
	VenueConfig conf.BBIN
}

func NewBBIN(venueConfig *conf.BBIN) IVenues {
	return &VenueBBIN{
		VenueConfig: *venueConfig,
	}
}

func (v VenueBBIN) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	LoginId := v.VenueConfig.LoginId
	appSecret := v.VenueConfig.Md5Key
	url := v.VenueConfig.Url
	apiUrl := url
	now := utils.ToString(time.Now().Unix())
	param := "loginId=" + LoginId + "&method=getUser&timestamp=" + now + "&userName=" + request.UserName + "&key=" + appSecret
	sign := strings.ToUpper(tool.MD5([]byte(param)))
	reqMap := map[string]string{
		"loginId":   LoginId,
		"method":    "getUser",
		"timestamp": now,
		"userName":  request.UserName,
		"sign":      sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	global.G_LOG.Infof("BBIN-CreateUser info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BBIN-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof("BBIN-CreateUser info :%v", apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BBINUserResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Result == "success" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}
	global.G_LOG.Infof("BBIN-CreateUser info :%v, error:%v", apiLog, respData.Msg)
	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueBBIN) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	LoginId := v.VenueConfig.LoginId
	appSecret := v.VenueConfig.Md5Key
	url := v.VenueConfig.Url
	apiUrl := url
	now := utils.ToString(time.Now().Unix())
	param := "loginId=" + LoginId + "&method=getBalance&timestamp=" + now + "&userName=" + request.UserName + "&walletType=1&key=" + appSecret
	sign := strings.ToUpper(tool.MD5([]byte(param)))
	reqMap := map[string]string{
		"loginId":    LoginId,
		"method":     "getBalance",
		"timestamp":  now,
		"userName":   request.UserName,
		"walletType": "1",
		"sign":       sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	global.G_LOG.Infof("BBIN-Balance info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BBIN-GetUserBalance-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.BBINBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Result == "success" {
		amount := respData.Balance
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	global.G_LOG.Infof("BBIN-CreateUser info :%v, error:%v", apiLog, respData.Msg)
	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueBBIN) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueBBIN) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueBBIN) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	LoginId := v.VenueConfig.LoginId
	appSecret := v.VenueConfig.Md5Key
	url := v.VenueConfig.Url
	apiUrl := url
	//transactionId := tool.RandomInt(100000000000, 9999999999999)
	transactionId := tool.SnowflakeIdByKey("bbin-transactionId")
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	now := utils.ToString(time.Now().Unix())
	param := "amount=" + amount + "&loginId=" + LoginId + "&method=getTransaction&timestamp=" + now + "&transactionId=" + transactionId + "&transactionType=deposit&userName=" + request.UserName + "&walletType=1&key=" + appSecret
	sign := strings.ToUpper(tool.MD5([]byte(param)))
	reqMap := map[string]string{
		"loginId":         LoginId,
		"method":          "getTransaction",
		"timestamp":       now,
		"userName":        request.UserName,
		"walletType":      "1",
		"transactionId":   transactionId,
		"transactionType": "deposit",
		"amount":          amount,
		"sign":            sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	global.G_LOG.Infof("BBIN-deposit info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BBIN-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BBINTransactionResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Result == "success" {
		amount2, err2 := strconv.ParseFloat(respData.Amount, 64)
		if err2 != nil {
			venueResp.Code = Deposit_FAIL_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = amount2
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}

// 转出
func (v VenueBBIN) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	LoginId := v.VenueConfig.LoginId
	appSecret := v.VenueConfig.Md5Key
	url := v.VenueConfig.Url
	apiUrl := url
	//transactionId := tool.RandomInt(100000000000, 9999999999999)
	transactionId := tool.SnowflakeIdByKey("bbin-transactionId")
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	now := utils.ToString(time.Now().Unix())
	param := "amount=" + amount + "&loginId=" + LoginId + "&method=getTransaction&timestamp=" + now + "&transactionId=" + transactionId + "&transactionType=withdrawal&userName=" + request.UserName + "&walletType=1&key=" + appSecret
	sign := strings.ToUpper(tool.MD5([]byte(param)))
	reqMap := map[string]string{
		"loginId":         LoginId,
		"method":          "getTransaction",
		"timestamp":       now,
		"userName":        request.UserName,
		"walletType":      "1",
		"transactionId":   transactionId,
		"transactionType": "withdrawal",
		"amount":          amount,
		"sign":            sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	global.G_LOG.Infof("BBIN-withdraw info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BBIN-withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BBINTransactionResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Result == "success" {
		amount2, err2 := strconv.ParseFloat(respData.Amount, 64)
		if err2 != nil {
			venueResp.Code = Deposit_FAIL_CODE
			venueResp.Msg = err2.Error()
			return &venueResp
		}
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = amount2
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenueBBIN) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	gameId := request.GameCode
	gameType := request.GType
	LoginId := v.VenueConfig.LoginId
	appSecret := v.VenueConfig.Md5Key
	url := v.VenueConfig.Url
	apiUrl := url
	platForm := ""
	switch strings.ToUpper(request.ClientType) {
	case "PC":
		platForm = "PC"
	default:
		platForm = "Mobile"
	}
	now := utils.ToString(time.Now().Unix())
	param := "device=" + platForm + "&gameType=" + gameType + "&loginId=" + LoginId + "&method=getGameUrl&timestamp=" + now + "&userName=" + request.UserName + "&key=" + appSecret
	if gameType == "62" || gameType == "66" || gameType == "63" {
		param = "device=" + platForm + "&gameId=" + gameId + "&gameType=" + gameType + "&loginId=" + LoginId + "&method=getGameUrl&timestamp=" + now + "&userName=" + request.UserName + "&key=" + appSecret
	}
	sign := strings.ToUpper(tool.MD5([]byte(param)))
	reqMap := map[string]string{
		"loginId":   LoginId,
		"method":    "getGameUrl",
		"timestamp": now,
		"device":    platForm,
		"userName":  request.UserName,
		"gameId":    gameId,
		"gameType":  gameType,
		"sign":      sign,
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	global.G_LOG.Infof("BBIN-Login info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("BBIN-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.BBINLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Result == "success" {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.GameUrl
		return &venueResp
	}
	global.G_LOG.Infof("BBIn-Logingame-----:%v, url:%v", apiLog, respData.GameUrl)
	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}
func (v VenueBBIN) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp
}
func (v VenueBBIN) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueBBIN) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueBBIN) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

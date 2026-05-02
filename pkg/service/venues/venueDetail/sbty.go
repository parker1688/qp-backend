package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type VenueSBTY struct {
	VenueConfig conf.SBTY
}

func NewSBTY(venueConfig *conf.SBTY) IVenues {
	return &VenueSBTY{
		VenueConfig: *venueConfig,
	}
}

func (v VenueSBTY) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	vendorId := v.VenueConfig.VendorId
	operatorId := v.VenueConfig.OperatorId
	currency := v.VenueConfig.Currency
	userName := operatorId + "_" + request.UserName
	url := v.VenueConfig.Url
	apiUrl := url + "/CreateMember"
	reqMap := map[string]string{
		"vendor_id":        vendorId,
		"vendor_member_id": userName,
		"operatorId":       operatorId,
		"username":         userName,
		"oddstype":         "2",      //2中国盘
		"currency":         currency, //20-测试    13-RMB
		"maxtransfer":      "9999999999",
		"mintransfer":      "0",
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	//global.G_LOG.Infof("PGDZ-Login info param:%v, reqMap:%v, head:%v", param, reqMap, headerParams)
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
	apiLog := fmt.Sprintf("SBTY-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.SBTYUserResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.ErrorCode == 0 {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueSBTY) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	vendorId := v.VenueConfig.VendorId
	operatorId := v.VenueConfig.OperatorId
	userName := operatorId + "_" + request.UserName
	url := v.VenueConfig.Url
	apiUrl := url + "/CheckUserBalance"
	reqMap := map[string]string{
		"vendor_id":         vendorId,
		"vendor_member_ids": userName,
		"wallet_id":         "1",
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
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
	apiLog := fmt.Sprintf("SBTY-GetUserBalance-apiUrl: %s response: %s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.SBTYBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v, err2:%v", apiLog, err.Error(), respData.Message)
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.ErrorCode == 0 {
		amount := float64(respData.Data[0].Balance)
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueSBTY) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueSBTY) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueSBTY) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	vendorId := v.VenueConfig.VendorId
	operatorId := v.VenueConfig.OperatorId
	currency := v.VenueConfig.Currency
	userName := operatorId + "_" + request.UserName
	url := v.VenueConfig.Url
	transId := operatorId + "_" + tool.SnowflakeIdByKey("sbty-transfer")
	apiUrl := url + "/FundTransfer"
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	reqMap := map[string]string{
		"vendor_id":        vendorId,
		"vendor_member_id": userName,
		"vendor_trans_id":  transId,
		"amount":           amount,
		"currency":         currency,
		"direction":        "1",
		"wallet_id":        "1",
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
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
	apiLog := fmt.Sprintf("SBTY-Deposit-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.SBTYDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.ErrorCode == 0 {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.AfterAmount
		return &venueResp
	}
	if respData.ErrorCode == 1 && respData.Data.Status == 2 {
		time.Sleep(10 * time.Second)
		request2 := &VenueTransferConfirmRequest{
			OrderSn: transId,
		}
		res := v.TransferConfirm(request2)
		if res.Code == TransferConfirm_SUCCESS_CODE {
			venueResp.Code = Deposit_SUCCESS_CODE
			venueResp.Data.Amount = request.Amount
			return &venueResp
		}
	}
	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

// 转出
func (v VenueSBTY) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	vendorId := v.VenueConfig.VendorId
	operatorId := v.VenueConfig.OperatorId
	currency := v.VenueConfig.Currency
	userName := operatorId + "_" + request.UserName
	url := v.VenueConfig.Url
	transId := operatorId + "_" + tool.SnowflakeIdByKey("sbty-transfer")
	apiUrl := url + "/FundTransfer"
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	reqMap := map[string]string{
		"vendor_id":        vendorId,
		"vendor_member_id": userName,
		"vendor_trans_id":  transId,
		"amount":           amount,
		"currency":         currency,
		"direction":        "0",
		"wallet_id":        "1",
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
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
	apiLog := fmt.Sprintf("SBTY-Withdraw-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)
	respBytes := []byte(respStr)

	var respData venuevo.SBTYWithdrawResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.ErrorCode == 0 {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Data.AfterAmount
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr

	return &venueResp
}
func (v VenueSBTY) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	vendorId := v.VenueConfig.VendorId
	operatorId := v.VenueConfig.OperatorId
	userName := operatorId + "_" + request.UserName
	url := v.VenueConfig.Url
	apiUrl := url + "/GetSabaUrl"
	platForm := ""
	switch strings.ToUpper(request.ClientType) {
	case "PC":
		platForm = "1"
	default:
		platForm = "2"
	}
	reqMap := map[string]string{
		"vendor_id":        vendorId,
		"vendor_member_id": userName,
		"platform":         platForm,
		//"skin_mode":         "8",    //3桌机版，7银河版，8轻量版
		//"deep_link_content": "CN/2", // lang/Type(1欧洲2中国3印尼4马来5美国）
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("SBTY-LoginGame-apiUrl:%s , %v, response:%s", apiUrl, reqMap, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.SBTYLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.ErrorCode != 0 {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = respStr
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = LoginGame_SUCCESS_CODE
	venueResp.Data.GameUrl = respData.Data
	return &venueResp
}
func (v VenueSBTY) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}
	vendorId := v.VenueConfig.VendorId
	url := v.VenueConfig.Url
	transId := request.OrderSn
	apiUrl := url + "/CheckFundTransfer"
	reqMap := map[string]string{
		"vendor_id":       vendorId,
		"vendor_trans_id": transId,
		"wallet_id":       "1",
	}
	headerParams := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headerParams).
		SetFormData(reqMap).
		Post(apiUrl)
	if err != nil {
		venueResp.Code = TransferConfirm_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("SBTY-CheckTrans-apiUrl:%s response:%s", apiUrl, respStr)
	global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.SBTYTransferConfirmResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = TransferConfirm_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.ErrorCode == 1 && respData.Data.Status == 2 {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		return v.TransferConfirm(request)
	}
	if respData.ErrorCode == 3 && respData.Data.Status == 2 {
		global.G_LOG.Errorf("%s  login-game error", apiLog)
		time.Sleep(10 * time.Second)
		return v.TransferConfirm(request)
	}
	if respData.ErrorCode == 0 {
		venueResp.Code = TransferConfirm_SUCCESS_CODE
		venueResp.Data = respData.Data.Amount
		return &venueResp
	}
	global.G_LOG.Infof(apiLog)
	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}
func (v VenueSBTY) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueSBTY) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueSBTY) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type VenueJDB struct {
	VenueConfig conf.JDB
}

func NewJDB(venueConfig *conf.JDB) IVenues {
	return &VenueJDB{
		VenueConfig: *venueConfig,
	}
}

func (v VenueJDB) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	venueResp := VenueResponse{}
	url := v.VenueConfig.Url
	iv := v.VenueConfig.IV
	key := v.VenueConfig.Key
	dc := v.VenueConfig.DCName
	parent := v.VenueConfig.AgentId
	uid := request.UserName
	Name := request.UserName
	now := time.Now().UnixMilli()
	data := `{"action":12,"ts":` + tool.String(now) + `,"parent":` + parent + `,"uid":` + uid + `,"name":` + Name + `}`

	//global.G_LOG.Infof("JDB-createrUser-----1:%v", data)
	// Encrypt data
	encData, err := tool.AesEncrypt([]byte(data), []byte(key), []byte(iv))
	if err != nil {
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiUrl := url + "/apiRequest.do?" + "dc=" + dc + "&x=" + encData
	//global.G_LOG.Infof("JDB-createrUser----2:%v", apiUrl)
	headerParams := map[string]string{
		"Content-Type": "application/json",
	}
	client := resty.New()

	resp, err2 := client.R().SetHeaders(headerParams).Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err2.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("JDB-CreateUser-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	// 该场馆有时候报错不会返回 status_code 字段
	respDataMap := map[string]interface{}{}
	err = tool.JsonUnmarshal(respBytes, &respDataMap)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = CreateUser_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respDataMap["status"] == "0000" {
		venueResp.Code = CreateUser_SUCCESS_CODE
		return &venueResp
	}

	venueResp.Code = CreateUser_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueJDB) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	venueResp := VenueGetUserBalanceResponse{}
	url := v.VenueConfig.Url
	iv := v.VenueConfig.IV
	key := v.VenueConfig.Key
	dc := v.VenueConfig.DCName
	parent := v.VenueConfig.AgentId
	uid := request.UserName
	now := time.Now().UnixMilli()
	data := `{"action":15,"ts":` + tool.String(now) + `,"parent":` + parent + `,"uid":` + uid + `}`

	// Encrypt data
	encData, err := tool.AesEncrypt([]byte(data), []byte(key), []byte(iv))
	if err != nil {
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiUrl := url + "/apiRequest.do?" + "dc=" + dc + "&x=" + encData
	headerParams := map[string]string{
		"Content-Type": "application/json",
	}
	client := resty.New()
	resp, err2 := client.R().SetHeaders(headerParams).Get(apiUrl)

	respStr := resp.String()
	apiLog := fmt.Sprintf("JDB-GetBalance-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)
	if err2 != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiLog, err2.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}
	respBytes := []byte(respStr)

	var respData venuevo.JDBBalanceResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = GetUserBalance_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Status == "0000" {
		amount := respData.Data
		venueResp.Code = GetUserBalance_SUCCESS_CODE
		venueResp.Data.Amount = amount[0].Balance
		return &venueResp
	}

	venueResp.Code = GetUserBalance_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueJDB) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueJDB) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}

// 转入
func (v VenueJDB) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	venueResp := VenueDepositResponse{}
	url := v.VenueConfig.Url
	iv := v.VenueConfig.IV
	key := v.VenueConfig.Key
	dc := v.VenueConfig.DCName
	parent := v.VenueConfig.AgentId
	uid := request.UserName
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(request.Amount, 2))
	now := time.Now().UnixMilli()
	serialNo := tool.SnowflakeIdByKey("jdb-serialNo")
	data := `{"action":19,"ts":` + tool.String(now) + `,"parent":` + parent + `,"uid":` + uid + `,"serialNo":` + serialNo + `,"amount":` + amount + `}`

	//global.G_LOG.Infof("JDB deposit params:%v", data)
	// Encrypt data
	encData, err := tool.AesEncrypt([]byte(data), []byte(key), []byte(iv))
	if err != nil {
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiUrl := url + "/apiRequest.do?" + "dc=" + dc + "&x=" + encData
	headerParams := map[string]string{
		"Content-Type": "application/json",
	}
	client := resty.New()
	resp, err2 := client.R().SetHeaders(headerParams).Get(apiUrl)

	if err2 != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err2.Error())
		venueResp.Code = Deposit_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}

	respStr := resp.String()
	apiLog := fmt.Sprintf("JDB-Deposit-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.JDBDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Deposit_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "0000" {
		venueResp.Code = Deposit_SUCCESS_CODE
		venueResp.Data.Amount = respData.Amount
		venueResp.Data.TransactionId = serialNo
		return &venueResp
	}

	venueResp.Code = Deposit_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

// 转出
func (v VenueJDB) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	venueResp := VenueWithdrawResponse{}
	url := v.VenueConfig.Url
	iv := v.VenueConfig.IV
	key := v.VenueConfig.Key
	dc := v.VenueConfig.DCName
	parent := v.VenueConfig.AgentId
	uid := request.UserName
	amount := fmt.Sprintf("%.2f", tool.TruncateFloat(-request.Amount, 2))
	now := time.Now().UnixMilli()
	serialNo := tool.SnowflakeIdByKey("jdb-serialNo")
	data := `{"action":19,"ts":` + tool.String(now) + `,"parent":` + parent + `,"uid":` + uid + `,"serialNo":` + serialNo + `,"amount":` + amount + `}`
	//global.G_LOG.Infof("JDB-Withdraw ---------------param:%v", data)
	// Encrypt data
	encData, err := tool.AesEncrypt([]byte(data), []byte(key), []byte(iv))
	if err != nil {
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiUrl := url + "/apiRequest.do?" + "dc=" + dc + "&x=" + encData
	headerParams := map[string]string{
		"Content-Type": "application/json",
	}
	client := resty.New()
	resp, err2 := client.R().SetHeaders(headerParams).Get(apiUrl)

	if err2 != nil {
		global.G_LOG.Errorf("%s httpReq err: %v", apiUrl, err2.Error())
		venueResp.Code = Withdraw_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("JDB-Withdraw-apiUrl: %s response: %s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)

	var respData venuevo.JDBDepositResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = Withdraw_Processing_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Status == "0000" {
		venueResp.Code = Withdraw_SUCCESS_CODE
		venueResp.Data.Amount = respData.Amount
		venueResp.Data.TransactionId = serialNo
		return &venueResp
	}

	venueResp.Code = Withdraw_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func (v VenueJDB) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	venueResp := VenueLoginGameResponse{}
	url := v.VenueConfig.Url
	iv := v.VenueConfig.IV
	key := v.VenueConfig.Key
	dc := v.VenueConfig.DCName
	uid := request.UserName
	mType := request.GameCode
	gType := request.GType
	now := time.Now().UnixMilli()
	data := `{"action":11,"ts":` + tool.String(now) + `,"uid":` + uid + `,"gType":` + gType + `,"mType":` + mType + `}`
	//global.G_LOG.Infof("JDB-LoginGame-params:%v", data)

	encData, err := tool.AesEncrypt([]byte(data), []byte(key), []byte(iv))
	if err != nil {
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	apiUrl := url + "/apiRequest.do?" + "dc=" + dc + "&x=" + encData
	headerParams := map[string]string{
		"Content-Type": "application/json",
	}
	client := resty.New()
	resp, err2 := client.R().SetHeaders(headerParams).Get(apiUrl)
	if err2 != nil {
		global.G_LOG.Errorf("login error :%v", apiUrl)
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err2.Error()
		return &venueResp
	}

	respStr := resp.String()
	apiLog := fmt.Sprintf("JDB-LoginGame-apiUrl:%s response:%s", apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	respBytes := []byte(respStr)
	var respData venuevo.JDBLoginResp
	err = tool.JsonUnmarshal(respBytes, &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshalMap err: %v", apiLog, err.Error())
		venueResp.Code = LoginGame_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}

	if respData.Status == "0000" {
		venueResp.Code = LoginGame_SUCCESS_CODE
		venueResp.Data.GameUrl = respData.Url
		return &venueResp
	}

	venueResp.Code = LoginGame_FAIL_CODE
	venueResp.Data.GameUrl = apiUrl
	return &venueResp
}

func (v VenueJDB) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	venueResp := VenueResponse{}

	venueResp.Code = TransferConfirm_SUCCESS_CODE
	return &venueResp

}

func (v VenueJDB) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueJDB) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueJDB) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

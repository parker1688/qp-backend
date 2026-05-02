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
	"github.com/tidwall/gjson"
)

// VenueFBTY
// @Description: IM体育实现
type VenueFBTY struct {
	MerchantCode string
	MerchantId   string
	Url          string
	Secret       string
}

func NewVenueFBTY(venueConfig *conf.Fbty) IVenues {
	return &VenueFBTY{
		MerchantCode: venueConfig.MerchantCode,
		MerchantId:   venueConfig.MerchantId,
		Secret:       venueConfig.Secret,
		Url:          venueConfig.Url,
	}
}

func (v VenueFBTY) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	apiUrl := fmt.Sprintf("%s/fb/data/api/v2/new/user/create", v.Url)
	postData := map[string]string{
		"merchantUserId": request.UserName,
	}

	postJson, err := tool.MarshalToString(postData)
	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	timeNow := time.Now().Unix()
	headers := map[string]string{
		"Content-Type": "application/json",
		"sign":         v.MakeSign(timeNow, postJson),
		"timestamp":    fmt.Sprintf("%d", timeNow),
		"merchantId":   v.MerchantId,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).SetBody(postJson).Post(apiUrl)

	//global.G_LOG.Info("FBTY-CreateUser-postData", postJson)
	//global.G_LOG.Info("FBTY-CreateUser-resp", resp.String())
	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}
	var fbtyRegisterResp venuevo.FbtyRegisterResp
	err = tool.JsonUnmarshal(resp.Body(), &fbtyRegisterResp)
	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}
	if fbtyRegisterResp.Success && fbtyRegisterResp.Message == "SUCCESS" && fbtyRegisterResp.Code == 0 {
		return &VenueResponse{Code: CreateUser_SUCCESS_CODE}
	}

	return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: ""}
}

func (v VenueFBTY) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	apiUrl := fmt.Sprintf("%s/fb/data/api/v2/user/detail", v.Url)
	postData := map[string]string{
		"merchantUserId": request.UserName,
	}

	postJson, err := tool.MarshalToString(postData)
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}

	timeNow := time.Now().Unix()
	headers := map[string]string{
		"Content-Type": "application/json",
		"sign":         v.MakeSign(timeNow, postJson),
		"timestamp":    fmt.Sprintf("%d", timeNow),
		"merchantId":   v.MerchantId,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).SetBody(postJson).Post(apiUrl)
	//global.G_LOG.Info(apiUrl)
	//global.G_LOG.Info(postJson)
	//global.G_LOG.Info(resp.String())
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}
	var fbtyBalanceResp venuevo.FbtyBalanceResp
	err = tool.JsonUnmarshal(resp.Body(), &fbtyBalanceResp)
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}
	if fbtyBalanceResp.Success && fbtyBalanceResp.Message == "SUCCESS" && fbtyBalanceResp.Code == 0 {
		amount, err := strconv.ParseFloat(fbtyBalanceResp.Data.Balance, 64)
		if err != nil {
			return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
		}
		venueGetUserBalanceResponse := &VenueGetUserBalanceResponse{}
		venueGetUserBalanceResponse.Code = GetUserBalance_SUCCESS_CODE
		venueGetUserBalanceResponse.Data.Amount = amount
		return venueGetUserBalanceResponse
	}

	return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: ""}
}

func (v VenueFBTY) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueFBTY) GetOrderNo() string {
	//TODO implement me
	return tool.MakeTransferrder()
}

func (v VenueFBTY) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
	apiUrl := fmt.Sprintf("%s/fb/data/api/v2/new/transfer/in", v.Url)

	postData := map[string]interface{}{
		"merchantUserId": request.UserName,
		"businessId":     request.OrderSn,
		"amount":         request.Amount,
	}

	postJson, err := tool.MarshalToString(postData)
	if err != nil {
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: err.Error()}
	}

	timeNow := time.Now().Unix()
	headers := map[string]string{
		"Content-Type": "application/json",
		"sign":         v.MakeSign(timeNow, postJson),
		"timestamp":    fmt.Sprintf("%d", timeNow),
		"merchantId":   v.MerchantId,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).SetBody(postJson).Post(apiUrl)

	global.G_LOG.Infof("FBTY-Deposit: url:%s parmas: %s  response: %s", apiUrl, postJson, resp.String())
	if err != nil {
		global.G_LOG.Errorf("FBTY deposit error:%v", err.Error())
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: err.Error()}
	}
	var fbtyTransferResp venuevo.FbtyTransferResp
	err = tool.JsonUnmarshal(resp.Body(), &fbtyTransferResp)
	if err != nil {
		global.G_LOG.Errorf("FBTY deposit error:%v", err.Error())
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: err.Error()}
	}
	if fbtyTransferResp.Success && fbtyTransferResp.Message == "SUCCESS" && fbtyTransferResp.Code == 0 {
		return &VenueDepositResponse{Code: Deposit_SUCCESS_CODE}
	}
	global.G_LOG.Errorf("FBTY deposit error:%v", resp.String())
	return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}
}

func (v VenueFBTY) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	apiUrl := fmt.Sprintf("%s/fb/data/api/v2/new/transfer/out", v.Url)

	postData := map[string]interface{}{
		"merchantUserId": request.UserName,
		"businessId":     request.OrderSn,
		"amount":         request.Amount,
	}

	postJson, err := tool.MarshalToString(postData)
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: err.Error()}
	}

	timeNow := time.Now().Unix()
	headers := map[string]string{
		"Content-Type": "application/json",
		"sign":         v.MakeSign(timeNow, postJson),
		"timestamp":    fmt.Sprintf("%d", timeNow),
		"merchantId":   v.MerchantId,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).SetBody(postJson).Post(apiUrl)

	//global.G_LOG.Infof("FBTY-Withdraw: url:%s parmas: %s  response: %s", apiUrl, postJson, resp.String())
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
	}
	var fbtyTransferResp venuevo.FbtyTransferResp
	err = tool.JsonUnmarshal(resp.Body(), &fbtyTransferResp)
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
	}
	if fbtyTransferResp.Success && fbtyTransferResp.Message == "SUCCESS" && fbtyTransferResp.Code == 0 {
		return &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
	}

	return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
}

func (v VenueFBTY) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	apiUrl := v.Url + "/fb/data/api/v2/token/get"
	platForm := ""
	switch strings.ToUpper(request.ClientType) {
	case "PC":
		platForm = "pc"
	case "H5":
		platForm = "h5"
	default:
		platForm = "mobile"
	}

	postMap := map[string]string{
		"merchantUserId": request.UserName,
		"platForm":       platForm,
		"ip":             request.IP,
	}

	postJson, err := tool.MarshalToString(postMap)
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}

	timeNow := time.Now().Unix()
	headers := map[string]string{
		"Content-Type": "application/json",
		"sign":         v.MakeSign(timeNow, postJson),
		"timestamp":    fmt.Sprintf("%d", timeNow),
		"merchantId":   v.MerchantId,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).SetBody(postJson).Post(apiUrl)
	//global.G_LOG.Infof("FBTY-LoginGame-postData:%v  response:%s", postJson, resp.String())
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}
	var fbtyTokenResp venuevo.FbtyTokenResp
	err = tool.JsonUnmarshal(resp.Body(), &fbtyTokenResp)
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}

	domain := ""
	switch platForm {
	case "pc":
		domain = fbtyTokenResp.Data.ServerInfo.PcAddress
	case "h5":
		domain = fbtyTokenResp.Data.ServerInfo.H5Address
	default:
		domain = fbtyTokenResp.Data.ServerInfo.H5Address
	}

	GameUrl := fmt.Sprintf("%s/index.html#/?token=%s&nickname=%s&apiSrc=%s&pushSrc=%s&platformName=%s&icoUrl=%s&handicap=%s&language=%s",
		domain,
		fbtyTokenResp.Data.Token,
		request.UserName,
		fbtyTokenResp.Data.ServerInfo.APIServerAddress,
		fbtyTokenResp.Data.ServerInfo.PushServerAddress,
		"",
		"",
		"",
		"CMN",
	)

	venueLoginGameResponse := &VenueLoginGameResponse{Code: LoginGame_SUCCESS_CODE}
	//if platForm == "mobile" && fbtyTokenResp.Success && fbtyTokenResp.Data.Token != "" {

	venueLoginGameResponse.Data.GameUrl = GameUrl
	venueLoginGameResponse.Data.Token = fbtyTokenResp.Data.Token
	venueLoginGameResponse.Data.ApiDomain = fbtyTokenResp.Data.ServerInfo.APIServerAddress
	return venueLoginGameResponse
	//}

	//return venueLoginGameResponse
}

func (v VenueFBTY) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	apiURL := fmt.Sprintf("%v/fb/data/api/v2/transfer/detail", v.Url)

	timestamp := time.Now().Unix()
	postData := fmt.Sprintf(`{"businessId":"%s","merchantUserId":"%s","transferType":"%s"}`, request.UserName, request.UserName, request.TransferType)
	sign := v.MakeSign(timestamp, postData)

	headers := map[string]string{
		"Content-Type": "application/json",
		"sign":         sign,
		"timestamp":    fmt.Sprintf("%d", timestamp),
		"merchantId":   v.MerchantId,
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).SetBody(postData).Post(apiURL)
	//global.G_LOG.Infof("FBTY-TransferConfirm-postData:%v  response:%s", postData, resp.String())
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}

	status := gjson.Get(resp.String(), "data.status")
	if status.Exists() {
		switch status.Int() {
		case 1:
			return &VenueResponse{Code: TransferConfirm_SUCCESS_CODE}
		case 0:
			return &VenueResponse{Code: TransferConfirm_FAIL_CODE}
		}
	}
	return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: resp.String()}
}

func (v VenueFBTY) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueFBTY) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueFBTY) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

func (v VenueFBTY) MakeSign(timestamp int64, body string) string {

	stringThatNeedsToBeSigned := fmt.Sprintf("%s.%s.%d.%s", body, v.MerchantId, timestamp, v.Secret)
	// 然后对stringThatNeedsToBeSigned字符串进行md5运算生成sign
	sign := tool.MD5([]byte(stringThatNeedsToBeSigned))
	return sign
}

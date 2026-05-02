package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"net/url"
)

type VenueMGDZ struct {
	VenueConfig conf.MGDZ
}

func NewVenueMGDZ(venueConfig *conf.MGDZ) IVenues {
	return &VenueMGDZ{
		VenueConfig: *venueConfig,
	}
}

func (v VenueMGDZ) CreateUser(request *VenueCreateUserRequest) *VenueResponse {

	apiUrl := fmt.Sprintf("%s/agents/%s/players",
		v.VenueConfig.Url, v.VenueConfig.AgentCode)

	params := map[string]string{
		"playerId": request.UserName,
	}
	token := mgGetToken(v.VenueConfig)

	if token == "" {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: "jwt token is nil"}
	}
	client := resty.New()
	resp, err := client.R().SetAuthToken(token).
		SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"}).
		SetFormData(params).
		Post(apiUrl)

	//global.G_LOG.Infof("MGDZ-CreateUser-param:%s apiUrl:%s status:%s response:%s", tool.String(params), apiUrl, resp.Status(), resp.String())

	if err != nil {
		global.G_LOG.Errorf("MGDZ-CreateUser-apiUrl:%s status:%s response:%s", apiUrl, resp.Status(), resp.String())
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	code := gjson.Get(resp.String(), "code").String()
	if code == "InputValidationError" {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: "The field PlayerId must be a string with a maximum length of 50. Using only numbers and english alphabets. PlayerId 必须不能超过 50 个字符。请只使用数字和英文字母。"}
	}

	errorCode := gjson.Get(resp.String(), "error.code").String()
	if errorCode == "GeneralError" {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: "服务器错误"}
	}
	playerId := gjson.Get(resp.String(), "playerId").String()
	if playerId != "" {
		return &VenueResponse{Code: CreateUser_SUCCESS_CODE}
	}
	return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: ""}
}
func (v VenueMGDZ) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	apiUrl := fmt.Sprintf("%s/agents/%s/players/%s?properties=balance",
		v.VenueConfig.Url, v.VenueConfig.AgentCode, request.UserName)

	token := mgGetToken(v.VenueConfig)

	if token == "" {
		global.G_LOG.Errorf("MGDZ-GetBalance-apiUrl:%s status:%s token:%s", apiUrl, token)
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: "jwt token is nil"}
	}
	client := resty.New()
	resp, err := client.R().SetAuthToken(token).
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		Get(apiUrl)
	//global.G_LOG.Infof("MGDZ-GetUserBalance apiUrl:%s  token:%s  response:%s", apiUrl, token, resp.String())
	if err != nil {
		global.G_LOG.Errorf("MGDZ-GetBalance-apiUrl:%s status:%s response:%s", apiUrl, resp.Status(), resp.String())
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}

	balance := gjson.Get(resp.String(), "balance.total")

	if balance.Exists() {

		venueGetUserBalanceResponse := &VenueGetUserBalanceResponse{}
		venueGetUserBalanceResponse.Code = GetUserBalance_SUCCESS_CODE
		venueGetUserBalanceResponse.Data.Amount = balance.Float()
		return venueGetUserBalanceResponse
	}

	return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: resp.String()}
}

func (v VenueMGDZ) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueMGDZ) GetOrderNo() string {
	return tool.SnowflakeIdByKey("mg")

}
func (v VenueMGDZ) Deposit(request *VenueDepositRequest) *VenueDepositResponse {

	apiUrl := fmt.Sprintf("%s/agents/%s/WalletTransactions",
		v.VenueConfig.Url, v.VenueConfig.AgentCode)

	token := mgGetToken(v.VenueConfig)

	if token == "" {
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: "jwt token is nil"}
	}

	params := url.Values{}
	params.Set("playerId", request.UserName)
	params.Set("type", "Deposit")
	params.Set("amount", fmt.Sprintf("%0.2f", request.Amount))
	params.Set("externalTransactionId", request.OrderSn)
	client := resty.New()
	resp, err := client.R().SetAuthToken(token).
		SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"}).
		SetFormDataFromValues(params).
		Post(apiUrl)

	//global.G_LOG.Infof("MGDZ-Deposit-param:%s apiUrl:%s response:%s", tool.String(params), apiUrl, resp.String())
	if err != nil {
		global.G_LOG.Errorf("MGDZ-Deposit-param:%s apiUrl:%s response:%s", tool.String(params), apiUrl, resp.String())
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: resp.String()}
	}
	status := gjson.Get(resp.String(), "status")

	if status.Exists() {
		switch status.String() {
		case "Succeeded":
			venueDepositResponse := &VenueDepositResponse{Code: Deposit_SUCCESS_CODE}
			return venueDepositResponse
		}
	}
	errorCode := gjson.Get(resp.String(), "error.code")
	if errorCode.Exists() {
		switch errorCode.String() {
		case "InsufficientFunds":
			return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}
		case "GeneralError":
			return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: resp.String()}
		}
	}

	code := gjson.Get(resp.String(), "code")
	if code.Exists() {
		if code.String() == "InputValidationError" {
			return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}
		}
	}
	return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}

}

func (v VenueMGDZ) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {

	apiUrl := fmt.Sprintf("%s/agents/%s/WalletTransactions",
		v.VenueConfig.Url, v.VenueConfig.AgentCode)

	token := mgGetToken(v.VenueConfig)

	if token == "" {
		return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: "jwt token is nil"}
	}
	params := url.Values{}
	params.Set("playerId", request.UserName)
	params.Set("type", "Withdraw")
	params.Set("amount", "")
	params.Set("externalTransactionId", request.OrderSn)
	client := resty.New()
	resp, err := client.R().SetAuthToken(token).
		SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"}).
		SetFormDataFromValues(params).
		Post(apiUrl)

	if err != nil {
		global.G_LOG.Errorf("MGDZ-Withdraw-param:%s apiUrl:%s response:%s", tool.String(params), apiUrl, resp.String())
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: resp.String()}
	}

	global.G_LOG.Infof("MGDZ-Withdraw-param:%s apiUrl:%s response:%s", tool.String(params), apiUrl, resp.String())
	status := gjson.Get(resp.String(), "status")

	if status.Exists() {
		switch status.String() {
		case "Succeeded":
			venueWithdrawResponse := &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
			return venueWithdrawResponse
		}
	}
	errorCode := gjson.Get(resp.String(), "error.code")
	if errorCode.Exists() {
		switch errorCode.String() {
		case "InsufficientFunds":
			return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
		case "GeneralError":
			return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: resp.String()}
		}
	}

	return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
}
func (v VenueMGDZ) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {

	apiUrl := fmt.Sprintf("%s/agents/%s/players/%s/sessions",
		v.VenueConfig.Url, v.VenueConfig.AgentCode, request.UserName)

	token := mgGetToken(v.VenueConfig)

	if token == "" {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: "jwt token is nil"}
	}
	//params := map[string]string{
	//
	//	"contentCode ": request.GameCode,
	//	"platform ":    "Mobile",
	//	"langCode ":    "zh-CN",
	//	//"homeUrl":      request.ReturnUrl,
	//	//"bankUrl":      request.CashierURL,
	//}
	params := url.Values{}
	params.Set("contentCode", request.GameCode)
	params.Set("platform", "Mobile")
	params.Set("langCode", "zh-CN")
	params.Set("homeUrl", request.ReturnUrl)
	client := resty.New()
	resp, err := client.R().SetAuthToken(token).
		SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"}).
		SetFormDataFromValues(params).
		Post(apiUrl)

	//global.G_LOG.Infof("MGDZ-LoginGame-param:%s  apiUrl:%s response:%s", tool.String(params), apiUrl, resp.String())
	if err != nil {
		global.G_LOG.Errorf("MGDZ-LoginGame-param:%s  apiUrl:%s response:%s", tool.String(params), apiUrl, resp.String())
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}

	gameUrl := gjson.Get(resp.String(), "url").String()
	if gameUrl != "" {
		venueLoginGameResponse := &VenueLoginGameResponse{Code: LoginGame_SUCCESS_CODE}
		venueLoginGameResponse.Data.GameUrl = gameUrl
		return venueLoginGameResponse
	}
	return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: resp.String()}
}

func (v VenueMGDZ) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	apiUrl := fmt.Sprintf("%s/agents/%s/WalletTransactions/%s",
		v.VenueConfig.Url, v.VenueConfig.AgentCode, request.OrderSn)

	token := mgGetToken(v.VenueConfig)

	if token == "" {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: "jwt token is nil"}
	}
	client := resty.New()
	resp, err := client.R().SetAuthToken(token).
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		Get(apiUrl)

	//global.G_LOG.Infof("MGDZ-TransferConfirm:  apiUrl:%s response:%s", apiUrl, resp.String())
	if err != nil {
		global.G_LOG.Errorf("MGDZ-TransferConfirm:  apiUrl:%s response:%s", apiUrl, resp.String())
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}
	status := gjson.Get(resp.String(), "status")
	code := gjson.Get(resp.String(), "code")

	if status.Exists() {
		switch status.String() {
		case "Succeeded":
			return &VenueResponse{Code: TransferConfirm_SUCCESS_CODE}
		case "Inprogress", "Unconfirmed":
			return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: resp.String()}
		case "Failed":
			return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: resp.String()}
		}
	}

	switch code.String() {
	case "TransactionDoesNotExist":
		return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: resp.String()}
	}

	return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: resp.String()}
}

func (v VenueMGDZ) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueMGDZ) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueMGDZ) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

func mgGetToken(venueConfig conf.MGDZ) string {
	apiUrl := fmt.Sprintf("%s/connect/token", venueConfig.TokenUrl)
	params := map[string]string{
		"client_id":     venueConfig.AgentCode,
		"client_secret": venueConfig.Secret,
		"grant_type":    "client_credentials",
	}
	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"}).
		SetFormData(params).
		Post(apiUrl)
	//global.G_LOG.Infof("MGDZ-mgGetToken:  apiUrl:%s status:%s response:%s", apiUrl, resp.Status(), resp.String())
	if err != nil {
		global.G_LOG.Errorf("MGDZ get token err: %s", err.Error())
		return ""
	}

	return gjson.Get(resp.String(), "access_token").String()
}

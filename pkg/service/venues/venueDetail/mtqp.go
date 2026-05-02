package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type VenueMTQP struct {
	VenueConfig conf.MTQP
}

func NewVenueMTQP(venueConfig *conf.MTQP) IVenues {
	return &VenueMTQP{
		VenueConfig: *venueConfig,
	}
}

func (v VenueMTQP) CreateUser(request *VenueCreateUserRequest) *VenueResponse {

	request.Password = strings.ToLower(tool.MD5([]byte(request.Password)))
	rawData := fmt.Sprintf("{\"nickname\":\"%s\"}", request.NickName)
	codeOrignStr := fmt.Sprintf("%s%s", v.VenueConfig.Secret, rawData)

	code := strings.ToLower(tool.MD5([]byte(codeOrignStr)))
	data := base64.StdEncoding.EncodeToString([]byte(rawData))

	apiUrl := fmt.Sprintf("%s/services/dg/player/playerCreate2/%s/%s/%s/%s/%s",
		v.VenueConfig.Url, request.UserName, v.VenueConfig.MerchantId, request.Password, code, data)

	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/json"}).Post(apiUrl)

	//global.G_LOG.Infof("MTQP-CreateUser-param:%s apiUrl:%s response:%s", rawData, apiUrl, resp.String())

	if err != nil {
		global.G_LOG.Errorf("MTQP-CreateUse apiUrl:%s response:%s", apiUrl, resp.String())
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	resultCode := gjson.Get(resp.String(), "resultCode").String()

	if resultCode == "1" {
		return &VenueResponse{Code: CreateUser_SUCCESS_CODE}
	}

	return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: "resultCode:" + resultCode}
}
func (v VenueMTQP) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
	request.Password = strings.ToLower(tool.MD5([]byte(request.Password)))
	rawData := fmt.Sprintf("{\"currency\":\"%s\"}", "CNY")
	data := base64.StdEncoding.EncodeToString([]byte(rawData))
	apiUrl := fmt.Sprintf("%s/services/dg/player/getPlayerBalance/%s/%s/%s",
		v.VenueConfig.Url, request.UserName, v.VenueConfig.MerchantId, data)

	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/json"}).Post(apiUrl)

	//global.G_LOG.Infof("MTQP-GetUserBalance-param:%s apiUrl:%s response:%s", rawData, apiUrl, resp.String())
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}

	resultCode := gjson.Get(resp.String(), "resultCode").String()
	coinBalance := gjson.Get(resp.String(), "coinBalance").String()

	if resultCode == "1" {
		balanceFloat, err := strconv.ParseFloat(coinBalance, 64)
		if err != nil {
			return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
		}
		venueGetUserBalanceResponse := &VenueGetUserBalanceResponse{}
		venueGetUserBalanceResponse.Code = GetUserBalance_SUCCESS_CODE
		venueGetUserBalanceResponse.Data.Amount = tool.Decimal(balanceFloat)
		return venueGetUserBalanceResponse
	}

	return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: resp.String()}
}

func (v VenueMTQP) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueMTQP) GetOrderNo() string {
	return tool.SnowflakeIdByKey("mt")

}
func (v VenueMTQP) Deposit(request *VenueDepositRequest) *VenueDepositResponse {

	request.Password = strings.ToLower(tool.MD5([]byte(request.Password)))
	rawData := fmt.Sprintf("{\"merchantId\":\"%s\",\"playerName\":\"%s\",\"extTransId\":\"%s\",\"coins\":\"%0.2f00\",\"currency\":\"%s\"}",
		v.VenueConfig.MerchantId,
		request.UserName,
		request.OrderSn,
		request.Amount,
		"CNY")
	codeOrignStr := fmt.Sprintf("%s%s", v.VenueConfig.Secret, rawData)

	code := strings.ToLower(tool.MD5([]byte(codeOrignStr)))
	data := base64.StdEncoding.EncodeToString([]byte(rawData))
	apiUrl := fmt.Sprintf("%s/services/dg/player/deposit2/%s/%s/%0.2f00/%s/%s/%s",
		v.VenueConfig.Url, v.VenueConfig.MerchantId, request.UserName, request.Amount, request.OrderSn, code, data)

	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/json"}).Post(apiUrl)

	//global.G_LOG.Infof("MTQP-Deposit-param:%s apiUrl:%s response:%s", rawData, apiUrl, resp.String())
	if err != nil {
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: err.Error()}
	}
	status := gjson.Get(resp.String(), "resultCode").String()

	switch status {
	case "1":
		venueDepositResponse := &VenueDepositResponse{Code: Deposit_SUCCESS_CODE}
		return venueDepositResponse
	case "0":
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: resp.String()}
	case "2", "3", "4", "6", "15", "21", "32", "40":
		//0	充值异常 1充值成功  2商户不存在 3商户无效 4商户用户不存在  6商户用户系统禁用  15	IP被限制  21  32	可选参数错误 40维护模式
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: "错误码" + status}
	}

	return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}

}

func (v VenueMTQP) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {

	request.Password = strings.ToLower(tool.MD5([]byte(request.Password)))
	rawData := fmt.Sprintf("{\"merchantId\":\"%s\",\"playerName\":\"%s\",\"extTransId\":\"%s\",\"coins\":\"%0.2f00\",\"currency\":\"%s\"}",
		v.VenueConfig.MerchantId,
		request.UserName,
		request.OrderSn,
		request.Amount,
		"CNY")
	codeOrignStr := fmt.Sprintf("%s%s", v.VenueConfig.Secret, rawData)

	code := strings.ToLower(tool.MD5([]byte(codeOrignStr)))
	data := base64.StdEncoding.EncodeToString([]byte(rawData))
	apiUrl := fmt.Sprintf("%s/services/dg/player/withdraw2/%s/%s/%0.2f00/%s/%s/%s",
		v.VenueConfig.Url, v.VenueConfig.MerchantId, request.UserName, request.Amount, request.OrderSn, code, data)

	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/json"}).Post(apiUrl)

	//global.G_LOG.Infof("MTQP-Withdraw-param:%s apiUrl:%s response:%s", rawData, apiUrl, resp.String())
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: err.Error()}
	}
	status := gjson.Get(resp.String(), "resultCode").String()

	switch status {
	case "1":
		venueDepositResponse := &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
		return venueDepositResponse
	case "0":
		venueDepositResponse := &VenueWithdrawResponse{Code: Withdraw_Processing_CODE}
		return venueDepositResponse
	case "2", "3", "4", "6", "9", "12", "15", "21", "32", "36", "40":
		//0	充值异常 1充值成功  2商户不存在 3商户无效 4商户用户不存在  6商户用户系统禁用 9	商户用户金币余额不足
		//12游戏未结算 15	IP被限制  21  32	可选参数错误 40维护模式 36	存在未处理订单
		venueDepositResponse := &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: "错误码" + status}
		return venueDepositResponse
	}

	return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
}
func (v VenueMTQP) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {

	request.Password = strings.ToLower(tool.MD5([]byte(request.Password)))
	rawData := fmt.Sprintf("{\"gameCode\":\"%s\",\"lang\":\"\",\"roomID\":\"\"}", request.GameCode)
	codeOrignStr := fmt.Sprintf("%s%s", v.VenueConfig.Secret, rawData)

	code := strings.ToLower(tool.MD5([]byte(codeOrignStr)))
	data := base64.StdEncoding.EncodeToString([]byte(rawData))
	apiUrl := fmt.Sprintf("%s/services/dg/player/playerPlatformUrl/%s/%s/%s/%s/%s",
		v.VenueConfig.Url, v.VenueConfig.MerchantId, request.UserName, request.Password, code, data)

	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/json"}).Post(apiUrl)

	//global.G_LOG.Infof("MTQP-LoginGame-param:%s  apiUrl:%s response:%s", rawData, apiUrl, resp.String())
	if err != nil {
		global.G_LOG.Errorf("MTQP-LoginGame-apiUrl:%s response:%s", apiUrl, resp.String())
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

func (v VenueMTQP) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
	request.Password = strings.ToLower(tool.MD5([]byte(request.Password)))
	rawData := fmt.Sprintf("{\"currency\":\"%s\"}", "CNY")

	data := base64.StdEncoding.EncodeToString([]byte(rawData))
	apiUrl := fmt.Sprintf("%s/services/dg/player/queryTransbyId/%s/%s/%s/%s",
		v.VenueConfig.Url, request.UserName, v.VenueConfig.MerchantId, request.OrderSn, data)

	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{"Content-Type": "application/json"}).Post(apiUrl)

	//global.G_LOG.Infof("MTQP-TransferConfirm-param:%s  apiUrl:%s response:%s", rawData, apiUrl, resp.String())
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}
	status := gjson.Get(resp.String(), "resultCode").String()
	switch status {
	case "1":
		venueResponse := &VenueResponse{Code: TransferConfirm_SUCCESS_CODE}
		return venueResponse
	case "0":
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: resp.String()}
	case "2", "3", "15", "40":
		//0	充值异常 1充值成功  2商户不存在 3商户无效 4商户用户不存在  6商户用户系统禁用 9	商户用户金币余额不足
		//12游戏未结算 15	IP被限制  21  32	可选参数错误 40维护模式 36	存在未处理订单
		return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: "错误码" + status}
	}

	return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: resp.String()}
}

func (v VenueMTQP) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueMTQP) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueMTQP) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

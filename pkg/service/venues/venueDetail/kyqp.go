package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/tool"
	"bootpkg/common/tool/crypt"
	"bootpkg/pkg/service/venues/venuevo"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"time"
)

type VenueKYQP struct {
	VenueConfig conf.KYQP
}

func NewKYQP(venueConfig *conf.KYQP) IVenues {
	return &VenueKYQP{
		VenueConfig: *venueConfig,
	}
}

func (v VenueKYQP) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	ip := "127.0.0.1"
	if request.Ip != "" {
		ip = request.Ip
	}

	timestamp := time.Now().UnixNano() / 1e6
	orderid := fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, request.UserName)
	paramStr := fmt.Sprintf("s=0&account=%s&money=0&orderid=%s&ip=%s&lineCode=1&KindID=0",
		request.UserName, orderid, ip)

	param := crypt.EncryptAes([]byte(v.VenueConfig.DesKey), []byte(paramStr))

	key := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, v.VenueConfig.Md5Key)))
	apiUrl := fmt.Sprintf("%s?agent=%s&timestamp=%d&param=%s&key=%s", v.VenueConfig.Url, v.VenueConfig.AgentId, timestamp, url.QueryEscape(param), key)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("KYQP-CreateUser-param:%s apiUrl:%s response:%s", paramStr, apiUrl, resp.String())

	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	var KYQPRegisterResp venuevo.KYQPSportRegisterResp
	err = tool.JsonUnmarshal(resp.Body(), &KYQPRegisterResp)
	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	if KYQPRegisterResp.S == 100 && KYQPRegisterResp.D.Code == 0 {
		return &VenueResponse{Code: CreateUser_SUCCESS_CODE}
	}

	return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: ""}
}
func (v VenueKYQP) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {

	timestamp := time.Now().UnixNano() / 1e6
	paramStr := fmt.Sprintf("s=7&account=%s",
		request.UserName)

	param := crypt.EncryptAes([]byte(v.VenueConfig.DesKey), []byte(paramStr))

	key := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, v.VenueConfig.Md5Key)))
	apiUrl := fmt.Sprintf("%s?agent=%s&timestamp=%d&param=%s&key=%s", v.VenueConfig.Url, v.VenueConfig.AgentId, timestamp, url.QueryEscape(param), key)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("KYQP-GetUserBalance-param:%v  response:%s", paramStr, resp.String())

	var KYQPBalanceResp venuevo.KYQPBalanceResp
	err = tool.JsonUnmarshal(resp.Body(), &KYQPBalanceResp)
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}
	if KYQPBalanceResp.S == 107 && KYQPBalanceResp.D.Code == 0 {
		venueGetUserBalanceResponse := &VenueGetUserBalanceResponse{}
		venueGetUserBalanceResponse.Code = GetUserBalance_SUCCESS_CODE
		venueGetUserBalanceResponse.Data.Amount = KYQPBalanceResp.D.TotalMoney
		return venueGetUserBalanceResponse
	}

	return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: ""}
}

func (v VenueKYQP) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueKYQP) GetOrderNo() string {
	timestamp := time.Now().UnixNano() / 1e6
	timestampStr := fmt.Sprintf("%d", timestamp)
	return v.VenueConfig.AgentId + time.Now().Format("20060102150405") + timestampStr[len(timestampStr)-3:]
	//TODO implement me
	//return tool.MakeTransferrder()
}
func (v VenueKYQP) Deposit(request *VenueDepositRequest) *VenueDepositResponse {

	timestamp := time.Now().UnixNano() / 1e6
	paramStr := fmt.Sprintf("s=2&account=%s&money=%0.2f&orderid=%s", request.UserName, request.Amount, request.OrderSn)

	param := crypt.EncryptAes([]byte(v.VenueConfig.DesKey), []byte(paramStr))

	key := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, v.VenueConfig.Md5Key)))
	apiUrl := fmt.Sprintf("%s?agent=%s&timestamp=%d&param=%s&key=%s", v.VenueConfig.Url, v.VenueConfig.AgentId, timestamp, url.QueryEscape(param), key)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("KYQP-Deposit-param:%s  response:%s", paramStr, resp.String())

	if err != nil {
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}
	var KYQPDepositResp venuevo.KYQPDepositResp
	err = tool.JsonUnmarshal(resp.Body(), &KYQPDepositResp)
	if err != nil {
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}
	if KYQPDepositResp.S == 102 && KYQPDepositResp.D.Code == 0 {
		venueDepositResponse := &VenueDepositResponse{Code: Deposit_SUCCESS_CODE}
		return venueDepositResponse
	}

	return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}
}

func (v VenueKYQP) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	timestamp := time.Now().UnixNano() / 1e6
	paramStr := fmt.Sprintf("s=3&account=%s&money=%0.2f&orderid=%s", request.UserName, request.Amount, request.OrderSn)

	param := crypt.EncryptAes([]byte(v.VenueConfig.DesKey), []byte(paramStr))

	key := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, v.VenueConfig.Md5Key)))
	apiUrl := fmt.Sprintf("%s?agent=%s&timestamp=%d&param=%s&key=%s", v.VenueConfig.Url, v.VenueConfig.AgentId, timestamp, url.QueryEscape(param), key)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("KYQP-Withdraw-param:%S  response:%s", paramStr, resp.String())

	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
	}
	var KYQPWithdrawResp venuevo.KYQPWithdrawResp
	err = tool.JsonUnmarshal(resp.Body(), &KYQPWithdrawResp)
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
	}
	if KYQPWithdrawResp.S == 103 && KYQPWithdrawResp.D.Code == 0 {
		venueWithdrawResponse := &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
		return venueWithdrawResponse
	}

	return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
}
func (v VenueKYQP) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	ip := "127.0.0.1"
	if request.IP != "" {
		ip = request.IP
	}

	timestamp := time.Now().UnixNano() / 1e6
	orderid := fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, request.UserName)
	paramStr := fmt.Sprintf("s=0&account=%s&money=0&orderid=%s&ip=%s&lineCode=text11&KindID=%s",
		request.UserName, orderid, ip, request.GameCode)

	param := crypt.EncryptAes([]byte(v.VenueConfig.DesKey), []byte(paramStr))

	key := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, v.VenueConfig.Md5Key)))
	apiUrl := fmt.Sprintf("%s/?agent=%s&timestamp=%d&param=%s&key=%s", v.VenueConfig.Url, v.VenueConfig.AgentId, timestamp, url.QueryEscape(param), key)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	//global.G_LOG.Infof("KYQP-LoginGame-param:%v apiUrl:%s response:%s httpstatus:%s", paramStr, apiUrl, resp.String(), resp.Status())

	var KYQPRegisterResp venuevo.KYQPSportRegisterResp
	err = tool.JsonUnmarshal(resp.Body(), &KYQPRegisterResp)
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}

	if KYQPRegisterResp.S == 100 && KYQPRegisterResp.D.Code == 0 {
		venueLoginGameResponse := &VenueLoginGameResponse{
			Code: LoginGame_SUCCESS_CODE,
		}
		venueLoginGameResponse.Data.GameUrl = KYQPRegisterResp.D.URL

		if request.ReturnUrl != "" {
			venueLoginGameResponse.Data.GameUrl = venueLoginGameResponse.Data.GameUrl + "&backUrl=" + request.ReturnUrl
		}

		return venueLoginGameResponse
	}
	return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: ""}
}

func (v VenueKYQP) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {

	timestamp := time.Now().UnixNano() / 1e6
	paramStr := fmt.Sprintf("s=4&orderid=%s", request.OrderSn)

	param := crypt.EncryptAes([]byte(v.VenueConfig.DesKey), []byte(paramStr))

	key := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.AgentId, timestamp, v.VenueConfig.Md5Key)))
	apiUrl := fmt.Sprintf("%s/channelHandle?agent=%s&timestamp=%d&param=%s&key=%s", v.VenueConfig.Url, v.VenueConfig.AgentId, timestamp, url.QueryEscape(param), key)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}
	//global.G_LOG.Infof("KYQP-TransferConfirm-param:%v  response:%s", paramStr, resp.String())

	var KYQPTransferConfirmResp venuevo.KYQPTransferConfirmResp
	err = tool.JsonUnmarshal(resp.Body(), &KYQPTransferConfirmResp)
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}

	if KYQPTransferConfirmResp.S == 104 {
		switch KYQPTransferConfirmResp.D.Status {
		case -1, 2:
			return &VenueResponse{Code: TransferConfirm_FAIL_CODE}
		case 0:
			return &VenueResponse{Code: TransferConfirm_SUCCESS_CODE}
		}
	}

	return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: resp.String()}
}

func (v VenueKYQP) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueKYQP) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueKYQP) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

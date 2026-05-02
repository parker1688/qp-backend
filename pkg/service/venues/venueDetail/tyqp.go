package venueDetail

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/common/tool/crypt"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues/venuevo"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type VenueTYQP struct {
	VenueConfig conf.TYQP
}

func NewTYQP(venueConfig *conf.TYQP) IVenues {
	return &VenueTYQP{
		VenueConfig: *venueConfig,
	}
}

func (v VenueTYQP) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
	timestamp := time.Now().UnixNano() / 1e6

	orderid := uuid.NewString()
	paramStr := fmt.Sprintf("username=%s&money=0&orderid=%s0&ip=127.0.0.1",
		request.UserName, orderid)
	param, err := tyAesEncrypt(paramStr, &v.VenueConfig)
	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.Userid, timestamp, v.VenueConfig.Md5Key)))

	apiUrl := fmt.Sprintf("%s/access/gameauth/?userid=%s&timestamp=%d&params=%s&sig=%s", v.VenueConfig.Url, v.VenueConfig.Userid, timestamp, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("TYQP-CreateUser-param:%s apiUrl:%s response:%s", paramStr, apiUrl, resp.String())

	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	var TYQPRegisterResp venuevo.TYQPSportRegisterResp
	err = tool.JsonUnmarshal(resp.Body(), &TYQPRegisterResp)
	if err != nil {
		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
	}

	if TYQPRegisterResp.Code == "0" {
		return &VenueResponse{Code: CreateUser_SUCCESS_CODE}
	}

	return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: ""}
}
func (v VenueTYQP) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {

	timestamp := time.Now().UnixNano() / 1e6

	paramStr := fmt.Sprintf("username=%s@%s&ip=127.0.0.1",
		request.UserName, v.VenueConfig.Apisuffix)
	param, err := tyAesEncrypt(paramStr, &v.VenueConfig)
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.Userid, timestamp, v.VenueConfig.Md5Key)))

	apiUrl := fmt.Sprintf("%s/access/getmoney/?userid=%s&timestamp=%d&params=%s&sig=%s", v.VenueConfig.Url, v.VenueConfig.Userid, timestamp, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}

	//global.G_LOG.Infof("TYQP-GetUserBalance-param:%v  response:%s", paramStr, resp.String())

	var TYQPBalanceResp venuevo.TYQPBalanceResp
	err = tool.JsonUnmarshal(resp.Body(), &TYQPBalanceResp)
	if err != nil {
		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
	}
	if TYQPBalanceResp.Code == "0" {
		amount, err := strconv.ParseFloat(TYQPBalanceResp.BagMoney, 64)
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

func (v VenueTYQP) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
	panic("implement me")
}

func (v VenueTYQP) GetOrderNo() string {

	//TODO implement me
	return tool.MakeTransferrder()
}
func (v VenueTYQP) Deposit(request *VenueDepositRequest) *VenueDepositResponse {

	timestamp := time.Now().UnixNano() / 1e6

	paramStr := fmt.Sprintf("username=%s@%s&money=%0.2f&orderid=%s0&ip=127.0.0.1",
		request.UserName, v.VenueConfig.Apisuffix, request.Amount, request.OrderSn)
	param, err := tyAesEncrypt(paramStr, &v.VenueConfig)
	if err != nil {
		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: err.Error()}
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.Userid, timestamp, v.VenueConfig.Md5Key)))

	apiUrl := fmt.Sprintf("%s/access/orderin/?userid=%s&timestamp=%d&params=%s&sig=%s", v.VenueConfig.Url, v.VenueConfig.Userid, timestamp, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("TYQP-Deposit-param:%s  response:%s", paramStr, resp.String())

	if err != nil {
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}
	var TYQPDepositResp venuevo.TYQPDepositResp
	err = tool.JsonUnmarshal(resp.Body(), &TYQPDepositResp)
	if err != nil {
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}

	amount, err := strconv.ParseFloat(TYQPDepositResp.BagMoney, 64)
	if err != nil {
		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
	}
	if TYQPDepositResp.Code == "0" && amount > 0 {
		venueDepositResponse := &VenueDepositResponse{Code: Deposit_SUCCESS_CODE}
		return venueDepositResponse
	}

	return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}
}

func (v VenueTYQP) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
	timestamp := time.Now().UnixNano() / 1e6

	paramStr := fmt.Sprintf("username=%s@%s&money=%0.2f&orderid=%s0&ip=127.0.0.1",
		request.UserName, v.VenueConfig.Apisuffix, request.Amount, request.OrderSn)
	param, err := tyAesEncrypt(paramStr, &v.VenueConfig)
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: err.Error()}
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.Userid, timestamp, v.VenueConfig.Md5Key)))

	apiUrl := fmt.Sprintf("%s/access/orderout/?userid=%s&timestamp=%d&params=%s&sig=%s", v.VenueConfig.Url, v.VenueConfig.Userid, timestamp, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)

	//global.G_LOG.Infof("TYQP-Withdraw-param:%S  response:%s", paramStr, resp.String())

	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
	}
	var TYQPWithdrawResp venuevo.TYQPWithdrawResp
	err = tool.JsonUnmarshal(resp.Body(), &TYQPWithdrawResp)
	if err != nil {
		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
	}
	if TYQPWithdrawResp.Code == "0" {
		venueWithdrawResponse := &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
		return venueWithdrawResponse
	}

	return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
}
func (v VenueTYQP) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
	timestamp := time.Now().UnixNano() / 1e6

	orderid := uuid.NewString()
	paramStr := fmt.Sprintf("username=%s&money=0&orderid=%s&ip=127.0.0.1&gametype=%s&showbackbutton=1&backurl=%s",
		request.UserName, orderid, request.GameCode, request.ReturnUrl)
	param, err := tyAesEncrypt(paramStr, &v.VenueConfig)
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.Userid, timestamp, v.VenueConfig.Md5Key)))

	apiUrl := fmt.Sprintf("%s/access/gameauth/?userid=%s&timestamp=%d&gametype=%s&params=%s&sig=%s", v.VenueConfig.Url, v.VenueConfig.Userid, timestamp, request.GameCode, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}
	//global.G_LOG.Infof("TYQP-LoginGame-param:%v apiUrl:%s response:%s httpstatus:%s", paramStr, apiUrl, resp.String(), resp.Status())

	var TYQPRegisterResp venuevo.TYQPSportRegisterResp
	err = tool.JsonUnmarshal(resp.Body(), &TYQPRegisterResp)
	if err != nil {
		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
	}

	if TYQPRegisterResp.Code == "0" {
		venueLoginGameResponse := &VenueLoginGameResponse{
			Code: LoginGame_SUCCESS_CODE,
		}
		venueLoginGameResponse.Data.GameUrl = TYQPRegisterResp.Loginurl
		return venueLoginGameResponse
	}
	return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: ""}
}

func (v VenueTYQP) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {

	timestamp := time.Now().UnixNano() / 1e6

	paramStr := fmt.Sprintf("orderid=%s&ip=127.0.0.1",
		request.OrderSn)
	param, err := tyAesEncrypt(paramStr, &v.VenueConfig)
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: err.Error()}
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", v.VenueConfig.Userid, timestamp, v.VenueConfig.Md5Key)))

	apiUrl := fmt.Sprintf("%s/access/orderinfo/?userid=%s&timestamp=%d&params=%s&sig=%s", v.VenueConfig.Url, v.VenueConfig.Userid, timestamp, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}
	//global.G_LOG.Infof("TYQP-TransferConfirm-param:%v  response:%s", paramStr, resp.String())

	var TYQPTransferConfirmResp venuevo.TYQPTransferConfirmResp
	err = tool.JsonUnmarshal(resp.Body(), &TYQPTransferConfirmResp)
	if err != nil {
		return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
	}

	if TYQPTransferConfirmResp.Code == "0" {
		switch TYQPTransferConfirmResp.Status {
		case -1, 2:
			return &VenueResponse{Code: TransferConfirm_FAIL_CODE}
		case 0:
			return &VenueResponse{Code: TransferConfirm_SUCCESS_CODE}
		}
	}

	return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: resp.String()}
}

func (v VenueTYQP) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueTYQP) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
	//TODO implement me
	panic("implement me")
}

func (v VenueTYQP) AmountLimitFix(amount float64, currency string) float64 {
	//TODO implement me
	return amount
}

// 获取 tyqp 回放录像
func GetTYQPPlayback(req *vo.VenuePlaybackReq) *vo.VenuePlaybackResp {
	venueResp := vo.VenuePlaybackResp{}
	vConf := global.CONFIG.Venue.TYQP
	timestamp := time.Now().UnixNano() / 1e6

	paramStr := fmt.Sprintf("desk_uuid=%s&ip=127.0.0.1",
		req.TableId)
	param, err := tyAesEncrypt(paramStr, &vConf)
	if err != nil {
		venueResp.Code = vo.Playback_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	sign := tool.MD5([]byte(fmt.Sprintf("%s%d%s", vConf.Userid, timestamp, vConf.Md5Key)))
	apiUrl := fmt.Sprintf("%s/access/deskrecordurl/?userid=%s&timestamp=%d&params=%s&sig=%s",
		vConf.Url, vConf.Userid, timestamp, param, sign)

	client := resty.New()
	resp, err := client.R().Get(apiUrl)
	if err != nil {
		venueResp.Code = vo.Playback_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	respStr := resp.String()
	apiLog := fmt.Sprintf("TYQP-Playback-param: %s apiUrl: %s response:%s", paramStr, apiUrl, respStr)
	//global.G_LOG.Infof(apiLog)

	var respData venuevo.TYQPPlaybackResp
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("%s JsonUnmarshal err: %v", apiLog, err.Error())
		venueResp.Code = vo.Playback_FAIL_CODE
		venueResp.Msg = err.Error()
		return &venueResp
	}
	if respData.Code == "0" {
		venueResp.Code = vo.Playback_SUCCESS_CODE
		venueResp.Data.PlaybackUrl = respData.DeskUuidUrl
		return &venueResp
	}

	global.G_LOG.Errorf("%s respCode failed", apiLog)
	venueResp.Code = vo.Playback_FAIL_CODE
	venueResp.Msg = respStr
	return &venueResp
}

func tyAesEncrypt(param string, venueConfig *conf.TYQP) (string, error) {
	h := sha256.New()
	h.Write([]byte(venueConfig.DesKey))
	sum := h.Sum(nil)
	vi := sum[0:16]

	p, err := crypt.AesCBCEncrypt([]byte(param), sum, vi)
	if err != nil {
		return "", err
	}
	param = base64.StdEncoding.EncodeToString(p)
	param = url.QueryEscape(param)
	return param, nil
}

package venueDetail

//import (
//	"bootpkg/common/global"
//	"bootpkg/common/tool"
//	"bootpkg/pkg/service/venues/venuevo"
//	"fmt"
//	"github.com/go-resty/resty/v2"
//	"github.com/tidwall/gjson"
//	"net/url"
//	"strings"
//)
//
//type VenueSBTY struct {
//	VenueConfig conf.ShabaTy
//}
//
//func NewSHABATY(venueConfig *conf.ShabaTy) IVenues {
//	return &VenueSHABATY{
//		VenueConfig: *venueConfig,
//	}
//}
//
//func (v VenueSHABATY) CreateUser(request *VenueCreateUserRequest) *VenueResponse {
//
//	postData := url.Values{}
//	postData.Add("vendor_id", v.VenueConfig.VendorId)
//	postData.Add("vendor_member_id", request.UserName)
//	postData.Add("operatorId", v.VenueConfig.OperatorId)
//	postData.Add("username", request.UserName)
//	postData.Add("oddstype", "2")
//	//postData.Add("currency", request.Currency)
//	postData.Add("currency", v.VenueConfig.Currency)
//	postData.Add("maxtransfer", "1000000")
//	postData.Add("mintransfer", "1")
//
//	apiUrl := fmt.Sprintf("%s/api/CreateMember", v.VenueConfig.Url)
//
//	client := resty.New()
//	headers := map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//
//	resp, err := client.R().SetHeaders(headers).SetFormDataFromValues(postData).Post(apiUrl)
//	global.G_LOG.Infof("SHABATY-CreateUser-apiUrl:%s  FormData:%v  response:%s", apiUrl, postData, resp.String())
//
//	if err != nil {
//		return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: err.Error()}
//	}
//
//	errCode := gjson.Get(resp.String(), "error_code")
//	if errCode.Exists() && errCode.Int() == 0 {
//		return &VenueResponse{Code: CreateUser_SUCCESS_CODE}
//	}
//
//	return &VenueResponse{Code: CreateUser_FAIL_CODE, Msg: resp.String()}
//}
//func (v VenueSHABATY) GetUserBalance(request *VenueGetUserBalanceRequest) *VenueGetUserBalanceResponse {
//	postData := url.Values{}
//	postData.Add("vendor_id", v.VenueConfig.VendorId)
//	postData.Add("vendor_member_ids", request.UserName)
//	postData.Add("wallet_id", "1")
//
//	apiUrl := fmt.Sprintf("%s/api/CheckUserBalance", v.VenueConfig.Url)
//
//	client := resty.New()
//	headers := map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//
//	resp, err := client.R().SetHeaders(headers).SetFormDataFromValues(postData).Post(apiUrl)
//
//	global.G_LOG.Infof("SHABATY-GetUserBalance-apiUrl:%s  FormData:%v  response:%s", apiUrl, postData, resp.String())
//
//	if err != nil {
//		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
//	}
//
//	var ShaBaTyBalanceResp venuevo.ShaBaTyBalanceResp
//	err = tool.JsonUnmarshal(resp.Body(), &ShaBaTyBalanceResp)
//	if err != nil {
//		return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: err.Error()}
//	}
//	if len(ShaBaTyBalanceResp.Data) > 0 && ShaBaTyBalanceResp.ErrorCode == 0 {
//		venueGetUserBalanceResponse := &VenueGetUserBalanceResponse{}
//		venueGetUserBalanceResponse.Code = GetUserBalance_SUCCESS_CODE
//		venueGetUserBalanceResponse.Data.Amount = ShaBaTyBalanceResp.Data[0].Balance
//		return venueGetUserBalanceResponse
//	}
//
//	return &VenueGetUserBalanceResponse{Code: GetUserBalance_FAIL_CODE, Msg: ""}
//}
//
//func (v VenueSHABATY) UpdateUserPassword(request *VenueUpdateUserPasswordRequest) *VenueResponse {
//	panic("implement me")
//}
//
//func (v VenueSHABATY) GetOrderNo() string {
//	//TODO implement me
//	return v.VenueConfig.OperatorId + "_" + tool.MakeTransferrder()
//}
//func (v VenueSHABATY) Deposit(request *VenueDepositRequest) *VenueDepositResponse {
//	postData := url.Values{}
//	postData.Add("vendor_id", v.VenueConfig.VendorId)
//	postData.Add("vendor_member_id", request.UserName)
//	postData.Add("vendor_trans_id", request.OrderSn)
//	postData.Add("amount", fmt.Sprintf("%0.2f", request.Amount))
//	postData.Add("currency", v.VenueConfig.Currency)
//	postData.Add("direction", "1")
//	postData.Add("wallet_id", "1")
//
//	apiUrl := fmt.Sprintf("%s/api/FundTransfer", v.VenueConfig.Url)
//
//	client := resty.New()
//	headers := map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//
//	resp, err := client.R().SetHeaders(headers).SetFormDataFromValues(postData).Post(apiUrl)
//
//	global.G_LOG.Infof("SHABATY-Deposit-apiUrl:%s  FormData:%v  response:%s", apiUrl, postData, resp.String())
//
//	if err != nil {
//		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
//	}
//	var ShaBaTyDepositResp venuevo.ShaBaTyDepositResp
//	err = tool.JsonUnmarshal(resp.Body(), &ShaBaTyDepositResp)
//	if err != nil {
//		return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: err.Error()}
//	}
//	if ShaBaTyDepositResp.Data.TransID > 0 && ShaBaTyDepositResp.ErrorCode == 0 && ShaBaTyDepositResp.Data.Status == 0 {
//		venueDepositResponse := &VenueDepositResponse{Code: Deposit_SUCCESS_CODE}
//		return venueDepositResponse
//	}
//	switch ShaBaTyDepositResp.ErrorCode {
//	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16:
//		return &VenueDepositResponse{Code: Deposit_FAIL_CODE, Msg: resp.String()}
//	}
//
//	return &VenueDepositResponse{Code: Deposit_Processing_CODE, Msg: resp.String()}
//}
//
//func (v VenueSHABATY) Withdraw(request *VenueWithdrawRequest) *VenueWithdrawResponse {
//	postData := url.Values{}
//	postData.Add("vendor_id", v.VenueConfig.VendorId)
//	postData.Add("vendor_member_id", request.UserName)
//	postData.Add("vendor_trans_id", request.OrderSn)
//	postData.Add("amount", fmt.Sprintf("%0.2f", request.Amount))
//	postData.Add("currency", v.VenueConfig.Currency)
//	postData.Add("direction", "0")
//	postData.Add("wallet_id", "1")
//
//	apiUrl := fmt.Sprintf("%s/api/FundTransfer", v.VenueConfig.Url)
//
//	client := resty.New()
//	headers := map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//
//	resp, err := client.R().SetHeaders(headers).SetFormDataFromValues(postData).Post(apiUrl)
//
//	global.G_LOG.Infof("SHABATY-Withdraw-apiUrl:%s  FormData:%v  response:%s", apiUrl, postData, resp.String())
//
//	if err != nil {
//		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
//	}
//	var ShaBaTyWithdrawResp venuevo.ShaBaTyWithdrawResp
//	err = tool.JsonUnmarshal(resp.Body(), &ShaBaTyWithdrawResp)
//	if err != nil {
//		return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: err.Error()}
//	}
//	if ShaBaTyWithdrawResp.Data.TransID > 0 && ShaBaTyWithdrawResp.ErrorCode == 0 && ShaBaTyWithdrawResp.Data.Status == 0 {
//		venueDepositResponse := &VenueWithdrawResponse{Code: Withdraw_SUCCESS_CODE}
//		return venueDepositResponse
//	}
//	switch ShaBaTyWithdrawResp.ErrorCode {
//	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16:
//		return &VenueWithdrawResponse{Code: Withdraw_FAIL_CODE, Msg: resp.String()}
//	}
//
//	return &VenueWithdrawResponse{Code: Withdraw_Processing_CODE, Msg: resp.String()}
//}
//
//func (v VenueSHABATY) LoginGame(request *VenueLoginGameRequest) *VenueLoginGameResponse {
//	platForm := "2"
//	switch strings.ToUpper(request.ClientType) {
//	case "PC":
//		platForm = "1"
//	case "H5":
//		platForm = "2"
//	default:
//		platForm = "2"
//	}
//
//	postData := url.Values{}
//	postData.Add("vendor_id", v.VenueConfig.VendorId)
//	postData.Add("vendor_member_id", request.UserName)
//	postData.Add("platform", platForm)
//
//	apiUrl := fmt.Sprintf("%s/api/GetSabaUrl", v.VenueConfig.Url)
//
//	client := resty.New()
//	headers := map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//
//	resp, err := client.R().SetHeaders(headers).SetFormDataFromValues(postData).Post(apiUrl)
//	if err != nil {
//		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
//	}
//	global.G_LOG.Infof("SHABATY-LoginGame-apiUrl:%s  FormData:%v  response:%s", apiUrl, postData, resp.String())
//
//	if resp.String() == "" {
//		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: resp.Status()}
//	}
//
//	var ShaBaTyLoginGameResp venuevo.ShaBaTyLoginGameResp
//	err = tool.JsonUnmarshal(resp.Body(), &ShaBaTyLoginGameResp)
//	if err != nil {
//		return &VenueLoginGameResponse{Code: LoginGame_FAIL_CODE, Msg: err.Error()}
//	}
//
//	venueLoginGameResponse := &VenueLoginGameResponse{
//		Code: LoginGame_SUCCESS_CODE,
//	}
//	//venueLoginGameResponse.Data.GameUrl = ShaBaTyLoginGameResp.Data + "&lang=th"
//	venueLoginGameResponse.Data.GameUrl = ShaBaTyLoginGameResp.Data
//	return venueLoginGameResponse
//}
//
//func (v VenueSHABATY) TransferConfirm(request *VenueTransferConfirmRequest) *VenueResponse {
//	postData := url.Values{}
//	postData.Add("vendor_id", v.VenueConfig.VendorId)
//	postData.Add("vendor_trans_id", request.OrderSn)
//	postData.Add("wallet_id", "1")
//
//	apiUrl := fmt.Sprintf("%s/api/CheckFundTransfer", v.VenueConfig.Url)
//
//	client := resty.New()
//	headers := map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//
//	resp, err := client.R().SetHeaders(headers).SetFormDataFromValues(postData).Post(apiUrl)
//	if err != nil {
//		global.G_LOG.Errorf("[TransferConfirm] Http post failed: %v", err.Error())
//		return &VenueResponse{Code: TransferConfirm_FAIL_CODE, Msg: err.Error()}
//	}
//
//	global.G_LOG.Infof("SHABATY-TransferConfirm-postData:%v  response:%s", postData, resp.String())
//
//	//var ShaBaTyTransferConfirmResp venuevo.ShaBaTyTransferConfirmResp
//	//err = tool.JsonUnmarshal(resp.Body(), &ShaBaTyTransferConfirmResp)
//	//if err != nil {
//	//	return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: err.Error()}
//	//}
//
//	//3040 ：交易不存在
//	errCode := gjson.Get(resp.String(), "error_code")
//	statusCode := gjson.Get(resp.String(), "Data.status")
//	if errCode.Exists() {
//		switch errCode.Int() {
//		case 0:
//			if statusCode.Exists() {
//				switch statusCode.Int() {
//				case 0:
//					return &VenueResponse{Code: TransferConfirm_SUCCESS_CODE}
//				case 1:
//					return &VenueResponse{Code: TransferConfirm_FAIL_CODE}
//				case 2:
//					return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: ""}
//				}
//			}
//
//		case 1, 2, 7, 8, 9, 10:
//			return &VenueResponse{Code: TransferConfirm_FAIL_CODE}
//		case 3:
//			return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: ""}
//		}
//	}
//
//	return &VenueResponse{Code: TransferConfirm_Processing_CODE, Msg: resp.String()}
//}
//
//func (v VenueSHABATY) PullOrder(request *VenuePullOrderRequest) *VenuePullOrderResponse {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (v VenueSHABATY) CallBackConfirm(request *VenueCallBackConfirmRequest) *VenueResponse {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (v VenueSHABATY) AmountLimitFix(amount float64, currency string) float64 {
//	//TODO implement me
//	return amount
//}
//
//func shaBaTyCurrency(currency string) string {
//	switch currency {
//	case "CNY":
//		return "124"
//	}
//	return "124"
//}

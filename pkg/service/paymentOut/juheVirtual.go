package paymentOut

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func init() {
	channelPayOut["Virtual"] = &JuHeVirtualPayOutGateway{}
}

// JuHeVirtualPayOutGateway
// @Description:
type JuHeVirtualPayOutGateway struct {
}

func (T JuHeVirtualPayOutGateway) PreWithdraw(param PreOtherPaymentOutParam) PrePaymentChannelOutResp {
	return PrePaymentChannelOutResp{
		Code: 200,
	}
}

func (T JuHeVirtualPayOutGateway) ChannelWithdraw(param OtherPaymentOutParam) PaymentChannelOutResp {

	conf := global.CONFIG.Payment.YinRunPay
	url := conf.APIURL + `/api/agentpay/apply`
	amount := param.Amount * 100
	amountStr := strconv.Itoa(int(amount))
	//extraStr := fmt.Sprintf("{\"uid\":\"%s\", \"oid\":\"%s\"}", param.UserId, param.OrderSn)
	body := map[string]string{
		"mchId":       fmt.Sprintf("%d", conf.MerchantCode),
		"mchOrderNo":  param.OrderSn,
		"amount":      amountStr,
		"accountName": param.WithdrawName,
		"accountNo":   param.VirtualAddress,
		"passageId":   param.PaymentCode,
		"reqTime":     time.Now().Format("20060102150405"),
		"extra":       param.UserId,
		"notifyUrl":   conf.VirtualNotifyURLOut,
		"remark":      fmt.Sprintf("付款%0.2f元", param.Amount),
	}
	if param.ThirdCode != "" {
		body["bankName"] = param.ThirdCode
	}
	signStr := tool.MapSortKeyAZString(body)
	signM := strings.ToUpper(tool.MD5([]byte(signStr + "&key=" + conf.Md5Key)))

	reqBodyStr := signStr + "&sign=" + signM
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(reqBodyStr).
		Post(url)

	global.G_LOG.Infof("JuHeVirtualPayOutGateway url:%v request:%v status:%v resp: %v err:%v", url, signStr, resp.StatusCode(), resp.String(), err)
	if err != nil {
		return PaymentChannelOutResp{}
	}

	status := gjson.Get(resp.String(), "status")
	//状态:0-待处理,1-处理中,2-成功,3-失败
	if status.Exists() && status.Int() == 0 || status.Exists() && status.Int() == 2 || status.Exists() && status.Int() == 1 {
		return PaymentChannelOutResp{ //已提交
			Code: 200,
		}
	}

	msg := gjson.Get(resp.String(), "transMsg").String()
	return PaymentChannelOutResp{
		Code:     500,
		ErrorMsg: msg,
	}
}

func (T JuHeVirtualPayOutGateway) ChannelCallBack(param CallBackPaymentOutParam) CallBackPaymentOutResp {
	conf := global.CONFIG.Payment.YinRunPay
	transMsg := gjson.Get(param.Raw, "transMsg").String()
	body := map[string]string{
		"agentpayOrderId": gjson.Get(param.Raw, "agentpayOrderId").String(),
		"amount":          gjson.Get(param.Raw, "amount").String(),
		"mchOrderNo":      gjson.Get(param.Raw, "mchOrderNo").String(),
		"status":          fmt.Sprintf("%d", gjson.Get(param.Raw, "status").Int()),
		"fee":             gjson.Get(param.Raw, "fee").String(),
		"extra":           gjson.Get(param.Raw, "extra").String(),
	}
	if transMsg != "" {
		body["transMsg"] = transMsg
	}
	userId := gjson.Get(param.Raw, "extra").String()
	sign := gjson.Get(param.Raw, "sign").String()
	signStr := tool.MapSortKeyAZString(body)
	signM := strings.ToUpper(tool.MD5([]byte(signStr + "&key=" + conf.Md5Key)))
	if sign != signM {
		global.G_LOG.Errorf("用户(%s)代付订单(%s)回调签名失败: %s", userId, body["mchOrderNo"], signStr)
		return CallBackPaymentOutResp{ReturnRaw: "FAIL"}
	}

	status := gjson.Get(param.Raw, "status").Int()
	if status == 2 {
		return CallBackPaymentOutResp{
			Code:      200,
			OrderSn:   body["mchOrderNo"],
			ReturnRaw: "SUCCESS",
		}
	}
	if status == 3 { //
		return CallBackPaymentOutResp{
			Code:      500,
			OrderSn:   body["orderCode"],
			ReturnRaw: "SUCCESS",
			ErrorMsg:  "",
		}
	}
	return CallBackPaymentOutResp{ReturnRaw: "FAIL"}
}

func (T JuHeVirtualPayOutGateway) FindOrderStatus(merchantOrderId string) FindOrderPaymentOutResp {
	conf := global.CONFIG.Payment.YinRunPay
	url := conf.APIURL + `/api/agentpay/query_order`

	body := map[string]string{
		"mchId":           conf.AppID,
		"agentpayOrderId": merchantOrderId,
		"reqTime":         time.Now().Format("20060102150405"),
	}

	signStr := tool.MapSortKeyAZString(body)
	signStr += "&key=" + conf.Md5Key
	signM := strings.ToUpper(tool.MD5([]byte(signStr)))
	body["sign"] = signM

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		SetFormData(body).
		Post(url)

	if err != nil {
		return FindOrderPaymentOutResp{
			Code: 0,
		}
	}
	status := gjson.Get(resp.String(), "return_code").String()
	if status == "SUCCESS" { //成功
		return FindOrderPaymentOutResp{
			Code: 200,
		}
	} else if status == "FAIL" { //失败
		return FindOrderPaymentOutResp{
			Code:     500,
			ErrorMsg: gjson.Get(resp.String(), "error_msg").String(),
		}
	} else if status == "PROCESSING" { //处理中
		return FindOrderPaymentOutResp{
			Code: 201,
		}
	}
	return FindOrderPaymentOutResp{
		Code: 0,
	}
}

func (T JuHeVirtualPayOutGateway) ChannelCallBackGetOrderSn(param CallBackPaymentOutParam) CallBackPaymentOutResp {
	/*extraData := gjson.Get(param.Raw, "extra").String() //系统订单号
	_, orderSn := GetOrderExtraData(extraData)
	if len(orderSn) == 0 {
		global.G_LOG.Errorf("[ChannelCallBackGetOrderSn] JuHeVirtualPayOutGateway Unmarshal extra failed: extra=%s", extraData)

		return CallBackPaymentOutResp{
			Code:      500,
			OrderSn:   "",
			ReturnRaw: "SUCCESS",
		}
	}

	return CallBackPaymentOutResp{
		Code:      200,
		OrderSn:   orderSn,
		ReturnRaw: "SUCCESS",
	}*/

	var reqData map[string]interface{}
	err := tool.JsonUnmarshalFromString(param.Raw, &reqData)
	if err != nil {
		global.G_LOG.Errorf("[ChannelCallBackGetOrderSn] JuHeVirtualPayOutGateway Unmarshal extra failed: raw=%s", param.Raw)

		return CallBackPaymentOutResp{
			Code:      500,
			OrderSn:   "",
			ReturnRaw: "SUCCESS",
		}
	}

	if v, ok := reqData["mchOrderNo"]; ok {
		return CallBackPaymentOutResp{
			Code:      200,
			OrderSn:   v.(string),
			ReturnRaw: "SUCCESS",
		}
	} else {
		global.G_LOG.Errorf("[ChannelCallBackGetOrderSn] JuHeVirtualPayOutGateway mchOrderNo null: raw=%s", param.Raw)
	}

	return CallBackPaymentOutResp{
		Code:      500,
		OrderSn:   "",
		ReturnRaw: "SUCCESS",
	}
}

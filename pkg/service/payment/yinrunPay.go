package payment

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/vo"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

var (
	ChannelProdIDMap = map[string]string{
		"bank":   "100010", // 100-20000
		"alipay": "100009", // 200-1000
		"wx":     "100008", // 100-20000
	}
)

const (
	Creat_Order_API = "/api/pay/create_order"
	Query_Order_API = "/api/pay/query_order"

	Recharge_CallBack = "/api/callback/pay/"
	YinRunPayType     = "yin_run"
)

func init() {
	channelPay[YinRunPayType] = &YinRunPay{}
}

type YinRunPay struct {
}

func (e YinRunPay) ChannelRecharge(payment *dos.FcPayment, param OtherPaymentParam) PaymentChannelResp {
	conf := global.CONFIG.Payment.YinRunPay
	url := conf.APIURL + Creat_Order_API

	if param.ReturnUrl == "" {
		param.ReturnUrl = conf.ReturnUrl
	}

	// 该通道单位为分
	amount := param.Amount * 100
	amountStr := strconv.Itoa(int(amount))
	conf.Currency = strings.ToLower(conf.Currency)
	subject := "paySubject"
	body := "payDescribe"

	//extra := YinRunExtra{}
	//extra.OpenId = param.UserID
	//extraBytes, _ := json.Marshal(extra)
	extraStr := param.UserInfo.UserId
	//prodID, ok := ChannelProdIDMap[payment.ChannelCode]
	//if !ok {
	//	err := fmt.Errorf("user=%s YinRunPay channelCode=%s not exist", param.UserInfo.UserName, payment.ChannelCode)
	//	return PaymentChannelResp{
	//		Success:  false,
	//		ErrorMsg: err.Error(),
	//	}
	//}
	prodID := payment.PayId
	notifyUrl := conf.ReturnUrl + Recharge_CallBack + YinRunPayType

	// 该参数会回调时返回给我们
	//param1 := map[string]interface{}{}
	//param1["amount"] = amount

	// 对数组进行从小到大排序, 空值不参与加密与排序
	bodyMap := map[string]string{
		"amount":     amountStr,
		"appId":      conf.AppID,
		"body":       body,
		"currency":   conf.Currency,
		"extra":      extraStr,
		"mchOrderNo": param.OrderSn,
		"mchId":      strconv.Itoa(conf.MerchantCode),
		"notifyUrl":  notifyUrl,
		//"param1":     "",
		"productId": prodID,
		"subject":   subject,
	}

	// map key 从小到大排序，value 为空不参与排序
	orgStr := tool.MapKeyLowToUpSortNotZero(bodyMap)
	orgStrKey := orgStr + "&key=" + conf.Md5Key
	sign := tool.MD5([]byte(orgStrKey))
	sign = strings.ToUpper(sign)
	reqBodyStr := orgStr + "&sign=" + sign
	//fmt.Println("reqBodyStr============> ", reqBodyStr)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(reqBodyStr).
		Post(url)

	if err != nil {
		errStr := fmt.Sprintf("YinRunPay username=%s  url=%s reqStr=%s err: %v", param.UserInfo.UserName, url, reqBodyStr, err)
		global.G_LOG.Errorf(errStr)
		return PaymentChannelResp{
			Success:  false,
			ErrorMsg: errStr,
		}
	}

	respStr := resp.String()
	reqRespStr := fmt.Sprintf("YinRunPay username: %s url: %v req: %v status: %v resp: %v", param.UserInfo.UserName, url, reqBodyStr, resp.StatusCode(), respStr)
	//global.G_LOG.Infof(reqRespStr)
	if gjson.Get(respStr, "retCode").String() != "SUCCESS" {
		global.G_LOG.Error(reqRespStr)
		return PaymentChannelResp{
			Success:  false,
			ErrorMsg: reqRespStr,
		}
	}

	return PaymentChannelResp{
		Success: true,
		Data: PaymentChannelDataResp{
			Style: STYLE_LINKE,
			Link:  gjson.Get(respStr, "payParams").Get("payUrl").String(),
		},
	}
}

func (e YinRunPay) ChannelCallBack(param CallBackPaymentParam) CallBackPaymentResp {
	global.G_LOG.Infof("YinRunPay ChannelCallBack param: %v merchantCode: %v", param.Raw, param.MerchantCode)

	conf := global.CONFIG.Payment.YinRunPay
	bodyMap := map[string]string{}
	bodyMap["payOrderId"] = gjson.Get(param.Raw, "payOrderId").String()
	bodyMap["mchId"] = gjson.Get(param.Raw, "mchId").String()
	bodyMap["appId"] = gjson.Get(param.Raw, "appId").String()
	bodyMap["productId"] = gjson.Get(param.Raw, "productId").String()
	bodyMap["mchOrderNo"] = gjson.Get(param.Raw, "mchOrderNo").String()
	bodyMap["amount"] = gjson.Get(param.Raw, "amount").String()
	bodyMap["income"] = gjson.Get(param.Raw, "income").String()
	bodyMap["extra"] = gjson.Get(param.Raw, "extra").String()
	bodyMap["status"] = gjson.Get(param.Raw, "status").String()
	bodyMap["channelOrderNo"] = gjson.Get(param.Raw, "channelOrderNo").String()
	bodyMap["channelAttach"] = gjson.Get(param.Raw, "channelAttach").String()
	bodyMap["param1"] = gjson.Get(param.Raw, "param1").String()
	bodyMap["param2"] = gjson.Get(param.Raw, "param2").String()
	bodyMap["paySuccTime"] = gjson.Get(param.Raw, "paySuccTime").String()
	bodyMap["backType"] = gjson.Get(param.Raw, "backType").String()

	// map key 从小到大排序，value 为空不参与排序
	orgStr := tool.MapKeyLowToUpSortNotZero(bodyMap)
	orgStrKey := orgStr + "&key=" + conf.Md5Key
	sign := tool.MD5([]byte(orgStrKey))
	sign = strings.ToUpper(sign)
	platSign := gjson.Get(param.Raw, "sign").String()
	if sign != platSign {
		tmpStr := fmt.Sprintf("YinRunPay ChannelCallBack mchOrderNo: %v sign: %v platSigin: %v not match", bodyMap["mchOrderNo"], sign, platSign)
		global.G_LOG.Errorf(tmpStr)
		return CallBackPaymentResp{ReturnRaw: "FAIL"}
	}

	statusStr := bodyMap["status"]
	if statusStr == "2" {
		global.G_LOG.Infof("YinRunPay ChannelCallBack mchOrderNo: %v success", bodyMap["mchOrderNo"])
		return CallBackPaymentResp{
			Code:      200,
			OrderSn:   bodyMap["mchOrderNo"],
			ReturnRaw: "SUCCESS",
		}
	}

	global.G_LOG.Infof("YinRunPay ChannelCallBack faild: %v", param.Raw)

	return CallBackPaymentResp{ReturnRaw: "FAIL"}
}

func (e YinRunPay) ChannelCallBackGetOrderSn(param CallBackPaymentParam) CallBackPaymentResp {
	conf := global.CONFIG.Payment.YinRunPay
	url := conf.APIURL + Query_Order_API
	merchantOrderId := gjson.Get(param.Raw, "mchOrderNo").String()
	payOrderId := gjson.Get(param.Raw, "payOrderId").String()
	conf.Currency = strings.ToLower(conf.Currency)
	orderID := merchantOrderId

	// 对数组进行从小到大排序, 空值不参与加密与排序
	bodyMap := map[string]string{
		"appId": conf.AppID,
		"mchId": strconv.Itoa(conf.MerchantCode),
	}

	// 商户订单号和支付订单号二传其一即可
	if merchantOrderId != "" {
		bodyMap["mchOrderNo"] = merchantOrderId
	} else {
		if payOrderId != "" {
			bodyMap["payOrderId"] = payOrderId
			orderID = payOrderId
		}
	}

	// map key 从小到大排序，value 为空不参与排序
	orgStr := tool.MapKeyLowToUpSortNotZero(bodyMap)
	orgStrKey := orgStr + "&key=" + conf.Md5Key
	sign := tool.MD5([]byte(orgStrKey))
	sign = strings.ToUpper(sign)
	reqBodyStr := orgStr + "&sign=" + sign

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(reqBodyStr).
		Post(url)

	if err != nil {
		errStr := fmt.Sprintf("YinRunPay_ChannelCallBackGetOrderSn url=%s reqStr=%s err: %v", url, reqBodyStr, err)
		global.G_LOG.Errorf(errStr)
		return CallBackPaymentResp{
			Code:      200,
			OrderSn:   orderID,
			ReturnRaw: "FAIL",
		}
	}

	respStr := resp.String()
	reqRespStr := fmt.Sprintf("YinRunPay_ChannelCallBackGetOrderSn url: %v req: %v status: %v resp: %v", url, reqBodyStr, resp.StatusCode(), respStr)
	global.G_LOG.Infof(reqRespStr)

	if gjson.Get(respStr, "retCode").String() != "SUCCESS" {
		global.G_LOG.Error(reqRespStr)
		return CallBackPaymentResp{
			Code:      200,
			OrderSn:   orderID,
			ReturnRaw: reqRespStr,
		}
	}

	return CallBackPaymentResp{
		Code:      200,
		OrderSn:   orderID,
		ReturnRaw: "success",
	}
}

func YinRunUpdateStatus(param vo.YRPayUpdateReq) (*vo.YRPayUpdateResp, error) {
	// map key 从小到大排序，value 为空不参与排序
	tmpMp := YRPayStructConvert(param)
	delete(tmpMp, "sign") // sgin字段不参与签名验证
	orgStr := tool.MapKeyLowToUpSortNotZero(tmpMp)
	orgStrKey := orgStr + "&key=" + global.CONFIG.Payment.YinRunPay.Md5Key
	sign := tool.MD5([]byte(orgStrKey))
	sign = strings.ToUpper(sign)

	if sign != param.Sign {
		global.G_LOG.Errorf("YinRunPay YinRunUpdateStatus signature not match: param=%v, productId=%v, sign=%v, platSigin=%v ", tmpMp, param.ProductId, sign, param.Sign)
		return &vo.YRPayUpdateResp{
			RetCode: "FAIL",
			RetMsg:  "sign not match",
		}, fmt.Errorf("sign not match")
	}

	return &vo.YRPayUpdateResp{
		RetCode: "SUCCESS",
		RetMsg:  "SUCCESS",
	}, nil
}

func YinRunQueryProductList(param vo.YRPayQueryProductListReq) (*vo.YRPayQueryProductListResp, error) {
	// map key 从小到大排序，value 为空不参与排序
	tmpMp := YRPayStructConvert(param)
	delete(tmpMp, "sign") // sgin字段不参与签名验证
	orgStr := tool.MapKeyLowToUpSortNotZero(tmpMp)
	orgStrKey := orgStr + "&key=" + global.CONFIG.Payment.YinRunPay.Md5Key
	sign := tool.MD5([]byte(orgStrKey))
	sign = strings.ToUpper(sign)

	if sign != param.Sign {
		global.G_LOG.Errorf("YinRunPay YinRunQueryProductList signature not match: param=%v, sign=%v, platSigin=%v ", tmpMp, sign, param.Sign)
		return &vo.YRPayQueryProductListResp{
			RetCode: "FAIL",
			RetMsg:  "sign not match",
		}, fmt.Errorf("sign not match")
	}

	return &vo.YRPayQueryProductListResp{
		RetCode: "SUCCESS",
		RetMsg:  "SUCCESS",
	}, nil
}

func YRPayStructConvert(obj interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		if fieldValue.CanInterface() {
			key := field.Tag.Get("json")
			if key == "" {
				key = field.Name
			}

			switch fieldValue.Kind() {
			case reflect.String:
				result[key] = fieldValue.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				result[key] = strconv.FormatInt(fieldValue.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				result[key] = strconv.FormatUint(fieldValue.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				result[key] = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
			case reflect.Bool:
				result[key] = strconv.FormatBool(fieldValue.Bool())
			default:
				result[key] = fmt.Sprintf("%v", fieldValue.Interface())
			}
		}
	}

	return result
}

// 货币汇率转换
func YRPayAmountFxCovert(action bool, amount float64) (float64, error) {
	if action {
		fxRateAdjustVal := 0.02 // 需要下调

		var result struct {
			Message   string  `json:"message"`
			RatePrice float64 `json:"ratePrice"`
			Status    string  `json:"status"`
		}

		_, err := resty.New().R().
			SetHeader("Accept", "application/json").
			SetResult(&result).
			Get("http://pay2.fulian.co/api/getRate?rateType=CNY")
		if err != nil {
			errStr := fmt.Sprintf("获取汇率失败: %v", err)
			global.G_LOG.Errorf(errStr)
			return amount, errors.New(errStr)
		}

		return (amount * (result.RatePrice - fxRateAdjustVal)), nil
	}

	return amount, nil
}

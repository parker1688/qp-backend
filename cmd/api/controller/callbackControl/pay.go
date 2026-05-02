package callbackControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/payment"
	"bootpkg/pkg/service/paymentOut"
	"bootpkg/pkg/service/userTransfer"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

func PaymentCallBack(c *gin.Context) {
	//global.G_LOG.Infof("PaymentCallBackIn start.........")
	paymentType := c.Param("paymentType")

	reqContentType := c.Request.Header.Get("Content-Type")
	isFormData := strings.Contains(reqContentType, "/form-data")
	isFormUrl := strings.Contains(reqContentType, "/x-www-form-urlencoded")
	var payload []byte
	var err error
	if isFormData || isFormUrl {
		_ = c.Request.ParseMultipartForm(32 << 20)
		postMap := make(map[string]interface{}, len(c.Request.Form))
		for k, v := range c.Request.Form {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
		payload = []byte(tool.String(postMap))
	} else {
		payload, err = c.GetRawData()
		if err != nil {
			c.String(200, "success")
			return
		}
	}

	//global.G_LOG.Infof("[PaymentCallBackIn] %v - raw: %v", c.Request.Host+c.FullPath()+paymentType, string(payload))
	m := payment.GetRechargeChannel(paymentType)
	if m == nil {
		c.String(200, "success")
		return
	}

	// 去三方查询订单信息
	pre := m.ChannelCallBackGetOrderSn(payment.CallBackPaymentParam{Raw: string(payload)})

	orderInfo := modules.FindByKeyFcOrderDepositFirst(&dos.FcOrderDeposit{OrderSn: pre.OrderSn})
	if orderInfo == nil || orderInfo.Id == "" {
		global.G_LOG.Errorf("PaymentCallBackIn orderSn: %v query err: %v", pre.OrderSn, err)
		c.String(200, "success")
		return
	}

	validator := payment.NewPaymentCallbackValidator()
	if err = validator.ComprehensiveCallbackValidation(
		orderInfo,
		paymentType,
		orderInfo.Amount,
		orderInfo.MerchantCode,
		c.ClientIP(),
		0.01,
	); err != nil {
		global.G_LOG.Warnf("[PaymentCallBack] validation failed: orderSn=%s, paymentType=%s, ip=%s, err=%v", pre.OrderSn, paymentType, c.ClientIP(), err)
		c.String(200, "success")
		return
	}

	res := m.ChannelCallBack(payment.CallBackPaymentParam{Raw: string(payload), MerchantCode: orderInfo.MerchantCode})
	//global.G_LOG.Infof("[PaymentCallBackInOut] %v - raw: %v", c.FullPath(), tool.String(res))
	if res.Code == 200 && orderInfo != nil { //成功
		if orderInfo.Status != enmus.ORDER_YES_STATUS {
			result, err := userTransfer.CallbackDepositSuccess(&dos.FcOrderDeposit{
				OrderSn: res.OrderSn,
				Status:  enmus.ORDER_YES_STATUS,
				BaseDos: dos.BaseDos{Id: orderInfo.Id},
			})
			//global.G_LOG.Infof("[PaymentCallBackUpStatus] orderSn:%s %v - err: %v", res.OrderSn, result, err)
			if err != nil {
				global.G_LOG.Errorf("[PaymentCallBack.CallbackDepositSuccess] Failed: orderSn=%s, err=%v", res.OrderSn, err.Error())
				c.String(200, "false")
				return
			}
			if !result {
				c.String(200, "false")
				return
			}

			// ============= 充值任务（累计充值）以下
			go modules.DoUserTaskAction(orderInfo.UserId, []modules.TaskActionParam{
				{
					Type:        enmus.DailyTaskType_Recharge,
					Subtype:     enmus.DailyTaskSubType_Pay,
					Amount:      orderInfo.Amount,
					ChannelCode: orderInfo.ChannelCode,
				},
			}, false)
			// ============= 充值任务（累计充值）以上
			//充值成功系统邮件提示
			go modules.SendSystemMail(orderInfo.UserId, orderInfo.MerchantCode, enmus.MailType_RechargeSuccess, orderInfo.Amount)
			// ============= 充值限单标识处理以下
			modules.DoUserPaymentStrategyResetAction(
				orderInfo.UserId,
				orderInfo.ChannelCode,
				orderInfo.MerchantCode,
			)
			// ============= 充值限单标识处理以上
		}
	}
	if !gjson.Valid(res.ReturnRaw) {
		c.String(200, res.ReturnRaw)
		return
	}
	c.Data(200, "application/json", []byte(res.ReturnRaw))
}

func PaymentCallBackOut(c *gin.Context) {
	paymentType := c.Param("paymentType")

	reqContentType := c.Request.Header.Get("Content-Type")
	isFormData := strings.Contains(reqContentType, "/form-data")
	isFormUrl := strings.Contains(reqContentType, "/x-www-form-urlencoded")
	var payload []byte
	var err error
	if isFormData || isFormUrl {
		_ = c.Request.ParseMultipartForm(32 << 20)
		postMap := make(map[string]interface{}, len(c.Request.Form))
		for k, v := range c.Request.Form {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
		payload = []byte(tool.String(postMap))
	} else {
		payload, err = c.GetRawData()
		if err != nil {
			c.String(200, "success")
			return
		}
	}

	//global.G_LOG.Infof("[PaymentCallBackOut] %v - raw: %v", c.Request.Host+c.FullPath()+paymentType, string(payload))
	m := paymentOut.GetPaymentOutChannel(paymentType)
	if m == nil {
		c.String(200, "success")
		return
	}
	pre := m.ChannelCallBackGetOrderSn(paymentOut.CallBackPaymentOutParam{Raw: string(payload)})
	if pre.Code != 200 {
		c.String(200, "success")
		return
	}
	order := modules.FindByKeyFcOrderWithdrawPaymentOutFirst(&dos.FcOrderWithdrawPaymentOut{
		//BaseDos: dos.BaseDos{Id: pre.OrderSn}
		OrderSn: pre.OrderSn,
	})
	res := m.ChannelCallBack(paymentOut.CallBackPaymentOutParam{Raw: string(payload), MerchantCode: order.MerchantCode})
	//global.G_LOG.Infof("[PaymentCallBackOut] %v - raw: %v", c.FullPath(), tool.String(res))
	if (res.Code == 200 || res.Code == 500) && order != nil {
		//处理业务
		if res.Code == 200 {
			//var ok bool

			var reqData map[string]interface{}
			err := tool.JsonUnmarshalFromString(string(payload), &reqData)
			if err != nil {
				global.G_LOG.Errorf("[PaymentCallBackOut] unmarshal json failed: %s, err=%v", string(payload), err.Error())
				ok := paymentOut.OrderWithdrawAnotherPayFailRemark(order.Id, res.ErrorMsg, enmus.OrderWithdrawPaymentOutStats_Progress, "system") //已经提交的状态
				if !ok {
					c.String(200, "false")
					//提款失败系统邮件提示
					go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
					return
				}
			}
			if v, ok := reqData["status"]; ok {
				switch v {
				case "2":
					// 真正支付成功
					ok, realOk := paymentOut.OrderWithdrawAnotherPaySuccess(order.Id, "system")

					if !ok {
						c.String(200, "false")
						//提款失败系统邮件提示
						go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
						return
					}
					if realOk {
						channelData.SendUserWithdrawal(&channelData.UserWithdrawalMessage{
							UserId:           order.UserId,
							UserName:         order.UserName,
							OrderSn:          order.OrderSn,
							WithdrawalAmount: order.Amount,
						})
						//
						//提款成功系统邮件提示
						go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawSuccess, order.Amount)
					}
				case "3":
					// 真正支付失败
					ok := paymentOut.OrderWithdrawAnotherPayFailRemark(order.Id, res.ErrorMsg, enmus.OrderWithdrawPaymentOutStats_Progress, "system") //已经提交的状态
					//提款失败系统邮件提示
					go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
					if !ok {
						c.String(200, "false")
						//提款失败系统邮件提示
						go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
						return
					}
					//提款失败系统邮件提示
					go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
				default:
					c.String(200, "false")
				}
			}
		} else if res.Code == 500 {
			//var ok bool
			ok := paymentOut.OrderWithdrawAnotherPayFailRemark(order.Id, res.ErrorMsg, enmus.OrderWithdrawPaymentOutStats_Progress, "system") //已经提交的状态

			if !ok {
				c.String(200, "false")
				//提款失败系统邮件提示
				go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
				return
			}
		}
	}

	if !gjson.Valid(res.ReturnRaw) {
		c.String(200, res.ReturnRaw)
		//提款失败系统邮件提示
		go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
		return
	}
	c.Data(200, "application/json", []byte(res.ReturnRaw))
}

// yinrun 支付回调更改通道状态
func YinRunPayUpdateStatusCallback(c *gin.Context) {
	var jsonp vo.YRPayUpdateReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		global.G_LOG.Errorf("YinRunPayUpdateStatusCallback err: %v", err.Error())
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// global.G_LOG.Infof("---- YinRunPayUpdateStatusCallback ---- req=%+v", jsonp)

	// 加密签名校验f
	resp := &vo.YRPayUpdateResp{}
	resp, err = payment.YinRunUpdateStatus(jsonp)
	if err != nil {
		global.G_LOG.Errorf("YinRunPayUpdateStatusCallback err: %v", err)
		c.JSON(200, resp)
		return
	}

	payStatus := 1         // 默认正常
	if jsonp.Status != 1 { // 三方状态 1: 开启，0: 关闭
		payStatus = 2
	}
	updatePayStatusMap := map[string]interface{}{}
	updatePayStatusMap["status"] = payStatus
	updatePayStatusMap["update_by"] = "system"
	updatePayStatusMap["pay_alias_name"] = jsonp.Title
	updatePayStatusMap["sort"] = jsonp.OrderBy
	updatePayStatusMap["amount_range"] = jsonp.AmountList
	newMinAmount, newMaxAmount := func() (float64, float64) {
		amountLis := strings.Split(jsonp.AmountList, ",")
		if len(amountLis) == 0 {
			global.G_LOG.Warnf("[YinRunPayUpdateStatusCallback] amount list is empty: %s",
				jsonp.AmountList)
			return 0, 0
		}

		minAmount := tool.StringToFloat64(amountLis[0])
		maxAmount := tool.StringToFloat64(amountLis[len(amountLis)-1])

		if minAmount == maxAmount {
			global.G_LOG.Warnf("[YinRunPayUpdateStatusCallback] Amount min is the same as max: %s",
				jsonp.AmountList)
		}

		return minAmount, maxAmount
	}()
	updatePayStatusMap["min_amount"] = newMinAmount
	updatePayStatusMap["max_amount"] = newMaxAmount
	newMinLevel, newMaxLevel := func() (int, int) {
		mapRangeFn := func(v int, isMin bool) int {
			rangeMp := map[int][]int{
				0:  {1, 1},
				1:  {2, 3},
				2:  {4, 6},
				3:  {7, 9},
				4:  {10, 12},
				5:  {13, 15},
				6:  {16, 18},
				7:  {19, 21},
				8:  {22, 24},
				9:  {25, 27},
				10: {28, 30},
			}

			if isMin {
				return rangeMp[v][0]
			} else {
				return rangeMp[v][1]
			}
		}

		vipLvLis := strings.Split(jsonp.VipLevel, ",")
		if len(vipLvLis) == 0 {
			global.G_LOG.Warnf("[YinRunPayUpdateStatusCallback] vip level is empty: %s",
				jsonp.VipLevel)
			return 1, 1
		}

		minLv := mapRangeFn(tool.Atoi(vipLvLis[0]), true)
		maxLv := mapRangeFn(tool.Atoi(vipLvLis[len(vipLvLis)-1]), false)

		if minLv == maxLv {
			global.G_LOG.Warnf("[YinRunPayUpdateStatusCallback] Vip level min is the same as max: %s",
				jsonp.VipLevel)
		}

		return minLv, maxLv
	}()
	updatePayStatusMap["min_level"] = newMinLevel
	updatePayStatusMap["max_level"] = newMaxLevel
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 查询通道总表
	paySumRow := dos.FcPaymentSum{}
	err = global.G_DB.WithContext(ctx).Model(&dos.FcPaymentSum{}).Where("payment_code = ? AND pay_id = ?", payment.YinRunPayType, jsonp.ProductId).First(&paySumRow).Error
	if err != nil {
		// 如果不存在则直接返回成功
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.G_LOG.Infof("YinRunPayUpdateStatusCallback not exist: payment_code=%v", jsonp.ProductId)
			c.JSON(200, resp)
			return
		}
	} else {
		//if paySumRow.Status != payStatus {
		// 更新总通道和商户通道表
		err = global.G_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			err = tx.Model(&dos.FcPaymentSum{}).
				Where("payment_code = ? AND pay_id = ?", payment.YinRunPayType, jsonp.ProductId).
				Updates(updatePayStatusMap).Error
			if err != nil {
				global.G_LOG.Errorf("YinRunPayUpdateStatusCallback FcPaymentSum update failed: req=%+v, status=%v, err=%v",
					jsonp, payStatus, err.Error())
				return err
			}

			err = tx.Model(&dos.FcPayment{}).
				Where("payment_code = ? AND pay_id = ?", payment.YinRunPayType, jsonp.ProductId).
				Updates(updatePayStatusMap).Error
			if err != nil {
				global.G_LOG.Errorf("YinRunPayUpdateStatusCallback FcPayment update failed: req=%+v, status=%v, err=%v",
					jsonp, payStatus, err.Error())
				return err
			}

			return nil
		})

		//}

	}

	c.JSON(200, resp)
}

// 支付获取通道列表
func YinRunPayQueryProductListCallback(c *gin.Context) {
	var jsonp vo.YRPayQueryProductListReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		global.G_LOG.Errorf("YinRunPayQueryProductListCallback err: %v", err.Error())
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	//global.G_LOG.Info("---- YinRunPayQueryProductListCallback ----", jsonp)

	// 加密签名校验f
	resp := &vo.YRPayQueryProductListResp{}
	resp, err = payment.YinRunQueryProductList(jsonp)
	if err != nil {
		global.G_LOG.Errorf("YinRunPayQueryProductListCallback err: %v", err)
		c.JSON(200, resp)
		return
	}

	var payments []*dos.FcPaymentSum
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	err = global.G_DB.WithContext(ctx).Model(&dos.FcPaymentSum{}).Find(&payments).Error
	if err != nil {
		global.G_LOG.Errorf("YinRunPayQueryProductListCallback query payments err: %v", err)
		resp.Data = "{}"
		c.JSON(200, resp)
		return
	}
	dataMap := map[string]int{}
	for _, v := range payments {
		if v.Status == 2 {
			dataMap[string(v.PayId)] = 0
		} else {
			dataMap[string(v.PayId)] = 1
		}
	}

	dataStr, err := tool.JsonMarshalString(dataMap)
	if err != nil {
		global.G_LOG.Errorf("YinRunPayQueryProductListCallback data string err: %v", err)
		dataStr = "{}"
	}
	resp.Data = dataStr
	c.JSON(200, resp)
}

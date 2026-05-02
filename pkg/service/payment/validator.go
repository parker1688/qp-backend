package payment

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"net"
	"strings"
)

// PaymentCallbackValidator 支付回调验证工具
type PaymentCallbackValidator struct {
	// 支持的支付渠道 IP 白名单 (应从配置文件读取)
	AllowedIPs map[string][]string // 按渠道存储允许的 IP
}

// NewPaymentCallbackValidator 创建验证器实例
func NewPaymentCallbackValidator() *PaymentCallbackValidator {
	allowedIPs := map[string][]string{}
	if ips := tool.GetPaymentCallbackIPs(); len(ips) > 0 {
		allowedIPs["*"] = ips
	}

	return &PaymentCallbackValidator{
		AllowedIPs: map[string][]string{
			"*": allowedIPs["*"],
			// TODO: 从配置文件读取各支付渠道的回调 IP 白名单
			// 示例：
			// "alipay": {"123.123.123.0/24", "124.124.124.0/24"},
			// "wechat": {"220.248.137.0/24"},
		},
	}
}

// VerifyCallbackIP 验证回调来源 IP 是否在白名单中
func (v *PaymentCallbackValidator) VerifyCallbackIP(channelCode string, clientIP string) bool {
	allowedIPs, exists := v.AllowedIPs[channelCode]
	if (!exists || len(allowedIPs) == 0) && len(v.AllowedIPs["*"]) > 0 {
		allowedIPs = v.AllowedIPs["*"]
		exists = true
	}
	if !exists || len(allowedIPs) == 0 {
		// 如果未配置白名单，则跳过 IP 检查（警告）
		global.G_LOG.Warnf("[PaymentCallback] No IP whitelist configured for channel: %s", channelCode)
		return true
	}
	
	clientIPObj := net.ParseIP(clientIP)
	if clientIPObj == nil {
		global.G_LOG.Errorf("[PaymentCallback] Invalid client IP: %s", clientIP)
		return false
	}
	
	for _, allowedIP := range allowedIPs {
		// 支持 CIDR 表示法或单个 IP
		if strings.Contains(allowedIP, "/") {
			_, cidr, err := net.ParseCIDR(allowedIP)
			if err != nil {
				global.G_LOG.Errorf("[PaymentCallback] Invalid CIDR: %s, err: %v", allowedIP, err)
				continue
			}
			if cidr.Contains(clientIPObj) {
				return true
			}
		} else {
			if clientIP == allowedIP {
				return true
			}
		}
	}
	
	global.G_LOG.Warnf("[PaymentCallback] Client IP %s not in whitelist for channel: %s", clientIP, channelCode)
	return false
}

// VerifyCallbackAmount 验证回调金额与订单金额是否一致
func (v *PaymentCallbackValidator) VerifyCallbackAmount(orderInfo *dos.FcOrderDeposit, callbackAmount float64, tolerance float64) bool {
	if orderInfo == nil {
		global.G_LOG.Errorf("[PaymentCallback] Order info is nil")
		return false
	}
	
	// 允许微小差异（如汇率变化，单位: 分）
	if tolerance == 0 {
		tolerance = 0.01 // 默认允许 0.01 元的差异
	}
	
	amountDiff := orderInfo.Amount - callbackAmount
	if amountDiff < -tolerance || amountDiff > tolerance {
		global.G_LOG.Errorf("[PaymentCallback] Amount mismatch: order=%f, callback=%f, diff=%f, tolerance=%f",
			orderInfo.Amount, callbackAmount, amountDiff, tolerance)
		return false
	}
	
	return true
}

// VerifyCallbackMerchantCode 验证商户代码
func (v *PaymentCallbackValidator) VerifyCallbackMerchantCode(orderInfo *dos.FcOrderDeposit, callbackMerchantCode string) bool {
	if orderInfo == nil || orderInfo.MerchantCode == "" {
		global.G_LOG.Errorf("[PaymentCallback] Order info or merchant code missing")
		return false
	}
	
	if orderInfo.MerchantCode != callbackMerchantCode {
		global.G_LOG.Errorf("[PaymentCallback] Merchant code mismatch: order=%s, callback=%s",
			orderInfo.MerchantCode, callbackMerchantCode)
		return false
	}
	
	return true
}

// VerifyDuplicateCallback 检查是否重复回调（订单已处理过）
func (v *PaymentCallbackValidator) VerifyDuplicateCallback(orderInfo *dos.FcOrderDeposit) bool {
	if orderInfo == nil {
		return false
	}
	
	// 如果订单已处理，则拒绝再次处理
	// 状态定义应在 enmus 中查看
	// 假设 0=未处理, 1=处理中, 2=已成功, 3=已失败
	if orderInfo.Status != 0 && orderInfo.Status != 1 {
		global.G_LOG.Warnf("[PaymentCallback] Duplicate callback detected for order: %s, current status: %d",
			orderInfo.OrderSn, orderInfo.Status)
		return false // 拒绝重复处理
	}
	
	return true
}

// ComprehensiveCallbackValidation 综合验证回调（一次性执行所有检查）
func (v *PaymentCallbackValidator) ComprehensiveCallbackValidation(
	orderInfo *dos.FcOrderDeposit,
	channelCode string,
	callbackAmount float64,
	callbackMerchantCode string,
	clientIP string,
	tolerance float64,
) error {
	if orderInfo == nil {
		return fmt.Errorf("order info is nil")
	}
	
	// 1. 检查重复回调
	if !v.VerifyDuplicateCallback(orderInfo) {
		return fmt.Errorf("duplicate callback or order already processed")
	}
	
	// 2. 检查商户代码
	if !v.VerifyCallbackMerchantCode(orderInfo, callbackMerchantCode) {
		return fmt.Errorf("merchant code mismatch")
	}
	
	// 3. 检查金额
	if !v.VerifyCallbackAmount(orderInfo, callbackAmount, tolerance) {
		return fmt.Errorf("amount mismatch")
	}
	
	// 4. 检查回调 IP
	if !v.VerifyCallbackIP(channelCode, clientIP) {
		return fmt.Errorf("callback IP not in whitelist")
	}
	
	return nil
}

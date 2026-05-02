package crontab

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/service/report"
	"time"
)

func init() {
	cronFunc = append(cronFunc, cronTabEvery{
		spec: "@every 20m", //10分钟一次
		cmd:  YesterdayComplexReport,
	})

	cronFunc = append(cronFunc, cronTabEvery{
		spec: "@every 10m", //10分钟一次
		cmd:  TodayComplexReport,
	})
}

func YesterdayComplexReport() {
	startAt, endAt := tool.DayStartEndDate(time.Now().AddDate(0, 0, -1))

	for _, m := range modules.FindByKeyFcMerchant(&dos.FcMerchant{Status: 1}, nil) {
		c := &report.ComplexReport{
			StartAt:      automaticType.Time(*tool.DateToTime(tool.TimeLayout, startAt)),
			EndAt:        automaticType.Time(*tool.DateToTime(tool.TimeLayout, endAt)),
			MerchantCode: m.MerchantCode,
			MerchantName: m.MerchantName,
		}
		ComplexReportDo(c)
	}

}

func TodayComplexReport() {
	startAt, endAt := tool.DayStartEndDate(time.Now())

	for _, m := range modules.FindByKeyFcMerchant(&dos.FcMerchant{Status: 1}, nil) {
		c := &report.ComplexReport{
			StartAt:      automaticType.Time(*tool.DateToTime(tool.TimeLayout, startAt)),
			EndAt:        automaticType.Time(*tool.DateToTime(tool.TimeLayout, endAt)),
			MerchantCode: m.MerchantCode,
			MerchantName: m.MerchantName,
		}
		ComplexReportDo(c)
	}
}

func ComplexReportDo(r *report.ComplexReport) {
	data := &dos.FcComplexReport{
		Day:                       r.StartAt.Timer().Format(tool.TimeDateLayout),
		MerchantCode:              r.MerchantCode,
		MerchantName:              r.MerchantName,
		BetWin:                    r.BetWin(),
		BetAmount:                 r.BetAmount(),
		RebateAmount:              r.RebateAmount(),
		RegisterNum:               r.RegisterNum(),
		FirstDepositNum:           r.FirstDepositNum(),
		DepositNum:                r.DepositNum(),
		DepositCount:              r.DepositCount(),
		FirstDepositAmount:        r.FirstDepositAmount(),
		NewUserDepositAmount:      r.NewUserDepositAmount(),
		NewUserDepositCount:       r.NewUserDepositCount(),
		NewUserFisrtDepositAmount: r.NewUserFirstDepositAmount(),
		LoginNum:                  r.LoginNum(),
		WithdrawNum:               r.WithdrawNum(),
		BetNum:                    r.BetNum(),
		PromotionAmount:           r.PromotionAmount(),
		AlipayDepositAmount:       r.AlipayDepositAmount(),
		WxDepositAmount:           r.WxDepositAmount(),
		BankDepositAmount:         r.BankDepositAmount(),
		WalletDepositAmount:       r.WalletDepositAmount(),
		NumCnyDepositAmount:       r.NumCnyDepositAmount(),
		UsdtDepositAmount:         r.UsdtDepositAmount(),
		AdminDepositAmount:        r.AdminDepositAmount(),
		TotalDepositAmount:        r.TotalDepositAmount(),
		AlipayWithdrawAmount:      r.AlipayWithdrawAmount(),
		BankWithdrawAmount:        r.BankWithdrawAmount(),
		WalletWithdrawAmount:      r.WalletWithdrawAmount(),
		UsdtWithdrawAmount:        r.UsdtWithdrawAmount(),
		TotalWithdrawAmount:       r.TotalWithdrawAmount(),
	}

	data.GameKillRate = r.GameKillRate(data)
	data.BetMultiple = r.BetMultiple(data)
	data.RegisterDepositRate = r.RegisterDepositRate(data)
	data.DepositWithdrawSubAmount = r.DepositWithdrawSubAmount(data)
	data.KillRate = r.KillRate(data)

	report := modules.FindByKeyFcComplexReportFirst(&dos.FcComplexReport{
		Day:          data.Day,
		MerchantCode: data.MerchantCode,
	})

	if len(report.Id) > 0 {
		data.Id = report.Id
		modules.UpdateFcComplexReport(data)
	} else {
		modules.SaveFcComplexReport(data)
	}
}

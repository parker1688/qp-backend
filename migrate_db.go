//go:build tools
// +build tools

package main

import (
	"bootpkg/common/conf"
	"bootpkg/common/database/mysqldb"
	"bootpkg/common/global"
	"bootpkg/common/logs"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"log"
	"strings"
)

func main() {
	// 初始化配置
	global.CONFIG = conf.NewOnceConfig("./conf.yaml")
	global.G_LOG = logs.NewLog(global.CONFIG.Log.Path)

	// 初始化数据库
	mysqldb.NewOnceMySql()

	if global.G_DB == nil {
		log.Fatal("数据库连接失败")
	}

	fmt.Println("数据库连接成功，开始建表...")

	// 执行迁移 - 只创建核心表
	err := global.G_DB.AutoMigrate(
		&dos.FcGlobal{},
		&dos.Department{},
		&dos.Dicts{},
		&dos.UserGroup{},
		&dos.WhiteIp{},
		&dos.FcMerchant{},
		&dos.FcMerchantLink{},
		&dos.FcAgentGroup{},
		&dos.FcMerchantVenue{},
		&dos.FcPayChannelSum{},
		&dos.FcPayChannelOut{},
		&dos.FcPaymentSum{},
		&dos.FcPaymentOut{},
		&dos.FcUserMaterial{},
		&dos.FcUserLogin{},
		&dos.FcUserWallet{},
		&dos.FcUserReport{},
		&dos.FcUserRebateRecords{},
		&dos.FcUserVipRecord{},
		&dos.FcUserNotify{},
		&dos.FcVenue{},
		&dos.FcVenueImg{},
		&dos.FcVenueUser{},
		&dos.FcVenueTransfer{},
		&dos.FcVenueMaintain{},
		&dos.FcChannelBankImg{},
		&dos.FcOrderDeposit{},
		&dos.FcOrderWithdraw{},
		&dos.FcOrderManageOpt{},
		&dos.FcOrderPromotion{},
		&dos.FcCustomerOrder{},
		&dos.FcWelfareManage{},
		&dos.FcTranscation{},
		&dos.FcComplexReport{},
		&dos.FcPromotionInfo{},
		&dos.FcVip{},
		&dos.FcVipWeekGift{},
		&dos.FcVipMonthGift{},
		&dos.FcPayChannel{},
		&dos.FcPayment{},
		&dos.FcVirtualCurrency{},
		&dos.FcVirtualCurrencyDetails{},
		&dos.FcVirtualCurrencyFx{},
		&dos.FcUserWithdrawBankBind{},
		&dos.FcVenueGame{},
		&dos.FcBetRecord{},
		&dos.FcSiteBanner{},
		&dos.FcSiteNotify{},
		&dos.FcSiteNotifyMarquee{},
		&dos.FcBulletin{},
		&dos.FcAgent{},
		&dos.DictsDetail{},
		&dos.FcLoginLog{},
		&dos.LoginLog{},
		&dos.FcClientLog{},
		&dos.DailyTask{},
		&dos.DailyBonus{},
		&dos.Blacklist{},
		&dos.AdsCarousel{},
		&dos.SplashScreen{},
		&dos.MailTemplate{},
	)

	if err != nil {
		log.Fatalf("建表失败: %v", err)
	}

	// 历史库中 FcGameRebate 可能存在 longtext 索引迁移差异，单独迁移避免阻断其他表。
	if err = global.G_DB.AutoMigrate(&dos.FcGameRebate{}); err != nil {
		if strings.Contains(err.Error(), "BLOB/TEXT column 'game_type' used in key specification without a key length") {
			log.Printf("警告: FcGameRebate 迁移跳过，保留现有结构: %v", err)
		} else {
			log.Fatalf("FcGameRebate 建表失败: %v", err)
		}
	}

	// 初始化全局自增键，避免依赖接口因缺少基础数据报错
	initGlobals := []dos.FcGlobal{
		{Key: "USER_ID_INCR", Value: "100000"},
		{Key: "INVITE_CODE_INCR", Value: "100000"},
	}
	for _, item := range initGlobals {
		var count int64
		global.G_DB.Model(&dos.FcGlobal{}).Where("`key` = ?", item.Key).Count(&count)
		if count == 0 {
			if createErr := global.G_DB.Create(&item).Error; createErr != nil {
				log.Fatalf("初始化全局键失败 key=%s err=%v", item.Key, createErr)
			}
		}
	}

	fmt.Println("建表完成！")
}

package versionControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func VersionInfo(c *gin.Context) {
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	//global.G_LOG.Infof("version_info ----------------0:%v", merchantCode)

	versionCode := global.G_REDIS.Get(context.Background(), fmt.Sprintf("client_version:%s", merchantCode)).Val()
	cdnUrl := global.G_REDIS.Get(context.Background(), fmt.Sprintf("client_version_cdn:%s", merchantCode)).Val()
	//global.G_LOG.Infof("version_info ----------------1:%v, %v", versionCode, cdnUrl)
	if versionCode == "" || cdnUrl == "" {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		query := global.G_DB.WithContext(ctx).Model(&dos.DictsDetail{}).Where("dicts_type_code = ? and dicts_tag =? ", "merchant_code", merchantCode)
		query.Select("dicts_value").Scan(&versionCode)
		query.Select("remarks").Scan(&cdnUrl)
		global.G_REDIS.Set(context.Background(), fmt.Sprintf("client_version:%s", merchantCode), versionCode, 10*time.Minute)
		global.G_REDIS.Set(context.Background(), fmt.Sprintf("client_version_cdn:%s", merchantCode), cdnUrl, 10*time.Minute)
	}
	//global.G_LOG.Infof("version_info ----------------2:%v, %v", versionCode, cdnUrl)
	data := vo.VersionInfoResp{Version: versionCode, CdnUrl: cdnUrl}
	response.SuccessJSON(c, data)
}

// 登录判断是否维护状态，是否在白名单
func LoginCheck(userName string) bool {
	//global.G_LOG.Infof("white list ----------------0:%v", userName)
	SeverStatus := global.G_REDIS.Get(context.Background(), "server_status").Val()
	//global.G_LOG.Infof("white list ----------------1:%v", SeverStatus)
	if SeverStatus == "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		query := global.G_DB.WithContext(ctx).Model(&dos.DictsDetail{}).Where("dicts_type_code = ?", "Service_Status")
		status := 0
		query.Select("dicts_value").Scan(&status)
		//global.G_LOG.Infof("white list ----------------2:%v", status)
		global.G_REDIS.Set(context.Background(), "server_status", status, time.Minute)
		SeverStatus = utils.ToString(status)
	}
	//global.G_LOG.Infof("white list ----------------3:%v", SeverStatus)
	if SeverStatus == "1" {
		whiteList := global.G_REDIS.Get(context.Background(), "server_white_list").Val()
		if whiteList == "" {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			query1 := global.G_DB.WithContext(ctx).Model(&dos.DictsDetail{}).Where("dicts_type_code = ?", "Service_Login_White")
			list := ""
			query1.Select("remarks").Scan(&list)
			//global.G_LOG.Infof("white list ----------------3:%v", list)
			global.G_REDIS.Set(context.Background(), "server_white_list", list, 3*time.Minute)
			whiteList = list
		}
		//global.G_LOG.Infof("white list ----------------4:%v, %v", whiteList, userName)
		return containsName(whiteList, userName)
	}
	return true
}

func containsName(namesStr string, target string) bool {
	names := strings.Split(namesStr, ",")
	for _, name := range names {
		//global.G_LOG.Infof("white list ----------------3:%v, %v", name, target)
		if strings.TrimSpace(name) == target {
			return true
		}
	}
	return false
}

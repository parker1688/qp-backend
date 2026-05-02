package router

import (
	"bootpkg/cmd/api/controller/userControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, UserInfoRouter)
}
func UserInfoRouter() {
	r := routers.Group("/api/userinfo/").Use(handler.AuthApiMiddleware())

	r.GET("/material", userControl.Material)                                //用户基本信息
	r.POST("/logout", userControl.Logout)                                   // 用户退出
	r.POST("/material/update", userControl.MaterialUpdate)                  //修改用户基本资料
	routers.POST("/api/userinfo/phone/veryCode", userControl.PhoneVeryCode) //获取手机验证码，触发（非token接口）
	r.POST("/phone/veryCodeSub", userControl.Verification)                  //获取手机验证码，提交验证码
	r.POST("/update/avatar", userControl.MaterialUpdateAvatar)              //修改用户头像
	r.POST("/update/sex", userControl.MaterialUpdateSex)                    //修改用户性别
	r.POST("/update/birthday", userControl.MaterialUpdateBirthday)          //修改用户生日
	r.POST("/update/nickname", userControl.MaterialUpdateNickName)          //修改用户昵称
	r.POST("/update/email", userControl.MaterialUpdateEmail)                //更新邮箱
	r.POST("/update/phone", userControl.MaterialUpdatePhone)                //更新手机号
	r.GET("/report", userControl.UserReport)                                //用户报表

	r.POST("/password/update", userControl.PasswordUpdate) //修改用户密码

	vipr := routers.Group("/api/vip/").Use(handler.AuthApiMiddleware())

	vipr.GET("/level/rebate", userControl.LevelrebateList)  //VIP返水配置列表
	vipr.GET("/level/vipProgress", userControl.VipProgress) //VIP返水配置列表

	viprS := routers.Group("/api/vip/level").Use(handler.AuthApiMiddleware())
	viprS.GET("/list", userControl.LevelList) //VIP列表

	routers.POST("/api/userinfo/password/reset", userControl.PasswordReset) //重置密码
	routers.POST("/api/userinfo/password/forgot", userControl.PasswordForget)
	routers.POST("/api/userinfo/password/forgot/update", userControl.PasswordForgotUpdate)

	rebateR := routers.Group("/api/rebate/").Use(handler.AuthApiMiddleware())
	rebateR.GET("/record/list", userControl.RebateRecordList)              // 用户返水记录
	rebateR.GET("/detail/record/list", userControl.RebateDetailRecordList) // 用户详情返水记录
	rebateR.GET("/intro", userControl.RebateIntro)                         // 洗码介绍
	rebateR.POST("/data", userControl.GetRebateData)                       // 获取洗码数据
	rebateR.POST("/apply", userControl.RebateApply)                        // 领取返水,洗码
}

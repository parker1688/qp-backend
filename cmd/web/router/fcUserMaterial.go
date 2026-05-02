package router

import (
	"bootpkg/cmd/web/controller/fcUserMaterial"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserMaterialRouter)
}

func fcUserMaterialRouter() {
	r := routers.Group("/api/fcUserMaterial").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserMaterial.FindPageFcUserMaterialControl)   //查询分页
	r.GET("/findByKey", fcUserMaterial.FindByKeyFcUserMaterialControl) //查询主键
	r.POST("/save", fcUserMaterial.SaveFcUserMaterialControl)          //保存
	r.POST("/update", fcUserMaterial.UpdateFcUserMaterialControl)      //修改
	r.POST("/delete", fcUserMaterial.DeleteFcUserMaterialControl)      //删除

	r.POST("/update/agent", fcUserMaterial.UpdateFcUserMaterialAgent) //修改代理信息

	r.GET("/detail", fcUserMaterial.UserMaterialDetail)                           //用户详情信息
	r.POST("/update/status", fcUserMaterial.UpdateFcUserMaterialIsFree)           //修改用户类型
	r.POST("/update/remark", fcUserMaterial.UpdateFcUserMaterialRemark)           //修改备注
	r.POST("/update/withdraw", fcUserMaterial.UpdateFcUserMaterialIsWithdraw)     // 修改用户是否可以提款
	r.POST("/update/bonus", fcUserMaterial.UpdateFcUserMaterialIsBonus)           // 修改用户是否可以领取红利
	r.POST("/update/realName", fcUserMaterial.UpdateFcUserMaterialRealName)       // 修改用户真实姓名
	r.POST("/clear/tel", fcUserMaterial.ClearFcUserMaterialTel)                   // 清空用户手机号码
	r.POST("/update/loginStatus", fcUserMaterial.UpdateFcUserMaterialLoginStatus) // 修改用户状态
	r.POST("/clear/walletPwd", fcUserMaterial.ClearFcUserMaterialWalletPwd)       // 清空用户二级密码

	r.POST("/simulateUser", fcUserMaterial.SimulateUser) //生成模拟账户

	r.GET("/repeat", fcUserMaterial.FindPageUserRepeatControl) // 用户查重

	r.GET("/sameIp/findPage", fcUserMaterial.FindPageFcUserMaterialSameControl) //生成模拟账户

	r.GET("/merchant", fcUserMaterial.GetFcUserMaterialMerchantControl) // 获取用户商户

}

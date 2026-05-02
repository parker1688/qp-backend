package router

import (
	"bootpkg/cmd/web/controller/adminUser"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, adminUserRouter)
}

func adminUserRouter() {
	r := routers.Group("/api/adminUser").Use(handler.AuthMiddleware())

	r.GET("/findPage", adminUser.FindPageAdminUserControl)                        //查询分页
	r.GET("/findByKey", adminUser.FindByKeyAdminUserControl)                      //查询主键
	r.POST("/save", adminUser.SaveAdminUserControl)                               //保存
	r.POST("/update", adminUser.UpdateAdminUserControl)                           //修改
	r.POST("/delete", adminUser.DeleteAdminUserControl)                           //删除
	r.POST("/updateDepartmentId", adminUser.UpdateAdminUserByDepartmentIdControl) //更新用户部门
	r.POST("/updateRoleIds", adminUser.UpdateAdminUserByRoleIdsControl)           //更新用户角色
	r.GET("/findUserAll", adminUser.FindAdminUserAllControl)                      //获取用户ID和用户名字典
	r.POST("/clearMFA", adminUser.ClearMAF)                                       //清除MAF码
	r.POST("/security", adminUser.AdminUserSecurity)                              // 账户安全
	r.POST("/updateMerchantCodes", adminUser.UpdateMerchantCodes)                 //更新用户商户数据
	r.POST("/banUser", adminUser.UpdateStatusBan)                                 //禁止用户登录
	r.POST("/resetPwd", adminUser.UpdateAdminUserPwd)                             //修改管理账号密码
}

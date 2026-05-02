package router

import (
	"bootpkg/cmd/web/controller/domainDistribution"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, domainDistributionRouter)
}

func domainDistributionRouter() {
	r := routers.Group("/api/domainDistribution").Use(handler.AuthMiddleware())

	r.GET("/findPage", domainDistribution.FindPageDomainDistributionControl)   //查询分页
	r.GET("/findByKey", domainDistribution.FindByKeyDomainDistributionControl) //查询主键
	r.POST("/save", domainDistribution.SaveDomainDistributionControl)          //保存
	r.POST("/update", domainDistribution.UpdateDomainDistributionControl)      //修改
	r.POST("/delete", domainDistribution.DeleteDomainDistributionControl)      //删除
}

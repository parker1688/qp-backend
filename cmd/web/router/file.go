package router

import (
	"bootpkg/cmd/web/controller/basecommon"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, filesRouter)
}

func filesRouter() {
	r := routers.Group("/api/filesUploads").Use(handler.AuthMiddleware())

	r.POST("/hash", basecommon.UpHashFile)         //断点上传-hash
	r.POST("/hashSave", basecommon.UpHashFileSave) //断点上传-保存块
	r.POST("/merge", basecommon.MergeHashFile)     //断点上传-hash文件合并

	r.POST("/upload/5cc8019d300000980a055e76", basecommon.UpFileSingle) //小文件上传-单个
}

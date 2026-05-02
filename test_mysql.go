//go:build tools
// +build tools

package main

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	global.CONFIG = conf.NewOnceConfig("./config/hy-backend-api/dev.conf.yaml")
	
	fmt.Println("DSN:", global.CONFIG.Mysql.DSN)
	
	db, err := gorm.Open(mysql.Open(global.CONFIG.Mysql.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	fmt.Println("连接成功!")
	
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Ping失败:", err)
		return
	}
	fmt.Println("Ping成功!")
}

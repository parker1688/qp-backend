package mysqldb

import (
	"bootpkg/common/global"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewOnceMySql() {
	sqlDB, err := sql.Open("mysql", global.CONFIG.Mysql.DSN)
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConn) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(global.CONFIG.Mysql.MaxIdleConn) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)
	err = sqlDB.Ping()
	fmt.Println("start db", err)

	gConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Silent),
		QueryFields: true,
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true,
		//},
	}
	if global.CONFIG.General.ENV == "Debug" {
		gConfig = &gorm.Config{
			Logger:      logger.Default.LogMode(logger.Info),
			QueryFields: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), gConfig)

	global.G_DB = gormDB
}

func NewOnceMySqlSharding() {
	sqlDB, err := sql.Open("mysql", global.CONFIG.MysqlSharding.DSN)
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(global.CONFIG.MysqlSharding.MaxOpenConn) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(global.CONFIG.MysqlSharding.MaxIdleConn) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)
	err = sqlDB.Ping()
	fmt.Println("start db sharding", err)

	gConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Silent),
		QueryFields: true,
		//NamingStrategy: schema.NamingStrategy{
		//	SingularTable: true,
		//},
	}
	if global.CONFIG.General.ENV == "Debug" {
		gConfig = &gorm.Config{
			Logger:      logger.Default.LogMode(logger.Info),
			QueryFields: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), gConfig)

	global.G_DB_SHARDING = gormDB
}

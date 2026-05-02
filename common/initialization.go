package common

import (
	"bootpkg/common/conf"
	"bootpkg/common/database/mysqldb"
	"bootpkg/common/database/redisdb"
	"bootpkg/common/global"
	"bootpkg/common/logs"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/channelData"
)

// TODO:初始化MySQL,Redis,Log,Config等
func Initialization(path string) {
	global.CONFIG = conf.NewOnceConfig(path)           //配置文件初始化
	global.G_LOG = logs.NewLog(global.CONFIG.Log.Path) //初始化日志
	tool.InitIPDB(global.CONFIG.General.Ipdb)          // 初始化 ip 数据库
	if global.CONFIG.Redis.IsInit {
		redisdb.NewOnceRedis() //初始化redis
	}
	if global.CONFIG.Mysql.IsInit {
		mysqldb.NewOnceMySql() //初始化db
	}
	if global.CONFIG.MysqlSharding.IsInit {
		mysqldb.NewOnceMySqlSharding() //初始化db
	}

	if global.CONFIG.Mq.IsInit {
		channelData.InitProducer()
	}

}

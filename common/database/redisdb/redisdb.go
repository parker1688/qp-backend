package redisdb

import (
	"bootpkg/common/global"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
)

func NewOnceRedis() {
	global.G_REDIS = createClient()
}

func createClient() redis.UniversalClient {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        strings.Split(global.CONFIG.Redis.Addr, ","),
		Password:     global.CONFIG.Redis.Password,
		DB:           global.CONFIG.Redis.Db,
		PoolSize:     global.CONFIG.Redis.PoolSize,
		MinIdleConns: global.CONFIG.Redis.MinIdleConn,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return client
}

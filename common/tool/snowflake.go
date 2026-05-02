package tool

import (
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyflakeGenMap map[string]*sonyflake.Sonyflake
	sonyflaskLock   sync.Mutex
)

func SnowflakeId() string {
	return SnowflakeIdByKey("")
}

func SnowflakeIdByKey(key string) string {
	sonyflakeGen, ok := sonyflakeGenMap[key]
	if !ok {
		sonyflaskLock.Lock()
		defer sonyflaskLock.Unlock()
		if sonyflakeGenMap == nil {
			sonyflakeGenMap = make(map[string]*sonyflake.Sonyflake, 1)
		}
		sonyflakeGen, ok = sonyflakeGenMap[key]
		if !ok {
			sonyflakeGen = sonyflake.NewSonyflake(sonyflake.Settings{
				StartTime: time.Date(2022, 10, 1, 1, 1, 1, 1, time.Local),
			})
			sonyflakeGenMap[key] = sonyflakeGen
		}
	}
	if sonyflakeGen == nil {
		return String(time.Now().UnixNano())
	}
	m, err := sonyflakeGen.NextID()
	if err != nil {
		return String(time.Now().UnixNano())
	}
	return String(m)
}

package automapper

import (
	"github.com/devfeel/mapper"
	"reflect"
	"sync"
)

var (
	mutex     sync.Mutex
	mapperMap = make(map[string]interface{})
)

func Register(model interface{}) {
	key := reflect.TypeOf(model).String()
	if _, ok := mapperMap[key]; !ok {
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok = mapperMap[key]; !ok {
			mapperMap[key] = struct{}{}
			mapper.Register(&model)
		}
	}
}

func Map(src, dst interface{}) (err error) {
	Register(src)
	Register(dst)
	mapper.Mapper(src, dst)
	return
}

func MapSlice(srcStruct, dstStruct, src, dst interface{}) (err error) {
	Register(srcStruct)
	Register(dstStruct)
	mapper.MapperSlice(src, dst)
	return
}

//go:build tools
// +build tools

package main

import (
	"bootpkg/common/conf"
	"bootpkg/common/global"
	"fmt"

	"github.com/kirinlabs/utils/encrypt"
)

func main() {
	global.CONFIG = conf.NewOnceConfig("./conf.yaml")
	password := "123456"
	salt := global.CONFIG.SHA256Salt
	hashed := encrypt.Sha256(password + salt)
	fmt.Printf("密码: %s\n盐值: %s\n哈希: %s\n", password, salt, hashed)
}

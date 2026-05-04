package main

import (
	"bootpkg/cmd/api/router"
	"bootpkg/common"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
)

var (
	port       = flag.String("port", "10072", "端口")
	cryptoPort = flag.String("cryptoPort", "10073", "端口")
	configDir  = flag.String("config", "./conf.yaml", "配置文件路径")
)

func init() {
	usage()
	flag.Parse()
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	common.Initialization(*configDir)
	tool.InitSecretManager()
	if err := tool.ValidateSecurityEnv(global.CONFIG.General.ENV == enmus.Release); err != nil {
		log.Fatalf("security config validation failed: %v", err)
	}
	tool.InitPasswordHasher(tool.GetGlobalSecrets().ApiSHA256Salt)
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			var buf [2048]byte
			n := runtime.Stack(buf[:], false)
			global.G_LOG.Errorf(string(buf[:n]))
		}
	}()
	r := router.NewRouter()
	//r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	server := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}

	go server.ListenAndServe()
	var cryptoServer *http.Server
	cryptoHttp := router.NewCryptoRouter()
	cryptoServer = &http.Server{
		Addr:    ":" + *cryptoPort,
		Handler: cryptoHttp,
	}
	go cryptoServer.ListenAndServe()

	InitValidator()

	global.G_LOG.Info("start port:" + *port + "," + *cryptoPort)
	Shutdwon(server, cryptoServer)
}

// **初始化验证器**
func InitValidator() {
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	global.LANG, _ = uni.GetTranslator("zh") // 确保翻译器使用中文

	global.VALIDATE = validator.New()

	// **注册中文翻译，只执行一次，避免多次注册导致冲突**
	err := zhtrans.RegisterDefaultTranslations(global.VALIDATE, global.LANG)
	if err != nil {
		global.G_LOG.Errorf("RegisterDefaultTranslations Error: %v", err)
	}
}

// 关机
func Shutdwon(server *http.Server, server2 *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-c
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	KickAllOnline()
	err := server.Shutdown(cxt)
	if err != nil {
		global.G_LOG.Errorf("shutdwon err: %v", err)
	}
	if server2 != nil {
		err = server2.Shutdown(cxt)
		if err != nil {
			global.G_LOG.Errorf("shutdwon2 err: %v", err)
		}
	}
	global.G_LOG.Info("exit.")
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `start: 
bootpkg -config=conf.yaml -port=10080`)
	if err != nil {
		panic(err.Error())
	}
}

func KickAllOnline() {
	list, err := global.G_REDIS.Keys(context.Background(), "LOGIN:MEMBER:TOKEN:*").Result()
	if err != nil {
		global.G_LOG.Errorf("redis error:%v", err.Error())
	}
	for _, key := range list {
		global.G_REDIS.Del(context.Background(), key)
	}
}

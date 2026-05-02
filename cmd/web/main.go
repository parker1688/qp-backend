package main

import (
	"bootpkg/cmd/web/router"
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
)

var (
	port      = flag.String("port", "10080", "端口")
	configDir = flag.String("config", "./conf.yaml", "配置文件路径")

	cryptoPort = flag.String("cryptoPort", "10081", "端口")
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
	tool.InitPasswordHasher(tool.GetGlobalSecrets().SHA256Salt)
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

	//var cryptoServer *http.Server
	//cryptoHttp := router.NewCryptoRouter()
	//cryptoServer = &http.Server{
	//	Addr:    ":" + *cryptoPort,
	//	Handler: cryptoHttp,
	//}
	//go cryptoServer.ListenAndServe()

	global.G_LOG.Info("start port:" + *port)
	Shutdwon(server)
}

// 关机
func Shutdwon(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-c
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("exit.")
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `start: 
bootpkg -config=conf.yaml -port=10080`)
	if err != nil {
		panic(err.Error())
	}
}

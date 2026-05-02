package main

import (
	"bootpkg/cmd/betconsumidor/controller"
	"bootpkg/common"
	"bootpkg/common/global"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	port      = flag.String("port", "10085", "端口")
	configDir = flag.String("config", "./conf.yaml", "配置文件路径")
)

func init() {
	usage()
	flag.Parse()
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	common.Initialization(*configDir)
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
	ctx, cancel := context.WithCancel(context.Background())
	controller.NewConsumers(ctx)
	global.G_LOG.Info("kafka start success.")
	Shutdwon(cancel)
}

// 关机
func Shutdwon(cancel context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	switch <-c {
	case syscall.SIGQUIT:
		global.G_LOG.Info("Shutdown quickly, bye...")
	case os.Interrupt, syscall.SIGTERM: // os.Interrupt==syscall.SIGINT
		global.G_LOG.Info("Shutdown gracefully, bye...")
	}
	cancel()
	<-time.After(2 * time.Second)
	os.Exit(0)
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `start: 
bootpkg -config=conf.yaml -port=10080`)
	if err != nil {
		panic(err.Error())
	}
}

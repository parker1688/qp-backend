package main

import (
	"bootpkg/cmd/cron/crontab"
	"bootpkg/common"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	port      = flag.String("port", "10082", "端口")
	configDir = flag.String("config", "./conf.yaml", "配置文件路径")
)

func init() {
	usage()
	flag.Parse()
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	common.Initialization(*configDir)
}

type CronRecoverLogger struct {
}

func (c CronRecoverLogger) Info(msg string, keysAndValues ...interface{}) {
	global.G_LOG.Info(msg, keysAndValues)
}

func (c CronRecoverLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	global.G_LOG.Error(err, msg, keysAndValues)
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
	logger := &CronRecoverLogger{}
	c := cron.New(cron.WithChain(cron.Recover(logger), cron.SkipIfStillRunning(cron.DefaultLogger)))
	crontab.NewCron(c)
	c.Start()
	for i, entry := range c.Entries() {
		fmt.Println(i, entry.Next.Format(tool.TimeLayout))
	}
	global.G_LOG.Info("cron start success.")
	Shutdwon(c)
}

// 关机
func Shutdwon(tab *cron.Cron) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	switch <-c {
	case syscall.SIGQUIT:
		global.G_LOG.Info("Shutdown quickly, bye...")
	case os.Interrupt, syscall.SIGTERM: // os.Interrupt==syscall.SIGINT
		global.G_LOG.Info("Shutdown gracefully, bye...")
	}
	if tab != nil {
		tab.Stop()
	}
	os.Exit(0)
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `start: 
bootpkg -config=conf.yaml -port=10080`)
	if err != nil {
		panic(err.Error())
	}
}

package crontab

import "github.com/robfig/cron/v3"

var (
	cronFunc []cronTabEvery
)

// cronTabEvery
// @Description: 定时任务配置信息
type cronTabEvery struct {
	spec string //定时任务
	cmd  func() //执行方法
}

func NewCron(c *cron.Cron) {
	for _, v := range cronFunc {
		c.AddFunc(v.spec, v.cmd)
	}
	cronFunc = nil
}

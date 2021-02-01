package main

import (
	"flag"
	"fmt"
	"github.com/hitokoto-osc/telegram_bot/build"
	"github.com/hitokoto-osc/telegram_bot/config"
	"github.com/hitokoto-osc/telegram_bot/event"
	"github.com/hitokoto-osc/telegram_bot/telegram"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var (
	v   bool   // 查看版本
	h   bool   // 查看帮助
	c   string // 指定的配置文件
	vvv bool
)

func parseFlags() {
	flag.BoolVar(&v, "v", false, "查看程序版本")
	flag.BoolVar(&h, "h", false, "查看程序帮助")
	flag.StringVar(&c, "c", "", "设定配置文件地址")
	flag.BoolVar(&vvv, "vvv", false, "启用调试模式")
	flag.Parse()
	if h {
		fmt.Printf(`一言电报机器人 v%s
使用: nginx [-hv] [-c filename]

选项:
`, build.Version)
		flag.PrintDefaults()
		os.Exit(0)
	}
	if v {
		fmt.Printf("一言电报机器人服务 \n版本: %s\nGitCommit: %s\n编译时间: %s\n编译环境: %s\n", build.Version, build.CommitTag, build.CommitTime, runtime.Version())
		os.Exit(0)
	}
}

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	if vvv { // 是否启用 DEBUG 日记输出
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func init() {
	parseFlags()
	initLogger()
	config.InitConfig(c)
}

func main() {
	log.Info("开始初始化机器人...")
	bot := telegram.InitBot()
	// 注册机器人事件
	event.RegisterEvent(bot)
	go bot.Start()
	log.Info("机器人初始化完成，开始接收信息。")
	select {} // revive:disable-line:empty-block 错误检测。 堵塞进程是必须的
}

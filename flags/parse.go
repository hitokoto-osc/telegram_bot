package flags

import (
	"flag"
	"fmt"
	"github.com/hitokoto-osc/telegram_bot/build"
	"github.com/hitokoto-osc/telegram_bot/config"
	"os"
	"runtime"
)

var (
	v bool // 查看版本
	h bool // 查看帮助
)

func init() {
	flag.BoolVar(&v, "v", false, "查看程序版本")
	flag.BoolVar(&h, "h", false, "查看程序帮助")
	flag.StringVar(&config.Path, "c", "", "设定配置文件地址")
	flag.BoolVar(&config.Debug, "vvv", false, "启用调试模式")
	flag.Parse()
}

func Do() {
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

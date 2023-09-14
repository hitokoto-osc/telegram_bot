package command

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hitokoto-osc/telegram_bot/build"
	"github.com/levigross/grequests"
	"github.com/shirou/gopsutil/v3/load"
	log "github.com/sirupsen/logrus"
)

// Status 用于响应获取统计信息的指令
func Status(b *telebot.Bot) {
	b.Handle("/status", func(ctx telebot.Context) error {
		response, err := grequests.Get("https://status.hitokoto.cn/v1/statistic", nil)
		if err != nil {
			log.Errorf("尝试获取统计数据时出现错误，错误信息： %s\n", err)
			_, err = b.Send(ctx.Chat(), "很抱歉，尝试获取数据时发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return err
		}
		result := &hitokotoStatusAPIV1Response{}
		err = response.JSON(result)
		if err != nil {
			log.Errorf("尝试解析统计数据时发生错误，错误信息： %s", err)
			_, err = b.Send(ctx.Chat(), "很抱歉，尝试解析数据时发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return err
		}

		// 读取系统负载
		lo, err := load.Avg()
		if err != nil {
			log.Errorf("尝试解析系统负载时发生错误，错误信息： %s", err)
			_, err = b.Send(ctx.Chat(), "很抱歉，尝试解析系统负载时发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return err
		}
		// log.Debug(data)
		_, err = b.Send(ctx.Chat(), fmt.Sprintf(`*[一言统计信息]*
句子总数： %s
现存分类： %s
服务负载： %s
内存占用： %s MB
每分请求： %s
每时请求： %s
当日请求： %s

*[调试信息]*
当前时间： %s
操作系统： %s
设备架构： %s
系统负载： %s
程序版本： v%s
运行环境： %s
编译时间： %s
编译哈希： %s
`,
			strconv.Itoa(result.Data.Status.Hitokoto.Total),
			strings.Join(result.Data.Status.Hitokoto.Category, ","),
			loadToString(result.Data.Status.Load[0])+","+loadToString(result.Data.Status.Load[1])+","+loadToString(result.Data.Status.Load[2]),
			loadToString(result.Data.Status.Memory),
			strconv.FormatUint(result.Data.Requests.All.PastMinute, 10),
			strconv.FormatUint(result.Data.Requests.All.PastHour, 10),
			strconv.FormatUint(result.Data.Requests.All.PastDay, 10),
			time.Now().Format("2006年1月2日 15:04:05"),
			runtime.GOOS,
			runtime.GOARCH,
			loadToString(lo.Load1)+","+loadToString(lo.Load5)+","+loadToString(lo.Load15),
			build.Version,
			runtime.Version(),
			build.CommitTime,
			build.CommitTag,
		),
			&telebot.SendOptions{
				ParseMode: "Markdown",
			},
		)
		return err
	})
}

type hitokotoStatusAPIV1Response struct { // 因为不需要使用全部数据，所以这里就只解析部分了
	Data struct {
		Status   status   `json:"status"`
		Requests requests `json:"requests"`
	} `json:"data"`
}

type status struct {
	Load     []float64 `json:"load"`
	Memory   float64   `json:"memory"`
	Hitokoto hitokoto  `json:"hitokoto"`
}

type hitokoto struct {
	Total    int      `json:"total"`
	Category []string `json:"category"`
}

type requests struct {
	All all `json:"all"`
}

type all struct {
	Total      uint64 `json:"total"`
	PastMinute uint64 `json:"past_minute"`
	PastHour   uint64 `json:"past_hour"`
	PastDay    uint64 `json:"past_day"`
}

func loadToString(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

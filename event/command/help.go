package command

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"time"

	"gopkg.in/telebot.v3"
)

// Help 返回使用帮助
func Help(b *telebot.Bot) {
	b.Handle("/help", func(ctx telebot.Context) error {
		err := ctx.Reply(fmt.Sprintf(`*以下是目前支持的指令：*
/about 关于机器人
/hitokoto [分类] 获取一条句子，[分类] 可以在“开发者中心”的“语句接口”部分找到详细定义。默认返回随机分类。
/help 获取机器人帮助信息。
/image [分类] 获取随机图片。[分类] 指图片的类别，目前只支持 Bing。默认返回 Bing。
/ping 测试机器人连通性。
/status 查看一言语句接口的状态，以及机器人的调试信息。
--------------
当前服务器时间：%s`, time.Now().Format("2006年1月2日 15:04:05")),
			&telebot.SendOptions{
				ParseMode: "markdown",
			},
		)
		if err != nil {
			return errors.Wrap(err, "无法发送消息")
		}
		return nil
	})
}

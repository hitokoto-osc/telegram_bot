package command

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"time"

	"gopkg.in/telebot.v3"
)

// About 返回关于信息
func About(b *telebot.Bot) {
	b.Handle("/about", func(ctx telebot.Context) error {
		err := ctx.Reply(fmt.Sprintf(`ヒトコト（一言） 官方 Telegram 机器人。 目前仅提供简体中文支持。 主要提供一句话服务。
* 官方网站: https://hitokoto.cn
* 官方 QQ 群组: 70029304
* 非官方 Telegram 群组: https://t.me/hitokoto
* 项目开源地址：https://github.com/hitokoto-osc
--------------
当前服务器时间：%s`, time.Now().Format("2006年1月2日 15:04:05")))
		if err != nil {
			return errors.Wrap(err, "无法发送消息")
		}
		return nil
	})
}

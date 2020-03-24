package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

func About(b *telebot.Bot) {
	b.Handle("/about", func(m *telebot.Message) {
		_, err := b.Send(m.Chat, fmt.Sprintf(`ヒトコト（一言） 官方 Telegram 机器人。 目前仅提供简体中文支持。 主要提供一句话服务。
* 官方网站: https://hitokoto.cn
* 官方 QQ 群组: 70029304
* 非官方 Telegram 群组: https://t.me/hitokoto
* 项目开源地址：https://github.com/hitokoto-osc
--------------
当前服务器时间：%s`, time.Now().Format("2006年1月2日 15:04:05")))
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n")
		}
	})
}

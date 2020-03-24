package command

import (
	"fmt"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

var supportedTypeList = []string{
	"a", // 动画
	"b", // 漫画
	"c", // 游戏
	"d", // 文学
	"e", // 原创
	"f", // 来自网络
	"g", // 其他
	"h", // 影视
	"i", // 诗词（主要是文言诗集）
	"j", // 网易云
	"k", // 哲学
	"l", // 抖机灵
}

func inStringSlice(haystack []string, needle string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}

	return false
}

func Hitokoto(b *telebot.Bot) {
	b.Handle("/hitokoto", func(m *telebot.Message) {
		payload := m.Payload // 指令：`/hitokoto <payload>` 这里提取 payload 用于提取参数
		url := "https://v1.hitokoto.cn/"
		if inStringSlice(supportedTypeList, payload) {
			url += "?c=" + payload
		}

		// 请求接口
		response, err := grequests.Get(url, nil)
		if err != nil {
			log.Errorf("尝试获取一言时出现错误，错误信息： %s\n", err)
			_, err = b.Send(m.Chat, "很抱歉，尝试获取发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return
		}
		data := &HitokotoSentenceApiV1Response{}
		err = response.JSON(data)
		if err != nil {
			log.Errorf("尝试解析一言时发生错误，错误信息： %s", err)
			_, err = b.Send(m.Chat, "很抱歉，尝试解析一言时发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return
		}
		_, err = b.Reply(m, fmt.Sprintf(`%s —— %s「%s」`, data.Hitokoto, data.FromWho, data.From))
		if err != nil {
			log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
		}
	})
}

type HitokotoSentenceApiV1Response struct {
	Id         uint32 `json:"id"`
	Hitokoto   string `json:"hitokoto"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUid int32  `json:"creator_uid"`
	Reviewer   int32  `json:"reviewer"`
	UUID       string `json:"uuid"`
	CreatedAt  string `json:"created_at"`
}

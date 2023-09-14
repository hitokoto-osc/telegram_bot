package command

import (
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/hitokoto-osc/telegram_bot/request"
	"github.com/samber/lo"
	"gopkg.in/telebot.v3"
	"strings"
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

// Hitokoto 获取一言
func Hitokoto(b *telebot.Bot) {
	b.Handle("/hitokoto", func(ctx telebot.Context) error {
		payload := strings.TrimSpace(ctx.Message().Payload) // 指令：`/hitokoto <payload>` 这里提取 payload 用于提取参数
		url := "https://v1.hitokoto.cn/"
		if lo.Contains(supportedTypeList, payload) {
			url += fmt.Sprintf("?c=%s", payload)
		}
		// 请求接口
		client := request.NewDefault()
		resp, err := client.Get(url)
		if err != nil {
			return errors.Wrap(err, "无法请求一言接口")
		}
		defer resp.Body.Close()
		// 解析 JSON
		data := &HitokotoSentenceAPIV1Response{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return errors.Wrap(err, "无法解析一言接口返回的 JSON 数据")
		}
		_, err = b.Reply(ctx.Message(), fmt.Sprintf(`%s —— %s「%s」`, data.Hitokoto, data.FromWho, data.From))
		return err
	})
}

// HitokotoSentenceAPIV1Response 定义了一言接口的结构
type HitokotoSentenceAPIV1Response struct {
	ID         uint32 `json:"id"`
	Hitokoto   string `json:"hitokoto"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int32  `json:"creator_uid"`
	Reviewer   int32  `json:"reviewer"`
	UUID       string `json:"uuid"`
	CreatedAt  string `json:"created_at"`
}

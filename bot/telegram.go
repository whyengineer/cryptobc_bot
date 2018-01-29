package bot

import (
	"gopkg.in/telegram-bot-api.v4"
)

func (b *Bot) SendMessage(msg *tgbotapi.Message, payload string, reply bool) error {
	content := tgbotapi.NewMessage(msg.Chat.ID, payload)
	if reply {
		content.ReplyToMessageID = msg.MessageID
	}
	ack, err := b.Bot.Send(content)
	b.log.Println(ack)
	return err
}

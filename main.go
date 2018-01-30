package main

import (
	"net/http"

	"github.com/whyengineer/cryptobc_bot/bot"
)

func main() {
	cbc := bot.NewBot("497425065:AAFDaMLuxdghQblsf6QG-ByH7YB4FvETlBs", nil)
	//cbc.bot.Debug = true
	updates := cbc.Bot.ListenForWebhook("/")
	go http.ListenAndServe("127.0.0.1:5000", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message != nil {
			cbc.Router(*update.Message)
		} else if update.EditedMessage != nil {
			cbc.Router(*update.EditedMessage)
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		//bot.Send(msg)
	}
}

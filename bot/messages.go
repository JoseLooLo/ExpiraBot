package bot

import (
	"log"
	security "github.com/JoseLooLo/ExpiraBot/security"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (e Event) SendStartMessage(r security.Request) {
	msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
	log.Printf(r.Update.Message.Text)
	msg.Text = r.Update.Message.Text
	r.Bot.Send(msg)
}

func (e Event) SendUpdateMessage(r security.Request) {
	msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
	log.Printf(r.Update.Message.Text)
	msg.Text = r.Update.Message.Text
	r.Bot.Send(msg)
}
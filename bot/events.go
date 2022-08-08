package bot

import (
	"log"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	security "github.com/JoseLooLo/ExpiraBot/security"
)



func (e Event) Start(r security.Requisition) {
	user := expiraBot.User{r.Update.Message.Chat.ID, false}
	r.Database.InsertUser(user)

	//TODO
	msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
	log.Printf(r.Update.Message.Text)
	msg.Text = r.Update.Message.Text
	r.Bot.Send(msg)
}

func (e Event) Insert(r security.Requisition) {
}

func (e Event) Update(r security.Requisition) {
}

func (e Event) Books(r security.Requisition) {
}

func (e Event) Help(r security.Requisition) {
}

func (e Event) ErrorCommand(r security.Requisition) {
}
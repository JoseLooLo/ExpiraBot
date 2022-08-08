package bot

import (
	"log"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r requisiton) Start() {
	user := expiraBot.User{r.update.Message.Chat.ID, false}
	r.database.InsertUser(user)

	//TODO
	msg := tgbotapi.NewMessage(r.update.Message.Chat.ID, "")
	log.Printf(r.update.Message.Text)
	msg.Text = r.update.Message.Text
	r.bot.Send(msg)
}

func (r requisiton) Insert() {
}

func (r requisiton) Update() {
}

func (r requisiton) Books() {
}

func (r requisiton) Help() {
}

func (r requisiton) ErrorCommand() {
}
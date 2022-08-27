package security

import (
	"log"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Request struct {
	Bot *tgbotapi.BotAPI
	Update tgbotapi.Update
	Database expiraBot.Database
}

type Security interface {
	Next(r Request)
}

type SecurityChain struct {
	Request Request
	NextChain Security
}

//Just call the Next method of the securityChain
func (s *SecurityChain) Execute() {
	s.Next(s.Request)
}

//Check if the user is blocked
//If is not blocked continue, otherwise just finish the request
func (s *SecurityChain) Next(r Request) {
	user := r.Database.GetUserInfoById(r.Update.Message.Chat.ID)
	if (user.Block) {
		log.Printf("[Security][SecurityChain][Next] - User %d is blocked", r.Update.Message.Chat.ID)
		return
	}
	s.NextChain.Next(r)
}



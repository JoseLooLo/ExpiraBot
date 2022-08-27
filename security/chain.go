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

func (s *SecurityChain) Execute() {
	log.Println("Security Chain")
	s.Next(s.Request)
}

func (s *SecurityChain) Next(r Request) {
	s.NextChain.Next(r)
}



package security

import (
	"log"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Requisition struct {
	Bot *tgbotapi.BotAPI
	Update tgbotapi.Update
	Database expiraBot.Database
}

type Security interface {
	Next(r Requisition)
}

type SecurityChain struct {
	Requisition Requisition
	NextChain Security
}

func (s *SecurityChain) Execute() {
	log.Println("Security Chain")
	s.Next(s.Requisition)
}

func (s *SecurityChain) Next(r Requisition) {
	s.NextChain.Next(r)
}



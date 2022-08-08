
package bot

import (
	"log"
	security "github.com/JoseLooLo/ExpiraBot/security"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Event struct {
	event string
}

func (e *Event) Next(r security.Requisition) {
	switch e.event {
		case "start":
			go e.Start(r)
		case "insert":
			go e.Insert(r)
		case "update":
			go e.Update(r)
		case "books":
			go e.Books(r)
		case "help":
			go e.Help(r)
		default:
			go e.ErrorCommand(r)
	}
}

func Start(key string, database expiraBot.Database, debug bool) {

	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			requisition := security.Requisition{bot, update, database}
			event := &Event{update.Message.Command()}
			flood := &security.FloodChain{event}
			chain := &security.SecurityChain{requisition, flood}
			go chain.Execute()
		}
	}
}

package bot

import (
	"log"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type requisiton struct {
	bot *tgbotapi.BotAPI
	update tgbotapi.Update
	database expiraBot.Database
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
			new_requisiton := requisiton{bot, update, database}
			switch update.Message.Command() {
				case "start":
					go new_requisiton.Start()
				case "insert":
					go new_requisiton.Insert()
				case "update":
					go new_requisiton.Update()
				case "books":
					go new_requisiton.Books()
				case "help":
					go new_requisiton.Help()
				default:
					go new_requisiton.ErrorCommand()
			}
		}
	}
}
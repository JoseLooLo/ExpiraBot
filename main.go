
package main

import (
	"log"
	"os"
	mongodb "github.com/JoseLooLo/ExpiraBot/database/mongodb"
	bot "github.com/JoseLooLo/ExpiraBot/bot"
)

func main() {
	log.Printf("Starting ExpiraBot...")

	telegram_bot_key := os.Getenv("TELEGRAM_BOT_KEY")
	database_url := os.Getenv("DATABASE_URL")

	log.Printf("%s=%s\n", "TELEGRAM_BOT_KEY", telegram_bot_key)
	log.Printf("%s=%s\n", "DATABASE_URL", database_url)

	database := &mongodb.Mongodb{}
	closedb := database.Start(database_url)
	defer closedb()

	bot.Start(telegram_bot_key, database, false)
}
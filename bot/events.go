package bot

import (
	"log"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	security "github.com/JoseLooLo/ExpiraBot/security"
	crawler "github.com/JoseLooLo/ExpiraBot/crawler"
)

//Start command
//Insert the user in the database and send back a welcome message
func (e Event) Start(r security.Request) {
	user := expiraBot.User{Id: r.Update.Message.Chat.ID, Block: false}
	r.Database.InsertUser(user)

	e.SendStartMessage(r)
}

func (e Event) Insert(r security.Request) {
}

//Update command
//Get books of the user on the bu website and insert in the database
func (e Event) Update(r security.Request) {
	args := GetArgs(r)
	if (len(args) != 2) {
		log.Printf("[Error][Bot][Update] - Invalid args. Expected 2 got %d", len(args))
		return
	}
	login := args[0]
	password := args[1]

	if (!ValidateLogin(login)) {
		log.Printf("[Error][Bot][Update] - Invalid Login Format %s", login)
		return
	}

	if (!ValidatePassword(password)) {
		log.Printf("[Error][Bot][Update] - Invalid Password Format %s", login)
		return
	}

	books, err_crawler := crawler.Crawler(login, password)
	if (err_crawler != nil) {
		log.Println(err_crawler.Error())
		return
	}

	books = InsertUserToBooks(r.Update.Message.Chat.ID, books)
	for index, value := range books {
		log.Println(index, value)
	}

	r.Database.Update(books)
	e.SendUpdateMessage(r)
}

func (e Event) Books(r security.Request) {
}

func (e Event) Help(r security.Request) {
}

func (e Event) ErrorCommand(r security.Request) {
}
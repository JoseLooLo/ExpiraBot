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
	log.Printf("[Info][Bot][Start] - User %d", r.Update.Message.Chat.ID)
	user := expiraBot.User{Id: r.Update.Message.Chat.ID, Block: true}
	r.Database.InsertUser(user)

	e.SendStartMessage(r)
}

//Insert command
//Insert a book that is passed by argument by the user
func (e Event) Insert(r security.Request) {
	//@TODO
	log.Printf("[Info][Bot][Insert] - User %d", r.Update.Message.Chat.ID)
	e.SendInsertMessage(r)
}

//Update command
//Get books of the user on the bu website and insert in the database
func (e Event) Update(r security.Request) {
	log.Printf("[Info][Bot][Update] - User %d", r.Update.Message.Chat.ID)
	args := GetArgs(r)
	if (len(args) != 2) {
		e.SendUpdateMessage(r, "Argumentos inválidos. Use /update [matricula] [senha]")
		log.Printf("[Error][Bot][Update] - Invalid args. Expected 2 got %d", len(args))
		return
	}
	login := args[0]
	password := args[1]

	if (!ValidateLogin(login)) {
		e.SendUpdateMessage(r, "Formato de matrícula inválido")
		log.Printf("[Error][Bot][Update] - Invalid Login Format %s", login)
		return
	}

	if (!ValidatePassword(password)) {
		e.SendUpdateMessage(r, "Formato de senha inválido")
		log.Printf("[Error][Bot][Update] - Invalid Password Format %s", login)
		return
	}

	books, err_crawler := crawler.Crawler(login, password)
	if (err_crawler != nil) {
		e.SendUpdateMessage(r, "Erro ao tentar acessar a sua conta na BU. Login ou senha inválidos.")
		log.Println(err_crawler.Error())
		return
	}

	books = InsertUserToBooks(r.Update.Message.Chat.ID, books)

	r.Database.Update(books)
	e.SendUpdateMessage(r, "Os livros foram atualizados!")
	e.Books(r)
}

//Books command
//Send a list of the user's books to the user
func (e Event) Books(r security.Request) {
	log.Printf("[Info][Bot][Books] - User %d", r.Update.Message.Chat.ID)
	user := expiraBot.User{Id: r.Update.Message.Chat.ID}
	books := r.Database.GetBooks(user)
	e.SendBooksMessage(r, books)
}

//Help command
//Send a help message to the user
func (e Event) Help(r security.Request) {
	log.Printf("[Info][Bot][Help] - User %d", r.Update.Message.Chat.ID)
	e.SendHelpMessage(r)
}

//ErrorCommand
//Send a message to the user when he try to use a invalid command
func (e Event) ErrorCommand(r security.Request) {
	log.Printf("[Info][Bot][ErrorCommand] - User %d", r.Update.Message.Chat.ID)
	e.SendErrorCommandMessage(r)
}
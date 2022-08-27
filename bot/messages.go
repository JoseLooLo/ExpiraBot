package bot

import (
	"log"
	security "github.com/JoseLooLo/ExpiraBot/security"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
)

func (e Event) SendStartMessage(r security.Request) {
	msg := tgbotapi.NewPhoto(r.Update.Message.Chat.ID, tgbotapi.FilePath("./docs/startCommandImage.jpg"))
	msg.Caption = "Alguém chamou a detetive Chika para descobrir quando os livros da BU precisam ser devolvidos?"
	_, err_msg := r.Bot.Send(msg)
	if err_msg != nil {
		log.Printf("[Error][Bot][SendStartMessage] - " + err_msg.Error())
		return
	}
}

func (e Event) SendUpdateMessage(r security.Request, message string) {
	msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
	msg.Text = message
	r.Bot.Send(msg)
}

func (e Event) SendBooksMessage(r security.Request, books []expiraBot.Books) {
	img := tgbotapi.NewPhoto(r.Update.Message.Chat.ID, tgbotapi.FilePath("./docs/booksCommandImage.jpg"))
	if (len(books) == 0) {
		img.Caption = "Não encontrei nenhum livro"
	} else {
		img.Caption = "Encontrei esses livros"
	}
	_, err_img := r.Bot.Send(img)

	if err_img != nil {
		log.Printf("[Error][Bot][SendBooksMessage] - " + err_img.Error())
	}

	//Is necessary send several messages because it don't work if the message is too long
	for _, value := range books {
		msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
		msg.Text = value.Book + "\nData de devolução: " + value.Date + "\n"
		_, err_msg := r.Bot.Send(msg)
		if err_msg != nil {
			log.Printf("[Error][Bot][SendBooksMessage] - " + err_msg.Error())
		}
	}
}

func (e Event) SendInsertMessage(r security.Request) {
	msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
	msg.Text = "Em construção..."
	r.Bot.Send(msg)
}

func (e Event) SendHelpMessage(r security.Request) {
	msg := tgbotapi.NewMessage(r.Update.Message.Chat.ID, "")
	msg.Text = "Lista de comandos válidos:\n"
	msg.Text += "/books - Lista os livros emprestados da BU\n"
	msg.Text += "/update [matricula] [senha] - Atualiza os livros emprestados da BU\n"
	msg.Text += "/help - Volte ao começo\n"
	_, err_msg := r.Bot.Send(msg)
	if err_msg != nil {
		log.Printf("[Error][Bot][SendHelpMessage] - " + err_msg.Error())
	}
}

func (e Event) SendErrorCommandMessage(r security.Request) {
	msg := tgbotapi.NewPhoto(r.Update.Message.Chat.ID, tgbotapi.FilePath("./docs/sendErrorCommandImage.jpg"))
	msg.Caption = "Você digitou um comando inválido! Digite /help para listar os comandos válidos."
	_, err_msg := r.Bot.Send(msg)
	if err_msg != nil {
		log.Printf("[Error][Bot][SendErrorCommandMessage] - " + err_msg.Error())
		return
	}
}


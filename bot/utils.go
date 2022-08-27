package bot

import (
	"strings"
	security "github.com/JoseLooLo/ExpiraBot/security"
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
)

//Split the command from the arguments
//Receive a user request
//Return a string slice that have every argument sended by the user
func GetArgs(r security.Request) []string {
	return strings.Split(r.Update.Message.Text, " ")[1:]
}

//Validate the login format
//Receive a login in a string format
//Return true if the format is correct, false otherwise
func ValidateLogin(login string) bool {
	//@TODO
	return true
}

//Validate the password format
//Receive a password in a string format
//Return true if the format is correct, false otherwise
func ValidatePassword(password string) bool {
	//@TODO
	return true
}

//The books slice received from crawler don't have the user id
//Receive as argument the user id and the books slice
//Return the books slice populate with the user id
func InsertUserToBooks(user int64, books []expiraBot.Books) []expiraBot.Books {
	for index, _ := range books {
		books[index].Userid = user;
	}
	return books
}
package database


type User struct {
	Id int64
	Block bool
}

type Books struct {
	Userid int64
	Book string
	Date string
}

type Database interface {
	InsertUser(user User) bool
	InsertBook(Books) bool
	DeleteBook(Books) bool
	DeleteAll(user User) bool
	GetBooks(user User) []Books
	GetUserInfoById(userId int64) User
	Update([]Books) bool
	Start(url string) func()
}
package database


type User struct {
	Id int64
	Block bool
}

type Books struct {
	Userid int64
	Book string
}

type Database interface {
	InsertUser(user User) bool
	// Delete() bool
	// Update() bool
	Start(url string) func()
}
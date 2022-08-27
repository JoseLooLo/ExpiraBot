package mongodb

import (
	expiraBot "github.com/JoseLooLo/ExpiraBot/database"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	user *mongo.Collection
	books *mongo.Collection
}

//Starts a connection with the mongodb database
//Returns a function that close the connection
func (db *Mongodb) Start(url string) func() {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("Error with the mongodb connection.")
		panic(err)
	}

	db.user = client.Database("expirabot").Collection("user")
	db.books = client.Database("expirabot").Collection("books")

	log.Printf("Connection with the mongodb was established.")

	return func (){
		client.Disconnect(ctx)
	}
}

//Update the books using the crawler on the bu website
//Return true if the operations was successful, false otherwise
func (db Mongodb) Update(books []expiraBot.Books) bool {
	if (len(books) == 0) {
		return true
	}
	user := expiraBot.User{Id: books[0].Userid, Block: false}
	if db.userExist(user) {
		//Delete all books before insert the news books
		db.DeleteAll(user)
		for _, value := range books {
			_, err := db.books.InsertOne(context.TODO(), value)
			if err != nil {
				log.Println(err.Error())
				return false
			}
		}
	}
	return true
}

//Delete all books from a user
//Return true if the operations was successful, false otherwise
func (db Mongodb) DeleteAll(user expiraBot.User) bool {
	filter := bson.M{"userid": user.Id}
    _, err := db.books.DeleteMany(context.TODO(), filter)
	return err != nil
}

//Delete all books from a user
//Return true if the operations was successful, false otherwise
func (db Mongodb) GetBooks(user expiraBot.User) []expiraBot.Books {
	var books []expiraBot.Books
	filter := bson.M{"userid": user.Id}
    res, err_find := db.books.Find(context.TODO(), filter)
	if (err_find != nil) {
		return books
	}

	err_all := res.All(context.TODO(), &books)
	if (err_all != nil) {
		return books
	}
	return books
}

//Insert a new user in the mongo database
//Return true if the operations was successful, false otherwise
func (db Mongodb) InsertUser(user expiraBot.User) bool {
	if !db.userExist(user) {
		_, err := db.user.InsertOne(context.TODO(), user)
		if err != nil {
			return false
		}
	}
	return true
}

//Check if a user already exists in the mongo database
//Return true if exists, false otherwise
func (db Mongodb) userExist(user expiraBot.User) bool {
	filter := bson.D{{Key: "id", Value: user.Id}}
	var res expiraBot.User
    err := db.user.FindOne(context.TODO(), filter).Decode(&res)
	return err != mongo.ErrNoDocuments
}

//Insert a new book in the mongo database
//Return true if the operations was successful, false otherwise
func (db Mongodb) InsertBook(book expiraBot.Books) bool {
	//@TODO
	return true
}

//Delete a specific book in the mongo database
//Return true if the operations was successful, false otherwise
func (db Mongodb) DeleteBook(book expiraBot.Books) bool {
	//@TODO
	return true
}

//Delete a specific book in the mongo database
//Return true if the operations was successful, false otherwise
func (db Mongodb) GetUserInfoById(userId int64) expiraBot.User {
	filter := bson.M{"id": userId}
	var res expiraBot.User
    db.user.FindOne(context.TODO(), filter).Decode(&res)
	return res
}
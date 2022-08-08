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

//Insert a new user in the mongo database
//Return true if the operations was successful, false otherwise
func (db Mongodb) InsertUser(user expiraBot.User) bool {
	if !db.userExist(user) {
		_, err := db.user.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println("[%s] - Error to insert the user.", user.Id)
			return false
		}
		log.Println("[%s] - New user inserted.", user.Id)
	}
	return true
}

//Check if a user already exists in the mongo database
//Return true if exists, false otherwise
func (db Mongodb) userExist(user expiraBot.User) bool {
	filter := bson.D{{"id", user.Id}}
	var res expiraBot.User
    err := db.user.FindOne(context.TODO(), filter).Decode(&res)
	return err != mongo.ErrNoDocuments
}
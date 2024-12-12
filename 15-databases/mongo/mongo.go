package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/

type DB struct {
	client   *mongo.Client
	database *mongo.Database
	coll     *mongo.Collection
}

type Person struct {
	FirstName string   `bson:"first_name"`
	Email     string   `bson:"email"`
	Age       int      `bson:"age"`
	Marks     int      `bson:"marks"`
	Hobbies   []string `bson:"hobbies"`
}

func NewDB(collection string) (*DB, error) {
	const uri = "mongodb://root:example@localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	database := client.Database("info")
	fmt.Printf("%+v\n", database)
	if database == nil {
		return nil, fmt.Errorf("failed to create DB: %w", err)
	}
	coll := database.Collection(collection)
	if coll == nil {
		return nil, fmt.Errorf("failed to create collection: %w", err)
	}

	return &DB{
		client:   client,
		database: database,
		coll:     coll,
	}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.client.Ping(ctx, nil)
}

func main() {

	db, err := NewDB("users")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.Ping(ctx)
	if err != nil {
		panic(err)
	}

	db.InsertOne(ctx)
	fmt.Println("Connected to MongoDB!")

}

func (db *DB) InsertOne(ctx context.Context) {
	u := Person{
		FirstName: "John",
		Email:     "john@email.com",
		Age:       30,
		Hobbies:   []string{"Sports", "Cooking"},
		Marks:     50,
	}

	res, err := db.coll.InsertOne(ctx, u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.InsertedID)
}

// create a function the inserts multiple users record in one go

//InsertMany

package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

//https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/

// DB struct is used to store MongoDB client, a reference to a database, and a collection.
type DB struct {
	client   *mongo.Client
	database *mongo.Database
	coll     *mongo.Collection
}

// Person struct represents the schema of a document in the MongoDB collection.
type Person struct {
	FirstName string   `bson:"first_name"`
	Email     string   `bson:"email"`
	Age       int      `bson:"age"`
	Marks     int      `bson:"marks"`
	Hobbies   []string `bson:"hobbies"`
}

// NewDB function initializes a connection to MongoDB, sets up the database and collection.
func NewDB(collection string) (*DB, error) {
	const uri = "mongodb://root:example@localhost:27017"
	//Create client options and set the connection URI
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB using the provided context and options
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	database := client.Database("info")
	fmt.Printf("%+v\n", database)
	if database == nil {
		return nil, fmt.Errorf("failed to get DB: %w", err)
	}
	coll := database.Collection(collection)
	if coll == nil {
		return nil, fmt.Errorf("failed to get collection: %w", err)
	}

	return &DB{
		client:   client,
		database: database,
		coll:     coll,
	}, nil
}

// Ping method is used to test if the MongoDB client can communicate with the database.
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

// InsertOne method inserts a single document into the MongoDB collection
func (db *DB) InsertOne(ctx context.Context) {
	u := Person{
		FirstName: "John",
		Email:     "john@email.com",
		Age:       30,
		Hobbies:   []string{"Sports", "Cooking"},
		Marks:     50,
	}

	// Insert the document into the collection
	res, err := db.coll.InsertOne(ctx, u)
	if err != nil {
		log.Println(err)
		return
	}
	//db.coll.InsertMany()

	//inserted id
	fmt.Println(res.InsertedID)
}

// create a function the inserts multiple users record in one go

//InsertMany

// Get retrieves a single document based on a filter
func (db *DB) Get() {
	var person Person
	ctx := context.Background()
	//
	filter := bson.D{{"first_name", bson.D{{"$eq", "John"}}}}
	//filter := bson.D{
	//   {"$and",
	//      bson.A{
	//         bson.D{{"marks", bson.D{{"$gt", 7}}}},
	//         bson.D{{"age", bson.D{{"$lte", 30}}}},
	//      },
	//   },
	//}
	//filter := bson.D{{"first_name", "John"}}

	err := db.coll.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Found a single document: %+v\n", person)
}

// FindAll retrieves all documents
func (db *DB) FindAll() {
	var results []Person
	ctx := context.Background()

	// get everything
	// or specify a specific condition in bson.D{}
	filter := bson.D{}
	cur, err := db.coll.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			log.Println(err)
			return
		}
		results = append(results, person)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return
	}
	for _, v := range results {
		fmt.Printf("%+v\n\n", v)
	}
	//fmt.Println("Found multiple documents: ", results)
}

// Update modifies a single document based on a filter
func (db *DB) Update() {
	filter := bson.D{{"email", "john@email.com"}}
	update := bson.D{
		{"$set", bson.D{
			{"age", 32},
		}},
	}

	ctx := context.Background()
	res, err := db.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
}

// Delete removes a single document based on a filter
func (db *DB) Delete() {
	filter := bson.D{{"email", "john2@email.com"}}

	ctx := context.Background()
	res, err := db.coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Deleted %v document(s)\n", res.DeletedCount)
}

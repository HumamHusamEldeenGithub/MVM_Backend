package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MVMRepository struct {
	mongoDBClient *mongo.Client
	ctx           context.Context

	usersCollection   *mongo.Collection
	friendsCollection *mongo.Collection
}

func NewMVMRepository(ctx context.Context, db, pass string) *MVMRepository {
	dbClient := initMongoDBConnection(ctx, db, pass)
	return &MVMRepository{
		ctx:               ctx,
		mongoDBClient:     dbClient,
		usersCollection:   dbClient.Database("public").Collection("users"),
		friendsCollection: dbClient.Database("public").Collection("friends"),
	}
}

func initMongoDBConnection(ctx context.Context, dbTitle, pass string) *mongo.Client {
	fmt.Println("Connecting to MongoDB...")

	mongoCtx := context.Background()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	uri := fmt.Sprintf("mongodb+srv://%s:%s@mvm.8o7anpd.mongodb.net/?retryWrites=true&w=majority", dbTitle, pass)
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(uri))
	// Handle potential errors
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the connection was succesful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	return db
}

// // Define the unique index model
// index := mongo.IndexModel{
//     Keys:    bson.M{"email": 1},
//     Options: options.Index().SetUnique(true),
// }

// // Create the index
// _, err = collection.Indexes().CreateOne(context.Background(), index)
// if err != nil {
//     log.Fatal(err)
// }

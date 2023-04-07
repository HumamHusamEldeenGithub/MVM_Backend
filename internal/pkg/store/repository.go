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
}

func NewMVMRepository(ctx context.Context, db, pass string) *MVMRepository {
	return &MVMRepository{
		ctx:           ctx,
		mongoDBClient: initMongoDBConnection(ctx, db, pass),
	}
}

func initMongoDBConnection(ctx context.Context, dbTitle, pass string) *mongo.Client {
	// Dev
	// 6EHO7HJ9Zr2bG1Uw

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
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

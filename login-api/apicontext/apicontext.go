package apicontext

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	util "github.com/serdarkalayci/go-grpc/login-api/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// APIContext is the struct holds both Db Context and Handlers
type APIContext struct {
	MongoClient  mongo.Client
	DatabaseName string
	HandlerFunc  http.HandlerFunc
}

// PrepareContext prepares the database connection, bundles it within an APIContext and returns it
func PrepareContext() *APIContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	connectionString := os.Getenv("ConnectionString")
	if connectionString == "" {
		connectionString = "mongodb://root:example@localhost:27017"
		util.Logger.INFO.Log(fmt.Sprintf("ConnectionString from Env not found, falling back to local DB"))
	} else {
		util.Logger.INFO.Log(fmt.Sprintf("ConnectionString from Env is used: '%s'", connectionString))
	}
	databaseName := os.Getenv("DatabaseName")
	if databaseName == "" {
		databaseName = "cookiemonster"
		util.Logger.INFO.Log(fmt.Sprintf("DatabaseName from Env not found, falling back to default"))
	} else {
		util.Logger.INFO.Log(fmt.Sprintf("DatabaseName from Env is used: '%s'", databaseName))
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	apiContext := &APIContext{MongoClient: *client, DatabaseName: databaseName}
	return apiContext
}

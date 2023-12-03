package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
*
* ConnectDB()
	Create a ConnectDB function that first configures the client to use the correct URI
	and check for errors.
	Secondly, we defined a timeout of 10 seconds we wanted to use when trying to connect.
	Thirdly, check if there is an error while connecting to the database and cancel the connection if the connecting period exceeds 10 seconds.
	Finally, we pinged the database to test our connection and returned the client
	instance.
*/

func ConnectDB() *mongo.Client {
	// create an instance of mongo client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(LoadEnvironmentVariables()))
	if err != nil {
		log.Fatal("Could not create mongo client: ", err)
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}

	log.Println("Connected to MongoDB!")

	return client
}

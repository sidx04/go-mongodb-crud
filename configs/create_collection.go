package configs

import "go.mongodb.org/mongo-driver/mongo"

var DB *mongo.Client = ConnectDB()

/*
Create a DB variable instance of the ConnectDB. This will come in handy when creating collections.
Create a GetCollection function to retrieve and create collections on the database.
*/
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-test").Collection(collectionName)
	return collection
}

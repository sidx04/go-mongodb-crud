package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-mongodb-crud/configs"
	"github.com/go-mongodb-crud/models"
	"github.com/go-mongodb-crud/responses"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Create userCollection and validate variables to create a collection and validate
models using the github.com/go-playground/validator/v10 library

Create a CreateUser function that returns an error.
Inside the function, we first defined a timeout of 10 seconds when inserting user into
the document, validating both the request body and required field using the validator
library. We returned the appropriate message and status code using the UserResponse
struct we created earlier. Secondly, we created a newUser variable, inserted it using
the userCollection.InsertOne function and check for errors if there are any. Finally
we returned the correct response if the insert was successful.
*/

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c *gin.Context) {
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userModel models.User

	// bind to request body
	if err := c.Bind(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to read request body...",
			Data:    map[string]interface{}{"data": err.Error()},
		})
	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&userModel); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to read validate struct...",
			Data:    map[string]interface{}{"data": validationErr.Error()},
		})
	}

	// create a new user
	user := models.User{
		Name:     userModel.Name,
		Location: userModel.Location,
		Title:    userModel.Title,
	}

	// push to collection
	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create a new user...",
			Data:    map[string]interface{}{"data": err.Error()},
		})
	}

	// respond
	c.JSON(http.StatusCreated, responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "Successfully created user...",
		Data:    map[string]interface{}{"data": result},
	})

}

/*
Create a GetAUser function that returns an error. Inside the function, we first
defined a timeout of 10 seconds when finding a user in the document, a userId variable
to get the user’s id from the URL parameter and a user variable. We converted the
userId from a string to a primitive.ObjectID type, a BSON type MongoDB uses. Secondly,
we searched for the user using the userCollection.FindOne, pass the objId as a filter
and use the Decode attribute method to get the corresponding object. Finally, we
returned the decoded response.
*/

func GetUserById(c *gin.Context) {
	var userModel models.User

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get `_id` off param body: it is a string at this stage
	userId := c.Param("id")
	log.Println(userId)

	// wrap it into a mongoId object, viz, ObjectID("<userId>")
	objectId, _ := primitive.ObjectIDFromHex(userId)
	log.Println(objectId)

	// if ObjectID exists, decode it into a User struct
	err := userCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&userModel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to find user...",
			Data:    map[string]interface{}{"data": err.Error()},
		})
		return
	}

	// respond
	c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Found user successfully...",
		Data:    map[string]interface{}{"data": userModel},
	})
}

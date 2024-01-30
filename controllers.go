package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golangcompany/restfulapui/database"
	"github.com/golangcompany/restfulapui/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "User2")

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var User models.User

		if err := c.BindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		User.ID = primitive.NewObjectID()
		_, inserterr := UserCollection.InsertOne(ctx, User)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			c.Abort()
		}
		c.IndentedJSON(200, "user created successfully")
	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var User models.User

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Query("id")
		if id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			c.Abort()

		}

		fmt.Println("heelol")
		fmt.Println(id)

		ID1, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")

		}
		fmt.Println(ID1)

		err = UserCollection.FindOne(ctx, bson.M{"_id": ID1}).Decode(&User)

		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
			fmt.Println("Am I Here?")
			c.Abort()
		}
		c.IndentedJSON(200, User)
	}
}
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Query("id")
		if id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
			c.Abort()
			return
		}

		ID1, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		_, err = UserCollection.DeleteOne(ctx, bson.M{"_id": ID1})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		c.IndentedJSON(200, "Successfully deleted")
	}
}

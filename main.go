package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Username string
	Password string
	Name     string
	Ci       int
	Typed    string
	Active   bool
}

func dbConnection() *mongo.Client {
	//MONGO
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func main() {
	//GIN
	r := gin.Default()
	client := dbConnection()
	user := r.Group("/user")
	{
		user.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			collection := client.Database("tracker").Collection("users")
			var result User
			filter := bson.D{{"username", name}}
			err := collection.FindOne(context.TODO(), filter).Decode(&result)
			if err != nil {
				fmt.Printf("Document not found: %+v\n", result)
				c.JSON(http.StatusNotFound, result)
			} else {
				fmt.Printf("Found a single document: %+v\n", result)
				c.JSON(http.StatusOK, result)
			}

		})
	}
	r.Run(":3000")
}

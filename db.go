package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string
	Password string
	Name     string
	Ci       int
	Typ      string
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
func GetUserByUsername(c *gin.Context) {
	client := dbConnection()
	defer client.Disconnect(context.TODO())
	collection := client.Database("tracker").Collection("users")

	name := c.Param("username")
	var result User
	filter := bson.M{"username": name}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("Document not found: %+v\n", result)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		c.JSON(http.StatusOK, result)
	}

}
func NewUser(c *gin.Context) {
	client := dbConnection()
	defer client.Disconnect(context.TODO())
	collection := client.Database("tracker").Collection("users")
	var user User
	user.Id = primitive.NewObjectID()
	user.Name = c.PostForm("name")
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Ci, _ = strconv.Atoi(c.PostForm("ci"))
	user.Typ = c.PostForm("type")
	user.Active, _ = strconv.ParseBool(c.PostForm("active"))
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("Document not found: %+v\n", result)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		c.JSON(http.StatusOK, result)
	}
}
func GetUsers(c *gin.Context) {
	client := dbConnection()
	defer client.Disconnect(context.TODO())
	collection := client.Database("tracker").Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusNotFound, nil)
	}
	c.JSON(http.StatusOK, users)
}
func UpdateUserByUsername(c *gin.Context) {
	client := dbConnection()
	defer client.Disconnect(context.TODO())
	collection := client.Database("tracker").Collection("users")
	name := c.Param("username")
	var user User
	user.Name = c.PostForm("name")
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Ci, _ = strconv.Atoi(c.PostForm("ci"))
	user.Typ = c.PostForm("type")
	user.Active, _ = strconv.ParseBool(c.PostForm("active"))
	updatedUser := bson.M{"$set": bson.M{"name": user.Name,
		"username": user.Username,
		"password": user.Password,
		"ci":       user.Ci,
		"type":     user.Typ,
		"active":   user.Active}}
	result, err := collection.UpdateOne(context.TODO(), bson.M{"username": name}, updatedUser)
	if err != nil {
		fmt.Printf("Document not found: %+v\n", err)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		c.JSON(http.StatusOK, result)
	}

}
func DeleteUserByUsername(c *gin.Context) {
	client := dbConnection()
	defer client.Disconnect(context.TODO())
	name := c.Param("username")
	filter := bson.M{"username": name}
	collection := client.Database("tracker").Collection("users")
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Document not found: %+v\n", err)
		c.JSON(http.StatusNotFound, nil)
	} else {
		fmt.Printf("Found a single document: %+v\n", result)
		c.JSON(http.StatusOK, result)
	}
}

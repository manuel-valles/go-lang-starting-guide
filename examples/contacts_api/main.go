package main

import (
	"context"
	"fmt"
	"log"

	"example.com/contacts_api/controllers"
	"example.com/contacts_api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      	*gin.Engine
	userService     services.UserService
	userController  controllers.UserController
	ctx         	context.Context
	userCollection  *mongo.Collection
	mongoClient 	*mongo.Client
	err         	error
)

func init() {
	ctx = context.TODO()

	mongoConnection := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConnection)

	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userCollection = mongoClient.Database("userdb").Collection("users")
	userService = services.NewUserServiceImpl(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	version := server.Group("/v1")
	userController.RegisterUserRoutes(version)

	log.Fatal(server.Run(":9000"))

}

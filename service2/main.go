package main

import (
	"context"
	"fmt"
	"log"

	controller "example.com/m/controllers"

	"example.com/m/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controller.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	productcollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
	productservice services.ProductService
   productController controller.ProductController
	
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	usercollection = mongoclient.Database("userdb").Collection("users")
	productcollection = mongoclient.Database("userdb").Collection("products")

	userservice = services.NewUserService(usercollection, ctx)
	productservice = services.NewProductService(productcollection ,ctx)
	usercontroller = controller.Newuser(userservice)
	productController = controller.Newproduct(productservice)

	server = gin.Default()
}

// v1/user/create
func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
    productController.RegisterProductRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}

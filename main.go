package main

import (
	"fmt"
	"os"
	"time"

	"iris-demo/controllers"

	"github.com/iris-contrib/graceful"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	//Init all the settings
	initialize(api)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "7000"
	}

	fmt.Println("Service is listening at:" + port)
	graceful.Run(":"+port, time.Duration(10)*time.Second, api)
}

func initialize(api *iris.Framework) {
	api.Use(logger.New())
	api.Use(recovery.Handler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "ACCEPT", "ORIGIN"},
		AllowCredentials: true,
		Debug:            true,
	})

	api.Use(c)

	//Init controller
	defController := &controllers.DefaultController{}
	userController := &controllers.UserController{}

	//Register the default API handler
	defController.Get("/", defController.DefaultHandler)

	//Use party to manage the prefix url
	user := api.Party("/user")
	user.Get("/profile", userController.GetUserProfile)
	user.Post("/login", userController.HandleLogin)
}

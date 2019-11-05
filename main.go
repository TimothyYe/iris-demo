package main

import (
	"os"

	"iris-demo/controllers"
	"iris-demo/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	api := iris.New()

	// Init all the settings
	initialize(api)

	// Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "7000"
	}

	// It will shutdown gracefully on CTRL/CMD+C.
	api.Run(iris.Addr(":" + port))
}

func initialize(api *iris.Application) {
	api.Use(logger.New())
	api.Use(recover.New())

	corsMiddleware := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, ACCEPT, Authorization")
		ctx.Next()
	} // or	"github.com/iris-contrib/middleware/cors"
	api.Use(corsMiddleware)
	api.AllowMethods(iris.MethodOptions)

	// Init custom controllers.
	//
	// Note: If you wish to take advandage of the Iris MVC feature
	// please navigate through https://github.com/kataras/iris/wiki/MVC
	userController := &controllers.UserController{}

	// Register the default API handler
	api.Get("/", controllers.DefaultHandler)

	// Use party to manage the prefix url
	user := api.Party("/user")
	user.Get("/profile", middleware.TokenMiddleware, userController.GetUserProfile)
	user.Post("/login", userController.HandleLogin)
}

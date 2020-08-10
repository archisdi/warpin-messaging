package main

import (
	"log"
	"os"
	"warpin/controllers"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading environment file")
	}

	// create new app instance
	app := iris.New()
	mvc.Configure(app.Party("/"), initController)

	// serve app
	app.Listen(":" + os.Getenv("APP_PORT"))
}

func initController(app *mvc.Application) {
	app.Handle(new(controllers.MessageController))
}

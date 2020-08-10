package main

import (
	"log"
	"net/url"
	"os"
	"warpin/controllers"
	mqtt "warpin/libs/mqtt"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initController(app *mvc.Application) {
	app.Handle(new(controllers.MessageController))
}

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading environment file")
	}

	// connect to mqtt server
	mqttURL, _ := url.Parse(os.Getenv("MQTT_URL"))
	mqtt.Initialize("", mqttURL)

	// create new app instance
	app := iris.New()
	mvc.Configure(app.Party("/"), initController)

	// serve app
	port := os.Getenv("APP_PORT")
	app.Run(iris.Addr(":" + port))
}

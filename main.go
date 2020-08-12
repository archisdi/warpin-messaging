package main

import (
	"log"
	"os"
	"warpin/controllers"
	mqtt "warpin/libs/mqtt"
	repo "warpin/repositories"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initialize(app *mvc.Application) {
	app.Handle(&controllers.MessageController{
		MessageRepo: &repo.MessageRepo{},
		Mqtt:        &mqtt.Mqtt{},
	})

	app.HandleError(func(ctx iris.Context, err error) {
		ctx.JSON(map[string]interface{}{
			"message": err.Error(),
		})
	})
}

func main() {
	// load environment variables
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("error while loading environment file")
		os.Exit(1)
	}

	// connect to mqtt server
	if mqttErr := mqtt.Initialize("", os.Getenv("MQTT_URL")); mqttErr != nil {
		log.Fatal("error while connecting to mqtt server")
		os.Exit(1)
	}

	// create new app instance
	app := iris.New()
	mvc.Configure(app.Party("/"), initialize)

	// serve app
	port := os.Getenv("APP_PORT")
	app.Run(iris.Addr(":"+port), iris.WithoutBodyConsumptionOnUnmarshal)
}

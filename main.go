package main

import (
	"log"
	"net/url"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading environment file")
	}

	// connect to mqtt server
	mqttURL, _ := url.Parse(os.Getenv("MQTT_URL"))
	mqtt.Initialize("", mqttURL)

	// create new app instance
	app := iris.New()
	mvc.Configure(app.Party("/"), initialize)

	// serve app
	port := os.Getenv("APP_PORT")
	app.Run(iris.Addr(":"+port), iris.WithoutBodyConsumptionOnUnmarshal)
}

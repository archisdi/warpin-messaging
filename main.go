package main

import (
	"warpin/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	mvc.Configure(app.Party("/"), initController)
	app.Listen(":8080")
}

func initController(app *mvc.Application) {
	app.Handle(new(controllers.MessageController))
}

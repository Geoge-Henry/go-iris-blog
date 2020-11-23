package route

import (
	"go-web/api"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.Configure(app.Party("/henry_blog"), BlogHandler)
	app.Get("/", api.HomePage)
	app.Post("/", api.PingPage)
}

func BlogHandler(app *mvc.Application) {
	app.Handle(api.MakeAuthController())
	app.Party("/user", api.AuthHandler).Handle(api.MakeUserController())
}

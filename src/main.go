package main

import (
	"go-web/api"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./public/templates", ".html").Reload(true))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.HandleDir("/", "./public/templates")
	app.HandleDir("/img", "./public/img")
	app.HandleDir("/css", "./public/css")
	app.HandleDir("/js", "./public/js")

	app.Post("/", func(ctx iris.Context) {
		ctx.StatusCode(200)
		ctx.WriteString("ok")
	})
	mvc.Configure(app.Party("/henry_blog"), api.BlogHandler)
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/config.yml")))
}

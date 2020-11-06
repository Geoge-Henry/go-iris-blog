package main

import (
	"go-web/api"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Hello Iris </h1>")
	})
	app.Post("/", func(ctx iris.Context) {
		ctx.StatusCode(200)
		ctx.WriteString("ok")
	})
	mvc.Configure(app.Party("/henry_blog"), api.BlogHandler)
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/config.yml")))
}

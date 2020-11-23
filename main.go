package main

import (
	"flag"
	"go-web/config"
	"go-web/route"

	"github.com/kataras/iris/v12"
)

func main() {
	flag.Parse()
	app := newApp()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+config.Setting.Port), iris.WithConfiguration(iris.YAML("./config/serverConfig.yml")))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)
	app.RegisterView(iris.HTML("./public/templates", ".html").Reload(true))
	app.HandleDir("/", "./public/templates")
	app.HandleDir("/img", "./public/img")
	app.HandleDir("/css", "./public/css")
	app.HandleDir("/js", "./public/js")
	return app
}

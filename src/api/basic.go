package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func myAuthMiddlewareHandler(ctx iris.Context) {
	ctx.WriteString("Authentication start")
}

func BlogHandler(app *mvc.Application) {
	app.Party("/rr").Handle(new(MyController))
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) GetHello(ctx iris.Context) string {
	return "Authentication failed"
}

func (m *MyController) AnyEe(ctx iris.Context) string {
	return "Authentication failed"
}

func (m *MyController) MyCustomHandler(id int64) string { return "MyCustomHandler says Hey" }

func userProfileHandler(ctx iris.Context) { //
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func userMessageHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

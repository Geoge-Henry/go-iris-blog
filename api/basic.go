package api

import (
	"github.com/kataras/iris/v12"
)

func HomePage(ctx iris.Context) {
	ctx.View("index.html")
}

func PingPage(ctx iris.Context) {
	ctx.StatusCode(200)
	ctx.WriteString("ok")
}


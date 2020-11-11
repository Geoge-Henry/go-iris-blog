package api

import (
	"go-web/service"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Service service.UserService
}

func MakeUserController() *UserController {
	return &UserController{Service: service.MakeUserService()}
}

func (self *UserController) AnyWe(ctx iris.Context) string {
	return "Authentication success"
}

// func userMessageHandler(ctx iris.Context) {
// 	id := ctx.Params().Get("id")
// 	ctx.WriteString(id)
// }

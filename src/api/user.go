package api

import (
	"github.com/kataras/iris/v12"
	"go-web/models"
	"go-web/service"
)

type UserController struct{
	Service service.UserService
}

func MakeUserController() *UserController {
	return &UserController{Service:service.MakeUserService()}
}

func (self *UserController) PostLogin(ctx iris.Context) string {
	user := self.userProfileHandler(ctx)
	if user != nil && len(user) > 0{
		return "Authentication success"
	}
	return "Authentication failed"
}

func (self *UserController) userProfileHandler(ctx iris.Context)[]models.User { //
	userName := ctx.PostValue("user_name")
	pwd := ctx.PostValue("pwd")
	ctx.WriteString(userName + pwd)
	return *self.Service.GetUserByLogin(userName, pwd)
}

// func userMessageHandler(ctx iris.Context) {
// 	id := ctx.Params().Get("id")
// 	ctx.WriteString(id)
// }

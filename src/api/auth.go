package api

import (
	"fmt"
	"go-web/models"
	"go-web/service"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "loginsessionid"
	SessionObj             = sessions.New(sessions.Config{
		Cookie:  cookieNameForSessionID,
		Expires: 60 * time.Minute})
	redisObj = &service.RedisService{}
	redisDb  = redisObj.GetSessionDb()
)

func init() {
	SessionObj.UseDatabase(redisDb)
}

func AuthHandler(ctx iris.Context) {
	fmt.Print("123")
	ctx.Next()
}

type AuthController struct {
	Service service.UserService
}

func MakeAuthController() *AuthController {
	return &AuthController{Service: service.MakeUserService()}
}

func (self *AuthController) PostLogin(ctx iris.Context) string {
	user := self.userProfileHandler(ctx)
	if user != nil && len(user) > 0 {
		currentUser := user[0]
		session := SessionObj.Start(ctx)
		session.Set("loginName", currentUser.Name)
		fmt.Print(session.Get("loginName"))
		return "Authentication success"
	}
	return "Authentication failed"
}

func (self *AuthController) PostLoginout(ctx iris.Context) string {
	user := self.userProfileHandler(ctx)
	if user != nil && len(user) > 0 {
		session := SessionObj.Start(ctx)
		session.Delete("loginName")
		return "Authentication success"
	}
	return "Authentication failed"
}

func (self *AuthController) userProfileHandler(ctx iris.Context) []models.User { //
	userName := ctx.PostValue("user_name")
	pwd := ctx.PostValue("pwd")
	ctx.WriteString(userName + pwd)
	return *self.Service.GetUserByLogin(userName, pwd)
}

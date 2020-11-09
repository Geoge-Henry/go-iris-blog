package service
 
import (
	"strconv"
	"go-web/datasource"
	"go-web/models"
)

type UserService interface {
	PostSaveUser(user models.User)map[string]string
	GetUserByLogin(name string, pwd string )*[]models.User
	// DelUser(id uint) map
}
 
type userService struct {}

var db = datasource.GetDB()
 
func MakeUserService() UserService{
	return &userService{}
}
 
func (self userService) PostSaveUser(user models.User) map[string]string{
	var result map[string]string
	code, msg := 0, ""

	err := db.Save(&user).Error

	if err != nil{
		code = -1
		msg = err.Error()
	}else{
		code = int(user.Id)
		msg = "SUCCESS"
	}
	result["code"] = strconv.Itoa(code)
	result["msg"] = msg
	return result
}
func (self userService) GetUserByLogin(name string, pwd string)*[]models.User{
	user:= new([]models.User)
	db.Raw(`select * FROM user where user.name = ? and user.password = ?`, name, pwd).Scan(&user)
	return user
}
// func (self userService) DelUser(id uint)(result models.Result){
// 	err := userRepo.DeleteUser(id)
// 	if err!= nil{
// 		result.Code = 400
// 		result.Msg = err.Error()
// 	}else{
// 		result.Code = 200
// 		result.Msg ="SUCCESS"
// 		list := userRepo.GetUserList()
// 		result.Data = list
 
// 	}
// 	return
// }
 
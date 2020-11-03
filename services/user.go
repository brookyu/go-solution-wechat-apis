package services

import (
	"github.com/brookyu/go-solution-wechat-apis/apiHelpers"
	models "github.com/brookyu/go-solution-wechat-apis/models"
	"github.com/brookyu/go-solution-wechat-apis/resources/api"
)

type UserService struct {
	User models.User
}

//UserList function returns the list of users
func (us *UserService) UserList() map[string]interface{} {
	user := us.User

	userData := api.UserResponse{
		ID:    user.ID,
		Name:  "test",
		Email: "test@gmail.com",
	}
	response := apiHelpers.Message(0, "This is from version 2 api")
	response["data"] = userData
	return response
}

package api

import (
	"encoding/json"
	helpers "github.com/brookyu/go-solution-wechat-apis/apiHelpers"
	services "github.com/brookyu/go-solution-wechat-apis/services"
	"github.com/gin-gonic/gin"
)
// UserList godoc
// @Summary Get user list
// @Description get user list
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} UserResponse
// @Router /User [get]
func UserList(c *gin.Context) {
	var userService services.UserService
	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		helpers.Respond(c.Writer, helpers.Message(1, "Invalid request"))
		return
	}
	resp := userService.UserList()
	helpers.Respond(c.Writer, resp)
}

package service

import (
	"fmt"
	"github.com/daviresio/financeiro_api/api_error"
	"github.com/daviresio/financeiro_api/config"
	"github.com/daviresio/financeiro_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(c *gin.Context) {
	var users []model.User

	config.DB.Preload("Address").Find(&users)

	for i, user := range users {
		user.RemovePassword()
		users[i] = user
	}

	token, _ := GenerateJWT()

	fmt.Println(c.GetHeader("authorization"))
	fmt.Println(token)

	c.JSON(http.StatusOK, users)
}

func FindUser(c *gin.Context) {
	var user model.User

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := config.DB.First(&user, id).RecordNotFound()

	if notFound == true {
		resErr := api_error.NewNotFoundError(fmt.Sprintf("user with id %s not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	user.RemovePassword()

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user model.User

	if err := CheckInvalidJson(c.ShouldBindJSON(&user), c); err != nil {
		return
	}

	config.DB.Create(&user)

	c.JSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {
	var user model.User

	if err := CheckInvalidJson(c.ShouldBindJSON(&user), c); err != nil {
		return
	}

	var originalUser model.User
	config.DB.First(&originalUser, user.Id)

	if user.Password != "" {
		user.Password = model.EncryptPassword(user.Password)
	} else {
		user.Password = originalUser.Password
	}

	config.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	user := model.User{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}
	
	config.DB.Delete(&user)

	c.Status(http.StatusOK)
}
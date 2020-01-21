package service

import (
	"fmt"
	"github.com/daviresio/financeiro_api/api_error"
	"github.com/daviresio/financeiro_api/config"
	"github.com/daviresio/financeiro_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListTag(c *gin.Context) {
	var tags []model.Tag

	err := config.DB.Preload("Notebooks").Find(&tags).GetErrors()
		fmt.Println(err)
	c.JSON(http.StatusOK, tags)
}

func FindTag(c *gin.Context) {
	var tag model.Tag

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := config.DB.Preload("Notebooks").First(&tag, id).RecordNotFound()

	if notFound == true {
		resErr := api_error.NewNotFoundError(fmt.Sprintf("tag with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, tag)

}

func CreateTag(c *gin.Context) {
	var tag model.Tag

	if err := CheckInvalidJson(c.ShouldBindJSON(&tag), c); err != nil {
		return
	}

	config.DB.Create(&tag)

	c.JSON(http.StatusOK, tag)

}

func UpdateTag(c *gin.Context) {
	var tag model.Tag

	if err := CheckInvalidJson(c.ShouldBindJSON(&tag), c); err != nil {
		return
	}

	var originalTag model.Tag
	config.DB.First(&originalTag, tag.Id)

	config.DB.Save(&tag)

	c.JSON(http.StatusOK, tag)
}

func DeleteTag(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	tag := model.Tag{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	config.DB.Delete(&tag)

	c.Status(http.StatusOK)
}
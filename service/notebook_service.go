package service

import (
	"fmt"
	"github.com/daviresio/financeiro_api/api_error"
	"github.com/daviresio/financeiro_api/config"
	"github.com/daviresio/financeiro_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListNotebook(c *gin.Context) {
	var books []model.Notebook

	config.DB.Preload("Book").Preload("Tags").Find(&books)

	c.JSON(http.StatusOK, books)
}

func FindNotebook(c *gin.Context) {
	var notebook model.Notebook

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := config.DB.Preload("Book").Preload("Tags").First(&notebook, id).RecordNotFound()

	if notFound == true {
		resErr := api_error.NewNotFoundError(fmt.Sprintf("notebook with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, notebook)

}

func CreateNotebook(c *gin.Context) {
	var notebook model.Notebook

	if err := CheckInvalidJson(c.ShouldBindJSON(&notebook), c); err != nil {
		return
	}

	config.DB.Create(&notebook)

	c.JSON(http.StatusOK, notebook)

}

func UpdateNotebook(c *gin.Context) {
	var notebook model.Notebook

	if err := CheckInvalidJson(c.ShouldBindJSON(&notebook), c); err != nil {
		return
	}

	var originalNotebook model.Notebook
	config.DB.First(&originalNotebook, notebook.Id)

	config.DB.Save(&notebook)

	c.JSON(http.StatusOK, notebook)
}

func DeleteNotebook(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notebook := model.Notebook{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	config.DB.Delete(&notebook)

	c.Status(http.StatusOK)
}
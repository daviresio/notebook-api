package service

import (
	"fmt"
	"github.com/daviresio/financeiro_api/api_error"
	"github.com/daviresio/financeiro_api/config"
	"github.com/daviresio/financeiro_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListBook(c *gin.Context) {
	var books []model.Book

	err := config.DB.Preload("Notebooks").Find(&books).GetErrors()
		fmt.Println(err)
	c.JSON(http.StatusOK, books)
}

func FindBook(c *gin.Context) {
	var book model.Book

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	notFound := config.DB.Preload("Notebooks").First(&book, id).RecordNotFound()

	if notFound == true {
		resErr := api_error.NewNotFoundError(fmt.Sprintf("book with id %d not found", id))
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, book)

}

func CreateBook(c *gin.Context) {
	var book model.Book

	if err := CheckInvalidJson(c.ShouldBindJSON(&book), c); err != nil {
		return
	}

	config.DB.Create(&book)

	c.JSON(http.StatusOK, book)

}

func UpdateBook(c *gin.Context) {
	var book model.Book

	if err := CheckInvalidJson(c.ShouldBindJSON(&book), c); err != nil {
		return
	}

	var originalBook model.Book
	config.DB.First(&originalBook, book.Id)

	config.DB.Save(&book)

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {

	id := GetIdParam("id", c)

	if id == 0 {
		return
	}

	book := model.Book{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}

	config.DB.Delete(&book)

	c.Status(http.StatusOK)
}
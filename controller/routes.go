package controller

import (
	"github.com/daviresio/financeiro_api/service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func MapUrls() {
	router.GET("/users", service.ListUser)
	router.GET("/users/:id", service.FindUser)
	router.POST("/users", service.CreateUser)
	router.PUT("/users", service.UpdateUser)
	router.DELETE("/users/:id", service.DeleteUser)

	router.GET("/notebooks", service.ListNotebook)
	router.GET("/notebooks/:id", service.FindNotebook)
	router.POST("/notebooks", service.CreateNotebook)
	router.PUT("/notebooks", service.UpdateNotebook)
	router.DELETE("/notebooks/:id", service.DeleteNotebook)

	router.GET("/books", service.ListBook)
	router.GET("/books/:id", service.FindBook)
	router.POST("/books", service.CreateBook)
	router.PUT("/books", service.UpdateBook)
	router.DELETE("/books/:id", service.DeleteBook)

	router.GET("/tags", service.ListTag)
	router.GET("/tags/:id", service.FindTag)
	router.POST("/tags", service.CreateTag)
	router.PUT("/tags", service.UpdateTag)
	router.DELETE("/tags/:id", service.DeleteTag)
}

func GetRouter() *gin.Engine {
	return router
}
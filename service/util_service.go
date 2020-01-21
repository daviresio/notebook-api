package service

import (
	"fmt"
	"github.com/daviresio/financeiro_api/api_error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckInvalidJson(err error, c *gin.Context) error {
	if err != nil {
		resErr := api_error.NewBadRequestError("invalid json body")
		c.JSON(resErr.Status, resErr)
		return err
	}
	return nil
}

func GetIdParam(key string, c *gin.Context) uint64 {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		resErr := api_error.NewBadRequestError(fmt.Sprintf("%s is not a valid number", idParam))
		c.JSON(resErr.Status, resErr)
		return 0
	}

	if id == 0 {
		resErr := api_error.NewNotFoundError("id cannot be 0")
		c.JSON(resErr.Status, resErr)
		return 0
	}

	return id
}
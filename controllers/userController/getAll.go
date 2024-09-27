package userController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *UserController) GetAll(ctx echo.Context) error {

	response, err := uc.Service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"is_success": false,
			"message":    "Failed to get all user",
			"error":      err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"is_success": true,
		"message":    "Success get all user",
		"data":       response,
	})
}

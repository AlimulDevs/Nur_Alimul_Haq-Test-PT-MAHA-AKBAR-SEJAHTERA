package userController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *UserController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	response, err := uc.Service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"is_success": false,
			"message":    "Failed to get user by id",
			"error":      err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"is_success": true,
		"message":    "Success get user by id",
		"data":       response,
	})
}

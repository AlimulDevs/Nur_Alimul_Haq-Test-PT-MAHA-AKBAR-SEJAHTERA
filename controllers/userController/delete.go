package userController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *UserController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := uc.Service.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"is_success": false,
			"message":    "Failed to delete user",
			"error":      err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"is_success": true,
		"message":    "Success delete user",
	})
}

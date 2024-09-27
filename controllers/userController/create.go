package userController

import (
	"maha-akbar-sejahtera/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *UserController) Create(ctx echo.Context) error {
	input := new(dto.UserRequest)
	err := ctx.Bind(&input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"is_success": false,
			"message":    "failed bind data",
			"error":      err.Error(),
		})
	}
	err = uc.Service.Create(*input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"is_success": false,
			"message":    "Failed to create user",
			"error":      err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"is_success": true,
		"message":    "Success create user",
	})
}

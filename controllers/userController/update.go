package userController

import (
	"maha-akbar-sejahtera/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *UserController) Update(ctx echo.Context) error {
	input := new(dto.UserRequest)
	id := ctx.Param("id")
	err := ctx.Bind(&input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"is_success": false,
			"message":    "failed bind data",
			"error":      err.Error(),
		})
	}
	err = uc.Service.Update(id, *input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"is_success": false,
			"message":    "Failed to update user",
			"error":      err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"is_success": true,
		"message":    "Success update user",
	})
}

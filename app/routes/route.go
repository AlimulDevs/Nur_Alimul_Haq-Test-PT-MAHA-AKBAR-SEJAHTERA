package routes

import (
	userController "maha-akbar-sejahtera/controllers/userController"
	"maha-akbar-sejahtera/repositories/userRepository"
	"maha-akbar-sejahtera/services/userService"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteService(db *gorm.DB) *echo.Echo {
	app := echo.New()

	api := app.Group("/api/v1")

	userRouteRepository := userRepository.NewUserRepositoy(db)
	userService := userService.NewUserService(userRouteRepository)
	userRouteController := userController.UserController{
		Service: userService,
	}
	api.GET("/users", userRouteController.GetAll)
	api.GET("/user/:id", userRouteController.GetByID)
	api.POST("/user", userRouteController.Create)
	api.PUT("/user/:id", userRouteController.Update)
	api.DELETE("/user/:id", userRouteController.Delete)

	return app
}

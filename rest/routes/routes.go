package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/motephyr/longcare/rest/controllers"
)

func LoadRoutes(app *fiber.App) {
	// api := app.Group("api").Use(middlewares.AuthApi())
	// web := app.Group("")
	// ApiRoutes(api)
	// WebRoutes(web)

	app.Get("/", controllers.Index)

}

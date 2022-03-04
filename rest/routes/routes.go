package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/motephyr/longapp/rest/controllers"
	"github.com/motephyr/longapp/rest/middlewares"
)

func LoadRoutes(app *fiber.App) {
	app.Get("/auth/login", controllers.AuthController.LoginPage)
	app.Post("/auth/login",
		middlewares.ValidateLoginPost, controllers.AuthController.Login)
	app.Post("/auth/logout",
		controllers.AuthController.Logout,
	)

	app.Get("/",
		controllers.IndexController.IndexPage)

	app.Get("/addNewWallet",
		controllers.IndexController.AddNewWalletPage)

}

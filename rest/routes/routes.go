package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/motephyr/longcare/rest/controllers"
	"github.com/motephyr/longcare/rest/middlewares"
)

func LoadRoutes(app *fiber.App) {
	// api := app.Group("api").Use(middlewares.AuthApi())
	// web := app.Group("")
	// ApiRoutes(api)
	// WebRoutes(web)

	users := app.Group("/manage/users")
	users.Get("/", controllers.UserController.Index)
	users.Get("/new", controllers.UserController.New)
	users.Get("/:id/edit", controllers.UserController.Edit)
	users.Get("/:id", controllers.UserController.Show)
	users.Post("/", controllers.UserController.Create)
	users.Post("/:id", controllers.UserController.Update)
	users.Delete("/:id", controllers.UserController.Delete)

	app.Get("/auth/login", controllers.AuthController.LoginPage)
	app.Post("/auth/login",
		middlewares.ValidateLoginPost, controllers.AuthController.Login)
	app.Post("/auth/logout",
		controllers.AuthController.Logout,
	)
	app.Use(middlewares.AuthWeb())

	app.Get("/", controllers.IndexController.IndexPage)
	app.Get("/list", controllers.IndexController.IndexPage)
	app.Get("/nurse", controllers.IndexController.Nurse)
	app.Get("/manage", controllers.IndexController.Manage)

	app.Get("/nurse/:datestring", controllers.NurseController.Nurse)
	app.Post("/nurse/:datestring/groups/:group_id", controllers.NurseController.SetOlderId)

	app.Get("/manage/:datestring", controllers.ManageController.Manage)
	app.Post("/manage/:datestring/createGroup", controllers.ManageController.CreateGroup)
	app.Post("/manage/:datestring/resetData", controllers.ManageController.ResetData)
	app.Post("/manage/:datestring/deleteData", controllers.ManageController.DeleteData)

	sources := app.Group("sources")
	sources.Get("/", controllers.SourceController.List)
	sources.Get("/:id", controllers.SourceController.Find)
	sources.Post("/", controllers.SourceController.Create)
	sources.Put("/:id", controllers.SourceController.Update)
	sources.Delete("/:id", controllers.SourceController.Delete)

	olders := app.Group("olders")
	olders.Get("/", controllers.OlderController.Index)
	olders.Get("/new", controllers.OlderController.New)
	olders.Get("/:id/edit", controllers.OlderController.Edit)
	olders.Get("/:id", controllers.OlderController.Show)
	olders.Post("/", controllers.OlderController.Create)
	olders.Post("/:id", controllers.OlderController.Update)
	olders.Delete("/:id", controllers.OlderController.Delete)

}

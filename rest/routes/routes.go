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

	app.Get("/", controllers.IndexController.Index)
	app.Get("/list", controllers.IndexController.Index)
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

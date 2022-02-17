package routes

import (
	"github.com/gofiber/fiber/v2"
	self "github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/rest/controllers"
	"github.com/motephyr/longcare/rest/middlewares"
)

func LoadRoutes(app *fiber.App) {
	// api := app.Group("api").Use(middlewares.AuthApi())
	// web := app.Group("")
	// ApiRoutes(api)
	// WebRoutes(web)
	app.Post("/sources/upload", controllers.SourceController.Upload)
	app.Get("mynotice", func(c *fiber.Ctx) error {
		self.Http.Websocket.Server.BroadcastToNamespace("", "notice_1", "AAA")
		return c.JSON("ok")
	})
	users := app.Group("/manage/users")
	users.Get("/", controllers.UserController.Index)
	users.Get("/new", controllers.UserController.New)
	users.Get("/:id/edit", controllers.UserController.Edit)

	users.Post("/", controllers.UserController.Create)
	users.Post("/:id", controllers.UserController.Update)
	users.Post("/:id/delete", controllers.UserController.Delete)

	groups := users.Group("/:id/groups")

	groups.Get("/", controllers.GroupController.Index)
	groups.Get("/:datestring", controllers.GroupController.Manage)
	groups.Post("/:datestring/createGroup", controllers.GroupController.CreateGroup)
	groups.Post("/:datestring/resetData", controllers.GroupController.ResetData)
	groups.Post("/:datestring/deleteData", controllers.GroupController.DeleteData)

	idstring := users.Group("/:id/idstrings")
	idstring.Get("/", controllers.UserIdstringController.Index)
	idstring.Get("/new", controllers.UserIdstringController.New)
	idstring.Post("/", controllers.UserIdstringController.Create)
	idstring.Post("/:idstring/delete", controllers.UserIdstringController.Delete)

	app.Get("/auth/login", controllers.AuthController.LoginPage)
	app.Post("/auth/login",
		middlewares.ValidateLoginPost, controllers.AuthController.Login)
	app.Post("/auth/logout",
		controllers.AuthController.Logout,
	)
	app.Use(middlewares.AuthWeb()) // 驗證token是否正確
	app.Use(middlewares.GetUser)   //

	app.Get("/",
		controllers.IndexController.IndexPage)
	app.Get("/list", controllers.IndexController.IndexPage)
	app.Get("/nurse", controllers.NurseController.Index)
	app.Get("/nurse/:datestring", controllers.NurseController.Nurse)
	app.Post("/nurse/:datestring/groups/:group_id", controllers.NurseController.SetOlderId)

	// sources := app.Group("sources")
	// sources.Get("/", controllers.SourceController.List)
	// sources.Get("/:id", controllers.SourceController.Find)
	// sources.Post("/", controllers.SourceController.Create)
	// sources.Post("/:id", controllers.SourceController.Update)
	// sources.Post("/:id/delete", controllers.SourceController.Delete)

	olders := app.Group("olders")
	olders.Get("/query", controllers.OlderController.Query)
	olders.Get("/query/:id", controllers.OlderController.Query)

	olders.Get("/", controllers.OlderController.Index)
	olders.Get("/new", controllers.OlderController.New)
	olders.Get("/:id/edit", middlewares.CanUserUpdateOlder, controllers.OlderController.Edit)
	olders.Get("/:id", middlewares.CanUserUpdateOlder, controllers.OlderController.Show)
	olders.Post("/", controllers.OlderController.Create)
	olders.Post("/:id", middlewares.CanUserUpdateOlder, controllers.OlderController.Update)
	olders.Post("/:id/delete", middlewares.CanUserUpdateOlder, controllers.OlderController.Delete)

}

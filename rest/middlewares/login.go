package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/models"
	"github.com/motephyr/longcare/utils"

	auth "github.com/motephyr/longcare/pkg/auth"
)

func RedirectToHomePageOnLogin(c *fiber.Ctx) error {
	// if auth.IsLoggedIn(c) {
	// 	return c.Redirect("/")
	// }
	return c.Next()
}

func ValidateLoginPost(c *fiber.Ctx) error {
	var login models.User
	if err := c.BodyParser(&login); err != nil {
		inertia.Share(fiber.Map{
			"flash": map[string]any{
				"message": err.Error(),
			},
		})
		return c.Redirect("/auth/login")
	}
	v := validate.Struct(login)
	if !v.Validate() {
		inertia.Share(fiber.Map{
			"flash": map[string]any{
				"message": v.Errors.One(),
			},
		})
		return c.Redirect("/auth/login")
	}
	user, err := auth.CheckLogin(login) //nolint:wsl

	if err != nil {
		inertia.Share(fiber.Map{
			"flash": map[string]any{
				"message": err.Error(),
			},
		})
		return c.Redirect("/auth/login")
	}

	c.Locals("user", user)

	return c.Next()
}

func GetUser(c *fiber.Ctx) error {
	user, err := auth.User(c)
	if err != nil {
		return c.Redirect("/auth/login")
	}
	c.Locals("user", user)
	inertia.Share(fiber.Map{
		"user": utils.StructToMap(*user),
	})

	GetFlashAndClean(c)

	return c.Next()
}

func GetFlashAndClean(c *fiber.Ctx) {
	store := app.Http.Session.Get(c) // get/create new session

	flash := store.Get("flash")
	inertia.Share(fiber.Map{
		"flash": flash,
	})
	store.Delete("flash")
	if err := store.Save(); err != nil {
		panic(err)
	}
}

func SetFlashMessage(c *fiber.Ctx, message string) {
	store := app.Http.Session.Get(c)

	store.Set("flash", map[string]any{
		"message": message,
	})
	if err := store.Save(); err != nil {
		panic(err)
	}
}

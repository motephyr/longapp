package controllers

import (
	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longapp/app"
	"github.com/motephyr/longapp/models"
	"github.com/motephyr/longapp/pkg/auth"
)

type authControler struct{}

var AuthController authControler

func (authControler) LoginPage(c *fiber.Ctx) error {
	return inertia.Render(c, "auth/Login.vue", inertia.Map{})
}

func (authControler) Login(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	_, err := auth.Login(c, user.ID, app.Http.Token.AppJwtSecret) //nolint:wsl
	if err != nil {
		inertia.Share(inertia.Map{
			"flash": map[string]any{
				"message": err.Error(),
			},
		})

		return c.Redirect("auth/login")
	}

	return c.Redirect("/")
}

func (authControler) Logout(c *fiber.Ctx) error { //nolint:nolintlint,wsl
	if auth.IsLoggedIn(c) {
		err := auth.Logout(c)
		if err != nil {
			panic(err)
		}
	}
	c.Set("X-DNS-Prefetch-Control", "off")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "Fri, 01 Jan 1990 00:00:00 GMT")
	c.Set("Cache-Control", "no-cache, must-revalidate, no-store, max-age=0, private")
	return c.Redirect("/auth/login")
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
)

type indexController struct{}

var IndexController indexController

func (indexController) IndexPage(c *fiber.Ctx) error {

	return inertia.Render(c,
		"Index.vue", // Will render component named as Main
		inertia.Map{},
	)
	// return c.SendStatus(200)

}

func (indexController) AddNewWalletPage(c *fiber.Ctx) error {

	return inertia.Render(c,
		"AddNewWallet.vue", // Will render component named as Main
		inertia.Map{},
	)
	// return c.SendStatus(200)

}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/utils"
)

func Index(c *fiber.Ctx) error {
	return inertia.Render(c,
		"Index.vue", // Will render component named as Main
		inertia.Map{
			"Hi-EN":       "Hello World",
			"name":        "ok",
			"important":   utils.Map{"name": "important", "notice": "noticeme"},
			"rooms":       utils.Map{"A": "5", "B": "3"},
			"datestrings": utils.Slice{"123", "456"},
			"sumetimes":   utils.Slice{utils.Map{"name": "ok", "sumtime": utils.Slice{"1"}}, utils.Map{"name": "ok2", "sumtime": utils.Slice{"123", "456"}}},
		},
	)
	// return c.SendStatus(200)

}

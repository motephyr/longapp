package controllers

import (
	"log"
	"strconv"

	"github.com/motephyr/longcare/models"

	"github.com/gofiber/fiber/v2"
	"github.com/motephyr/longcare/pkg/modelhelper"
)

type sourceController struct{}

var SourceController sourceController

func (sourceController) List(c *fiber.Ctx) error {
	sources := modelhelper.SourceAll()
	return c.JSON(sources)
}

func (sourceController) Create(c *fiber.Ctx) error {
	// Validate input
	var sourceRequest models.Source
	c.BodyParser(&sourceRequest)

	source := modelhelper.SourceCreate(&sourceRequest)

	return c.JSON(source)
}

func (sourceController) Find(c *fiber.Ctx) error {

	uid, _ := strconv.Atoi(c.Params("id"))

	source := modelhelper.SourceFind(uid)
	log.Println(source)

	return c.JSON(source)
}

func (sourceController) Update(c *fiber.Ctx) error {

	var source models.Source
	uid, _ := strconv.Atoi(c.Params("id"))
	c.BodyParser(&source)
	result := modelhelper.SourceUpdate(uid, &source)
	return c.JSON(result)
}

func (sourceController) Delete(c *fiber.Ctx) error {
	// Get model if exist
	var source models.Source
	uid, _ := strconv.Atoi(c.Params("id"))
	result := modelhelper.SourceDelete(uid, &source)

	return c.JSON(result)
}

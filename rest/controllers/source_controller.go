package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/motephyr/longcare/models"

	"github.com/gofiber/fiber/v2"
	"github.com/motephyr/longcare/pkg/modelhelper"
)

type sourceController struct{}

var SourceController sourceController

func (sourceController) List(c *fiber.Ctx) error {
	sources := modelhelper.Source.All()
	return c.JSON(sources)
}

func (sourceController) Find(c *fiber.Ctx) error {

	uid, _ := strconv.Atoi(c.Params("id"))

	source := modelhelper.Source.Find(uid)

	return c.JSON(source)
}

func (sourceController) Create(c *fiber.Ctx) error {
	// Validate input
	var sourceRequest models.Source
	c.BodyParser(&sourceRequest)

	source := modelhelper.Source.Create(&sourceRequest)

	return c.JSON(source)
}

func (sourceController) Update(c *fiber.Ctx) error {

	var source models.Source
	uid, _ := strconv.Atoi(c.Params("id"))
	c.BodyParser(&source)
	result := modelhelper.Source.Update(uid, &source)
	return c.JSON(result)
}

func (sourceController) Delete(c *fiber.Ctx) error {
	// Get model if exist
	var source models.Source
	uid, _ := strconv.Atoi(c.Params("id"))
	result := modelhelper.Source.Delete(uid, &source)

	return c.JSON(result)
}

func (sourceController) Upload(c *fiber.Ctx) error {
	start := time.Now()
	log.SetOutput(ioutil.Discard)
	var err error
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "documents" key:
		files := form.File["upload"]

		for _, file := range files {
			err = modelhelper.HandleUploadIndividualFile(c, file)
		}
	}
	fmt.Printf("\n%2fs", time.Since(start).Seconds())
	return err
}

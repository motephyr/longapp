package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longapp/models"
	"github.com/motephyr/longapp/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userIdstringController struct{}

var UserIdstringController userIdstringController

func (userIdstringController) Index(c *fiber.Ctx) error {
	uid, _ := strconv.Atoi(c.Params("id"))

	user, err := models.Users(Load("UserIdstrings"), Where("id = ?", uid)).OneG()
	if err != nil {
		log.Println(err)
	}

	userIdstrings := utils.SliceStructToSliceMap(user.R.UserIdstrings)

	return inertia.Render(c,
		"user/idstring/Index.vue", // Will render component named as Main
		inertia.Map{
			"user":          user,
			"userIdstrings": userIdstrings,
		},
	)
}
func (userIdstringController) New(c *fiber.Ctx) error {
	uid, _ := strconv.Atoi(c.Params("id"))

	user, err := models.Users(Where("id = ?", uid)).OneG()
	if err != nil {
		log.Println(err)
	}

	return inertia.Render(c,
		"user/idstring/New.vue", // Will render component named as Main
		inertia.Map{
			"user": user,
		},
	)
}

func (userIdstringController) Create(c *fiber.Ctx) error {
	// Validate input

	id := c.Params("id")

	uid, _ := strconv.Atoi(id)
	userIdstring := models.UserIdstring{
		Idstring: null.StringFrom(c.FormValue("idstring")),
		UserID:   null.IntFrom(uid),
	}

	userIdstring.InsertG(boil.Infer())

	return c.Redirect("/manage/users/" + id + "/idstrings")

}

func (userIdstringController) Delete(c *fiber.Ctx) error {
	// Get model if exist
	id := c.Params("id")
	idstring, _ := strconv.Atoi(c.Params("idstring"))

	userIdstring := models.UserIdstring{
		ID: idstring,
	}

	userIdstring.DeleteG()

	return c.Redirect("/manage/users/" + id + "/idstrings")

}

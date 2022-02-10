package controllers

import (
	"log"
	"strconv"

	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/models"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/gofiber/fiber/v2"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userController struct{}

var UserController userController

func (userController) Index(c *fiber.Ctx) error {
	users, _ := models.Users().AllG()
	return inertia.Render(c,
		"user/Index.vue", // Will render component named as Main
		inertia.Map{
			"users": users,
		},
	)
}

func (userController) New(c *fiber.Ctx) error {

	return inertia.Render(c,
		"user/New.vue", // Will render component named as Main
		inertia.Map{
			"user": nil,
		},
	)
}

func (userController) Edit(c *fiber.Ctx) error {

	uid, _ := strconv.Atoi(c.Params("id"))

	user, err := models.Users(Where("id = ?", uid)).OneG()
	if err != nil {
		log.Println(err)
	}

	return inertia.Render(c,
		"user/Edit.vue", // Will render component named as Main
		inertia.Map{
			"user": user,
		},
	)
}
func (userController) Create(c *fiber.Ctx) error {
	// Validate input
	var userRequest models.User

	userRequest.Username = c.FormValue("username")
	userRequest.Password, _ = app.Http.Hash.Create(c.FormValue("password"))

	userRequest.InsertG(boil.Infer())

	return c.Redirect("/manage/users")

}

func (userController) Update(c *fiber.Ctx) error {

	var user models.User
	uid, _ := strconv.Atoi(c.Params("id"))
	user.ID = uid
	user.Username = c.FormValue("username")
	user.Password, _ = app.Http.Hash.Create(c.FormValue("password"))

	user.UpdateG(boil.Infer())
	return c.Redirect("/manage/users")
}

func (userController) Delete(c *fiber.Ctx) error {
	// Get model if exist
	uid, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		ID: uid,
	}

	user.DeleteG()

	return c.Redirect("/manage/users")

}

package controllers

import (
	"log"
	"strconv"
	"time"

	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/models"
	"github.com/motephyr/longcare/utils"
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

func (userController) Show(c *fiber.Ctx) error {

	uid, _ := strconv.Atoi(c.Params("id"))

	user, err := models.Users(Where("id = ?", uid)).OneG()
	if err != nil {
		log.Println(err)
	}
	today := time.Now().Format("20060102")
	datestring := time.Now().AddDate(0, 0, -230).Format("20060102")

	groups, err := models.Groups(Load("Sources", OrderBy("timestring")), Where("user_id = ? and datestring > ?", uid, datestring), OrderBy("datestring desc")).AllG()
	if err != nil {
		log.Println(err)
	}
	todaytimes := 0
	todaysumtime := 0
	var groupslice []any

	for _, x := range groups {
		if x.Datestring.String == today {
			todaytimes += 1
			i, _ := strconv.Atoi(x.Duringtime.String)
			todaysumtime += i

		}
		groupobj := utils.StructToMap(*x)
		sources := utils.SliceStructToSliceMap(x.R.Sources)
		groupobj["sources"] = sources
		groupslice = append(groupslice, groupobj)
	}

	return inertia.Render(c,
		"user/Show.vue", // Will render component named as Main
		inertia.Map{
			"user":         user,
			"groups":       groupslice,
			"todaytimes":   todaytimes,
			"todaysumtime": todaysumtime,
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
	var user models.User
	uid, _ := strconv.Atoi(c.Params("id"))
	user.ID = uid

	rowsAff, _ := user.DeleteG()
	if rowsAff == 0 {
		return c.JSON(false)
	} else {
		return c.JSON(true)
	}
}

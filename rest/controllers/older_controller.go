package controllers

import (
	"log"
	"strconv"
	"time"

	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/models"
	"github.com/motephyr/longcare/utils"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/gofiber/fiber/v2"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type olderController struct{}

var OlderController olderController

func (olderController) Index(c *fiber.Ctx) error {
	olders, _ := models.Olders().AllG()
	return inertia.Render(c,
		"older/Index.vue", // Will render component named as Main
		inertia.Map{
			"olders": olders,
		},
	)
}

func (olderController) Show(c *fiber.Ctx) error {

	uid, _ := strconv.Atoi(c.Params("id"))

	older, err := models.Olders(Where("id = ?", uid)).OneG()
	if err != nil {
		log.Println(err)
	}
	today := time.Now().Format("20060102")
	datestring := time.Now().AddDate(0, 0, -230).Format("20060102")

	groups, err := models.Groups(Load("Sources", OrderBy("timestring")), Where("older_id = ? and datestring > ?", uid, datestring), OrderBy("datestring desc")).AllG()
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
		"older/Show.vue", // Will render component named as Main
		inertia.Map{
			"older":        older,
			"groups":       groupslice,
			"todaytimes":   todaytimes,
			"todaysumtime": todaysumtime,
		},
	)
}

func (olderController) New(c *fiber.Ctx) error {

	return inertia.Render(c,
		"older/New.vue", // Will render component named as Main
		inertia.Map{
			"older": nil,
		},
	)
}

func (olderController) Edit(c *fiber.Ctx) error {

	uid, _ := strconv.Atoi(c.Params("id"))

	older, err := models.Olders(Where("id = ?", uid)).OneG()
	if err != nil {
		log.Println(err)
	}

	return inertia.Render(c,
		"older/Edit.vue", // Will render component named as Main
		inertia.Map{
			"older": older,
		},
	)
}
func (olderController) Create(c *fiber.Ctx) error {
	// Validate input
	var olderRequest models.Older
	c.BodyParser(&olderRequest)

	olderRequest.InsertG(boil.Infer())

	return c.Redirect("/olders/" + strconv.Itoa(olderRequest.ID))

}

func (olderController) Update(c *fiber.Ctx) error {

	var older models.Older
	uid, _ := strconv.Atoi(c.Params("id"))
	c.BodyParser(&older)
	older.ID = uid
	older.UpdateG(boil.Infer())
	return c.Redirect("/olders/" + strconv.Itoa(older.ID))
}

func (olderController) Delete(c *fiber.Ctx) error {
	// Get model if exist
	var older models.Older
	uid, _ := strconv.Atoi(c.Params("id"))
	older.ID = uid

	rowsAff, _ := older.DeleteG()
	if rowsAff == 0 {
		return c.JSON(false)
	} else {
		return c.JSON(true)
	}
}

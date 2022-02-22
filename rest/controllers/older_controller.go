package controllers

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longapp/app"
	"github.com/motephyr/longapp/models"
	"github.com/motephyr/longapp/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/gofiber/fiber/v2"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type olderController struct{}

var OlderController olderController

func (olderController) Query(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	query := utils.GetSqlQuery(c.Query("query"))
	if c.Params("id") != "" {
		query = append(query, And("id = ?", c.Params("id")))
	} else {
		query = append(query, InnerJoin("user_olders c on c.older_id = olders.id and c.user_id=?", user.ID))
	}
	olders, _ := models.Olders(query...).AllG()
	log.Println(olders)
	return c.JSON(olders)
}

func (olderController) Index(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	query := utils.GetSqlQuery(c.Query("query"))
	query = append(query, InnerJoin("user_olders c on c.older_id = olders.id and c.user_id=?", user.ID))

	olders, _ := models.Olders(query...).AllG()

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
	user := c.Locals("user").(*models.User)

	var userOlder models.UserOlder
	// Validate input
	var olderRequest models.Older
	c.BodyParser(&olderRequest)

	err := utils.Tx(app.Http.Database.DB, func(tx *sql.Tx) error {
		olderRequest.InsertG(boil.Infer())
		userOlder.UserID = null.IntFrom(user.ID)
		userOlder.OlderID = null.IntFrom(olderRequest.ID)
		userOlder.InsertG(boil.Infer())
		return nil
	})

	if err != nil { // Load clients from file
		c.JSON("Transaction fail.")
	}

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

	userOlder, err := models.UserOlders(Where("older_id = ?", c.Params("id"))).OneG()

	userOlder.DeleteG()

	_, err = older.DeleteG()
	if err != nil { // Load clients from file
		log.Println(err)
	}
	return c.Redirect("/olders/")

}

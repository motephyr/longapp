package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/models"
	"github.com/motephyr/longcare/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type manageController struct{}

var ManageController manageController

func (manageController) Manage(c *fiber.Ctx) error {
	datestring := c.Params("datestring")
	sources, err := models.Sources(Where("datestring = ?", datestring), OrderBy("id")).AllG()
	if err != nil {
		log.Println(err)
	}

	return inertia.Render(c,
		"manage/Manage.vue", // Will render component named as Main
		inertia.Map{
			"datestring": datestring,
			"sources":    sources,
		},
	)
	// return c.SendStatus(200)
}

func (manageController) CreateGroup(c *fiber.Ctx) error {
	datestring := c.Params("datestring")

	payload := struct {
		SourceId []string `json:"sourceId"`
		Remark   string   `json:"remark"`
	}{}
	err := c.BodyParser(&payload)
	if err != nil {
		return err
	}

	source1, err := models.Sources(Where("id = ?", payload.SourceId[0])).OneG()
	source2, err := models.Sources(Where("id = ?", payload.SourceId[1])).OneG()
	a, _ := strconv.Atoi(source1.Timestring.String)
	b, _ := strconv.Atoi(source2.Timestring.String)

	duringtime := utils.Abs(a - b)

	var g models.Group
	g.Remark = null.StringFrom(payload.Remark)
	g.Duringtime = null.StringFrom(strconv.Itoa(duringtime))
	g.Datestring = null.StringFrom(datestring)

	g.InsertG(boil.Infer())

	source1.GroupID = null.IntFrom(g.ID)
	source2.GroupID = null.IntFrom(g.ID)
	source1.UpdateG(boil.Infer())
	source2.UpdateG(boil.Infer())

	return c.JSON("ok")

}

func (manageController) ResetData(c *fiber.Ctx) error {
	datestring := c.Params("datestring")
	models.Sources(Where("datestring = ?", datestring)).UpdateAllG(models.M{"group_id": nil})
	models.Groups(Where("datestring = ?", datestring)).DeleteAllG()

	return c.JSON("ok")
}

func (manageController) DeleteData(c *fiber.Ctx) error {
	datestring := c.Params("datestring")
	models.Sources(Where("datestring = ?", datestring)).DeleteAllG()
	models.Groups(Where("datestring = ?", datestring)).DeleteAllG()

	return c.JSON("ok")
}

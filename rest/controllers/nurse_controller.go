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

type nurseController struct{}

var NurseController nurseController

func (nurseController) Index(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	groupsquery, err := models.Groups(
		InnerJoin("user_olders c on c.older_id = groups.older_id and c.user_id=?", user.ID),
		OrderBy("datestring")).AllG()
	if err != nil {
		log.Println(err)
	}
	datestrings := map[string]map[string]int{}
	for _, x := range groupsquery {
		obj := map[string]int{"time": 0, "longesttime": 0}
		if datestrings[x.Datestring.String] != nil {
			obj["time"] = datestrings[x.Datestring.String]["time"] + 1
		} else {
			obj["time"] = 1
		}

		duringtime, _ := strconv.Atoi(x.Duringtime.String)

		if datestrings[x.Datestring.String] != nil {
			if datestrings[x.Datestring.String]["longesttime"] > duringtime {
				obj["longesttime"] = datestrings[x.Datestring.String]["longesttime"]
			} else {
				obj["longesttime"] = duringtime
			}
		} else {
			obj["longesttime"] = duringtime
		}

		datestrings[x.Datestring.String] = obj

	}

	return inertia.Render(c,
		"nurse/Index.vue", // Will render component named as Main
		inertia.Map{
			"datestrings": datestrings,
		},
	)
}

func (nurseController) Nurse(c *fiber.Ctx) error {
	datestring := c.Params("datestring")
	groups, err := models.Groups(Load("Sources"), Where("datestring = ?", datestring), OrderBy("id")).AllG()
	if err != nil {
		log.Println(err)
	}

	olders, err := models.Olders().AllG()

	var groupslice []any
	for _, obj := range groups {
		groupobj := utils.StructToMap(*obj)
		sources := utils.SliceStructToSliceMap(obj.R.Sources)
		groupobj["sources"] = sources
		groupslice = append(groupslice, groupobj)
	}

	return inertia.Render(c,
		"nurse/Nurse.vue", // Will render component named as Main
		inertia.Map{
			"datestring": datestring,
			"olders":     olders,
			"groups":     groupslice,
		},
	)
	// return c.SendStatus(200)

}

func (nurseController) SetOlderId(c *fiber.Ctx) error {
	var payload struct {
		GroupId int `param:"group_id"`
		OlderId int `form:"older_id"`
	}

	utils.GetFiberParams(c, &payload)

	group, err := models.Groups(Where("id = ?", payload.GroupId)).OneG()
	if err != nil {
		log.Println(err)
	}
	group.OlderID = null.IntFrom(payload.OlderId)
	_, err = group.UpdateG(boil.Infer())
	if err != nil {
		log.Println(err)
	}

	return c.Redirect("/nurse/" + c.Params("datestring"))

}

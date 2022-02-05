package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longcare/models"
	"github.com/motephyr/longcare/utils"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type nurseController struct{}

var NurseController nurseController

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

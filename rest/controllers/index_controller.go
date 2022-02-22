package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	inertia "github.com/motephyr/fiber-inertia"
	"github.com/motephyr/longapp/models"
	"github.com/motephyr/longapp/utils"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type indexController struct{}

var IndexController indexController

func (indexController) IndexPage(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	today := time.Now().AddDate(0, 0, -7).Format("20060102")
	olders, err := models.Olders(
		InnerJoin("user_olders c on c.older_id = olders.id and c.user_id=?", user.ID),
		Load("Groups", Where("datestring > ?", today), OrderBy("datestring"))).AllG()

	if err != nil {
		log.Println(err)
	}

	groups := []*models.Group{}
	for _, x := range olders {
		groups = append(groups, x.R.Groups...)
	}

	// groups := utils.Reduce(olders, func(acc []*models.Group, i *models.Older) []*models.Group {
	// 	return append(acc, i.R.Groups...)
	// })

	datestrings := []string{}
	for _, x := range groups {
		datestrings = append(datestrings, x.Datestring.String)
	}

	// datestrings := utils.Map(groups, func(x *models.Group) string {
	// 	return x.Datestring.String
	// })

	datestrings = utils.Unique(datestrings)

	sumetimes := utils.Map(olders, func(x *models.Older) map[string]any {
		result := utils.StructToMap(*x)
		result["sumtime"] = utils.Map(datestrings, func(y string) int {
			a := utils.Filter(x.R.Groups, func(z *models.Group) bool {
				return z.Datestring.String == y
			})
			return utils.Reduce(a, func(acc int, w *models.Group) int {
				i, _ := strconv.Atoi(w.Duringtime.String)
				return acc + i
			})

		})

		filterGroups := utils.Filter(x.R.Groups, func(z *models.Group) bool {
			return z.Datestring.String == today
		})

		result["todaytime"] = utils.Reduce(filterGroups, func(acc int, w *models.Group) int {
			i, _ := strconv.Atoi(w.Duringtime.String)
			return acc + i
		})
		result["todaynum"] = utils.Reduce(filterGroups, func(acc int, w *models.Group) int {
			return acc + 1
		})
		return result
	})
	// log.Println(sumetimes)
	important := map[string]any{}

	if len(sumetimes) != 0 {

		importantolders := utils.SortBy(sumetimes, func(i, j int) bool {
			a := sumetimes[i]["todaytime"].(int)
			b := sumetimes[j]["todaytime"].(int)
			return a > b
		})[0]

		if importantolders["todaynum"] == 0 {
			important = map[string]any{
				"name":      "無",
				"room":      "",
				"todaynum":  nil,
				"todaytime": nil,
				"notice":    "無",
			}
		} else {
			important = importantolders
		}
	} else {
		important = map[string]any{
			"name":      "無",
			"room":      "",
			"todaynum":  nil,
			"todaytime": nil,
			"notice":    "無",
		}
	}

	rooms := utils.ReduceWithInitialValue(olders, map[string]any{}, func(sum map[string]any, x *models.Older) map[string]any {
		room := x.Room.String

		if sum[room] != nil {
			sum[room] = sum[room].(int) + 1
		} else {
			sum[room] = 1
		}

		return sum
	})

	return inertia.Render(c,
		"Index.vue", // Will render component named as Main
		inertia.Map{
			"datestrings": datestrings,
			"sumetimes":   sumetimes,
			"rooms":       rooms,
			"important":   important,
		},
	)
	// return c.SendStatus(200)

}

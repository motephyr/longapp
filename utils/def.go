package utils

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Tx(db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetFiberParams(c *fiber.Ctx, payload interface{}) error {

	v := reflect.ValueOf(payload).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {

		file := t.Field(i)

		// name := file.Name
		typeValue := file.Type.String()
		// tag := string(file.Tag)

		if jt, ok := file.Tag.Lookup("param"); ok {
			paramsName := strings.Split(jt, ",")[0]
			value := c.Params(paramsName)

			fieldVal := v.Field(i)

			if typeValue == "int" {
				value, _ := strconv.Atoi(value)
				fieldVal.Set(reflect.ValueOf(value))
			} else if typeValue == "string" {

				fieldVal.Set(reflect.ValueOf(value))

			}
		}
	}
	err := c.BodyParser(payload)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func GetSqlQuery(query string) []qm.QueryMod {
	where := gjson.Get(query, "where")
	with := gjson.Get(query, "with")
	sort := gjson.Get(query, "sort")
	page := gjson.Get(query, "page")
	perPage := gjson.Get(query, "perPage")

	var result = []qm.QueryMod{}
	result = handleWhereClause(where, result)

	result = handleWithClause(with, result)

	result = handleSortClause(sort, result)

	result = handlePageClause(page, perPage, result)

	return result

}

func handlePageClause(page gjson.Result, perPage gjson.Result, result []qm.QueryMod) []qm.QueryMod {
	if page.Raw != "" {
		limit := 10
		if perPage.Raw != "" {
			limit = int(perPage.Num)
		}
		result = append(result, qm.Limit(limit))

		if int(page.Num) == 1 {
			result = append(result, qm.Offset(-1))
		} else {
			result = append(result, qm.Offset(limit*(int(page.Num)-1)))
		}
	}
	return result
}

func handleSortClause(sort gjson.Result, result []qm.QueryMod) []qm.QueryMod {
	if sort.Raw != "" {
		result = append(result, qm.OrderBy(sort.Str))
	}
	return result
}

func handleWithClause(with gjson.Result, result []qm.QueryMod) []qm.QueryMod {
	if with.Array() != nil {
		for _, x := range with.Array() {
			result = append(result, qm.Load(x.Str))
		}
	}
	return result
}

func handleWhereClause(where gjson.Result, result []qm.QueryMod) []qm.QueryMod {
	if where.Map() != nil {
		for k, v := range where.Map() {
			if v.Type.String() == "String" || v.Type.String() == "Number" {
				result = append(result, qm.Where(k+" = ?", v))
			} else if v.Type.String() == "JSON" {
				result = append(result, qm.Where(k+" "+v.Array()[0].Str+" ?", v.Array()[1]))
			}
		}
	}
	return result
}

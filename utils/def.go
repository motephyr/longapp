package utils

import (
	"database/sql"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func Reduce[T1, T2 any](input []T1, f func(T2, T1) T2) T2 {
	var acc T2

	for _, v := range input {
		acc = f(acc, v)
	}

	return acc
}

func ReduceWithInitialValue[T1, T2 any](input []T1, acc T2, f func(T2, T1) T2) T2 {
	for _, v := range input {
		acc = f(acc, v)
	}

	return acc
}

func Map[T1, T2 any](input []T1, f func(T1) T2) (output []T2) {
	output = make([]T2, 0, len(input))
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}

func Unique[T Ordered](ss []T) []T {
	return SortedUnique[T](Sort(ss))
}

func SortedUnique[T Ordered](ss []T) []T {
	if ss == nil {
		return nil
	}
	result := []T{}
	last := *new(T)
	for i, s := range ss {
		if i != 0 && last == s {
			continue
		}
		result = append(result, s)
		last = s
	}
	return result
}

// Sort returns a new slice that is the sorted copy of the slice it was called on. Unlike sort.Strings, it does not mutate the original slice
func Sort[T Ordered](ss []T) []T {
	if ss == nil {
		return nil
	}
	ss2 := make([]T, len(ss))
	copy(ss2, ss)
	sort.Slice(ss2, func(i int, j int) bool {
		return ss2[i] <= ss2[j]
	})
	return ss2
}

func SortBy[T any](ss []T, sortFunc func(i, j int) bool) []T {
	if ss == nil {
		return nil
	}
	ss2 := make([]T, len(ss))
	copy(ss2, ss)
	sort.Slice(ss2, func(i, j int) bool {
		return sortFunc(i, j)
	})
	return ss2
}

func StructToMap(obj any) map[string]any {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		name := ToSnakeCase(obj1.Field(i).Name)
		value := obj2.Field(i).Interface()
		data[name] = value
	}
	return data
}

func SliceStructToSliceMap[T any](objs []*T) []any {

	var data []any
	for _, obj := range objs {
		data = append(data, StructToMap(*obj))
	}
	return data
}

func Filter[T any](input []T, pred func(T) bool) (output []T) {
	for _, v := range input {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}

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

package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestGetSqlQuery(t *testing.T) {

	name := `{"q":{"search_string":""},"where":{},"page":1,"with":["Groups"],"sort":"created_at desc"}`
	result := []qm.QueryMod{
		qm.Load("Groups"),
		qm.OrderBy("created_at desc"),
		qm.Limit(10),
		qm.Offset(-1),
	}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery2(t *testing.T) {

	name := `{"q":{"search_string":""},"where":{},"with":["Groups"],"sort":"created_at desc"}`
	result := []qm.QueryMod{
		qm.Load("Groups"),
		qm.OrderBy("created_at desc"),
	}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery3(t *testing.T) {

	name := `{"q":{"search_string":""}}`
	result := []qm.QueryMod{}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery4(t *testing.T) {

	name := `{"where":{}}`
	result := []qm.QueryMod{}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery5(t *testing.T) {

	name := `{"where":{"id": 1}}`
	result := []qm.QueryMod{
		qm.Where("id = ?", 1),
	}
	msg := GetSqlQuery(name)

	eq := fmt.Sprintln(msg) == fmt.Sprintln(result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery6(t *testing.T) {

	name := `{"where":{"name": "john"}}`
	result := []qm.QueryMod{
		qm.Where("name = ?", "john"),
	}
	msg := GetSqlQuery(name)

	eq := fmt.Sprintln(msg) == fmt.Sprintln(result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery7(t *testing.T) {

	name := `{"where":{"id": [">", 9]}}`
	result := []qm.QueryMod{
		qm.Where("id > ?", 9),
	}
	msg := GetSqlQuery(name)

	eq := fmt.Sprintln(msg) == fmt.Sprintln(result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQuery8(t *testing.T) {

	name := `{"where":{"id": [">", 9], "age": [">", 50]}}`
	result := []qm.QueryMod{
		qm.Where("id > ?", 9),
		qm.Where("age > ?", 50),
	}
	msg := GetSqlQuery(name)

	eq := fmt.Sprintln(msg) == fmt.Sprintln(result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQueryPage1(t *testing.T) {

	name := `{"page":1}`
	result := []qm.QueryMod{
		qm.Limit(10),
		qm.Offset(-1),
	}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQueryPage2(t *testing.T) {

	name := `{"page":2}`
	result := []qm.QueryMod{
		qm.Limit(10),
		qm.Offset(10),
	}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQueryPage3(t *testing.T) {

	name := `{"page":2, "perPage": 6}`
	result := []qm.QueryMod{
		qm.Limit(6),
		qm.Offset(6),
	}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

func TestGetSqlQueryPage4(t *testing.T) {

	name := `{"page":1, "perPage": 6}`
	result := []qm.QueryMod{
		qm.Limit(6),
		qm.Offset(-1),
	}
	msg := GetSqlQuery(name)

	eq := reflect.DeepEqual(msg, result)

	if !eq {
		t.Fatalf(`%v is not equal %v`, msg, result)
	} else {
		t.Log("succeess")
	}
}

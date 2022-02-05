// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOlders(t *testing.T) {
	t.Parallel()

	query := Olders()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOldersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOldersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Olders().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOldersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OlderSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOldersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OlderExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Older exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OlderExists to return true, but got false.")
	}
}

func testOldersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	olderFound, err := FindOlder(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if olderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOldersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Olders().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testOldersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Olders().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOldersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	olderOne := &Older{}
	olderTwo := &Older{}
	if err = randomize.Struct(seed, olderOne, olderDBTypes, false, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}
	if err = randomize.Struct(seed, olderTwo, olderDBTypes, false, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = olderOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = olderTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Olders().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOldersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	olderOne := &Older{}
	olderTwo := &Older{}
	if err = randomize.Struct(seed, olderOne, olderDBTypes, false, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}
	if err = randomize.Struct(seed, olderTwo, olderDBTypes, false, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = olderOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = olderTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func olderBeforeInsertHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderAfterInsertHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderAfterSelectHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderBeforeUpdateHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderAfterUpdateHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderBeforeDeleteHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderAfterDeleteHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderBeforeUpsertHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func olderAfterUpsertHook(e boil.Executor, o *Older) error {
	*o = Older{}
	return nil
}

func testOldersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Older{}
	o := &Older{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, olderDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Older object: %s", err)
	}

	AddOlderHook(boil.BeforeInsertHook, olderBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	olderBeforeInsertHooks = []OlderHook{}

	AddOlderHook(boil.AfterInsertHook, olderAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	olderAfterInsertHooks = []OlderHook{}

	AddOlderHook(boil.AfterSelectHook, olderAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	olderAfterSelectHooks = []OlderHook{}

	AddOlderHook(boil.BeforeUpdateHook, olderBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	olderBeforeUpdateHooks = []OlderHook{}

	AddOlderHook(boil.AfterUpdateHook, olderAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	olderAfterUpdateHooks = []OlderHook{}

	AddOlderHook(boil.BeforeDeleteHook, olderBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	olderBeforeDeleteHooks = []OlderHook{}

	AddOlderHook(boil.AfterDeleteHook, olderAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	olderAfterDeleteHooks = []OlderHook{}

	AddOlderHook(boil.BeforeUpsertHook, olderBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	olderBeforeUpsertHooks = []OlderHook{}

	AddOlderHook(boil.AfterUpsertHook, olderAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	olderAfterUpsertHooks = []OlderHook{}
}

func testOldersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOldersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(olderColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOlderToManyGroups(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Older
	var b, c Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.OlderID, a.ID)
	queries.Assign(&c.OlderID, a.ID)
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Groups().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.OlderID, b.OlderID) {
			bFound = true
		}
		if queries.Equal(v.OlderID, c.OlderID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OlderSlice{&a}
	if err = a.L.LoadGroups(tx, false, (*[]*Older)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Groups); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Groups = nil
	if err = a.L.LoadGroups(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Groups); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOlderToManyAddOpGroups(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Older
	var b, c, d, e Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, olderDBTypes, false, strmangle.SetComplement(olderPrimaryKeyColumns, olderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Group{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Group{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGroups(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.OlderID) {
			t.Error("foreign key was wrong value", a.ID, first.OlderID)
		}
		if !queries.Equal(a.ID, second.OlderID) {
			t.Error("foreign key was wrong value", a.ID, second.OlderID)
		}

		if first.R.Older != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Older != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Groups[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Groups[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Groups().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOlderToManySetOpGroups(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Older
	var b, c, d, e Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, olderDBTypes, false, strmangle.SetComplement(olderPrimaryKeyColumns, olderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Group{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetGroups(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Groups().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetGroups(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Groups().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.OlderID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.OlderID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.OlderID) {
		t.Error("foreign key was wrong value", a.ID, d.OlderID)
	}
	if !queries.Equal(a.ID, e.OlderID) {
		t.Error("foreign key was wrong value", a.ID, e.OlderID)
	}

	if b.R.Older != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Older != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Older != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Older != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Groups[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Groups[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testOlderToManyRemoveOpGroups(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Older
	var b, c, d, e Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, olderDBTypes, false, strmangle.SetComplement(olderPrimaryKeyColumns, olderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Group{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddGroups(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Groups().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveGroups(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Groups().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.OlderID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.OlderID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Older != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Older != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Older != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Older != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Groups) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Groups[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Groups[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testOldersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOldersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OlderSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testOldersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Olders().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	olderDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `Room`: `character varying`, `Birthday`: `character varying`, `Medicalrecord`: `character varying`, `Medicine`: `character varying`, `Notice`: `character varying`, `Pictureurl`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_            = bytes.MinRead
)

func testOldersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(olderPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(olderAllColumns) == len(olderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, olderDBTypes, true, olderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOldersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(olderAllColumns) == len(olderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Older{}
	if err = randomize.Struct(seed, o, olderDBTypes, true, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, olderDBTypes, true, olderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(olderAllColumns, olderPrimaryKeyColumns) {
		fields = olderAllColumns
	} else {
		fields = strmangle.SetComplement(
			olderAllColumns,
			olderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OlderSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOldersUpsert(t *testing.T) {
	t.Parallel()

	if len(olderAllColumns) == len(olderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Older{}
	if err = randomize.Struct(seed, &o, olderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Older: %s", err)
	}

	count, err := Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, olderDBTypes, false, olderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Older: %s", err)
	}

	count, err = Olders().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
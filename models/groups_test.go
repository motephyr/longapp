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

func testGroups(t *testing.T) {
	t.Parallel()

	query := Groups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGroupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
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

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Groups().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GroupSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GroupExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Group exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GroupExists to return true, but got false.")
	}
}

func testGroupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	groupFound, err := FindGroup(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if groupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGroupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Groups().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testGroupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Groups().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGroupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupOne := &Group{}
	groupTwo := &Group{}
	if err = randomize.Struct(seed, groupOne, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}
	if err = randomize.Struct(seed, groupTwo, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = groupOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = groupTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Groups().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGroupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	groupOne := &Group{}
	groupTwo := &Group{}
	if err = randomize.Struct(seed, groupOne, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}
	if err = randomize.Struct(seed, groupTwo, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = groupOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = groupTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func groupBeforeInsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterInsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterSelectHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupBeforeUpdateHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterUpdateHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupBeforeDeleteHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterDeleteHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupBeforeUpsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterUpsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func testGroupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Group{}
	o := &Group{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, groupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Group object: %s", err)
	}

	AddGroupHook(boil.BeforeInsertHook, groupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	groupBeforeInsertHooks = []GroupHook{}

	AddGroupHook(boil.AfterInsertHook, groupAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	groupAfterInsertHooks = []GroupHook{}

	AddGroupHook(boil.AfterSelectHook, groupAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	groupAfterSelectHooks = []GroupHook{}

	AddGroupHook(boil.BeforeUpdateHook, groupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	groupBeforeUpdateHooks = []GroupHook{}

	AddGroupHook(boil.AfterUpdateHook, groupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	groupAfterUpdateHooks = []GroupHook{}

	AddGroupHook(boil.BeforeDeleteHook, groupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	groupBeforeDeleteHooks = []GroupHook{}

	AddGroupHook(boil.AfterDeleteHook, groupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	groupAfterDeleteHooks = []GroupHook{}

	AddGroupHook(boil.BeforeUpsertHook, groupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	groupBeforeUpsertHooks = []GroupHook{}

	AddGroupHook(boil.AfterUpsertHook, groupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	groupAfterUpsertHooks = []GroupHook{}
}

func testGroupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGroupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(groupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGroupToManySources(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Group
	var b, c Source

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, sourceDBTypes, false, sourceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sourceDBTypes, false, sourceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.GroupID, a.ID)
	queries.Assign(&c.GroupID, a.ID)
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Sources().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.GroupID, b.GroupID) {
			bFound = true
		}
		if queries.Equal(v.GroupID, c.GroupID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GroupSlice{&a}
	if err = a.L.LoadSources(tx, false, (*[]*Group)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sources); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Sources = nil
	if err = a.L.LoadSources(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Sources); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testGroupToManyAddOpSources(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Group
	var b, c, d, e Source

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Source{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sourceDBTypes, false, strmangle.SetComplement(sourcePrimaryKeyColumns, sourceColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Source{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSources(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.GroupID) {
			t.Error("foreign key was wrong value", a.ID, first.GroupID)
		}
		if !queries.Equal(a.ID, second.GroupID) {
			t.Error("foreign key was wrong value", a.ID, second.GroupID)
		}

		if first.R.Group != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Group != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Sources[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Sources[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Sources().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testGroupToManySetOpSources(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Group
	var b, c, d, e Source

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Source{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sourceDBTypes, false, strmangle.SetComplement(sourcePrimaryKeyColumns, sourceColumnsWithoutDefault)...); err != nil {
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

	err = a.SetSources(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Sources().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSources(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Sources().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.GroupID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.GroupID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.GroupID) {
		t.Error("foreign key was wrong value", a.ID, d.GroupID)
	}
	if !queries.Equal(a.ID, e.GroupID) {
		t.Error("foreign key was wrong value", a.ID, e.GroupID)
	}

	if b.R.Group != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Group != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Group != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Group != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Sources[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Sources[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testGroupToManyRemoveOpSources(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Group
	var b, c, d, e Source

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Source{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sourceDBTypes, false, strmangle.SetComplement(sourcePrimaryKeyColumns, sourceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSources(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Sources().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSources(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Sources().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.GroupID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.GroupID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Group != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Group != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Group != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Group != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Sources) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Sources[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Sources[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testGroupToOneOlderUsingOlder(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Group
	var foreign Older

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, olderDBTypes, false, olderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Older struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.OlderID, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Older().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := GroupSlice{&local}
	if err = local.L.LoadOlder(tx, false, (*[]*Group)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Older == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Older = nil
	if err = local.L.LoadOlder(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Older == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGroupToOneSetOpOlderUsingOlder(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Group
	var b, c Older

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, olderDBTypes, false, strmangle.SetComplement(olderPrimaryKeyColumns, olderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, olderDBTypes, false, strmangle.SetComplement(olderPrimaryKeyColumns, olderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Older{&b, &c} {
		err = a.SetOlder(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Older != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Groups[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.OlderID, x.ID) {
			t.Error("foreign key was wrong value", a.OlderID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OlderID))
		reflect.Indirect(reflect.ValueOf(&a.OlderID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.OlderID, x.ID) {
			t.Error("foreign key was wrong value", a.OlderID, x.ID)
		}
	}
}

func testGroupToOneRemoveOpOlderUsingOlder(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Group
	var b Older

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, olderDBTypes, false, strmangle.SetComplement(olderPrimaryKeyColumns, olderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOlder(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOlder(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Older().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Older != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.OlderID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Groups) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testGroupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
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

func testGroupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GroupSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testGroupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Groups().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	groupDBTypes = map[string]string{`ID`: `integer`, `Remark`: `character varying`, `Duringtime`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `Datestring`: `character varying`, `OlderID`: `integer`}
	_            = bytes.MinRead
)

func testGroupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(groupAllColumns) == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, groupDBTypes, true, groupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGroupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(groupAllColumns) == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Group{}
	if err = randomize.Struct(seed, o, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, groupDBTypes, true, groupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(groupAllColumns, groupPrimaryKeyColumns) {
		fields = groupAllColumns
	} else {
		fields = strmangle.SetComplement(
			groupAllColumns,
			groupPrimaryKeyColumns,
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

	slice := GroupSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGroupsUpsert(t *testing.T) {
	t.Parallel()

	if len(groupAllColumns) == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Group{}
	if err = randomize.Struct(seed, &o, groupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Group: %s", err)
	}

	count, err := Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, groupDBTypes, false, groupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Group: %s", err)
	}

	count, err = Groups().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

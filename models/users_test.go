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

func testUsers(t *testing.T) {
	t.Parallel()

	query := Users()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
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

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Users().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if User exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserExists to return true, but got false.")
	}
}

func testUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userFound, err := FindUser(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Users().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Users().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userOne := &User{}
	userTwo := &User{}
	if err = randomize.Struct(seed, userOne, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err = randomize.Struct(seed, userTwo, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = userOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Users().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userOne := &User{}
	userTwo := &User{}
	if err = randomize.Struct(seed, userOne, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err = randomize.Struct(seed, userTwo, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = userOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userBeforeInsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterInsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterSelectHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeUpdateHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterUpdateHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeDeleteHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterDeleteHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeUpsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterUpsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func testUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &User{}
	o := &User{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userDBTypes, false); err != nil {
		t.Errorf("Unable to randomize User object: %s", err)
	}

	AddUserHook(boil.BeforeInsertHook, userBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userBeforeInsertHooks = []UserHook{}

	AddUserHook(boil.AfterInsertHook, userAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userAfterInsertHooks = []UserHook{}

	AddUserHook(boil.AfterSelectHook, userAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userAfterSelectHooks = []UserHook{}

	AddUserHook(boil.BeforeUpdateHook, userBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userBeforeUpdateHooks = []UserHook{}

	AddUserHook(boil.AfterUpdateHook, userAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userAfterUpdateHooks = []UserHook{}

	AddUserHook(boil.BeforeDeleteHook, userBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userBeforeDeleteHooks = []UserHook{}

	AddUserHook(boil.AfterDeleteHook, userAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userAfterDeleteHooks = []UserHook{}

	AddUserHook(boil.BeforeUpsertHook, userBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userBeforeUpsertHooks = []UserHook{}

	AddUserHook(boil.AfterUpsertHook, userAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userAfterUpsertHooks = []UserHook{}
}

func testUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(userColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserToManyUserIdstrings(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c UserIdstring

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userIdstringDBTypes, false, userIdstringColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userIdstringDBTypes, false, userIdstringColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.UserID, a.ID)
	queries.Assign(&c.UserID, a.ID)
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserIdstrings().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.UserID, b.UserID) {
			bFound = true
		}
		if queries.Equal(v.UserID, c.UserID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UserSlice{&a}
	if err = a.L.LoadUserIdstrings(tx, false, (*[]*User)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserIdstrings); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserIdstrings = nil
	if err = a.L.LoadUserIdstrings(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserIdstrings); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testUserToManyUserOlders(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c UserOlder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userOlderDBTypes, false, userOlderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userOlderDBTypes, false, userOlderColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.UserID, a.ID)
	queries.Assign(&c.UserID, a.ID)
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserOlders().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.UserID, b.UserID) {
			bFound = true
		}
		if queries.Equal(v.UserID, c.UserID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UserSlice{&a}
	if err = a.L.LoadUserOlders(tx, false, (*[]*User)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserOlders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserOlders = nil
	if err = a.L.LoadUserOlders(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserOlders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testUserToManyAddOpUserIdstrings(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c, d, e UserIdstring

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserIdstring{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userIdstringDBTypes, false, strmangle.SetComplement(userIdstringPrimaryKeyColumns, userIdstringColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserIdstring{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserIdstrings(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.UserID) {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if !queries.Equal(a.ID, second.UserID) {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserIdstrings[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserIdstrings[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserIdstrings().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testUserToManySetOpUserIdstrings(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c, d, e UserIdstring

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserIdstring{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userIdstringDBTypes, false, strmangle.SetComplement(userIdstringPrimaryKeyColumns, userIdstringColumnsWithoutDefault)...); err != nil {
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

	err = a.SetUserIdstrings(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserIdstrings().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUserIdstrings(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserIdstrings().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.UserID) {
		t.Error("foreign key was wrong value", a.ID, d.UserID)
	}
	if !queries.Equal(a.ID, e.UserID) {
		t.Error("foreign key was wrong value", a.ID, e.UserID)
	}

	if b.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.User != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.User != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.UserIdstrings[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UserIdstrings[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testUserToManyRemoveOpUserIdstrings(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c, d, e UserIdstring

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserIdstring{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userIdstringDBTypes, false, strmangle.SetComplement(userIdstringPrimaryKeyColumns, userIdstringColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUserIdstrings(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserIdstrings().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUserIdstrings(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserIdstrings().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.User != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.User != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.UserIdstrings) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UserIdstrings[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UserIdstrings[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testUserToManyAddOpUserOlders(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c, d, e UserOlder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserOlder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userOlderDBTypes, false, strmangle.SetComplement(userOlderPrimaryKeyColumns, userOlderColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserOlder{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserOlders(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.UserID) {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if !queries.Equal(a.ID, second.UserID) {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserOlders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserOlders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserOlders().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testUserToManySetOpUserOlders(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c, d, e UserOlder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserOlder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userOlderDBTypes, false, strmangle.SetComplement(userOlderPrimaryKeyColumns, userOlderColumnsWithoutDefault)...); err != nil {
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

	err = a.SetUserOlders(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserOlders().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUserOlders(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserOlders().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.UserID) {
		t.Error("foreign key was wrong value", a.ID, d.UserID)
	}
	if !queries.Equal(a.ID, e.UserID) {
		t.Error("foreign key was wrong value", a.ID, e.UserID)
	}

	if b.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.User != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.User != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.UserOlders[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UserOlders[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testUserToManyRemoveOpUserOlders(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a User
	var b, c, d, e UserOlder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserOlder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userOlderDBTypes, false, strmangle.SetComplement(userOlderPrimaryKeyColumns, userOlderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUserOlders(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserOlders().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUserOlders(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserOlders().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.User != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.User != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.UserOlders) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UserOlders[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UserOlders[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
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

func testUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Users().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userDBTypes = map[string]string{`ID`: `integer`, `Username`: `character varying`, `Password`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testUsersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userAllColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userDBTypes, true, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userAllColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userDBTypes, true, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userAllColumns, userPrimaryKeyColumns) {
		fields = userAllColumns
	} else {
		fields = strmangle.SetComplement(
			userAllColumns,
			userPrimaryKeyColumns,
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

	slice := UserSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(userAllColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := User{}
	if err = randomize.Struct(seed, &o, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert User: %s", err)
	}

	count, err := Users().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userDBTypes, false, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert User: %s", err)
	}

	count, err = Users().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

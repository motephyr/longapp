// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserOlder is an object representing the database table.
type UserOlder struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    null.Int  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	OlderID   null.Int  `boil:"older_id" json:"older_id,omitempty" toml:"older_id" yaml:"older_id,omitempty"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *userOlderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userOlderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserOlderColumns = struct {
	ID        string
	UserID    string
	OlderID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	OlderID:   "older_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserOlderTableColumns = struct {
	ID        string
	UserID    string
	OlderID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "user_olders.id",
	UserID:    "user_olders.user_id",
	OlderID:   "user_olders.older_id",
	CreatedAt: "user_olders.created_at",
	UpdatedAt: "user_olders.updated_at",
}

// Generated where

var UserOlderWhere = struct {
	ID        whereHelperint
	UserID    whereHelpernull_Int
	OlderID   whereHelpernull_Int
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"user_olders\".\"id\""},
	UserID:    whereHelpernull_Int{field: "\"user_olders\".\"user_id\""},
	OlderID:   whereHelpernull_Int{field: "\"user_olders\".\"older_id\""},
	CreatedAt: whereHelpernull_Time{field: "\"user_olders\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"user_olders\".\"updated_at\""},
}

// UserOlderRels is where relationship names are stored.
var UserOlderRels = struct {
	Older string
	User  string
}{
	Older: "Older",
	User:  "User",
}

// userOlderR is where relationships are stored.
type userOlderR struct {
	Older *Older `boil:"Older" json:"Older" toml:"Older" yaml:"Older"`
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userOlderR) NewStruct() *userOlderR {
	return &userOlderR{}
}

// userOlderL is where Load methods for each relationship are stored.
type userOlderL struct{}

var (
	userOlderAllColumns            = []string{"id", "user_id", "older_id", "created_at", "updated_at"}
	userOlderColumnsWithoutDefault = []string{}
	userOlderColumnsWithDefault    = []string{"id", "user_id", "older_id", "created_at", "updated_at"}
	userOlderPrimaryKeyColumns     = []string{"id"}
	userOlderGeneratedColumns      = []string{}
)

type (
	// UserOlderSlice is an alias for a slice of pointers to UserOlder.
	// This should almost always be used instead of []UserOlder.
	UserOlderSlice []*UserOlder
	// UserOlderHook is the signature for custom UserOlder hook methods
	UserOlderHook func(boil.Executor, *UserOlder) error

	userOlderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userOlderType                 = reflect.TypeOf(&UserOlder{})
	userOlderMapping              = queries.MakeStructMapping(userOlderType)
	userOlderPrimaryKeyMapping, _ = queries.BindMapping(userOlderType, userOlderMapping, userOlderPrimaryKeyColumns)
	userOlderInsertCacheMut       sync.RWMutex
	userOlderInsertCache          = make(map[string]insertCache)
	userOlderUpdateCacheMut       sync.RWMutex
	userOlderUpdateCache          = make(map[string]updateCache)
	userOlderUpsertCacheMut       sync.RWMutex
	userOlderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userOlderAfterSelectHooks []UserOlderHook

var userOlderBeforeInsertHooks []UserOlderHook
var userOlderAfterInsertHooks []UserOlderHook

var userOlderBeforeUpdateHooks []UserOlderHook
var userOlderAfterUpdateHooks []UserOlderHook

var userOlderBeforeDeleteHooks []UserOlderHook
var userOlderAfterDeleteHooks []UserOlderHook

var userOlderBeforeUpsertHooks []UserOlderHook
var userOlderAfterUpsertHooks []UserOlderHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserOlder) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserOlder) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserOlder) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserOlder) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserOlder) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserOlder) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserOlder) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserOlder) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserOlder) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userOlderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserOlderHook registers your hook function for all future operations.
func AddUserOlderHook(hookPoint boil.HookPoint, userOlderHook UserOlderHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userOlderAfterSelectHooks = append(userOlderAfterSelectHooks, userOlderHook)
	case boil.BeforeInsertHook:
		userOlderBeforeInsertHooks = append(userOlderBeforeInsertHooks, userOlderHook)
	case boil.AfterInsertHook:
		userOlderAfterInsertHooks = append(userOlderAfterInsertHooks, userOlderHook)
	case boil.BeforeUpdateHook:
		userOlderBeforeUpdateHooks = append(userOlderBeforeUpdateHooks, userOlderHook)
	case boil.AfterUpdateHook:
		userOlderAfterUpdateHooks = append(userOlderAfterUpdateHooks, userOlderHook)
	case boil.BeforeDeleteHook:
		userOlderBeforeDeleteHooks = append(userOlderBeforeDeleteHooks, userOlderHook)
	case boil.AfterDeleteHook:
		userOlderAfterDeleteHooks = append(userOlderAfterDeleteHooks, userOlderHook)
	case boil.BeforeUpsertHook:
		userOlderBeforeUpsertHooks = append(userOlderBeforeUpsertHooks, userOlderHook)
	case boil.AfterUpsertHook:
		userOlderAfterUpsertHooks = append(userOlderAfterUpsertHooks, userOlderHook)
	}
}

// OneG returns a single userOlder record from the query using the global executor.
func (q userOlderQuery) OneG() (*UserOlder, error) {
	return q.One(boil.GetDB())
}

// One returns a single userOlder record from the query.
func (q userOlderQuery) One(exec boil.Executor) (*UserOlder, error) {
	o := &UserOlder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_olders")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all UserOlder records from the query using the global executor.
func (q userOlderQuery) AllG() (UserOlderSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all UserOlder records from the query.
func (q userOlderQuery) All(exec boil.Executor) (UserOlderSlice, error) {
	var o []*UserOlder

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserOlder slice")
	}

	if len(userOlderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all UserOlder records in the query, and panics on error.
func (q userOlderQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all UserOlder records in the query.
func (q userOlderQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_olders rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q userOlderQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q userOlderQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_olders exists")
	}

	return count > 0, nil
}

// Older pointed to by the foreign key.
func (o *UserOlder) Older(mods ...qm.QueryMod) olderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OlderID),
	}

	queryMods = append(queryMods, mods...)

	query := Olders(queryMods...)
	queries.SetFrom(query.Query, "\"olders\"")

	return query
}

// User pointed to by the foreign key.
func (o *UserOlder) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadOlder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userOlderL) LoadOlder(e boil.Executor, singular bool, maybeUserOlder interface{}, mods queries.Applicator) error {
	var slice []*UserOlder
	var object *UserOlder

	if singular {
		object = maybeUserOlder.(*UserOlder)
	} else {
		slice = *maybeUserOlder.(*[]*UserOlder)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userOlderR{}
		}
		if !queries.IsNil(object.OlderID) {
			args = append(args, object.OlderID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userOlderR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.OlderID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.OlderID) {
				args = append(args, obj.OlderID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`olders`),
		qm.WhereIn(`olders.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Older")
	}

	var resultSlice []*Older
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Older")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for olders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for olders")
	}

	if len(userOlderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Older = foreign
		if foreign.R == nil {
			foreign.R = &olderR{}
		}
		foreign.R.UserOlders = append(foreign.R.UserOlders, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OlderID, foreign.ID) {
				local.R.Older = foreign
				if foreign.R == nil {
					foreign.R = &olderR{}
				}
				foreign.R.UserOlders = append(foreign.R.UserOlders, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userOlderL) LoadUser(e boil.Executor, singular bool, maybeUserOlder interface{}, mods queries.Applicator) error {
	var slice []*UserOlder
	var object *UserOlder

	if singular {
		object = maybeUserOlder.(*UserOlder)
	} else {
		slice = *maybeUserOlder.(*[]*UserOlder)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userOlderR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userOlderR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userOlderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserOlders = append(foreign.R.UserOlders, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserOlders = append(foreign.R.UserOlders, local)
				break
			}
		}
	}

	return nil
}

// SetOlderG of the userOlder to the related item.
// Sets o.R.Older to related.
// Adds o to related.R.UserOlders.
// Uses the global database handle.
func (o *UserOlder) SetOlderG(insert bool, related *Older) error {
	return o.SetOlder(boil.GetDB(), insert, related)
}

// SetOlder of the userOlder to the related item.
// Sets o.R.Older to related.
// Adds o to related.R.UserOlders.
func (o *UserOlder) SetOlder(exec boil.Executor, insert bool, related *Older) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_olders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"older_id"}),
		strmangle.WhereClause("\"", "\"", 2, userOlderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.OlderID, related.ID)
	if o.R == nil {
		o.R = &userOlderR{
			Older: related,
		}
	} else {
		o.R.Older = related
	}

	if related.R == nil {
		related.R = &olderR{
			UserOlders: UserOlderSlice{o},
		}
	} else {
		related.R.UserOlders = append(related.R.UserOlders, o)
	}

	return nil
}

// RemoveOlderG relationship.
// Sets o.R.Older to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *UserOlder) RemoveOlderG(related *Older) error {
	return o.RemoveOlder(boil.GetDB(), related)
}

// RemoveOlder relationship.
// Sets o.R.Older to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *UserOlder) RemoveOlder(exec boil.Executor, related *Older) error {
	var err error

	queries.SetScanner(&o.OlderID, nil)
	if _, err = o.Update(exec, boil.Whitelist("older_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Older = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.UserOlders {
		if queries.Equal(o.OlderID, ri.OlderID) {
			continue
		}

		ln := len(related.R.UserOlders)
		if ln > 1 && i < ln-1 {
			related.R.UserOlders[i] = related.R.UserOlders[ln-1]
		}
		related.R.UserOlders = related.R.UserOlders[:ln-1]
		break
	}
	return nil
}

// SetUserG of the userOlder to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserOlders.
// Uses the global database handle.
func (o *UserOlder) SetUserG(insert bool, related *User) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUser of the userOlder to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserOlders.
func (o *UserOlder) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_olders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userOlderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &userOlderR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserOlders: UserOlderSlice{o},
		}
	} else {
		related.R.UserOlders = append(related.R.UserOlders, o)
	}

	return nil
}

// RemoveUserG relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *UserOlder) RemoveUserG(related *User) error {
	return o.RemoveUser(boil.GetDB(), related)
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *UserOlder) RemoveUser(exec boil.Executor, related *User) error {
	var err error

	queries.SetScanner(&o.UserID, nil)
	if _, err = o.Update(exec, boil.Whitelist("user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.User = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.UserOlders {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.UserOlders)
		if ln > 1 && i < ln-1 {
			related.R.UserOlders[i] = related.R.UserOlders[ln-1]
		}
		related.R.UserOlders = related.R.UserOlders[:ln-1]
		break
	}
	return nil
}

// UserOlders retrieves all the records using an executor.
func UserOlders(mods ...qm.QueryMod) userOlderQuery {
	mods = append(mods, qm.From("\"user_olders\""))
	return userOlderQuery{NewQuery(mods...)}
}

// FindUserOlderG retrieves a single record by ID.
func FindUserOlderG(iD int, selectCols ...string) (*UserOlder, error) {
	return FindUserOlder(boil.GetDB(), iD, selectCols...)
}

// FindUserOlder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserOlder(exec boil.Executor, iD int, selectCols ...string) (*UserOlder, error) {
	userOlderObj := &UserOlder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_olders\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, userOlderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_olders")
	}

	if err = userOlderObj.doAfterSelectHooks(exec); err != nil {
		return userOlderObj, err
	}

	return userOlderObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UserOlder) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserOlder) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_olders provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userOlderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userOlderInsertCacheMut.RLock()
	cache, cached := userOlderInsertCache[key]
	userOlderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userOlderAllColumns,
			userOlderColumnsWithDefault,
			userOlderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userOlderType, userOlderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userOlderType, userOlderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_olders\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_olders\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_olders")
	}

	if !cached {
		userOlderInsertCacheMut.Lock()
		userOlderInsertCache[key] = cache
		userOlderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single UserOlder record using the global executor.
// See Update for more documentation.
func (o *UserOlder) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the UserOlder.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserOlder) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userOlderUpdateCacheMut.RLock()
	cache, cached := userOlderUpdateCache[key]
	userOlderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userOlderAllColumns,
			userOlderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_olders, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_olders\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userOlderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userOlderType, userOlderMapping, append(wl, userOlderPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update user_olders row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_olders")
	}

	if !cached {
		userOlderUpdateCacheMut.Lock()
		userOlderUpdateCache[key] = cache
		userOlderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q userOlderQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userOlderQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_olders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_olders")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserOlderSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserOlderSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userOlderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_olders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userOlderPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userOlder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userOlder")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UserOlder) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserOlder) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_olders provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userOlderColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userOlderUpsertCacheMut.RLock()
	cache, cached := userOlderUpsertCache[key]
	userOlderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userOlderAllColumns,
			userOlderColumnsWithDefault,
			userOlderColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userOlderAllColumns,
			userOlderPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_olders, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userOlderPrimaryKeyColumns))
			copy(conflict, userOlderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_olders\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userOlderType, userOlderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userOlderType, userOlderMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert user_olders")
	}

	if !cached {
		userOlderUpsertCacheMut.Lock()
		userOlderUpsertCache[key] = cache
		userOlderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single UserOlder record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UserOlder) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single UserOlder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserOlder) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserOlder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userOlderPrimaryKeyMapping)
	sql := "DELETE FROM \"user_olders\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_olders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_olders")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q userOlderQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q userOlderQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userOlderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_olders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_olders")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserOlderSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserOlderSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userOlderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userOlderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_olders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userOlderPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userOlder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_olders")
	}

	if len(userOlderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UserOlder) ReloadG() error {
	if o == nil {
		return errors.New("models: no UserOlder provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserOlder) Reload(exec boil.Executor) error {
	ret, err := FindUserOlder(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserOlderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty UserOlderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserOlderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserOlderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userOlderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_olders\".* FROM \"user_olders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userOlderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserOlderSlice")
	}

	*o = slice

	return nil
}

// UserOlderExistsG checks if the UserOlder row exists.
func UserOlderExistsG(iD int) (bool, error) {
	return UserOlderExists(boil.GetDB(), iD)
}

// UserOlderExists checks if the UserOlder row exists.
func UserOlderExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_olders\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_olders exists")
	}

	return exists, nil
}

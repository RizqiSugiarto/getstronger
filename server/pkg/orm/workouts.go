// Code generated by SQLBoiler 4.17.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Workout is an object representing the database table.
type Workout struct {
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	FinishedAt time.Time `boil:"finished_at" json:"finished_at" toml:"finished_at" yaml:"finished_at"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	StartedAt  time.Time `boil:"started_at" json:"started_at" toml:"started_at" yaml:"started_at"`

	R *workoutR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L workoutL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WorkoutColumns = struct {
	ID         string
	UserID     string
	FinishedAt string
	CreatedAt  string
	Name       string
	StartedAt  string
}{
	ID:         "id",
	UserID:     "user_id",
	FinishedAt: "finished_at",
	CreatedAt:  "created_at",
	Name:       "name",
	StartedAt:  "started_at",
}

var WorkoutTableColumns = struct {
	ID         string
	UserID     string
	FinishedAt string
	CreatedAt  string
	Name       string
	StartedAt  string
}{
	ID:         "workouts.id",
	UserID:     "workouts.user_id",
	FinishedAt: "workouts.finished_at",
	CreatedAt:  "workouts.created_at",
	Name:       "workouts.name",
	StartedAt:  "workouts.started_at",
}

// Generated where

var WorkoutWhere = struct {
	ID         whereHelperstring
	UserID     whereHelperstring
	FinishedAt whereHelpertime_Time
	CreatedAt  whereHelpertime_Time
	Name       whereHelperstring
	StartedAt  whereHelpertime_Time
}{
	ID:         whereHelperstring{field: "\"getstronger\".\"workouts\".\"id\""},
	UserID:     whereHelperstring{field: "\"getstronger\".\"workouts\".\"user_id\""},
	FinishedAt: whereHelpertime_Time{field: "\"getstronger\".\"workouts\".\"finished_at\""},
	CreatedAt:  whereHelpertime_Time{field: "\"getstronger\".\"workouts\".\"created_at\""},
	Name:       whereHelperstring{field: "\"getstronger\".\"workouts\".\"name\""},
	StartedAt:  whereHelpertime_Time{field: "\"getstronger\".\"workouts\".\"started_at\""},
}

// WorkoutRels is where relationship names are stored.
var WorkoutRels = struct {
	User string
	Sets string
}{
	User: "User",
	Sets: "Sets",
}

// workoutR is where relationships are stored.
type workoutR struct {
	User *User    `boil:"User" json:"User" toml:"User" yaml:"User"`
	Sets SetSlice `boil:"Sets" json:"Sets" toml:"Sets" yaml:"Sets"`
}

// NewStruct creates a new relationship struct
func (*workoutR) NewStruct() *workoutR {
	return &workoutR{}
}

func (r *workoutR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *workoutR) GetSets() SetSlice {
	if r == nil {
		return nil
	}
	return r.Sets
}

// workoutL is where Load methods for each relationship are stored.
type workoutL struct{}

var (
	workoutAllColumns            = []string{"id", "user_id", "finished_at", "created_at", "name", "started_at"}
	workoutColumnsWithoutDefault = []string{"user_id", "finished_at", "name", "started_at"}
	workoutColumnsWithDefault    = []string{"id", "created_at"}
	workoutPrimaryKeyColumns     = []string{"id"}
	workoutGeneratedColumns      = []string{}
)

type (
	// WorkoutSlice is an alias for a slice of pointers to Workout.
	// This should almost always be used instead of []Workout.
	WorkoutSlice []*Workout
	// WorkoutHook is the signature for custom Workout hook methods
	WorkoutHook func(context.Context, boil.ContextExecutor, *Workout) error

	workoutQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	workoutType                 = reflect.TypeOf(&Workout{})
	workoutMapping              = queries.MakeStructMapping(workoutType)
	workoutPrimaryKeyMapping, _ = queries.BindMapping(workoutType, workoutMapping, workoutPrimaryKeyColumns)
	workoutInsertCacheMut       sync.RWMutex
	workoutInsertCache          = make(map[string]insertCache)
	workoutUpdateCacheMut       sync.RWMutex
	workoutUpdateCache          = make(map[string]updateCache)
	workoutUpsertCacheMut       sync.RWMutex
	workoutUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var workoutAfterSelectMu sync.Mutex
var workoutAfterSelectHooks []WorkoutHook

var workoutBeforeInsertMu sync.Mutex
var workoutBeforeInsertHooks []WorkoutHook
var workoutAfterInsertMu sync.Mutex
var workoutAfterInsertHooks []WorkoutHook

var workoutBeforeUpdateMu sync.Mutex
var workoutBeforeUpdateHooks []WorkoutHook
var workoutAfterUpdateMu sync.Mutex
var workoutAfterUpdateHooks []WorkoutHook

var workoutBeforeDeleteMu sync.Mutex
var workoutBeforeDeleteHooks []WorkoutHook
var workoutAfterDeleteMu sync.Mutex
var workoutAfterDeleteHooks []WorkoutHook

var workoutBeforeUpsertMu sync.Mutex
var workoutBeforeUpsertHooks []WorkoutHook
var workoutAfterUpsertMu sync.Mutex
var workoutAfterUpsertHooks []WorkoutHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Workout) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Workout) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Workout) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Workout) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Workout) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Workout) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Workout) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Workout) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Workout) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range workoutAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWorkoutHook registers your hook function for all future operations.
func AddWorkoutHook(hookPoint boil.HookPoint, workoutHook WorkoutHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		workoutAfterSelectMu.Lock()
		workoutAfterSelectHooks = append(workoutAfterSelectHooks, workoutHook)
		workoutAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		workoutBeforeInsertMu.Lock()
		workoutBeforeInsertHooks = append(workoutBeforeInsertHooks, workoutHook)
		workoutBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		workoutAfterInsertMu.Lock()
		workoutAfterInsertHooks = append(workoutAfterInsertHooks, workoutHook)
		workoutAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		workoutBeforeUpdateMu.Lock()
		workoutBeforeUpdateHooks = append(workoutBeforeUpdateHooks, workoutHook)
		workoutBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		workoutAfterUpdateMu.Lock()
		workoutAfterUpdateHooks = append(workoutAfterUpdateHooks, workoutHook)
		workoutAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		workoutBeforeDeleteMu.Lock()
		workoutBeforeDeleteHooks = append(workoutBeforeDeleteHooks, workoutHook)
		workoutBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		workoutAfterDeleteMu.Lock()
		workoutAfterDeleteHooks = append(workoutAfterDeleteHooks, workoutHook)
		workoutAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		workoutBeforeUpsertMu.Lock()
		workoutBeforeUpsertHooks = append(workoutBeforeUpsertHooks, workoutHook)
		workoutBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		workoutAfterUpsertMu.Lock()
		workoutAfterUpsertHooks = append(workoutAfterUpsertHooks, workoutHook)
		workoutAfterUpsertMu.Unlock()
	}
}

// One returns a single workout record from the query.
func (q workoutQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Workout, error) {
	o := &Workout{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for workouts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Workout records from the query.
func (q workoutQuery) All(ctx context.Context, exec boil.ContextExecutor) (WorkoutSlice, error) {
	var o []*Workout

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Workout slice")
	}

	if len(workoutAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Workout records in the query.
func (q workoutQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count workouts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q workoutQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if workouts exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Workout) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Sets retrieves all the set's Sets with an executor.
func (o *Workout) Sets(mods ...qm.QueryMod) setQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"getstronger\".\"sets\".\"workout_id\"=?", o.ID),
	)

	return Sets(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (workoutL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWorkout interface{}, mods queries.Applicator) error {
	var slice []*Workout
	var object *Workout

	if singular {
		var ok bool
		object, ok = maybeWorkout.(*Workout)
		if !ok {
			object = new(Workout)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeWorkout)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeWorkout))
			}
		}
	} else {
		s, ok := maybeWorkout.(*[]*Workout)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeWorkout)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeWorkout))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &workoutR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &workoutR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`getstronger.users`),
		qm.WhereIn(`getstronger.users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
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

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
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
		foreign.R.Workouts = append(foreign.R.Workouts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Workouts = append(foreign.R.Workouts, local)
				break
			}
		}
	}

	return nil
}

// LoadSets allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (workoutL) LoadSets(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWorkout interface{}, mods queries.Applicator) error {
	var slice []*Workout
	var object *Workout

	if singular {
		var ok bool
		object, ok = maybeWorkout.(*Workout)
		if !ok {
			object = new(Workout)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeWorkout)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeWorkout))
			}
		}
	} else {
		s, ok := maybeWorkout.(*[]*Workout)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeWorkout)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeWorkout))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &workoutR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &workoutR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`getstronger.sets`),
		qm.WhereIn(`getstronger.sets.workout_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sets")
	}

	var resultSlice []*Set
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sets")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sets")
	}

	if len(setAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Sets = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &setR{}
			}
			foreign.R.Workout = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.WorkoutID {
				local.R.Sets = append(local.R.Sets, foreign)
				if foreign.R == nil {
					foreign.R = &setR{}
				}
				foreign.R.Workout = local
				break
			}
		}
	}

	return nil
}

// SetUser of the workout to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Workouts.
func (o *Workout) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"getstronger\".\"workouts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, workoutPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &workoutR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Workouts: WorkoutSlice{o},
		}
	} else {
		related.R.Workouts = append(related.R.Workouts, o)
	}

	return nil
}

// AddSets adds the given related objects to the existing relationships
// of the workout, optionally inserting them as new records.
// Appends related to o.R.Sets.
// Sets related.R.Workout appropriately.
func (o *Workout) AddSets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Set) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.WorkoutID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"getstronger\".\"sets\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"workout_id"}),
				strmangle.WhereClause("\"", "\"", 2, setPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.WorkoutID = o.ID
		}
	}

	if o.R == nil {
		o.R = &workoutR{
			Sets: related,
		}
	} else {
		o.R.Sets = append(o.R.Sets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &setR{
				Workout: o,
			}
		} else {
			rel.R.Workout = o
		}
	}
	return nil
}

// Workouts retrieves all the records using an executor.
func Workouts(mods ...qm.QueryMod) workoutQuery {
	mods = append(mods, qm.From("\"getstronger\".\"workouts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"getstronger\".\"workouts\".*"})
	}

	return workoutQuery{q}
}

// FindWorkout retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWorkout(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Workout, error) {
	workoutObj := &Workout{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"getstronger\".\"workouts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, workoutObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from workouts")
	}

	if err = workoutObj.doAfterSelectHooks(ctx, exec); err != nil {
		return workoutObj, err
	}

	return workoutObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Workout) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no workouts provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(workoutColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	workoutInsertCacheMut.RLock()
	cache, cached := workoutInsertCache[key]
	workoutInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			workoutAllColumns,
			workoutColumnsWithDefault,
			workoutColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(workoutType, workoutMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(workoutType, workoutMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"getstronger\".\"workouts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"getstronger\".\"workouts\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "orm: unable to insert into workouts")
	}

	if !cached {
		workoutInsertCacheMut.Lock()
		workoutInsertCache[key] = cache
		workoutInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Workout.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Workout) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	workoutUpdateCacheMut.RLock()
	cache, cached := workoutUpdateCache[key]
	workoutUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			workoutAllColumns,
			workoutPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update workouts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"getstronger\".\"workouts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, workoutPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(workoutType, workoutMapping, append(wl, workoutPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update workouts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for workouts")
	}

	if !cached {
		workoutUpdateCacheMut.Lock()
		workoutUpdateCache[key] = cache
		workoutUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q workoutQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for workouts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for workouts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WorkoutSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), workoutPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"getstronger\".\"workouts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, workoutPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in workout slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all workout")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Workout) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("orm: no workouts provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(workoutColumnsWithDefault, o)

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

	workoutUpsertCacheMut.RLock()
	cache, cached := workoutUpsertCache[key]
	workoutUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			workoutAllColumns,
			workoutColumnsWithDefault,
			workoutColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			workoutAllColumns,
			workoutPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert workouts, could not build update column list")
		}

		ret := strmangle.SetComplement(workoutAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(workoutPrimaryKeyColumns) == 0 {
				return errors.New("orm: unable to upsert workouts, could not build conflict column list")
			}

			conflict = make([]string, len(workoutPrimaryKeyColumns))
			copy(conflict, workoutPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"getstronger\".\"workouts\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(workoutType, workoutMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(workoutType, workoutMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "orm: unable to upsert workouts")
	}

	if !cached {
		workoutUpsertCacheMut.Lock()
		workoutUpsertCache[key] = cache
		workoutUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Workout record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Workout) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Workout provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), workoutPrimaryKeyMapping)
	sql := "DELETE FROM \"getstronger\".\"workouts\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from workouts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for workouts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q workoutQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no workoutQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from workouts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for workouts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WorkoutSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(workoutBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), workoutPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"getstronger\".\"workouts\" WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, workoutPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from workout slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for workouts")
	}

	if len(workoutAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Workout) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWorkout(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WorkoutSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WorkoutSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), workoutPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"getstronger\".\"workouts\".* FROM \"getstronger\".\"workouts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, workoutPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in WorkoutSlice")
	}

	*o = slice

	return nil
}

// WorkoutExists checks if the Workout row exists.
func WorkoutExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"getstronger\".\"workouts\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if workouts exists")
	}

	return exists, nil
}

// Exists checks if the Workout row exists.
func (o *Workout) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return WorkoutExists(ctx, exec, o.ID)
}

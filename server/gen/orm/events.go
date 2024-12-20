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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Event is an object representing the database table.
type Event struct {
	ID        string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Topic     EventTopic `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	Payload   types.JSON `boil:"payload" json:"payload" toml:"payload" yaml:"payload"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *eventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventColumns = struct {
	ID        string
	Topic     string
	Payload   string
	CreatedAt string
}{
	ID:        "id",
	Topic:     "topic",
	Payload:   "payload",
	CreatedAt: "created_at",
}

var EventTableColumns = struct {
	ID        string
	Topic     string
	Payload   string
	CreatedAt string
}{
	ID:        "events.id",
	Topic:     "events.topic",
	Payload:   "events.payload",
	CreatedAt: "events.created_at",
}

// Generated where

type whereHelperEventTopic struct{ field string }

func (w whereHelperEventTopic) EQ(x EventTopic) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperEventTopic) NEQ(x EventTopic) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperEventTopic) LT(x EventTopic) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperEventTopic) LTE(x EventTopic) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperEventTopic) GT(x EventTopic) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperEventTopic) GTE(x EventTopic) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperEventTopic) IN(slice []EventTopic) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperEventTopic) NIN(slice []EventTopic) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var EventWhere = struct {
	ID        whereHelperstring
	Topic     whereHelperEventTopic
	Payload   whereHelpertypes_JSON
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"getstronger\".\"events\".\"id\""},
	Topic:     whereHelperEventTopic{field: "\"getstronger\".\"events\".\"topic\""},
	Payload:   whereHelpertypes_JSON{field: "\"getstronger\".\"events\".\"payload\""},
	CreatedAt: whereHelpertime_Time{field: "\"getstronger\".\"events\".\"created_at\""},
}

// EventRels is where relationship names are stored.
var EventRels = struct {
}{}

// eventR is where relationships are stored.
type eventR struct {
}

// NewStruct creates a new relationship struct
func (*eventR) NewStruct() *eventR {
	return &eventR{}
}

// eventL is where Load methods for each relationship are stored.
type eventL struct{}

var (
	eventAllColumns            = []string{"id", "topic", "payload", "created_at"}
	eventColumnsWithoutDefault = []string{"topic", "payload"}
	eventColumnsWithDefault    = []string{"id", "created_at"}
	eventPrimaryKeyColumns     = []string{"id"}
	eventGeneratedColumns      = []string{}
)

type (
	// EventSlice is an alias for a slice of pointers to Event.
	// This should almost always be used instead of []Event.
	EventSlice []*Event
	// EventHook is the signature for custom Event hook methods
	EventHook func(context.Context, boil.ContextExecutor, *Event) error

	eventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventType                 = reflect.TypeOf(&Event{})
	eventMapping              = queries.MakeStructMapping(eventType)
	eventPrimaryKeyMapping, _ = queries.BindMapping(eventType, eventMapping, eventPrimaryKeyColumns)
	eventInsertCacheMut       sync.RWMutex
	eventInsertCache          = make(map[string]insertCache)
	eventUpdateCacheMut       sync.RWMutex
	eventUpdateCache          = make(map[string]updateCache)
	eventUpsertCacheMut       sync.RWMutex
	eventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventAfterSelectMu sync.Mutex
var eventAfterSelectHooks []EventHook

var eventBeforeInsertMu sync.Mutex
var eventBeforeInsertHooks []EventHook
var eventAfterInsertMu sync.Mutex
var eventAfterInsertHooks []EventHook

var eventBeforeUpdateMu sync.Mutex
var eventBeforeUpdateHooks []EventHook
var eventAfterUpdateMu sync.Mutex
var eventAfterUpdateHooks []EventHook

var eventBeforeDeleteMu sync.Mutex
var eventBeforeDeleteHooks []EventHook
var eventAfterDeleteMu sync.Mutex
var eventAfterDeleteHooks []EventHook

var eventBeforeUpsertMu sync.Mutex
var eventBeforeUpsertHooks []EventHook
var eventAfterUpsertMu sync.Mutex
var eventAfterUpsertHooks []EventHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Event) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Event) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Event) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Event) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Event) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Event) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Event) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Event) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Event) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventHook registers your hook function for all future operations.
func AddEventHook(hookPoint boil.HookPoint, eventHook EventHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventAfterSelectMu.Lock()
		eventAfterSelectHooks = append(eventAfterSelectHooks, eventHook)
		eventAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		eventBeforeInsertMu.Lock()
		eventBeforeInsertHooks = append(eventBeforeInsertHooks, eventHook)
		eventBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		eventAfterInsertMu.Lock()
		eventAfterInsertHooks = append(eventAfterInsertHooks, eventHook)
		eventAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		eventBeforeUpdateMu.Lock()
		eventBeforeUpdateHooks = append(eventBeforeUpdateHooks, eventHook)
		eventBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		eventAfterUpdateMu.Lock()
		eventAfterUpdateHooks = append(eventAfterUpdateHooks, eventHook)
		eventAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		eventBeforeDeleteMu.Lock()
		eventBeforeDeleteHooks = append(eventBeforeDeleteHooks, eventHook)
		eventBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		eventAfterDeleteMu.Lock()
		eventAfterDeleteHooks = append(eventAfterDeleteHooks, eventHook)
		eventAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		eventBeforeUpsertMu.Lock()
		eventBeforeUpsertHooks = append(eventBeforeUpsertHooks, eventHook)
		eventBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		eventAfterUpsertMu.Lock()
		eventAfterUpsertHooks = append(eventAfterUpsertHooks, eventHook)
		eventAfterUpsertMu.Unlock()
	}
}

// One returns a single event record from the query.
func (q eventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Event, error) {
	o := &Event{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for events")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Event records from the query.
func (q eventQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventSlice, error) {
	var o []*Event

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Event slice")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Event records in the query.
func (q eventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if events exists")
	}

	return count > 0, nil
}

// Events retrieves all the records using an executor.
func Events(mods ...qm.QueryMod) eventQuery {
	mods = append(mods, qm.From("\"getstronger\".\"events\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"getstronger\".\"events\".*"})
	}

	return eventQuery{q}
}

// FindEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEvent(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Event, error) {
	eventObj := &Event{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"getstronger\".\"events\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from events")
	}

	if err = eventObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventObj, err
	}

	return eventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Event) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no events provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventInsertCacheMut.RLock()
	cache, cached := eventInsertCache[key]
	eventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventType, eventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"getstronger\".\"events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"getstronger\".\"events\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into events")
	}

	if !cached {
		eventInsertCacheMut.Lock()
		eventInsertCache[key] = cache
		eventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Event.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Event) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventUpdateCacheMut.RLock()
	cache, cached := eventUpdateCache[key]
	eventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"getstronger\".\"events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, eventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, append(wl, eventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for events")
	}

	if !cached {
		eventUpdateCacheMut.Lock()
		eventUpdateCache[key] = cache
		eventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"getstronger\".\"events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, eventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all event")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Event) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("orm: no events provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

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

	eventUpsertCacheMut.RLock()
	cache, cached := eventUpsertCache[key]
	eventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert events, could not build update column list")
		}

		ret := strmangle.SetComplement(eventAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(eventPrimaryKeyColumns) == 0 {
				return errors.New("orm: unable to upsert events, could not build conflict column list")
			}

			conflict = make([]string, len(eventPrimaryKeyColumns))
			copy(conflict, eventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"getstronger\".\"events\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventType, eventMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert events")
	}

	if !cached {
		eventUpsertCacheMut.Lock()
		eventUpsertCache[key] = cache
		eventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Event record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Event) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Event provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventPrimaryKeyMapping)
	sql := "DELETE FROM \"getstronger\".\"events\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for events")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no eventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"getstronger\".\"events\" WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, eventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for events")
	}

	if len(eventAfterDeleteHooks) != 0 {
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
func (o *Event) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"getstronger\".\"events\".* FROM \"getstronger\".\"events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in EventSlice")
	}

	*o = slice

	return nil
}

// EventExists checks if the Event row exists.
func EventExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"getstronger\".\"events\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if events exists")
	}

	return exists, nil
}

// Exists checks if the Event row exists.
func (o *Event) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventExists(ctx, exec, o.ID)
}

// Code generated by SQLBoiler 4.17.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("orm: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type NotificationType string

// Enum values for NotificationType
const (
	NotificationTypeFollow         NotificationType = "Follow"
	NotificationTypeWorkoutComment NotificationType = "WorkoutComment"
)

func AllNotificationType() []NotificationType {
	return []NotificationType{
		NotificationTypeFollow,
		NotificationTypeWorkoutComment,
	}
}

func (e NotificationType) IsValid() error {
	switch e {
	case NotificationTypeFollow, NotificationTypeWorkoutComment:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e NotificationType) String() string {
	return string(e)
}

func (e NotificationType) Ordinal() int {
	switch e {
	case NotificationTypeFollow:
		return 0
	case NotificationTypeWorkoutComment:
		return 1

	default:
		panic(errors.New("enum is not valid"))
	}
}

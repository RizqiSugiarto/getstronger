package repo

import (
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/crlssn/getstronger/server/gen/orm"
)

type ModelItem interface {
	*orm.Workout | *orm.Exercise | *orm.User | *orm.Routine | *orm.Set | *orm.WorkoutComment | *orm.Notification
}

type ModelSlice[T any] interface {
	~[]T
}

type Pagination[Item ModelItem, Slice ModelSlice[Item]] struct {
	Items         Slice
	NextPageToken []byte
}

func PaginateSlice[Item ModelItem, Slice ModelSlice[Item]](
	items Slice, limit int, createdAt func(Item) time.Time,
) (*Pagination[Item, Slice], error) {
	if len(items) <= limit {
		return &Pagination[Item, Slice]{
			Items:         items,
			NextPageToken: nil,
		}, nil
	}

	items = items[:limit]
	nextPageToken, err := json.Marshal(PageTokenCreatedAt(createdAt(items[len(items)-1])))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal page token: %w", err)
	}

	return &Pagination[Item, Slice]{
		Items:         items,
		NextPageToken: nextPageToken,
	}, nil
}

func PageTokenCreatedAt(t time.Time) PageToken {
	return PageToken{
		// Truncate to microseconds to unify precision across different databases.
		CreatedAt: t.Truncate(time.Microsecond),
	}
}

type PageToken struct {
	CreatedAt time.Time `json:"createdAt"`
}

type updateOpt interface {
	UpdateRoutineOpt | UpdateAuthOpt | UpdateExerciseOpt | UpdateWorkoutOpt
}

var (
	ErrUpdateNoColumns       = fmt.Errorf("update opt: no columns")
	ErrUpdateRowsAffected    = fmt.Errorf("update opt: rows affected")
	ErrUpdateDuplicateColumn = fmt.Errorf("update opt: duplicate column")
)

func updateColumnsFromOpts[T updateOpt](opts []T) (orm.M, error) {
	if len(opts) == 0 {
		return nil, ErrUpdateNoColumns
	}

	columns := make(orm.M, len(opts))
	for _, opt := range opts {
		column, err := opt()
		if err != nil {
			return nil, fmt.Errorf("update opt: %w", err)
		}

		for key, value := range column {
			if columns[key] != nil {
				return nil, fmt.Errorf("%w: %s", ErrUpdateDuplicateColumn, key)
			}

			columns[key] = value
		}
	}

	return columns, nil
}

func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return hash, nil
}

func MustHashPassword(password string) []byte {
	hash, err := hashPassword(password)
	if err != nil {
		panic(err)
	}

	return hash
}

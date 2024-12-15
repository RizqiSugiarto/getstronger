package testdb

import (
	"context"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/crlssn/getstronger/server/pkg/orm"
)

type UserOpt func(event *orm.User)

func (f *Factory) NewUser(opts ...UserOpt) *orm.User {
	m := &orm.User{
		AuthID:    "",
		FirstName: f.faker.FirstName(),
		LastName:  f.faker.LastName(),
		CreatedAt: time.Time{},
	}

	for _, opt := range opts {
		opt(m)
	}

	if m.AuthID == "" {
		m.AuthID = f.NewAuth().ID
	}

	boil.DebugMode = f.debug
	if err := m.Insert(context.Background(), f.db, boil.Infer()); err != nil {
		panic(fmt.Errorf("failed to insert user: %w", err))
	}
	boil.DebugMode = false

	return m
}

func UserID(id string) UserOpt {
	return func(m *orm.User) {
		m.ID = id
	}
}

func UserAuthID(authID string) UserOpt {
	return func(m *orm.User) {
		m.AuthID = authID
	}
}

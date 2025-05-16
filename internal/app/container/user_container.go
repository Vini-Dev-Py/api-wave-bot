package container

import (
	"database/sql"

	"github.com/google/wire"
)

type UserContainer struct {
}

var UserSet = wire.NewSet(
	NewUserContainer,
)

func NewUserContainer(db *sql.DB) *UserContainer {

	return &UserContainer{}
}

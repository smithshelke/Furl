package core

import (
	"github.com/smithshelke/flur/internal/core/graphs"
	"github.com/smithshelke/flur/internal/core/users"
)

type Core struct {
	UserAPI  *users.UserAPI
	GraphAPI *graphs.GraphsAPI
}

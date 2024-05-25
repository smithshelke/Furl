package ports

import "github.com/smithshelke/flur/internal/ports/db"

type Db struct {
	Users     db.Users
	Blocks    db.Blocks
	Workflows db.Workflows
}

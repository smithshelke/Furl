package main

import (
	"context"

	"github.com/smithshelke/flur/internal/adapters/config"
	"github.com/smithshelke/flur/internal/adapters/neo4j"
	"github.com/smithshelke/flur/internal/adapters/web"
	"github.com/smithshelke/flur/internal/core"
	"github.com/smithshelke/flur/internal/core/graphs"
	"github.com/smithshelke/flur/internal/core/users"
	"github.com/smithshelke/flur/internal/ports"
)

func main() {
	config := config.Initialize()
	neo4jInstance := neo4j.Initialize(context.Background(), config.Neo4j)
	repos := neo4jInstance.GetRepo()
	core := InitializeCore(repos)
	app := web.App{
		Config: config,
		Core:   core,
	}
	app.Initialize()
}

func InitializeCore(repos ports.Db) *core.Core {
	userAPI := users.NewUserAPI(&repos.Users)
	graphsAPI := graphs.NewGraphsAPI(&repos.Workflows, &repos.Blocks)
	return &core.Core{
		UserAPI:  userAPI,
		GraphAPI: graphsAPI,
	}
}

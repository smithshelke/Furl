package neo4j

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/smithshelke/flur/internal/adapters/config"
	"github.com/smithshelke/flur/internal/adapters/neo4j/repos"
	"github.com/smithshelke/flur/internal/ports"
)

type Neo4jInstance struct {
	Driver neo4j.DriverWithContext
}

func Initialize(ctx context.Context, config config.Neo4jConfig) *Neo4jInstance {
	driver, err := neo4j.NewDriverWithContext(
		config.Addr,
		neo4j.BasicAuth(config.Username, config.Password, ""))

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Neo4j connected")
	return &Neo4jInstance{
		Driver: driver,
	}
}

func (n *Neo4jInstance) GetRepo() ports.Db {
	driver := n.Driver
	return ports.Db{
		Users:     repos.NewUsersRepo(driver),
		Workflows: repos.NewWorkflowsRepo(driver),
		Blocks:    repos.NewBlocksRepo(driver),
	}
}

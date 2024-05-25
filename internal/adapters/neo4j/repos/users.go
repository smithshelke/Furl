package repos

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/smithshelke/flur/internal/core/domain"
)

type UsersRepo struct {
	Driver neo4j.DriverWithContext
}

func NewUsersRepo(driver neo4j.DriverWithContext) *UsersRepo {
	return &UsersRepo{
		Driver: driver,
	}
}

func (r *UsersRepo) Create(ctx context.Context, users []*domain.User) ([]*domain.User, error) {
	for _, u := range users {
		_, err := neo4j.ExecuteQuery(ctx, r.Driver,
			"MERGE (u:User {name: $name}) RETURN u",
			map[string]any{
				"name": u.Name,
			}, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))
		if err != nil {
			log.Printf("Failed to create user: %+v\n", err)
			return nil, err
		}
	}
	return nil, nil
}

func (r *UsersRepo) List(ctx context.Context) ([]*domain.User, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.Driver,
		"MATCH (u:User) RETURN u AS user",
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}

	// Loop through results and do something with them
	users := []*domain.User{}
	for _, record := range result.Records {
		key := "user"
		user, ok := record.Get(key)
		if !ok {
			return nil, fmt.Errorf("%s not present in return statement of query", key)
		}
		userNode, _ := user.(neo4j.Node)
		users = append(users, &domain.User{
			ID:   userNode.ElementId,
			Name: userNode.GetProperties()["name"].(string),
		})
	}
	return users, nil
}

func (r *UsersRepo) GetUserByID(ID string) (*domain.User, error) {
	return nil, nil
}

package user

import "github.com/smithshelke/flur/internal/core/domain"

func CreateUserRequestToUsers(r *CreateUserRequest) []*domain.User {
	users := []*domain.User{}
	for _, u := range r.Users {
		users = append(users, &domain.User{
			Name: u.Name,
		})
	}
	return users
}

func UsersToCreateUserResponse(d []*domain.User) *CreateUserResponse {
	users := CreateUserResponse{Users: []UserResponse{}}
	for _, u := range d {
		users.Users = append(users.Users, UserResponse{
			Name: u.Name,
		})
	}
	return &users
}

func UsersToListUserResponse(d []*domain.User) *ListUserResponse {
	users := ListUserResponse{Users: []UserResponse{}}
	for _, u := range d {
		users.Users = append(users.Users, UserResponse{
			ID:   u.ID,
			Name: u.Name,
		})
	}
	return &users
}

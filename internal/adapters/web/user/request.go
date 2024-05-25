package user

type UserRequest struct {
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Users []UserRequest `json:"users"`
}

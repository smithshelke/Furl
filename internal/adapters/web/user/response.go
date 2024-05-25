package user

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Users []UserResponse `json:"users"`
}

type ListUserResponse struct {
	Users []UserResponse `json:"users"`
}

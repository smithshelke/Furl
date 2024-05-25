package user

import (
	"encoding/json"
	"net/http"

	"github.com/smithshelke/flur/internal/adapters/web/common"
	"github.com/smithshelke/flur/internal/core/users"
)

type UserController struct {
	UserAPI *users.UserAPI
}

func NewUserController(userAPI *users.UserAPI) *UserController {
	return &UserController{
		UserAPI: userAPI,
	}
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	request := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}
	users, _ := u.UserAPI.Create(r.Context(), CreateUserRequestToUsers(request))
	response := UsersToCreateUserResponse(users)
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (u *UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserAPI.List(r.Context())
	if err != nil {
		common.WriteError(w, err, common.BAD_REQUEST)
		return
	}
	response := UsersToListUserResponse(users)
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

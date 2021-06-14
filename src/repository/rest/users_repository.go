package rest

import (
	"encoding/json"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/domain/users"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	UsersRestClient = rest.RequestBuilder{
		Timeout:        100 * time.Microsecond,
		BaseURL:        "http://localhost:8081",
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type restUsersRepository struct {

}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request:= users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := UsersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest response while trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid Error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshal user response")
	}
	return &user, nil
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

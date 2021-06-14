package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: "https://api.bookstore.com/users/login",
		HTTPMethod: http.MethodPost,
		ReqBody: `{"email" :"email@gmail.com","password":"the-password", "status": "active"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
	})
	repository := restUsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest response while trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: "https://api.bookstore.com/users/login",
		HTTPMethod: http.MethodPost,
		ReqBody: `{"email" :"email@gmail.com","password":"the-password", "status": "active"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message":"Invalid Credentials!!!", "status": "404", "error": "not_found"}`,
	})
	repository := restUsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid Error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: "https://api.bookstore.com/users/login",
		HTTPMethod: http.MethodPost,
		ReqBody: `{"email" :"email@gmail.com","password":"the-password", "status": "active"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message":"Invalid Credentials!!!", "status": 404, "error": "not_found"}`,
	})
	repository := restUsersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid Credentials!!!", err.Message)
}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Fede", "last_name": "León", "email": "fedeleon.cba@gmail.com"}`,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error when trying to unmarshal user response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Fede", "last_name": "León", "email": "fedeleon.cba@gmail.com"}`,
	})

	repository := restUsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Fede", user.FirstName)
	assert.EqualValues(t, "León", user.LastName)
	assert.EqualValues(t, "fedeleon.cba@gmail.com", user.Email)
}
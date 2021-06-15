package db

import (
	"errors"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/clients/cassandra"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/domain/access_token"
	"github.com/abhilashdk2016/bookstore_utils_go/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?)"
	queryUpdateExpires = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(token access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTIme(token access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("No access token found for given id")
		}
		return nil, rest_errors.NewInternalServerError(err.Error(), errors.New("cassandra error"))
	}

	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, token.AccessToken, token.UserId, token.ClientId, token.Expires).Exec();
	err != nil {
		return rest_errors.NewInternalServerError(err.Error(), errors.New("cassandra error"))
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTIme(token access_token.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires, token.Expires, token.AccessToken).Exec();
		err != nil {
		return rest_errors.NewInternalServerError(err.Error(), errors.New("cassandra error"))
	}
	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
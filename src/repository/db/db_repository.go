package db

import (
	"github.com/abhilashdk2016/bookstore_oauth_api/src/clients/cassandra"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/domain/access_token"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?)"
	queryUpdateExpires = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTIme(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFound("No access token found for given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, token.AccessToken, token.UserId, token.ClientId, token.Expires).Exec();
	err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTIme(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires, token.Expires, token.AccessToken).Exec();
		err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
package access_token

import (
	"github.com/abhilashdk2016/bookstore_oauth_api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTIme(token AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTIme(token AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func (s *service) GetById(accessTokenId string ) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid Access Token Id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(token AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.Create(token)
}

func (s *service) UpdateExpirationTIme(token AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTIme(token)
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

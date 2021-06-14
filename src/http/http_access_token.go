package http

import (
	atDomain "github.com/abhilashdk2016/bookstore_oauth_api/src/domain/access_token"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/services/access_token"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler{
	return &accessTokenHandler{
		service: service,
	}
}
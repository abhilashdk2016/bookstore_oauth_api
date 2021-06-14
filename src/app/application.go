package app

import (
	"github.com/abhilashdk2016/bookstore_oauth_api/src/http"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/repository/db"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/repository/rest"
	"github.com/abhilashdk2016/bookstore_oauth_api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)
	router.Run(":8080")
}

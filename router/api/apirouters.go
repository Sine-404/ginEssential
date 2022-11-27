package api

import (
	"WebProject/ginEssential/controller/api"
	"WebProject/ginEssential/middel"
	"github.com/gin-gonic/gin"
)

func GetApiRounters(g *gin.Engine) {

	ApiGroup := g.Group("/api")
	{
		ApiGroup.POST("/auth/register", api.ApiAuth{}.Register)
		ApiGroup.POST("/auth/login", api.ApiAuth{}.Login)
		ApiGroup.GET("/auth/info", middel.AuthMiddleware(), api.ApiAuth{}.UserInfo)
	}

}

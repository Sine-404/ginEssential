package router

import (
	"WebProject/ginEssential/router/api"
	"github.com/gin-gonic/gin"
)

func DisPatchRouters(g *gin.Engine) {
	api.GetApiRounters(g)
}

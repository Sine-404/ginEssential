package main

import (
	"WebProject/ginEssential/config"
	"WebProject/ginEssential/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.AutoConfig()
	defer db.Close()
	r := gin.Default()
	router.DisPatchRouters(r)
	r.Run(":8081") // listen and serve on 0.0.0.0:8200
}

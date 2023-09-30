package initRouter

import (
	"chatgin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	apiGroup := Router.Group("")
	router.SysBaseUser(apiGroup)
	return Router
}

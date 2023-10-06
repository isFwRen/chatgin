package router

import (
	"chatgin/api"
	"github.com/gin-gonic/gin"
)

func SysBaseUser(Router *gin.RouterGroup) {
	SysBasicUser := Router.Group("/user-basic")
	{
		SysBasicUser.POST("/add", api.UserBasicAdd)
	}
}

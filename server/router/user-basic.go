package router

import (
	"chatgin/api/v1/system"
	"github.com/gin-gonic/gin"
)

func SysBaseUser(Router *gin.RouterGroup) {
	sysBasicRouter := Router.Group("/user-basic")
	{
		sysBasicRouter.POST("/add", system.UserBasicAdd)         //用户信息添加
		sysBasicRouter.DELETE("/delete", system.UserBasicDelete) //用户信息删除
		sysBasicRouter.PUT("/update", system.UserBasicUpdate)    //用户信息修改
		sysBasicRouter.POST("/query", system.UserBasicQuery)     //用户信息查询
	}
}

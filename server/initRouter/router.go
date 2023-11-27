package initRouter

import (
	"chatgin/initialize"
	"chatgin/router"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(initialize.Cors())
	apiGroup := Router.Group("")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("欢迎使用swagAPI文档:http://localhost:8080/swagger/index.html")
	router.SysBaseUser(apiGroup)
	return Router
}

package main

//go:generate go mod tidy
//go:generate swag init

import (
	_ "chatgin/docs"
	"chatgin/global"
	"chatgin/initRouter"
	"chatgin/initialize"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	initialize.InitConfig()
	global.DB = initialize.InitMysql()
	router := initRouter.Routers()
	err := router.Run()
	if err != nil {

	}
}

package main

//go:generate go mod tidy
//go:generate swag init

import (
	"chatgin/initRouter"
)

func main() {
	router := initRouter.Routers()
	router.Run()
}

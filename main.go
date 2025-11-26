package main

import (
	"get-getById-with-clean-architecture/handler"
	"get-getById-with-clean-architecture/routes"
	"get-getById-with-clean-architecture/repository"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	
	userRepository := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepository)

	router.RegisterRoutes(r, userHandler)

	// run server on port 8080
	r.Run(":8080")
}
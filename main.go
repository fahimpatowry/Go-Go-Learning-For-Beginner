package main

import (
	"fmt"
	"log"

	"get-getById-with-clean-architecture/model"
	"get-getById-with-clean-architecture/handler"
	"get-getById-with-clean-architecture/repository"
	"get-getById-with-clean-architecture/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Gin router
	r := gin.Default()

	// Repository
	userRepo := repository.NewUserRepository(db)

	// Handler
	userHandler := handler.NewUserHandler(userRepo)

	routes.RegisterRoutes(r, userHandler)

	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}

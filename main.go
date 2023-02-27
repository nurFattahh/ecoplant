package main

import (
	"ecoplant/database"
	"ecoplant/handler"
	"ecoplant/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env file")
	}
	port := os.Getenv("PORT")
	// Membuat Gin Engine
	r := gin.Default()

	db := database.InitDB()

	userRepo := repository.NewUserRepository(db)
	//handler
	userHandler := handler.NewUserHandler(&userRepo)

	r.POST("/register", userHandler.CreateUser)

	r.Run(":" + port)
}

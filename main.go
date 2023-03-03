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
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalln("auto migrate error,", err)
	}

	//repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	//handler
	userHandler := handler.NewUserHandler(&userRepo)
	productHandler := handler.NewProductRepository(&productRepo)

	r.POST("/register", userHandler.CreateUser)
	r.POST("/login", userHandler.LoginUser)
	r.GET("/products", productHandler.GetAllProduct)
	r.POST("/product", productHandler.CreateProduct)
	r.GET("/product/:id", productHandler.GetProductByID)
	r.GET("/product/search/:name", productHandler.GetProductByName)

	// r.GET("listproduct", productHandler.GetListProduct)

	r.Run(":" + port)
}

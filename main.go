package main

import (
	"ecoplant/database"
	"ecoplant/handler"
	middleware "ecoplant/middleware"
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
	// cartRepo := repository.NewCartRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	//handler
	userHandler := handler.NewUserHandler(&userRepo)
	productHandler := handler.NewProductRepository(&productRepo)
	// cartHandler := handler.NewCartHandler(&cartRepo)
	transactionHandler := handler.NewTransactionHandler(&transactionRepo)

	//user
	r.POST("/user/register", userHandler.CreateUser)
	r.POST("/user/login", userHandler.LoginUser)
	r.POST("/user/:id", userHandler.GetUserById)

	//product
	r.GET("/products", productHandler.GetAllProduct)
	r.POST("/product/", productHandler.CreateProduct)
	r.GET("/product/:id", productHandler.GetProductByID)
	r.GET("/product/search/", productHandler.GetProductByName)
	r.DELETE("/product/:id", productHandler.DeleteProductById)

	//cart
	r.POST("/transaction", middleware.JwtMiddleware(), transactionHandler.CreateTransaction)

	r.Run(":" + port)
}

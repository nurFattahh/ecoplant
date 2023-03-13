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

	r := gin.Default()

	db := database.InitDB()
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalln("auto migrate error,", err)
	}

	//repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	addressRepo := repository.NewAddressRepository(db)

	//handler
	userHandler := handler.NewUserHandler(&userRepo)
	productHandler := handler.NewProductRepository(&productRepo)
	cartHandler := handler.NewCartHandler(&cartRepo)
	transactionHandler := handler.NewTransactionHandler(&transactionRepo)
	addressHandler := handler.NewAddressHandler(&addressRepo)

	//user
	r.POST("/user/register/", userHandler.CreateUser)
	r.POST("/user/login/", userHandler.LoginUser)
	r.POST("/user/", userHandler.GetUserById)
	r.GET("/user/bearer/", middleware.JwtMiddleware(), userHandler.GetUserByBearer)

	//product
	r.GET("/products", productHandler.GetAllProduct)
	r.POST("/product/", productHandler.CreateProduct)
	r.GET("/product/:id", productHandler.GetProductByID)
	r.GET("/product/search/", productHandler.GetProductByName)
	r.DELETE("/product/:id", productHandler.DeleteProductById)
	r.PATCH("/product/update/:id", productHandler.UpdateLocation)

	//cart
	r.POST("/cart/add/", middleware.JwtMiddleware(), cartHandler.CreateProductForCart)
	r.GET("/carts/", middleware.JwtMiddleware(), cartHandler.GetAllProductInCart)

	//transaction
	r.POST("/transaction/", middleware.JwtMiddleware(), transactionHandler.CreateTransaction)
	r.GET("/transaction/bearer/", middleware.JwtMiddleware(), transactionHandler.GetAllTransactionByBearer)

	//address
	r.PUT("/transaction/shipping/", middleware.JwtMiddleware(), addressHandler.ShippingAddress)

	r.Run(":" + port)
}

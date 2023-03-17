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
	communityRepo := repository.NewCommunityRepository(db)
	donationRepo := repository.NewDonationRepository(db)

	//handler
	userHandler := handler.NewUserHandler(&userRepo)
	productHandler := handler.NewProductHandler(&productRepo)
	cartHandler := handler.NewCartHandler(&cartRepo)
	transactionHandler := handler.NewTransactionHandler(&transactionRepo)
	addressHandler := handler.NewAddressHandler(&addressRepo)
	communityHandler := handler.NewCommunityHandler(&communityRepo)
	donationHandler := handler.NewDonationHandler(&donationRepo)

	//user
	r.POST("/user/register/", userHandler.CreateUser)
	r.POST("/user/login/", userHandler.LoginUser)
	r.POST("/user/", userHandler.GetUserById)
	r.GET("/user/bearer/", middleware.JwtMiddleware(), userHandler.GetUserByBearer)
	r.PATCH("/user/update/", middleware.JwtMiddleware(), userHandler.UpdateUser)
	r.PUT("/user/update/picture/", middleware.JwtMiddleware(), userHandler.UpdateProfilePicture)

	//product
	r.GET("/products/", productHandler.GetAllProduct)
	r.GET("/product/:id", productHandler.GetProductByID)
	r.GET("/product/search/", productHandler.GetProductByName)
	r.POST("/product/", productHandler.CreateProduct)
	r.DELETE("/product/:id", productHandler.DeleteProductById)
	r.PATCH("/product/update/:id", productHandler.UpdateLocation)

	//cart
	r.POST("/cart/add/", middleware.JwtMiddleware(), cartHandler.AddProductToCart)
	r.GET("/carts/", middleware.JwtMiddleware(), cartHandler.GetAllProductInCart)
	r.DELETE("/cart/delete/", middleware.JwtMiddleware(), cartHandler.DeleteItemInCartByID)

	//transaction
	r.POST("/transaction/", middleware.JwtMiddleware(), transactionHandler.CreateTransaction)
	r.GET("/transaction/bearer/", middleware.JwtMiddleware(), transactionHandler.GetAllTransactionByBearer)

	//address
	r.PUT("/transaction/shipping/", middleware.JwtMiddleware(), addressHandler.ShippingAddress)

	//Community
	r.POST("/community/", communityHandler.CreateCommunity)
	r.GET("/community/:id", communityHandler.GetCommunityByID)
	r.GET("/communities/", communityHandler.GetAllCommunity)
	r.GET("/community/name/", communityHandler.GetCommunityByName)

	//Donation
	r.POST("/donation/:id", donationHandler.CreateDonation)
	r.GET("/donations/", donationHandler.GetAllDonation)
	r.GET("/donation/:id", donationHandler.GetDonationByID)
	r.PATCH("/donation/update/:id", donationHandler.UpdatePlanAndNewsDonation)
	r.GET("/donation/regency/", donationHandler.GetDonationByRegency)
	r.POST("/user/donation/", middleware.JwtMiddleware(), donationHandler.UserDonation)
	r.GET("/user/donations/", middleware.JwtMiddleware(), donationHandler.GetAllUserDonation)

	r.Run(":" + port)

}

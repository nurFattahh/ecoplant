package main

import (
	"ecoplant/database"
	"ecoplant/handler"
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

	//handler
	postHandler := handler.NewPostHandler(db)

	// Membuat route "/helloworld"
	r.GET("/helloworld", func(c *gin.Context) {
		// Mengirimkan string "hello world" sebagai response
		c.String(200, "hello world")
	})

	r.POST("/post", postHandler.CreatePost)

	// Menjalankan Gin Engine
	r.Run(":" + port)
}

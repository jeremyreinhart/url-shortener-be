package main

import (
	"log"
	"os"

	"url-shortener/config"
	"url-shortener/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	if err := config.ConnectDatabase(); err != nil {
		log.Fatal("Database connection failed")
	}

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.RedirectURL)
	r.GET("/stats/:code", handlers.GetStats)
	r.GET("/qr/:code", handlers.GenerateQR)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
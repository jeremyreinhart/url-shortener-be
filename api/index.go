package handler

import (
	"net/http"
	"url-shortener/config"
	"url-shortener/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	config.ConnectDatabase()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.RedirectURL)
	r.GET("/stats/:code", handlers.GetStats)
	r.GET("/qr/:code", handlers.GenerateQR)

	app = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
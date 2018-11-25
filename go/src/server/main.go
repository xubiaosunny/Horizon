package main

import (
	"github.com/gin-gonic/gin"
	"server/wallpaper"
)


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/wallpaper/bing/", wallpaper.Bing)

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var port string

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/health", health)

	log.Fatal(router.Run(fmt.Sprintf("0.0.0.0:%s", port)))
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"message": "api is working",
	})
}
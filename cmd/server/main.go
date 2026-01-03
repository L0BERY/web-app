package main

import (
	"application/internal/handler"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := handler.InitRouter()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "sneakernews",
		})
	})

	logger := log.New(os.Stdout, "[SNEAKERNEWS] ", log.LstdFlags)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Printf("Server starting on :8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server failed: %v", err)
	}
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("web/template/*")

	router.GET("/", home)
	router.GET("/:path", notFound)

	return router
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{})
}

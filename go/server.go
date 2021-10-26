package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", RequestHandler)
	r.NoRoute(NotFoundHandler)

	r.Run("0.0.0.0:8080")
}

func RequestHandler(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Hellgo"})
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "404"})
}

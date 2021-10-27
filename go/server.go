package main

import (
	"crypto/sha256"

	"github.com/gin-gonic/gin"

	"context"

	"github.com/go-redis/redis/v8"

	b64 "encoding/base64"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	r := gin.Default()

	r.POST("/sha256", PutHandler)
	r.GET("/sha256", GetHandler)
	r.NoRoute(NotFoundHandler)

	r.Run(":8080")
}

func PutHandler(c *gin.Context) {
	str := c.PostForm("str")
	if len(str) < 8 {
		c.JSON(400, gin.H{"error": "400"})
		return
	}
	h := sha256.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	sha := b64.StdEncoding.EncodeToString(b)
	rdb.Set(ctx, sha, str, 0)
	c.JSON(200, gin.H{"result": sha})
}

func GetHandler(c *gin.Context) {
	sha := c.Query("sha")
	if len(sha) < 44 {
		c.JSON(400, gin.H{"error": "400"})
		return
	}
	str, err := rdb.Get(ctx, sha).Result()
	if err == redis.Nil {
		c.JSON(200, gin.H{"result": nil})
		return
	}
	c.JSON(200, gin.H{"result": str})
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "404"})
}

package main

import (
	"EzGin/gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello miniGin!</h1>")
	})
	r.GET("/hello/*", func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

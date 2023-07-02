package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello world")
	})

	router.Run()
}

package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kmr-ankitt/titly/handler"
	"github.com/kmr-ankitt/titly/store"
)

func main()  {
	router := gin.Default()
	router.Use(cors.Default())
	
	router.GET("/", func(ctx * gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Welcome to Titly!",
		})
	})
	
	router.POST("/create-short-url", func(ctx * gin.Context){
		handler.CreateShortUrl(ctx)
	})
	
	router.GET("/:short-url", func(ctx * gin.Context){
		handler.HandleShortUrlRedirect(ctx)
	})
	
	store.InitaliseStore()
	
	err := router.Run(":4000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

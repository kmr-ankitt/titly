package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.GET("/", func(ctx * gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	
	err := router.Run(":4000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

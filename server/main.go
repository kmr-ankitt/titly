package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kmr-ankitt/titly/handler"
	"github.com/kmr-ankitt/titly/store"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"

)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10,
	}
	storee := memory.NewStore()

	// Instantiate the limiter
	instance := limiter.New(storee, rate)
	middleware := ginlimiter.NewMiddleware(instance)
	router.Use(middleware)


	db := store.InitialiseStore()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to Titly!",
		})
	})

	router.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx, db)
	})

	router.GET("/:short-url", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx, db)
	})

	err := router.Run(":4000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

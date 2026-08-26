package main

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kmr-ankitt/titly/handler"
	"github.com/kmr-ankitt/titly/store"
)

func main()  {
	router := gin.Default()
	router.Use(cors.Default())

	db := store.InitialiseStore()
	
	router.GET("/", func(ctx * gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Welcome to Titly!",
		})
	})
	
	router.POST("/create-short-url", func(ctx * gin.Context){
		handler.CreateShortUrl(ctx, db)
	})
	
	router.GET("/:short-url", func(ctx * gin.Context){
		handler.HandleShortUrlRedirect(ctx, db)
	})

	//TODO: Just for testing
	router.GET("/test/sqlite", func(ctx * gin.Context){
		mappings, err := store.GetAllMappingsFromSqliteStore(db.SqliteClient)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to retrieve mappings"})
			return
		}
		ctx.JSON(200, gin.H{"mappings": mappings})
	})

	//TODO: Just for testing
	router.GET("/test/redis", func(ctx * gin.Context){
		mappings, err := store.GetAllMappingsFromRedisStore(context.Background(), db.RedisClient)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to retrieve mappings"})
			return
		}
		ctx.JSON(200, gin.H{"mappings": mappings})
	})

	err := router.Run(":4000")	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

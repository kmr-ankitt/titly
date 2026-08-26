package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmr-ankitt/titly/generator"
	"github.com/kmr-ankitt/titly/store"
)

type ShortUrlRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context, db *store.StoreService) {
	var req ShortUrlRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	longUrl := req.LongUrl

	// check if the long URL already exists in the database
	if store.ExistsInSqliteStore(db.SqliteClient, longUrl) {
		shortUrl := store.GetShortUrlFromSqliteStore(db.SqliteClient, longUrl)
		ctx.JSON(200, gin.H{"short_url": shortUrl})
		return
	}

	shortUrl := generator.GenerateShortURL(longUrl)
	id := store.StoreMappingInSqliteStore(db.SqliteClient, longUrl, shortUrl)

	ctx.JSON(200, gin.H{
		"id":        id,
		"long_url":  longUrl,
		"short_url": shortUrl,
	})
}

/*
If a user hits a short URL, we will be checking if it exists in cache,
then in database, if it exists in database, then we will redirect the user to the long URL
and also store mapping in cache for future requests.
*/
func HandleShortUrlRedirect(ctx *gin.Context, db *store.StoreService) {
	shortUrl := ctx.Param("short-url")

	var longUrl string
	// check if the short URL exists in cache
	longUrl, err := store.GetLongUrlFromRedisStore(ctx, db.RedisClient, shortUrl)
	if err == nil {
		ctx.Redirect(http.StatusFound, longUrl)
		return
	}

	// check if the short URL exists in Database
	longUrl, err = store.GetLongUrlFromSqliteStore(ctx, db.SqliteClient, shortUrl)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Short URL not found"})
		return
	}

	// store mapping in cache for future requests
	err = store.PutUrlMappingInRedisStore(ctx, db.RedisClient, shortUrl, longUrl)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to store mapping in cache"})
		return
	}

	ctx.Redirect(http.StatusFound, longUrl)
}
package handler

import (
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
	// store.StoreMappingInSqliteStore(db.SqliteClient, longUrl, shortUrl)

	ctx.JSON(200, gin.H{"short_url": shortUrl})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
}

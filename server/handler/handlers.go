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
	id := store.StoreMappingInSqliteStore(db.SqliteClient, longUrl, shortUrl)

	ctx.JSON(200, gin.H{
		"id":        id,
		"long_url":  longUrl,
		"short_url": shortUrl,
	})
}

//TODO: Just for testing purpose
func ShowAllMappings(ctx *gin.Context, db * store.StoreService) {
	mappings, err := store.GetAllMappingsFromSqliteStore(db.SqliteClient)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, mappings)
}

func HandleShortUrlRedirect(ctx *gin.Context) {
}

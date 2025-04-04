package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmr-ankitt/titly/shortner"
	"github.com/kmr-ankitt/titly/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx * gin.Context){
	var creationRequest UrlCreationRequest
	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	shortUrl := shortner.GenerateShortURL(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	
	host := "http://localhost:4000/"
	ctx.JSON(200, gin.H{
		"message": "Short URL generated successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(ctx * gin.Context){
	shortUrl := ctx.Param("short-url")
	longUrl := store.RetriveInitialUrl(shortUrl)
	ctx.Redirect(302, longUrl)
}

package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitaliseStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink)

	retrivedUrl := RetriveInitialUrl(shortURL)
	assert.Equal(t, initialLink, retrivedUrl)
}

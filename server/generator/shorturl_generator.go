package generator

import (
	"encoding/base64"
	"uuid"
)

/*
Using uuid gives us URL of length 36 characters which is still long for a short URL.
*/
func generateUniqueId() string {
	return uuid.New().String()
}

/*
We choose base64 encoding to generate a short URL
because it provides 64 unique characters that can be used to represent the hash value.
Gives 64^8 = 281,474,976,710,656 unique combinations for an 8-character short URL which
decreases the chances of collision to stronomically low and practically zero.
*/
func GenerateBase64Hash(uniqueId string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(uniqueId))
}

/*
But, Base64 encoding still gives us length of 48 character which is still too long.

Hence, we take the first 8 characters of the base64 encoded string to generate a short URL.
Tradeoff: **Truncating hash could increase chance of Collision Probability**
*/
func GenerateShortURL(longUrl string) string {
	uniqueId := generateUniqueId()
	shortUrl := GenerateBase64Hash(uniqueId)

	return shortUrl[:8]
}

package shortner

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha2560f(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// We chose base58 instead of base64 because it provides a more compact
// representation of the hash value. The characters 0, O, I, and l are highly confusing
// when used in certain fonts and are even harder to differentiate for people with visual impairments.
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoder, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoder)
}

// The userId is included to ensure that different users receive unique shortened URLs
// even if they attempt to shorten the same link.
func GenerateShortURL(initialLink string, userId string) string {
	urlHashBytes := sha2560f(initialLink + userId)
	generateNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Appendf(nil, "%d", generateNumber)))
	return finalString[:8]
}

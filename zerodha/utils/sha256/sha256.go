package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// The data to be hashed (can be a string or byte slice)
	data := "sc7lei36xyotmrq5Yw2h2G7SCn5pdL4zTtWeYIIYv2ysUdEKwtga2ouzugdauow4itdqvenicd6huarq"
	hexString := GenerateSha256(data)
	// Print the SHA256 hash
	fmt.Println("SHA256 Hash:", hexString)

}

func GenerateSha256(data string) string {
	// Create a new SHA256 hash calculator
	hasher := sha256.New()

	// Write the data to the hash calculator
	hasher.Write([]byte(data))

	// Calculate the SHA256 hash
	sha256Hash := hasher.Sum(nil)

	// Convert the hash to a hex string (common representation)
	hexString := hex.EncodeToString(sha256Hash)

	return hexString
}

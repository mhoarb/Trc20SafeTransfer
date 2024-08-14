package keygen

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// GenerateAddress generates a TRC-20 compatible address from the public key using SHA3-256
func GenerateAddress(pubKeyBytes []byte) string {
	// Perform SHA-256 hashing on the public key
	sha256Hash := sha256.New()
	sha256Hash.Write(pubKeyBytes)
	sha256HashedPubKey := sha256Hash.Sum(nil)

	// Perform SHA3-256 hashing on the SHA-256 result
	sha3Hash := sha3.New256()
	sha3Hash.Write(sha256HashedPubKey)
	sha3HashedPubKey := sha3Hash.Sum(nil)

	// Convert to hexadecimal format
	address := hex.EncodeToString(sha3HashedPubKey)

	// Prepend '41' (the Tron Network identifier) to the address
	return "41" + address
}

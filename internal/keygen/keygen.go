package keygen

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

// GenerateKeyPair generates a new ECDSA key pair and returns the private and public keys
func GenerateKeyPair() (*ecdsa.PrivateKey, []byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Marshal the public key into a byte slice
	publicKeyBytes := elliptic.Marshal(elliptic.P256(), privateKey.PublicKey.X, privateKey.PublicKey.Y)

	return privateKey, publicKeyBytes, nil
}

// PrivateKeyToHex converts a private key to a hex string
func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(privateKey.D.Bytes())
}

// PublicKeyToHex converts a public key to a hex string
func PublicKeyToHex(publicKeyBytes []byte) string {
	return hex.EncodeToString(publicKeyBytes)
}

package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// SHA256 creates a hex-encoded SHA256 checksum for a string input.
func SHA256(value string) string {
	b := sha256.Sum256([]byte(value))
	hash := []byte{}
	for _, byte := range b {
		hash = append(hash, byte)
	}
	return hex.EncodeToString(hash)
}

// SHA512 creates a hex-encoded SHA512 checksum for a string input.
func SHA512(value string) string {
	b := sha512.Sum512([]byte(value))
	hash := []byte{}
	for _, byte := range b {
		hash = append(hash, byte)
	}
	return hex.EncodeToString(hash)
}

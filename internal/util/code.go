package util

import (
	"crypto/rand"
	"encoding/base32"
)

// GenerateImportCode provides a cryptographically secure and random 8-character string.
// The returned code is uppercase, alphanumeric, and designed to be human-readable,
// making it ideal for use as a temporary, unique identifier.
func GenerateImportCode() (string, error) {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base32.StdEncoding.
		WithPadding(base32.NoPadding).
		EncodeToString(b), nil
}

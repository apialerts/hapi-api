package util

import (
	"regexp"
	"testing"
)

func TestGenerateImportCode(t *testing.T) {
	t.Run("generates 8-character code without error", func(t *testing.T) {
		code, err := GenerateImportCode()
		if err != nil {
			t.Fatalf("GenerateImportCode() returned an unexpected error: %v", err)
		}

		if len(code) != 8 {
			t.Errorf("expected code length to be 8, but got %d", len(code))
		}
	})

	t.Run("contains only valid Base32 characters", func(t *testing.T) {
		code, err := GenerateImportCode()
		if err != nil {
			t.Fatalf("GenerateImportCode() returned an unexpected error: %v", err)
		}

		// Base32 alphabet (unpadded) consists of A-Z and 2-7.
		validChars := regexp.MustCompile(`^[A-Z2-7]+$`)
		if !validChars.MatchString(code) {
			t.Errorf("generated code '%s' contains invalid characters", code)
		}
	})

	t.Run("is highly unlikely to produce duplicates", func(t *testing.T) {
		// This test generates a large number of codes to check for collisions.
		// A collision is statistically improbable with a proper random source,
		// so finding one indicates a serious problem.
		const numCodes = 100_000
		generated := make(map[string]bool, numCodes)

		for i := 0; i < numCodes; i++ {
			code, err := GenerateImportCode()
			if err != nil {
				t.Fatalf("GenerateImportCode() returned an unexpected error on iteration %d: %v", i, err)
			}

			if generated[code] {
				t.Fatalf("collision detected: code '%s' was generated more than once", code)
			}
			generated[code] = true
		}
	})
}

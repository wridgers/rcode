package rcode

import (
	"fmt"
	"math/rand"
)

// Generate a A-Za-z0-9 code of the specified length
func Generate(length int) (string, error) {
	var source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	return GenerateFromSource(source, length)
}

// GenerateFromSource returns a random code of specified length sourced from the input string
func GenerateFromSource(source string, length int) (string, error) {
	runes := []rune(source)

	if length < 0 {
		return "", fmt.Errorf("Length must >= 0")
	}

	code := make([]rune, length)

	for i := range code {
		code[i] = runes[rand.Intn(len(runes))]
	}

	return string(code), nil
}

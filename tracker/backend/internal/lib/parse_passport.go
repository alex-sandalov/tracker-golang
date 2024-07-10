package lib

import (
	"fmt"
	"strings"
	"unicode"
)

// ParsePassport parses a passport string into its two parts.
//
// The input passport string is expected to be in the format of two words separated by a space.
// If the format is not as expected, it returns an error.
//
// The function iterates through each word in the passport string and checks if all characters in the word are digits.
// If a non-digit character is found, it returns an error.
//
// Parameters:
// - ctx: The context.Context for the function.
// - passport: The passport string to parse.
// Returns:
// - The first word of the passport string.
// - The second word of the passport string.
// - An error if the format of the passport string is invalid.
func ParsePassport(passport string) (string, string, error) {
	passportSplit := strings.Split(passport, " ")
	if len(passportSplit) != 2 {
		return "", "", fmt.Errorf("invalid passport format")
	}

	for _, elem := range passportSplit {
		for _, char := range elem {
			if !unicode.IsDigit(char) {
				return "", "", fmt.Errorf("invalid passport format")
			}
		}
	}

	return passportSplit[0], passportSplit[1], nil
}

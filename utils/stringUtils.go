package utils

import "unicode"

func StripNonLetters(message string) string {
	result := ""

	for _, char := range message {
		if unicode.IsLetter(char) {
			result += string(char)
		}
	}

	return result
}

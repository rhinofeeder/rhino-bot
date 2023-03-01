package conditional

import (
	"rhino-bot/behavior"
	"strings"
	"unicode"
)

type SpongemockConditional struct {
}

func spongemockMessage(message string) (string, error) {
	result := ""
	currentRule := 50
	for _, char := range message {
		if unicode.IsLetter(char) {
			if behavior.GenerateBool(currentRule) {
				result += string(unicode.ToUpper(char))
				currentRule = 20
			} else {
				result += string(unicode.ToLower(char))
				currentRule = 80
			}
		} else {
			result += string(char)
		}
	}
	return result, nil
}

func (sc *SpongemockConditional) Handle(message string) (string, error) {
	if strings.Contains(strings.ToLower(message), "silksong") {
		return spongemockMessage(message)
	}

	if behavior.GenerateBool(5) {
		return spongemockMessage(message)
	}

	return "", nil
}

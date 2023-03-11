package conditional

import (
	"strings"
	"unicode"
)

type SpongemockConditional struct {
	RngFunc func(int) bool
}

func (sc *SpongemockConditional) spongemockMessage(message string) (string, error) {
	result := ""
	currentRule := 50
	for _, char := range message {
		if unicode.IsLetter(char) {
			if sc.RngFunc(currentRule) {
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
		return sc.spongemockMessage(message)
	}

	if sc.RngFunc(5) {
		return sc.spongemockMessage(message)
	}

	return "", nil
}

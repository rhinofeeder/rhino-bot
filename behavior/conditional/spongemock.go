package conditional

import (
	"rhino-bot/behavior"
	"strings"
	"unicode"
)

type SpongemockConditional struct {
}

func (sc *SpongemockConditional) Handle(message string) (string, error) {
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

func (sc *SpongemockConditional) RequiresMod() bool {
	return false
}

func (sc *SpongemockConditional) ShouldHandle(message string) bool {
	if strings.Contains(strings.ToLower(message), "silksong") {
		return true
	}
	return behavior.GenerateBool(5)
}

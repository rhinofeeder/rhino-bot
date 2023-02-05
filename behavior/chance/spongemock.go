package chance

import (
	"rhino-bot/behavior"
	"unicode"
)

type SpongemockCommand struct {
}

func (sc *SpongemockCommand) Handle(message string) (string, error) {
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

func (sc *SpongemockCommand) RequiresMod() bool {
	return false
}

func (sc *SpongemockCommand) ShouldHandle() bool {
	return behavior.GenerateBool(10)
}

package conditional

import (
	"fmt"
	"rhino-bot/behavior"
	"rhino-bot/utils"
	"strings"
)

type WhatsupConditional struct {
}

func findUpWord(message string) string {
	words := strings.Fields(message)

	for _, word := range words {
		strippedWord := utils.StripNonLetters(word)
		if strings.HasPrefix(word, "up") && len(strippedWord) > 3 {
			return strippedWord
		}
	}

	return ""
}

func (wc *WhatsupConditional) Handle(message string) (string, error) {
	upWord := findUpWord(message)
	if upWord != "" && behavior.GenerateBool(50) {
		result := fmt.Sprintf("What's %v?", upWord)
		return result, nil
	}

	return "", nil
}

package conditional

import (
	"fmt"
	"rhino-bot/utils"
	"strings"
)

type WhatsupConditional struct {
	RngFunc func(int) bool
}

func (wc *WhatsupConditional) findUpWord(message string) string {
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
	upWord := wc.findUpWord(message)
	if upWord != "" && wc.RngFunc(50) {
		result := fmt.Sprintf("What's %v?", upWord)
		return result, nil
	}

	return "", nil
}

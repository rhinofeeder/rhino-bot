package conditional

import (
	"errors"
	"fmt"
	"rhino-bot/behavior"
	"rhino-bot/utils"
	"strings"
)

type WhatsupConditional struct {
	upWord string
}

func findUpWord(message string) string {
	words := strings.Fields(message)

	for _, word := range words {
		if strings.HasPrefix(word, "up") && word != "up" {
			return utils.StripNonLetters(word)
		}
	}

	return ""
}

func (wc *WhatsupConditional) Handle(string) (string, error) {
	if wc.upWord == "" {
		return "", errors.New("WhatsupConditional Handle called with no word")
	}
	result := fmt.Sprintf("What's %v?", wc.upWord)
	wc.upWord = ""
	return result, nil
}

func (wc *WhatsupConditional) ShouldHandle(message string) bool {
	upWord := findUpWord(message)
	wc.upWord = upWord
	if upWord != "" {
		return behavior.GenerateBool(50)
	}

	return false
}

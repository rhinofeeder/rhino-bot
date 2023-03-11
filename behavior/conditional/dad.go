package conditional

import (
	"fmt"
	"strings"
	"unicode"
)

type DadConditional struct {
	RngFunc func(int) bool
}

func (dc *DadConditional) Handle(message string) (string, error) {
	stopIndex := len(message)
	for i, ch := range message {
		if unicode.IsPunct(ch) && ch != '\'' {
			stopIndex = i
			break
		}
	}
	words := strings.Fields(message[:stopIndex])
	if len(words) > 1 && len(words) < 20 && strings.EqualFold(words[0], "I'm") && dc.RngFunc(50) {
		return fmt.Sprintf("Hi %v, I'm rhinobot!", strings.Join(words[1:], " ")), nil
	}
	return "", nil
}

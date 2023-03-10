package conditional

import (
	"errors"
	"fmt"
	"rhino-bot/behavior"
	"strings"
)

type DadConditional struct {
}

func (dc *DadConditional) Handle(message string) (string, error) {
	words := strings.Fields(message)
	if len(words) > 20 {
		return "", errors.New("dad command limited to 20 words or fewer. tighten it up")
	}

	if len(words) > 1 && strings.EqualFold(words[0], "I'm") && behavior.GenerateBool(50) {
		return fmt.Sprintf("Hi %v, I'm rhinobot!", strings.Join(words[1:], " ")), nil
	}
	return "", nil
}

package conditional

import (
	"fmt"
	"rhino-bot/behavior"
	"strings"
)

type DadConditional struct {
}

func (dc *DadConditional) Handle(message string) (string, error) {
	words := strings.Fields(message)

	if len(words) > 1 && len(words) < 20 && strings.EqualFold(words[0], "I'm") && behavior.GenerateBool(50) {
		return fmt.Sprintf("Hi %v, I'm rhinobot!", strings.Join(words[1:], " ")), nil
	}
	return "", nil
}

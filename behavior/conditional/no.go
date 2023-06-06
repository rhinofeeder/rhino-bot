package conditional

import "strings"

type NoConditional struct{}

func (nc *NoConditional) Handle(message string) (string, error) {
	if strings.EqualFold(message, "rhinobot no") {
		return "rhinof1Sad", nil
	}
	return "", nil
}

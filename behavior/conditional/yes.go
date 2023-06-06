package conditional

import "strings"

type YesConditional struct{}

func (yc *YesConditional) Handle(message string) (string, error) {
	if strings.EqualFold(message, "rhinobot yes") {
		return "rhinof1Bongo", nil
	}
	return "", nil
}

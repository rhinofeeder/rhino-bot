package command

import (
	"errors"
	"fmt"
	"regexp"
)

type SoCommand struct{}

var UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]\w{3,24}$`)

func (sc *SoCommand) Name() string {
	return "so"
}

func (sc *SoCommand) Handle(message string) (string, error) {
	if UsernameRegex.MatchString(message) {
		return fmt.Sprintf("/me shoutouts to %v at https://twitch.tv/%v", message, message), nil
	}
	return "", errors.New("input is not a valid twitch username")
}

func (sc *SoCommand) RequiresMod() bool {
	return true
}

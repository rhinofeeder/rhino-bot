package command

import (
	"fmt"
)

type SoCommand struct{}

func (sc *SoCommand) Name() string {
	return "so"
}

func (sc *SoCommand) Handler(message string) (string, error) {
	return fmt.Sprintf("/me shoutouts to %v at https://twitch.tv/%v", message, message), nil
}

func (sc *SoCommand) RequiresMod() bool {
	return true
}

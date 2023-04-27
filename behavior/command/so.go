package command

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type SoCommand struct {
	lastCalled time.Time
}

var UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]\w{3,24}$`)

func (sc *SoCommand) Name() string {
	return "so"
}

func (sc *SoCommand) Handle(message string) (string, error) {
	sc.lastCalled = time.Now()
	if message[0] == '@' {
		message = message[1:]
	}
	if UsernameRegex.MatchString(message) {
		return fmt.Sprintf("/me shoutouts to %v at https://twitch.tv/%v", message, message), nil
	}
	return "", errors.New("input is not a valid twitch username")
}

func (sc *SoCommand) RequiresMod() bool {
	return true
}

func (sc *SoCommand) OnCooldown() bool {
	return !sc.lastCalled.IsZero() && time.Since(sc.lastCalled) < 5*time.Second
}

func (sc *SoCommand) Help() string {
	return "Mods only. Takes an input of another user's Twitch handle and gives them a text shout out in chat."
}

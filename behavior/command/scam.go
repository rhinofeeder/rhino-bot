package command

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type ScamCommand struct {
	lastCalled time.Time
}

func (sc *ScamCommand) Name() string {
	return "scam"
}

func (sc *ScamCommand) Handle(msg string) (string, error) {
	sc.lastCalled = time.Now()
	msg = strings.TrimSpace(msg)
	if msg == "" {
		return "", nil
	}

	msgWords := strings.Split(msg, " ")

	for i, word := range msgWords {
		msgRune := []rune(word)
		msgRune[0] = unicode.ToUpper(msgRune[0])
		msgWords[i] = string(msgRune)
	}

	return fmt.Sprintf("%v is a scam created by Big %v to sell more %v", msg, strings.Join(msgWords, " "), msg), nil
}

func (sc *ScamCommand) RequiresMod() bool {
	return false
}

func (sc *ScamCommand) OnCooldown() bool {
	return !sc.lastCalled.IsZero() && time.Since(sc.lastCalled) < 5*time.Second
}

func (sc *ScamCommand) Help() string {
	return "Takes an input and prints a message in the following format: \"[input] is a scam created by Big [input] to sell more [input].\""
}

package command

import (
	"fmt"
	"rhino-bot/bot"
	"time"
)

type HelpCommand struct {
	lastCalled      time.Time
	GetRhinoBotFunc func() *bot.RhinoBot
}

func (hc *HelpCommand) Handle(message string) (string, error) {
	hc.lastCalled = time.Now()

	registeredCommands := hc.GetRhinoBotFunc().Commands
	targetCommand, ok := registeredCommands[message]
	if !ok {
		return fmt.Sprintf("Command \"%v\" not found.", message), nil
	}

	return targetCommand.Help(), nil
}

func (hc *HelpCommand) Name() string { return "help" }

func (hc *HelpCommand) RequiresMod() bool { return false }

func (hc *HelpCommand) OnCooldown() bool {
	return !hc.lastCalled.IsZero() && time.Since(hc.lastCalled) < 5*time.Second
}

func (hc *HelpCommand) Help() string {
	return "Takes a command name as input and prints out a message explaining what it does."
}

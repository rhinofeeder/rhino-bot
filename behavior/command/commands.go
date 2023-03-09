package command

import (
	"rhino-bot/bot"
	"sort"
	"strings"
	"time"
)

type CommandsCommand struct {
	lastCalled      time.Time
	GetRhinoBotFunc func() *bot.RhinoBot
}

func (cc *CommandsCommand) Handle(string) (string, error) {
	cc.lastCalled = time.Now()

	registeredCommands := cc.GetRhinoBotFunc().Commands
	commandNames := make([]string, 0, len(registeredCommands))
	for commandName := range registeredCommands {
		commandNames = append(commandNames, "!"+commandName)
	}

	sort.Strings(commandNames)

	return strings.Join(commandNames, ", "), nil
}

func (cc *CommandsCommand) Name() string {
	return "commands"
}

func (cc *CommandsCommand) RequiresMod() bool {
	return false
}

func (cc *CommandsCommand) OnCooldown() bool {
	return !cc.lastCalled.IsZero() && time.Since(cc.lastCalled) < 5*time.Second
}

func (cc *CommandsCommand) Help() string {
	return "Prints a list of all commands available."
}

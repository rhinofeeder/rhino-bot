package command

import "time"

type CommandsCommand struct {
	lastCalled time.Time
}

func (cc *CommandsCommand) Handle(message string) (string, error) {
	cc.lastCalled = time.Now()

	// TODO
	panic("implement me")
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

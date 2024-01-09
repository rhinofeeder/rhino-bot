package command

import (
	"time"
)

type DiscordCommand struct {
	lastCalled time.Time
}

func (dc *DiscordCommand) Name() string { return "discord" }

func (dc *DiscordCommand) Handle(string) (string, error) {
	dc.lastCalled = time.Now()
	return "/me https://discord.com/invite/mrzNnq6", nil
}

func (dc *DiscordCommand) RequiresMod() bool { return false }

func (dc *DiscordCommand) OnCooldown() bool {
	return !dc.lastCalled.IsZero() && time.Since(dc.lastCalled) < 5*time.Second
}

func (dc *DiscordCommand) Help() string {
	return "Prints a permanent link to join the RhinoFeeder Discord server."
}

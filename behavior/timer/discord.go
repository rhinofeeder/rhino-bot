package timer

import (
	"time"
)

type Discord struct {
}

func (d *Discord) Name() string {
	return "discord"
}

func (d *Discord) Handle(message string) (string, error) {
	return "/me https://discord.com/invite/mrzNnq6", nil
}

func (d *Discord) RequiresMod() bool {
	return false
}

func (d *Discord) Trigger() string {
	return "timer"
}

func (d *Discord) Duration() time.Duration {
	return 2 * time.Second
}

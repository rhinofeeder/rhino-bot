package timer

import (
	"time"
)

type DiscordTimer struct {
}

func (dt *DiscordTimer) Handle(string) (string, error) {
	return "/me https://discord.com/invite/mrzNnq6", nil
}

func (dt *DiscordTimer) Duration() time.Duration {
	return 15 * time.Minute
}

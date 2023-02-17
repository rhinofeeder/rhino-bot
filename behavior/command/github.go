package command

import "time"

type GitHubCommand struct {
	lastCalled time.Time
}

func (gc *GitHubCommand) Name() string {
	return "github"
}

func (gc *GitHubCommand) Handle(_ string) (string, error) {
	gc.lastCalled = time.Now()
	return "/me https://github.com/rhinofeeder/rhino-bot", nil
}

func (gc *GitHubCommand) RequiresMod() bool {
	return false
}

func (gc *GitHubCommand) OnCooldown() bool {
	return !gc.lastCalled.IsZero() && time.Since(gc.lastCalled) < 5*time.Second
}

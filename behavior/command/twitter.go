package command

import "time"

type TwitterCommand struct {
	lastCalled time.Time
}

func (tc *TwitterCommand) Name() string {
	return "twitter"
}

func (tc *TwitterCommand) Handle(_ string) (string, error) {
	tc.lastCalled = time.Now()
	return "/me https://twitter.com/RhinoFeeder", nil
}

func (tc *TwitterCommand) RequiresMod() bool {
	return false
}

func (tc *TwitterCommand) OnCooldown() bool {
	return !tc.lastCalled.IsZero() && time.Since(tc.lastCalled) < 5*time.Second
}

package command

import "time"

type CoopCommand struct {
	lastCalled time.Time
}

func (cc *CoopCommand) Name() string { return "coop" }

func (cc *CoopCommand) Handle(string) (string, error) {
	cc.lastCalled = time.Now()
	return "Does your mom know you're asking strange men on the Internet to play with you?", nil
}

func (cc *CoopCommand) RequiresMod() bool { return false }

func (cc *CoopCommand) OnCooldown() bool {
	return !cc.lastCalled.IsZero() && time.Since(cc.lastCalled) < 5*time.Second
}

func (cc *CoopCommand) Help() string {
	return "Explains my co-op policy for viewers."
}

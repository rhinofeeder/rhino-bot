package command

import "time"

type NameCommand struct {
	lastCalled time.Time
}

func (nc *NameCommand) Name() string { return "name" }

func (nc *NameCommand) Handle(string) (string, error) {
	nc.lastCalled = time.Now()
	return "/me https://youtu.be/R22zSrpeSA4?t=127", nil
}

func (nc *NameCommand) RequiresMod() bool { return false }

func (nc *NameCommand) OnCooldown() bool {
	return !nc.lastCalled.IsZero() && time.Since(nc.lastCalled) < 5*time.Second
}

func (nc *NameCommand) Help() string {
	return "Prints a timestamped link to where the name 'RhinoFeeder' comes from."
}

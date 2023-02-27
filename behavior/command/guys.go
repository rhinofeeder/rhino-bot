package command

import "time"

type GuysCommand struct {
	lastCalled time.Time
}

func (gc *GuysCommand) Name() string {
	return "guys"
}

func (gc *GuysCommand) Handle(string) (string, error) {
	gc.lastCalled = time.Now()
	return "Hi! 'Guys' is a gendered pronoun. We recommend alternatives like 'folks', 'all', 'everyone', 'y'all', 'team', 'crew' etc. We appreciate your help in building an inclusive community at twitch.tv/rhinofeeder .", nil
}

func (gc *GuysCommand) RequiresMod() bool {
	return false
}

func (gc *GuysCommand) OnCooldown() bool {
	return !gc.lastCalled.IsZero() && time.Since(gc.lastCalled) < 5*time.Second
}

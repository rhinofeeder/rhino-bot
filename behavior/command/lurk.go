package command

import "time"

type LurkCommand struct {
	lastCalled time.Time
}

func (lc *LurkCommand) Handle(string) (string, error) {
	lc.lastCalled = time.Now()
	return "rhinof1Hi Enjoy your lurk!", nil
}

func (lc *LurkCommand) RequiresMod() bool {
	return false
}

func (lc *LurkCommand) Name() string {
	return "lurk"
}

func (lc *LurkCommand) OnCooldown() bool {
	return !lc.lastCalled.IsZero() && time.Since(lc.lastCalled) < 5*time.Second
}

func (lc *LurkCommand) Help() string {
	return "Tells the chat that you'll be lurking."
}

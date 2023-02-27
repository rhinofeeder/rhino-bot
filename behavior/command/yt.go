package command

import "time"

type YtCommand struct {
	lastCalled time.Time
}

func (yt *YtCommand) Name() string {
	return "yt"
}

func (yt *YtCommand) Handle(string) (string, error) {
	yt.lastCalled = time.Now()
	return "/me https://www.youtube.com/channel/UCXs2LBSCBb2gPhqka9HKdmg", nil
}

func (yt *YtCommand) RequiresMod() bool {
	return false
}

func (yt *YtCommand) OnCooldown() bool {
	return !yt.lastCalled.IsZero() && time.Since(yt.lastCalled) < 5*time.Second
}

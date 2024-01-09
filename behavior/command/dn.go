package command

import (
	"rhino-bot/singleton"
	"time"
)

type DnCommand struct {
	lastCalled time.Time
}

var setups = []string{
	"Are you going to SawCon?",
	"Do you like candies?",
	"Have you played Hades?",
	"Do you like stroganoff?",
	"Do you like pudding?",
	"Have you been to Dunkin' D's?",
	"Have you seen the SNL Putin parodies?",
	"Do you prefer e-girls or iLadies?",
	"Are you going to E10?",
	"Are you familiar with Eden?",
	"Have you gotten to Good Time Dragon?",
	"Are you a fan of Chef Boyardee's?",
	"Do you speak Sugondese?",
	"Is it true that you caught ligma?",
	"Hey can you give me back my CD's?",
	"Do you know about the Greek Warrior Bophades?",
	"Have you been to Room 40?",
}

func (dc *DnCommand) Name() string { return "dn" }

func (dc *DnCommand) Handle(string) (string, error) {
	dc.lastCalled = time.Now()
	return setups[singleton.GetRandom().Intn(len(setups))], nil
}

func (dc *DnCommand) RequiresMod() bool { return false }

func (dc *DnCommand) OnCooldown() bool {
	return !dc.lastCalled.IsZero() && time.Since(dc.lastCalled) < 10*time.Second
}

func (dc *DnCommand) Help() string {
	return "Tells the caller if they're infected with Sugma."
}

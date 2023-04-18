package singleton

import (
	"rhino-bot/bot"
	"time"
)

var rhinoBot *bot.RhinoBot

func newRhinoBot() *bot.RhinoBot {
	rhinoBot = &bot.RhinoBot{
		Channel: "rhinofeeder",
		MsgRate: time.Duration(30/100) * time.Millisecond,
		Name:    "RhinoFeederBot",
	}
	return rhinoBot
}

func GetRhinoBot() *bot.RhinoBot {
	if rhinoBot == nil {
		return newRhinoBot()
	}
	return rhinoBot
}

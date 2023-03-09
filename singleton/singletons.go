package singleton

import (
	"rhino-bot/bot"
	"time"
)

var rhinoBot *bot.RhinoBot

func NewRhinoBot() *bot.RhinoBot {
	rhinoBot = &bot.RhinoBot{
		Channel:     "rhinofeeder",
		MsgRate:     time.Duration(30/100) * time.Millisecond,
		Name:        "RhinoFeederBot",
		Port:        "6667",
		PrivatePath: "./private/oauth",
		Server:      "irc.chat.twitch.tv",
	}
	return rhinoBot
}

func GetRhinoBot() *bot.RhinoBot {
	return rhinoBot
}

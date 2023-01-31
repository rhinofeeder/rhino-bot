package main

import (
	"os"
	"rhino-bot/behavior"
	"rhino-bot/behavior/chance"
	"rhino-bot/behavior/command"
	"rhino-bot/behavior/timer"
	"rhino-bot/bot"
	"time"
)

var behaviors = []behavior.Behavior{
	&command.DiscordCommand{},
	&command.FeedCommand{
		ReadFile:  os.ReadFile,
		WriteFile: os.WriteFile,
	},
	&command.NameCommand{},
	&command.OoooCommand{},
	&command.SoCommand{},
	&command.TwitterCommand{},
	&command.YtCommand{},
	&chance.SpongemockCommand{},
	&timer.Discord{},
}

func main() {
	rhinoBot := bot.RhinoBot{
		Channel:     "rhinofeeder",
		MsgRate:     time.Duration(100/30) * time.Millisecond,
		Name:        "RhinoFeederBot",
		Port:        "6667",
		PrivatePath: "./private/oauth",
		Server:      "irc.chat.twitch.tv",
	}
	rhinoBot.RegisterBehavior(behaviors...)
	rhinoBot.Start()
}

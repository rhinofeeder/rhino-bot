package main

import (
	"os"
	"rhino-bot/bot"
	"rhino-bot/command"
	"time"
)

var commands = []command.Command{
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
	rhinoBot.RegisterCommand(commands...)
	rhinoBot.Start()
}

func unused() {
	return
}

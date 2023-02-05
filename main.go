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

var commands = []behavior.Command{
	&command.DiscordCommand{},
	&command.FeedCommand{
		ReadFile:  os.ReadFile,
		WriteFile: os.WriteFile,
	},
	&command.LurkCommand{},
	&command.NameCommand{},
	&command.OoooCommand{},
	&command.SoCommand{},
	&command.StratPlsCommand{},
	&command.TwitterCommand{},
	&command.YtCommand{},
}

var timers = []behavior.Timer{
	&timer.Discord{},
}

var chances = []behavior.Chance{
	&chance.SpongemockCommand{},
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
	rhinoBot.RegisterCommands(commands...)
	rhinoBot.RegisterTimers(timers...)
	rhinoBot.RegisterChances(chances...)
	rhinoBot.Start()
}

package main

import (
	"math/rand"
	"os"
	"rhino-bot/behavior"
	"rhino-bot/behavior/command"
	"rhino-bot/behavior/conditional"
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
	&command.GitHubCommand{},
	&command.GuysCommand{},
	&command.LurkCommand{},
	&command.NameCommand{},
	&command.OoooCommand{},
	&command.ScamCommand{},
	&command.SoCommand{},
	&command.StratPlsCommand{},
	&command.TwitterCommand{},
	&command.YtCommand{},
}

var timers = []behavior.Timer{
	&timer.Discord{},
}

var conditionals = []behavior.Conditional{
	&conditional.SpongemockConditional{},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rhinoBot := bot.RhinoBot{
		Channel:     "rhinofeeder",
		MsgRate:     time.Duration(30/100) * time.Millisecond,
		Name:        "RhinoFeederBot",
		Port:        "6667",
		PrivatePath: "./private/oauth",
		Server:      "irc.chat.twitch.tv",
	}
	rhinoBot.RegisterCommands(commands...)
	rhinoBot.RegisterTimers(timers...)
	rhinoBot.RegisterConditionals(conditionals...)
	rhinoBot.Start()
}

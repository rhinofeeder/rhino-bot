package main

import (
	"math/rand"
	"os"
	"rhino-bot/behavior"
	"rhino-bot/behavior/command"
	"rhino-bot/behavior/conditional"
	"rhino-bot/behavior/timer"
	"rhino-bot/singleton"
	"time"
)

var commands = []behavior.Command{
	&command.CommandsCommand{GetRhinoBotFunc: singleton.GetRhinoBot},
	&command.DiscordCommand{},
	&command.DnCommand{},
	&command.FeedCommand{ReadFileFunc: os.ReadFile, WriteFileFunc: os.WriteFile},
	&command.GitHubCommand{},
	&command.HelpCommand{GetRhinoBotFunc: singleton.GetRhinoBot},
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
	&timer.DiscordTimer{},
}

var conditionals = []behavior.Conditional{
	&conditional.DadConditional{RngFunc: behavior.GenerateBool},
	&conditional.SpongemockConditional{RngFunc: behavior.GenerateBool},
	&conditional.WhatsupConditional{RngFunc: behavior.GenerateBool},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rhinoBot := singleton.NewRhinoBot()

	rhinoBot.RegisterCommands(commands...)
	rhinoBot.RegisterTimers(timers...)
	rhinoBot.RegisterConditionals(conditionals...)
	rhinoBot.Start()
}

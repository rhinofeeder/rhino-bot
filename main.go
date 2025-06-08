package main

import (
	"os"
	"rhino-bot/behavior"
	"rhino-bot/behavior/command"
	"rhino-bot/behavior/conditional"
	"rhino-bot/behavior/timer"
	"rhino-bot/singleton"
	"rhino-bot/utils"
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
	&conditional.DadConditional{RngFunc: utils.RandomBool},
	&conditional.NoConditional{},
	&conditional.SpongemockConditional{RngFunc: utils.RandomBool},
	&conditional.WhatsupConditional{RngFunc: utils.RandomBool},
	&conditional.YesConditional{},
}

func main() {
	rhinoBot := singleton.GetRhinoBot()

	rhinoBot.RegisterCommands(commands...)
	rhinoBot.RegisterTimers(timers...)
	rhinoBot.RegisterConditionals(conditionals...)
	rhinoBot.Start()
}

package singleton

import (
	"math/rand"
	"rhino-bot/bot"
	"time"
)

var rhinoBot *bot.RhinoBot

func newRhinoBot() *bot.RhinoBot {
	return &bot.RhinoBot{
		Channel: "rhinofeeder",
		MsgRate: time.Duration(30/100) * time.Millisecond,
		Name:    "RhinoFeederBot",
	}
}

func GetRhinoBot() *bot.RhinoBot {
	if rhinoBot == nil {
		rhinoBot = newRhinoBot()
	}
	return rhinoBot
}

var random *rand.Rand

func newRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandom() *rand.Rand {
	if random == nil {
		random = newRandom()
	}
	return random
}

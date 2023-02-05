package command

import (
	"fmt"
	"math/rand"
	"time"
)

type StratPlsCommand struct{}

type Strat struct {
	Template    string
	Replacement func() string
}

func getLeftRight() string {
	options := []string{"left", "right"}
	return options[rand.Intn(len(options))]
}

func getOrdinalDirection() string {
	options := []string{"left", "right", "up", "down", "up-left", "up-right", "down-left", "down-right"}
	return options[rand.Intn(len(options))]
}

func getEnemy() string {
	options := []string{"tiktik", "vengefly", "durandoo", "aspid", "spikes", "sawblades", "mantis", "bluggsac", "soul totem"}
	return options[rand.Intn(len(options))]
}

func getObject() string {
	options := []string{"geo", "relic", "journal entry", "ability", "eggu"}
	return options[rand.Intn(len(options))]
}

func getFloorObject() string {
	options := []string{"acid", "water", "spikes", "void pool"}
	return options[rand.Intn(len(options))]
}

var strats = []Strat{
	{Template: "dash %v", Replacement: getLeftRight},
	{Template: "upslash the %v", Replacement: getEnemy},
	{Template: "pogo the %v", Replacement: getEnemy},
	{Template: "cdash %v", Replacement: getLeftRight},
	{Template: "walk %v", Replacement: getLeftRight},
	{Template: "wings %v", Replacement: getOrdinalDirection},
	{Template: "jump %v", Replacement: getOrdinalDirection},
	{Template: "claw up the %v wall", Replacement: getLeftRight},
	{Template: "hit the %v", Replacement: getEnemy},
	{Template: "pick up the %v", Replacement: getObject},
	{Template: "dreamnail the %v", Replacement: getEnemy},
	{Template: "jump in the %v", Replacement: getFloorObject},
	{Template: "fall in the %v", Replacement: getFloorObject},
}

const minStrats = 3
const maxStrats = 5

func (spc *StratPlsCommand) Name() string {
	return "stratpls"
}

func (spc *StratPlsCommand) Handle(message string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	// Copy the strats slice so we can mutate it and remove elements
	stratsCopy := strats
	numStrats := rand.Intn(maxStrats-minStrats) + minStrats
	result := "Ok so, "
	for i := 0; i < numStrats; i++ {
		if i > 0 {
			result += ", then "
		}
		lenStrats := len(stratsCopy)
		stratIndex := rand.Intn(lenStrats)
		strat := stratsCopy[stratIndex]
		stratsCopy = append(stratsCopy[:stratIndex], stratsCopy[stratIndex+1:]...)
		result += fmt.Sprintf(strat.Template, strat.Replacement())
	}
	return result, nil
}

func (spc *StratPlsCommand) RequiresMod() bool {
	return false
}

func (sc *StratPlsCommand) Trigger() string {
	return "command"
}

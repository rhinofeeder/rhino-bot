package command

import (
	"rhino-bot/behavior"
	"rhino-bot/bot"
	"testing"
)

func TestHelpCommand_Handle(t *testing.T) {
	tests := []struct {
		name     string
		commands map[string]behavior.Command
		message  string
		want     string
	}{
		{
			name: "Empty message",
			want: "Command \"\" not found.",
		},
		{
			name:     "Invalid command",
			commands: map[string]behavior.Command{"foo": nil},
			message:  "bar",
			want:     "Command \"bar\" not found.",
		},
		{
			name:     "Valid command",
			commands: map[string]behavior.Command{"foo": &HelpCommand{}},
			message:  "foo",
			want:     "Takes a command name as input and prints out a message explaining what it does.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := &HelpCommand{
				GetRhinoBotFunc: func() *bot.RhinoBot {
					return &bot.RhinoBot{Commands: tt.commands}
				},
			}
			got, _ := hc.Handle(tt.message)
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelpCommand_Help(t *testing.T) {
	hc := &HelpCommand{}
	expected := "Takes a command name as input and prints out a message explaining what it does."

	if result := hc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestHelpCommand_Name(t *testing.T) {
	hc := &HelpCommand{}
	expected := "help"

	if result := hc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestHelpCommand_RequiresMod(t *testing.T) {
	hc := &HelpCommand{}
	if result := hc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

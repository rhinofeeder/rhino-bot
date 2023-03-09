package command

import (
	"rhino-bot/behavior"
	"rhino-bot/bot"
	"testing"
)

func TestCommandsCommand_Handle(t *testing.T) {
	tests := []struct {
		name     string
		commands map[string]behavior.Command
		message  string
		want     string
	}{
		{
			name: "No registered commands",
		},
		{
			name:     "One registered command",
			commands: map[string]behavior.Command{"foo": nil},
			want:     "!foo",
		},
		{
			name:     "Multiple registered commands ordered",
			commands: map[string]behavior.Command{"bar": nil, "foo": nil},
			want:     "!bar, !foo",
		},
		{
			name:     "Multiple registered commands unordered",
			commands: map[string]behavior.Command{"foo": nil, "bar": nil},
			want:     "!bar, !foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommandsCommand{
				GetRhinoBotFunc: func() *bot.RhinoBot {
					return &bot.RhinoBot{Commands: tt.commands}
				},
			}
			got, _ := cc.Handle(tt.message)
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandsCommand_Name(t *testing.T) {
	cc := &CommandsCommand{}
	expected := "commands"

	if result := cc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestCommandsCommand_RequiresMod(t *testing.T) {
	cc := &CommandsCommand{}
	if result := cc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestCommandsCommand_Help(t *testing.T) {
	cc := &CommandsCommand{}
	expected := "Prints a list of all commands available."

	if result := cc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

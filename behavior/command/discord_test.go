package command

import (
	"testing"
)

func TestDiscordCommand_Handle(t *testing.T) {
	dc := &DiscordCommand{}
	expected := "/me https://discord.com/invite/mrzNnq6"

	if result, _ := dc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestDiscordCommand_Name(t *testing.T) {
	dc := &DiscordCommand{}
	expected := "discord"

	if result := dc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestDiscordCommand_RequiresMod(t *testing.T) {
	dc := &DiscordCommand{}

	if result := dc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestDiscordCommand_Help(t *testing.T) {
	dc := &DiscordCommand{}
	expected := "Prints a permanent link to join the RhinoFeeder Discord server."

	if result := dc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

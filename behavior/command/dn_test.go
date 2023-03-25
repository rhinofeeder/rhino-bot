package command

import (
	"testing"
)

func TestDnCommand_Name(t *testing.T) {
	dc := &DnCommand{}
	expected := "dn"

	if result := dc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestDnCommand_RequiresMod(t *testing.T) {
	dc := &DiscordCommand{}

	if result := dc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestDnCommand_Help(t *testing.T) {
	dc := &DnCommand{}
	expected := "Tells the caller if they're infected with Sugma."

	if result := dc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

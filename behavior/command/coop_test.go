package command

import (
	"testing"
)

func TestCoopCommand_Handle(t *testing.T) {
	cc := &CoopCommand{}
	expected := "Does your mom know you're asking strange men on the Internet to play with you?"

	if result, _ := cc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestCoopCommand_Name(t *testing.T) {
	cc := &CoopCommand{}
	expected := "coop"

	if result := cc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestCoopCommand_RequiresMod(t *testing.T) {
	cc := &CoopCommand{}

	if result := cc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestCoopCommand_Help(t *testing.T) {
	cc := &CoopCommand{}
	expected := "Explains my co-op policy for viewers."

	if result := cc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

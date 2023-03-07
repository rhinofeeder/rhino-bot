package command

import (
	"testing"
)

func TestLurkCommand_Handle(t *testing.T) {
	lc := &LurkCommand{}
	expected := "rhinof1Hi Enjoy your lurk!"

	if result, _ := lc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestLurkCommand_Name(t *testing.T) {
	lc := &LurkCommand{}
	expected := "lurk"

	if result := lc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestLurkCommand_RequiresMod(t *testing.T) {
	lc := &LurkCommand{}

	if result := lc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestLurkCommand_Help(t *testing.T) {
	lc := &LurkCommand{}
	expected := "Tells the chat that you'll be lurking."

	if result := lc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

package command

import "testing"

func TestStratPlsCommand_Handle(t *testing.T) {
	sp := &StratPlsCommand{}

	if _, err := sp.Handle(""); err != nil {
		t.Errorf("Handle() error = %v, want nil", err)
	}
}

func TestStratPlsCommand_Name(t *testing.T) {
	sc := &StratPlsCommand{}
	expected := "stratpls"

	if result := sc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestStratPlsCommand_RequiresMod(t *testing.T) {
	sc := &StratPlsCommand{}

	if result := sc.RequiresMod(); result != false {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

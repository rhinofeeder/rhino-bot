package command

import "testing"

func TestNameCommand_Handle(t *testing.T) {
	nc := &NameCommand{}
	expected := "/me https://youtu.be/R22zSrpeSA4?t=127"

	if result, _ := nc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestNameCommand_Name(t *testing.T) {
	nc := &NameCommand{}
	expected := "name"

	if result := nc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestNameCommand_RequiresMod(t *testing.T) {
	nc := &NameCommand{}

	if result := nc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

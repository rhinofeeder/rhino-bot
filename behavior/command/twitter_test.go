package command

import "testing"

func TestTwitterCommand_Handle(t *testing.T) {
	tc := &TwitterCommand{}
	expected := "/me https://twitter.com/RhinoFeeder"

	if result, _ := tc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestTwitterCommand_Name(t *testing.T) {
	tc := &TwitterCommand{}
	expected := "twitter"

	if result := tc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestTwitterCommand_RequiresMod(t *testing.T) {
	tc := &TwitterCommand{}

	if result := tc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

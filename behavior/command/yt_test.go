package command

import "testing"

func TestYtCommand_Handle(t *testing.T) {
	yt := &YtCommand{}
	expected := "/me https://www.youtube.com/channel/UCXs2LBSCBb2gPhqka9HKdmg"

	if result, _ := yt.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestYtCommand_Name(t *testing.T) {
	yt := &YtCommand{}
	expected := "yt"

	if result := yt.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestYtCommand_RequiresMod(t *testing.T) {
	yt := &YtCommand{}

	if result := yt.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

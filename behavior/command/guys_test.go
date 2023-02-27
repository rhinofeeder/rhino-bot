package command

import "testing"

func TestGuysCommand_Handle(t *testing.T) {
	gc := &GuysCommand{}
	expected := "Hi! 'Guys' is a gendered pronoun. We recommend alternatives like 'folks', 'all', 'everyone', 'y'all', 'team', 'crew' etc. We appreciate your help in building an inclusive community at twitch.tv/rhinofeeder ."

	if result, _ := gc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestGuysCommand_Name(t *testing.T) {
	gc := &GuysCommand{}
	expected := "guys"

	if result := gc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestGuysCommand_RequiresMod(t *testing.T) {
	gc := &GuysCommand{}

	if result := gc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

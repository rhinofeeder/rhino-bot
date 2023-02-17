package command

import (
	"testing"
)

func TestGitHubCommand_Handle(t *testing.T) {
	gc := &GitHubCommand{}
	expected := "/me https://github.com/rhinofeeder/rhino-bot"

	if result, _ := gc.Handle(""); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

func TestGitHubCommand_Name(t *testing.T) {
	gc := &GitHubCommand{}
	expected := "github"

	if result := gc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestGitHubCommand_RequiresMod(t *testing.T) {
	gc := &GitHubCommand{}
	if result := gc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

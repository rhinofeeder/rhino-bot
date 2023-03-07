package command

import (
	"testing"
	"time"
)

func TestCommandsCommand_Handle(t *testing.T) {
	type fields struct {
		lastCalled time.Time
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommandsCommand{
				lastCalled: tt.fields.lastCalled,
			}
			got, err := cc.Handle(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandsCommand_Name(t *testing.T) {
	cc := &CommandsCommand{}
	expected := "commands"

	if result := cc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestCommandsCommand_RequiresMod(t *testing.T) {
	cc := &CommandsCommand{}
	if result := cc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestCommandsCommand_Help(t *testing.T) {
	cc := &CommandsCommand{}
	expected := "Prints a list of all commands available."

	if result := cc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

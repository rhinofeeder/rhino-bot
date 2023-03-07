package command

import (
	"testing"
)

func TestScamCommand_Handle(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "Empty message",
			msg:  "",
			want: "",
		},
		{
			name: "Blank message",
			msg:  " ",
			want: "",
		},
		{
			name: "Invalid message single word",
			msg:  "!foo",
			want: "!foo is a scam created by Big !foo to sell more !foo",
		},
		{
			name: "Invalid message multiple words",
			msg:  "!foo bar",
			want: "!foo bar is a scam created by Big !foo Bar to sell more !foo bar",
		},
		{
			name: "Valid message",
			msg:  "cars",
			want: "cars is a scam created by Big Cars to sell more cars",
		},
		{
			name: "Valid message multiple words",
			msg:  "cars dogs bikes",
			want: "cars dogs bikes is a scam created by Big Cars Dogs Bikes to sell more cars dogs bikes",
		},
		{
			name: "Invalid message non-first word in message",
			msg:  "foo !bar",
			want: "foo !bar is a scam created by Big Foo !bar to sell more foo !bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &ScamCommand{}
			got, err := sc.Handle(tt.msg)
			if (err != nil) != false {
				t.Errorf("Handle() error = %v, wantErr %v", err, false)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScamCommand_Name(t *testing.T) {
	sc := &ScamCommand{}
	expected := "scam"

	if result := sc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestScamCommand_RequiresMod(t *testing.T) {
	sc := &ScamCommand{}
	if result := sc.RequiresMod(); result != false {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestScamCommand_Help(t *testing.T) {
	sc := &ScamCommand{}
	expected := "Takes an input and prints a message in the following format: \"[input] is a scam created by Big [input] to sell more [input].\""

	if result := sc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

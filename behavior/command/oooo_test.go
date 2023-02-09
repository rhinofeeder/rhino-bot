package command

import "testing"

func TestOoooCommand_Handle(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "contains o",
			message: "rhino",
			want:    "rhin OOOO",
		},
		{
			name:    "does not contain o",
			message: "thistle",
			want:    "th OOOO stle",
		},
		{
			name:    "multiple non-o vowel candidates favors the earlier one",
			message: "jaden",
			want:    "j OOOO den",
		},
		{
			name:    "multiple non-o vowels adjacent to each other",
			message: "cheese",
			want:    "ch OOOO se",
		},
		{
			name:    "non-o vowel at the end of string",
			message: "hi",
			want:    "h OOOO",
		},
		{
			name:    "short message",
			message: "i",
			want:    "OOOO",
		},
		{
			name:    "does not contain vowels",
			message: "xkcd",
			want:    "Can't find any vowels to replace MMMM",
		},
		{
			name:    "empty message",
			message: "",
			want:    "OOOO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oc := &OoooCommand{}
			got, err := oc.Handle(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err != nil, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOoooCommand_Name(t *testing.T) {
	oc := &OoooCommand{}
	expected := "oooo"

	if result := oc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestOoooCommand_RequiresMod(t *testing.T) {
	oc := &OoooCommand{}

	if result := oc.RequiresMod(); result != false {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

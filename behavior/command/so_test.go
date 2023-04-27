package command

import "testing"

func TestSoCommand_Handle(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "all letters",
			message: "rhinofeeder",
			want:    "/me shoutouts to rhinofeeder at https://twitch.tv/rhinofeeder",
		},
		{
			name:    "has numbers",
			message: "rhin0f33d3r",
			want:    "/me shoutouts to rhin0f33d3r at https://twitch.tv/rhin0f33d3r",
		},
		{
			name:    "valid special characters",
			message: "rhino_feeder",
			want:    "/me shoutouts to rhino_feeder at https://twitch.tv/rhino_feeder",
		},
		{
			name:    "minimum valid length",
			message: "abcd",
			want:    "/me shoutouts to abcd at https://twitch.tv/abcd",
		},
		{
			name:    "maximum valid length",
			message: "aaaaaaaaaaaaaaaaaaaaaaaaa",
			want:    "/me shoutouts to aaaaaaaaaaaaaaaaaaaaaaaaa at https://twitch.tv/aaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		{
			name:    "contains invalid character",
			message: "rhino-feeder",
			wantErr: true,
		},
		{
			name:    "leading underscore",
			message: "_rhinofeeder",
			wantErr: true,
		},
		{
			name:    "under minimum length",
			message: "abc",
			wantErr: true,
		},
		{
			name:    "exceeds maximum length",
			message: "aaaaaaaaaaaaaaaaaaaaaaaaaa",
			wantErr: true,
		},
		{
			name:    "contains an @",
			message: "@rhinofeeder",
			want:    "/me shoutouts to rhinofeeder at https://twitch.tv/rhinofeeder",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SoCommand{}
			got, err := sc.Handle(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, true)
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoCommand_Name(t *testing.T) {
	sc := &SoCommand{}
	expected := "so"

	if result := sc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestSoCommand_RequiresMod(t *testing.T) {
	sc := &SoCommand{}

	if result := sc.RequiresMod(); !result {
		t.Errorf("RequiresMod() = %v, want %v", result, true)
	}
}

func TestSoCommand_Help(t *testing.T) {
	sc := &SoCommand{}
	expected := "Mods only. Takes an input of another user's Twitch handle and gives them a text shout out in chat."

	if result := sc.Help(); result != expected {
		t.Errorf("Handle() = %v, want %v", result, expected)
	}
}

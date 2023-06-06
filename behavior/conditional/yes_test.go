package conditional

import "testing"

func TestYesConditional_Handle(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "Message is not `rhinobot yes`",
			message: "foo bar",
			want:    "",
		},
		{
			name:    "Matching text in all lowercase",
			message: "rhinobot yes",
			want:    "rhinof1Bongo",
		},
		{
			name:    "Matching text in all uppercase",
			message: "RHINOBOT YES",
			want:    "rhinof1Bongo",
		},
		{
			name:    "Matching text mixed case",
			message: "RhiNoBot YeS",
			want:    "rhinof1Bongo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yc := &YesConditional{}
			got, err := yc.Handle(tt.message)
			if err != nil {
				t.Errorf("Handle() error = %v, wantErr %v", err, false)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

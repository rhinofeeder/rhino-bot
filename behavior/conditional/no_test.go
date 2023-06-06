package conditional

import "testing"

func TestNoConditional_Handle(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "Message is not `rhinobot no`",
			message: "foo bar",
			want:    "",
		},
		{
			name:    "Matching text in all lowercase",
			message: "rhinobot no",
			want:    "rhinof1Sad",
		},
		{
			name:    "Matching text in all uppercase",
			message: "RHINOBOT NO",
			want:    "rhinof1Sad",
		},
		{
			name:    "Matching text mixed case",
			message: "RhiNoBot nO",
			want:    "rhinof1Sad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NoConditional{}
			got, err := nc.Handle(tt.message)
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

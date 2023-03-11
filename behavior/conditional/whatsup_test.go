package conditional

import "testing"

func TestWhatsupConditional_Handle(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "empty string",
			message: "",
			want:    "",
		},
		{
			name:    `no word that starts with "up"`,
			message: "This is a sentence",
			want:    "",
		},
		{
			name:    `sentence has the word "up" by itself`,
			message: "What's up",
			want:    "",
		},
		{
			name:    `sentence has the word "up" with punctuation`,
			message: "What's up?",
			want:    "",
		},
		{
			name:    `sentence has a word that starts with "up" and contains punctuation`,
			message: "Is that an upslash?",
			want:    "What's upslash?",
		},
		{
			name:    `sentence has word that starts with "up"`,
			message: "You should upslash there",
			want:    "What's upslash?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &WhatsupConditional{
				RngFunc: func(int) bool { return true },
			}
			got, err := wc.Handle(tt.message)
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

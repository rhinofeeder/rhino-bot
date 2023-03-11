package conditional

import "testing"

func TestDadConditional_Handle(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "empty message",
			message: "",
			want:    "",
		},
		{
			name:    `one word message not starting with "I'm"`,
			message: "foo",
			want:    "",
		},
		{
			name:    `message is "I'm" with nothing else`,
			message: "I'm",
			want:    "",
		},
		{
			name:    "message is longer than 20 words",
			message: "I'm a b c d e f g h i j k l m n o p q r s t u v w x y z",
			want:    "",
		},
		{
			name:    "message is longer than 20 words with puncutation after 20th word",
			message: "I'm a b c d e f g h i j k l m n o p q r s t u v w x y, z",
			want:    "",
		},
		{
			name:    "message is longer than 20 words with punctuation before 20th word",
			message: "I'm a b c d, e f g h i j k l m n o p q r s t u v w x y z",
			want:    "Hi a b c d, I'm rhinobot!",
		},
		{
			name:    "valid message with no punctuation",
			message: "I'm really tired",
			want:    "Hi really tired, I'm rhinobot!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &DadConditional{
				RngFunc: func(int) bool { return true },
			}
			got, err := dc.Handle(tt.message)
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

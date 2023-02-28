package conditional

import "testing"

func TestWhatsupConditional_ShouldHandle(t *testing.T) {
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
			name: "Message with a keyword",
			msg:  "this sentence randomly has the word upslash",
			want: "upslash",
		},
		{
			name: "Message without keyword",
			msg:  "this sentence has no usable words",
			want: "",
		},
		{
			name: "Message with symbols",
			msg:  "what's updog?!",
			want: "updog",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &WhatsupConditional{}
			wc.ShouldHandle(tt.msg)
			if wc.upWord != tt.want {
				t.Errorf("Handle() got = %v, want %v", wc.upWord, tt.want)
			}
		})
	}
}

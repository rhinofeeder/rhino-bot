package conditional

import "testing"

func TestSpongemockConditional_Handle(t *testing.T) {
	sc := &SpongemockConditional{}

	if result, _ := sc.Handle("Will you play Silksong?"); result == "" {
		t.Errorf("Handle(message) = %v, want non-empty result", result)
	}
}

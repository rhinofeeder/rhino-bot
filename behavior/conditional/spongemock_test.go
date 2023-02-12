package conditional

import "testing"

func TestSpongemockConditional_RequiresMod(t *testing.T) {
	sc := &SpongemockConditional{}

	if result := sc.RequiresMod(); result != false {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestSpongemockConditional_ShouldHandle(t *testing.T) {
	sc := &SpongemockConditional{}

	if result := sc.ShouldHandle("Will you play Silksong?"); result != true {
		t.Errorf("ShouldHandle(message) = %v, want %v", result, true)
	}
}

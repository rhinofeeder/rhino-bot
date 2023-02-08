package chance

import "testing"

func TestSpongemockChance_RequiresMod(t *testing.T) {
	sc := &SpongemockCommand{}

	if result := sc.RequiresMod(); result != false {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

func TestSpongemockChance_ShouldHandle(t *testing.T) {
	sc := &SpongemockCommand{}

	if result := sc.ShouldHandle("Will you play Silksong?"); result != true {
		t.Errorf("ShouldHandle(message) = %v, want %v", result, true)
	}
}

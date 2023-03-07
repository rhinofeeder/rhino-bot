package behavior

type Command interface {
	Behavior
	Name() string
	RequiresMod() bool
	OnCooldown() bool
	Help() string
}

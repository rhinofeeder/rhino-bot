package behavior

type Behavior interface {
	Handle(message string) (string, error)
	RequiresMod() bool
}

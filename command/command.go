package command

type Command interface {
	Name() string
	Handle(message string) (string, error)
	RequiresMod() bool
}

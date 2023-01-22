package command

type Command interface {
	Name() string
	Handler(message string) (string, error)
	RequiresMod() bool
}

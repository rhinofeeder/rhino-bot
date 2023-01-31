package behavior

type Command interface {
	Behavior
	Name() string
}

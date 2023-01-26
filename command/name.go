package command

type NameCommand struct{}

func (nc *NameCommand) Name() string {
	return "name"
}

func (nc *NameCommand) Handle(_ string) (string, error) {
	return "/me https://youtu.be/R22zSrpeSA4?t=127", nil
}

func (nc *NameCommand) RequiresMod() bool {
	return false
}

package command

type TwitterCommand struct{}

func (tc *TwitterCommand) Name() string {
	return "twitter"
}

func (tc *TwitterCommand) Handle(_ string) (string, error) {
	return "/me https://twitter.com/RhinoFeeder", nil
}

func (tc *TwitterCommand) RequiresMod() bool {
	return false
}

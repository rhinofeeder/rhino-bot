package command

type DiscordCommand struct{}

func (dc *DiscordCommand) Name() string {
	return "discord"
}

func (dc *DiscordCommand) Handler(_ string) (string, error) {
	return "/me https://discord.com/invite/mrzNnq6", nil
}

func (dc *DiscordCommand) RequiresMod() bool {
	return false
}

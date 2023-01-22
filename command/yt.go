package command

type YtCommand struct{}

func (yt *YtCommand) Name() string {
	return "yt"
}

func (yt *YtCommand) Handler(_ string) (string, error) {
	return "/me https://www.youtube.com/channel/UCXs2LBSCBb2gPhqka9HKdmg", nil
}

func (yt *YtCommand) RequiresMod() bool {
	return false
}

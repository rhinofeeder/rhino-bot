package command

type OoooCommand struct{}

func (oc *OoooCommand) Name() string {
	return "oooo"
}

func (oc *OoooCommand) Handler(message string) (string, error) {
	oooo := ""
	for _, ch := range message {
		if ch == 'o' || ch == 'O' {
			oooo += " OOOO "
		} else {
			oooo += string(ch)
		}
	}
	return oooo, nil
}

func (oc *OoooCommand) RequiresMod() bool {
	return false
}

package command

import (
	"fmt"
	"os"
	"strconv"
)

const (
	filePath = "persistence/rhinosfed"
)

type FeedCommand struct {
	ReadFile  func(name string) ([]byte, error)
	WriteFile func(name string, data []byte, perm os.FileMode) error
}

func (fc *FeedCommand) Name() string {
	return "feed"
}

func (fc *FeedCommand) Handle(_ string) (string, error) {
	file, err := fc.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	rhinosFed, _ := strconv.Atoi(string(file))
	rhinosFed++

	if err = fc.WriteFile(filePath, []byte(strconv.Itoa(rhinosFed)), 0644); err != nil {
		return "", err
	}
	return fmt.Sprintf("Rhinos fed: %v", rhinosFed), nil
}

func (fc *FeedCommand) RequiresMod() bool {
	return false
}

func (fc *FeedCommand) Trigger() string {
	return "command"
}
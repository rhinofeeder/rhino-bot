package command

import (
	"fmt"
	"os"
	"strconv"
)

const (
	filePath = "persistence/rhinosfed"
)

type FeedCommand struct{}

func (fc FeedCommand) Name() string {
	return "feed"
}

func (fc FeedCommand) Handler(_ string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	rhinosFed, _ := strconv.Atoi(string(file))
	rhinosFed++

	if err = os.WriteFile(filePath, []byte(strconv.Itoa(rhinosFed)), 0644); err != nil {
		return "", err
	}
	return fmt.Sprintf("Rhinos fed: %v", rhinosFed), nil
}

func (fc FeedCommand) RequiresMod() bool {
	return false
}

package command

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const filePath = "persistence/rhinosfed"

type FeedCommand struct {
	ReadFileFunc  func(name string) ([]byte, error)
	WriteFileFunc func(name string, data []byte, perm os.FileMode) error
	lastCalled    time.Time
}

func (fc *FeedCommand) Name() string { return "feed" }

func (fc *FeedCommand) Handle(string) (string, error) {
	fc.lastCalled = time.Now()

	file, err := fc.ReadFileFunc(filePath)
	if err != nil {
		return "", err
	}

	rhinosFed, _ := strconv.Atoi(string(file))
	rhinosFed++

	if err = fc.WriteFileFunc(filePath, []byte(strconv.Itoa(rhinosFed)), 0644); err != nil {
		return "", err
	}
	return fmt.Sprintf("Rhinos fed: %v", rhinosFed), nil
}

func (fc *FeedCommand) RequiresMod() bool { return false }

func (fc *FeedCommand) OnCooldown() bool {
	return !fc.lastCalled.IsZero() && time.Since(fc.lastCalled) < 5*time.Second
}

func (fc *FeedCommand) Help() string {
	return "Increments the number of rhinos fed by the community."
}

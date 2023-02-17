package command

import (
	"errors"
	"os"
	"testing"
)

func TestFeedCommand_Handle(t *testing.T) {
	tests := []struct {
		name      string
		readFile  func(name string) ([]byte, error)
		writeFile func(name string, data []byte, perm os.FileMode) error
		want      string
		wantErr   bool
	}{
		{
			name: "happy path",
			readFile: func(string) ([]byte, error) {
				return []byte("0"), nil
			},
			writeFile: func(string, []byte, os.FileMode) error {
				return nil
			},
			want: "Rhinos fed: 1",
		},
		{
			name: "error reading file",
			readFile: func(string) ([]byte, error) {
				return nil, errors.New("")
			},
			wantErr: true,
		},
		{
			name: "error writing to file",
			readFile: func(string) ([]byte, error) {
				return []byte("0"), nil
			},
			writeFile: func(string, []byte, os.FileMode) error {
				return errors.New("")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := FeedCommand{
				ReadFile:  tt.readFile,
				WriteFile: tt.writeFile,
			}
			got, err := fc.Handle("")
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFeedCommand_Name(t *testing.T) {
	fc := &FeedCommand{}
	expected := "feed"

	if result := fc.Name(); result != expected {
		t.Errorf("Name() = %v, want %v", result, expected)
	}
}

func TestFeedCommand_RequiresMod(t *testing.T) {
	fc := &FeedCommand{}

	if result := fc.RequiresMod(); result {
		t.Errorf("RequiresMod() = %v, want %v", result, false)
	}
}

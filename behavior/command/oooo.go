package command

import (
	"strings"
	"time"
)

func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

type OoooCommand struct {
	lastCalled time.Time
}

func (oc *OoooCommand) Name() string {
	return "oooo"
}

// Handle OOOO-ifying a provided string
// - replace any "o"s if they exist
// - if no "o"s exist, try to replace the vowel(s) closest to the center of the word
func (oc *OoooCommand) Handle(message string) (string, error) {
	oc.lastCalled = time.Now()
	if message == "" {
		return "OOOO", nil
	}
	oooo := ""
	if strings.Contains(strings.ToLower(message), "o") {
		for _, ch := range message {
			if ch == 'o' || ch == 'O' {
				oooo += " OOOO "
			} else {
				oooo += string(ch)
			}
		}
	} else {
		messageLength := len(message)
		midpoint := messageLength / 2
		vowelIndex := -1
		if isVowel(message[midpoint]) {
			vowelIndex = midpoint
		} else {
			// Search outwards to find the vowel closest to the center
			for i := 1; i < messageLength-midpoint; i++ {
				if beforeIndex := midpoint - i; beforeIndex > 0 && isVowel(message[beforeIndex]) {
					vowelIndex = beforeIndex
					break
				} else if afterIndex := midpoint + i; isVowel(message[afterIndex]) {
					vowelIndex = afterIndex
					break
				}
			}
		}

		if vowelIndex == -1 {
			return "Can't find any vowels to replace MMMM", nil
		}

		// Check if adjacent letters are vowels
		lowerIndex := vowelIndex
		upperIndex := vowelIndex
		if vowelIndex > 0 && isVowel(message[vowelIndex-1]) {
			lowerIndex = vowelIndex - 1
		} else if vowelIndex+1 < messageLength && isVowel(message[vowelIndex+1]) {
			upperIndex = vowelIndex + 1
		}

		oooo = message[:lowerIndex] + " OOOO "
		if upperIndex+1 < messageLength {
			oooo += message[upperIndex+1:]
		}
	}

	return strings.TrimSpace(oooo), nil
}

func (oc *OoooCommand) RequiresMod() bool {
	return false
}

func (oc *OoooCommand) OnCooldown() bool {
	return !oc.lastCalled.IsZero() && time.Since(oc.lastCalled) < 5*time.Second
}

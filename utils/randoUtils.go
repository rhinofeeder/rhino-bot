package utils

import "rhino-bot/singleton"

// RandomBool generates a random boolean with a `chance` percent likelihood to return true.
func RandomBool(chance int) bool {
	return singleton.GetRandom().Intn(100) < chance
}

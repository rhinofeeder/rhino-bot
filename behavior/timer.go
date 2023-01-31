package behavior

import "time"

type Timer interface {
	Behavior
	Duration() time.Duration
}

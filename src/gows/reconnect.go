package gows

import "time"

const sessionReconnectQuiesce = 2 * time.Second

func reconnectQuiesceRemaining(stoppedAt, now time.Time) time.Duration {
	if stoppedAt.IsZero() {
		return 0
	}
	remaining := sessionReconnectQuiesce - now.Sub(stoppedAt)
	if remaining < 0 {
		return 0
	}
	return remaining
}

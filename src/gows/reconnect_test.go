package gows

import (
	"testing"
	"time"
)

func TestReconnectQuiesceRemaining(t *testing.T) {
	now := time.Now()
	remaining := reconnectQuiesceRemaining(now.Add(-500*time.Millisecond), now)
	if remaining < 1490*time.Millisecond || remaining > 1510*time.Millisecond {
		t.Fatalf("expected about 1.5s remaining, got %v", remaining)
	}
	if reconnectQuiesceRemaining(now.Add(-3*time.Second), now) != 0 {
		t.Fatal("expired reconnect window must not delay a later start")
	}
	if reconnectQuiesceRemaining(time.Time{}, now) != 0 {
		t.Fatal("a first start must not be delayed")
	}
}

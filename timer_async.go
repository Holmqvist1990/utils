package utils

import (
	"context"
	"time"
)

// Runs f after seconds.
func After(seconds float64, f func()) {
	go func() {
		time.Sleep(time.Millisecond * time.Duration(seconds*1000))
		f()
	}()
}

// Runs f every seconds.
func Every(seconds float64, f func()) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond * time.Duration(seconds*1000))
				f()
			}
		}
	}()
	return cancel
}

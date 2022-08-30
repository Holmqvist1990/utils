package utils

import (
	"errors"
	"fmt"
)

// Generic timer. Tick, TickDelta.
type Timer[T Number] struct {
	Duration    T
	CurrentTime T
}

// Duration must be positive.
func NewTimer[T Number](duration T) (*Timer[T], error) {
	if !(duration > 0) {
		return nil, errors.New(
			"failed to create timer: duration must be greater than zero")
	}

	return &Timer[T]{
		Duration:    duration,
		CurrentTime: 0,
	}, nil
}

func (t *Timer[T]) SetCurrentTime(time T) {
	t.CurrentTime = time
}

// Increments by one.
func (t *Timer[T]) Tick() {
	t.CurrentTime++
}

// Increments by delta, which must be positive.
func (t *Timer[T]) TickDelta(delta T) error {
	if !(delta > 0) {
		return fmt.Errorf("must be greater than zero: %v", delta)
	}

	t.CurrentTime += delta
	return nil
}

// Resets timer on true.
func (t *Timer[T]) Finished() bool {
	finished := t.CurrentTime >= t.Duration
	if finished {
		t.CurrentTime = 0
	}

	return finished
}

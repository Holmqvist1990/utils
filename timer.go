package utils

import (
	"errors"
	"fmt"
)

type Timer[T Number] struct {
	Duration    T
	CurrentTime T
}

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

func (t *Timer[T]) Tick() {
	t.CurrentTime++
}

func (t *Timer[T]) TickDelta(delta T) error {
	if !(delta > 0) {
		return fmt.Errorf("must be greater than zero: %v", delta)
	}
	t.CurrentTime += delta
	return nil
}

func (t *Timer[T]) Finished() bool {
	finished := t.CurrentTime >= t.Duration
	if finished {
		t.CurrentTime = 0
	}
	return finished
}

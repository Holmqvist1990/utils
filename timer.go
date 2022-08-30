package utils

import "errors"

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

func (t *Timer[T]) Tick() {
	t.CurrentTime++
}

func (t *Timer[T]) TickDelta(delta T) {
	t.CurrentTime += delta
}

func (t *Timer[T]) Finished() bool {
	finished := t.CurrentTime >= t.Duration
	if finished {
		t.CurrentTime = 0
	}
	return finished
}

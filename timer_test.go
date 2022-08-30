package utils

import (
	"fmt"
	"testing"
)

func TestTimerTick(t *testing.T) {
	tests := []struct {
		Name       string
		Duration   int
		FinishedIn int
		WantError  bool
	}{
		{
			Name:      "minus one",
			Duration:  -1,
			WantError: true,
		},
		{
			Name:      "zero",
			Duration:  0,
			WantError: true,
		},
		{
			Name:       "two seconds",
			Duration:   2,
			FinishedIn: 2,
		},
	}
	for _, test := range tests {
		timer, err := NewTimer(test.Duration)
		ok, err := checkError(test.Name, test.WantError, err)
		if !ok {
			if test.WantError {
				continue
			}
			t.Fatal(err)
		}

		var finishedIn int
		for !timer.Finished() {
			timer.Tick()
			finishedIn++
		}

		if finishedIn != test.FinishedIn {
			t.Fatalf("%v: wanted: %v, got: %v",
				test.Name, test.FinishedIn, finishedIn)
		}
	}
}

func TestTimerTickDelta(t *testing.T) {
	tests := []struct {
		Name       string
		Duration   float32
		Delta      float32
		WantError  bool
		FinishedIn int
	}{
		{
			Name:      "negative delta",
			Duration:  1.0,
			Delta:     -1.0,
			WantError: true,
		},
		{
			Name:      "zero delta",
			Duration:  1.0,
			WantError: true,
		},
		{
			Name:       "two point five seconds",
			Duration:   2.5,
			Delta:      0.5,
			FinishedIn: 5,
		},
	}
	for _, test := range tests {
		timer, _ := NewTimer(test.Duration)
		var finishedIn int
		for !timer.Finished() {
			err := timer.TickDelta(test.Delta)
			ok, err := checkError(test.Name, test.WantError, err)
			if !ok {
				if test.WantError {
					return
				}
				t.Fatal(err)
			}

			finishedIn++
		}

		if finishedIn != test.FinishedIn {
			t.Fatalf("%v: wanted: %v, got: %v",
				test.Name, test.FinishedIn, finishedIn)
		}
	}
}

func checkError(name string, want bool, err error) (bool, error) {
	if err != nil && !want {
		return false, fmt.Errorf("%v: did not expect error: %v",
			name, err)
	}
	if err == nil && want {
		return false, fmt.Errorf("%v: expected error: %v", name, err)
	}
	if err != nil && want {
		return false, nil
	}
	return true, nil
}

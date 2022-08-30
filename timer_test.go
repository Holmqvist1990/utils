package utils

import "testing"

func TestTimer(t *testing.T) {
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
		if test.WantError && err == nil {
			t.Fatalf("%v: expected error", test.Name)
		}
		if err != nil {
			if test.WantError {
				continue
			} else {
				t.Fatalf("%v: did not want error, got: %v", test.Name, err)
			}
		}
		var finishedIn int
		for !timer.Finished() {
			timer.Tick()
			finishedIn++
		}
		if finishedIn != test.FinishedIn {
			t.Fatalf("%v: wanted: %v, got: %v", test.Name, test.FinishedIn, finishedIn)
		}
	}
}

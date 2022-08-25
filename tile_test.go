package utils

import "testing"

var MAP = []int{
	0, 1, 2, 3,
	4, 5, 6, 7,
	8, 9, 10, 11,
	12, 13, 14, 15,
}

var MAP_XY = [][]int{
	{0, 1, 2, 3},
	{4, 5, 6, 7},
	{8, 9, 10, 11},
	{12, 13, 14, 15},
}

const (
	MAP_WIDTH  = 4
	MAP_HEIGHT = 4
)

func TestClampXYtoMap(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{-100, 0, 0},
		{0, -1, 0},
		{0, 0, 0},
		{1, 0, 1},
		{4, 0, 3},
		{0, 1, 4},
		{0, 4, 12},
		{4, 4, 15},
		{5, 5, 15},
		{5, 100, 15},
	}
	for _, test := range tests {
		x, y := ClampXYToMap(test.x, test.y, MAP_WIDTH, MAP_HEIGHT)
		got := MAP_XY[y][x]
		if got != test.want {
			t.Fatalf("{x: %d, y: %d}: wanted %d, got %d",
				test.x, test.y, test.want, got)
		}
	}
}

func TestClampIndexToMap(t *testing.T) {
	tests := []struct {
		index int
		want  int
	}{
		{-100, 0},
		{-1, 0},
		{0, 0},
		{1, 1},
		{3, 3},
		{4, 4},
		{12, 12},
		{15, 15},
		{16, 15},
		{100, 15},
	}
	for _, test := range tests {
		index := ClampIndexToMap(test.index, MAP_WIDTH, MAP_HEIGHT)
		got := MAP[index]
		if got != test.want {
			t.Fatalf("{index: %d}: wanted %d, got %d",
				test.index, test.want, got)
		}
	}
}

func TestIsXYInRange(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want bool
	}{
		{-1, 0, false},
		{0, -1, false},
		{0, 0, true},
		{3, 0, true},
		{0, 3, true},
		{4, 0, false},
		{0, 5, false},
	}
	for _, test := range tests {
		got := IsXYInRange(test.x, test.y, MAP_WIDTH, MAP_HEIGHT)
		if got != test.want {
			t.Fatalf("{x: %d, y: %d}: wanted %v, got %v",
				test.x, test.y, test.want, got)
		}
	}
}

func TestIsIndexInRange(t *testing.T) {
	tests := []struct {
		index int
		want  bool
	}{
		{-1, false},
		{0, true},
		{15, true},
		{16, false},
		{17, false},
	}
	for _, test := range tests {
		got := IsIndexInRange(test.index, MAP_WIDTH, MAP_HEIGHT)
		if got != test.want {
			t.Fatalf("{index: %d}: wanted %v, got %v",
				test.index, test.want, got)
		}
	}
}

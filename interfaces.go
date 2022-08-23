package utils

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Integer interface {
	constraints.Signed | constraints.Unsigned
}

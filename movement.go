package utils

import "fmt"

// Normalized, lerped movement with drag.
type Mover[T Float] struct {
	PosX T
	PosY T

	DeltaX T
	DeltaY T

	MinimumVelocity T

	Acceleration T
	Drag         T
}

func NewMover[T Float](x, y, minimumVel, acc, drag T) (*Mover[T], error) {
	if !(acc > 0) {
		return nil, fmt.Errorf("acc must be positive: %v", acc)
	}
	return &Mover[T]{
		PosX:            x,
		PosY:            y,
		MinimumVelocity: minimumVel,
		Acceleration:    acc,
		Drag:            drag,
	}, nil
}

func (m *Mover[T]) AddMovement(x, y T) {
	nx, ny := Normalize(x, y)
	m.DeltaX += nx * m.Acceleration
	m.DeltaY += ny * m.Acceleration
}

func (m *Mover[T]) Update() {
	m.PosX += m.DeltaX
	m.PosY += m.DeltaY
	m.drag()
}

func (m *Mover[T]) UpdateDelta(delta T) {
	m.PosX += m.DeltaX * delta
	m.PosY += m.DeltaY * delta
	m.drag()
}

func (m *Mover[T]) drag() {
	m.DeltaX *= 1 - m.Drag
	m.DeltaX = MinOrZero(m.DeltaX, m.MinimumVelocity)
	m.DeltaY *= 1 - m.Drag
	m.DeltaY = MinOrZero(m.DeltaY, m.MinimumVelocity)
}

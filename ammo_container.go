package utils

type AmmoContainer[T comparable] map[T]uint

func (a AmmoContainer[T]) CurrentAmmo(t T) uint {
	return a[t]
}

func (a AmmoContainer[T]) Take(t T, hasAlready, capacity uint) uint {
	toTake := capacity - hasAlready
	if toTake == 0 {
		return hasAlready
	}

	ammo := a[t]

	if toTake > ammo {
		a[t] = 0
		return ammo
	}

	a[t] = ammo - toTake

	return toTake
}

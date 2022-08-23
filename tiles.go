package utils

/*
	1D index (sprite) to tileset rectangle.
	Assumes a square tileset.

		sprite = 3, tileSize = N, tilesetWidth = 4
		. . x x .
		. . x x .
		. . . . .
		. . . . .
*/
func TilesetRect[T Integer](sprite, tileSize, tilesetWidth T) (x1, y1, x2, y2 T) {
	x1 = (sprite % tilesetWidth) * tileSize
	y1 = (sprite / tilesetWidth) * tileSize
	x2 = x1 + tileSize
	y2 = y2 + tileSize
	return x1, y1, x2, y2
}

// 1D to 2D coordinate conversion.
func IndexToXY[T Integer](idx, tilesetWidth T) (x, y T) {
	x = idx % tilesetWidth
	y = idx / tilesetWidth
	return x, y
}

// 2D to 1D coordinate conversion.
func XYToIndex[T Integer](x, y, tilesetWidth T) T {
	return x + y*tilesetWidth
}

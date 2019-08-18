package rotatematrix

type matrix = [][]uint

func rotateMatrix(input matrix) matrix {
	return input
}

type direction int

const (
	right direction = iota
	down  direction = iota
	left  direction = iota
	up    direction = iota
)

func replace(m matrix, x uint, y uint, v uint) uint {
	var capture = m[y][x]
	m[y][x] = v
	return capture
}

type nextPoint struct {
	x int
	y int
	d direction
}

func getNextPoint(x uint, y uint, size uint, d direction) nextPoint {
	return nextPoint{0, 0, right}
}

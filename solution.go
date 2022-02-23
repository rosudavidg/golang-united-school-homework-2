package square

import (
	"math"
)

// Custom int type which defines a shape
type ShapeSides int

// Constants that represents the number of shapes for a shape
const (
	SidesCircle   ShapeSides = 0
	SidesTriangle ShapeSides = 3
	SidesSquare   ShapeSides = 4
)

func calcCircleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func calcTriangleArea(sideLen float64) float64 {
	return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
}

func calcSquareArea(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func CalcSquare(sideLen float64, sidesNum ShapeSides) float64 {
	switch sidesNum {
	case SidesCircle:
		return calcCircleArea(sideLen)
	case SidesTriangle:
		return calcTriangleArea(sideLen)
	case SidesSquare:
		return calcSquareArea(sideLen)
	}

	return 0
}

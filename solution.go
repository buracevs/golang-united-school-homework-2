package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type Shape int

const (
	SidesCircle   Shape = iota
	SidesTriangle       = iota + 2
	SidesSquare
	SidesUndefined
)

func CalcSquare(sideLen float64, sidesNum Shape) float64 {
	var result float64
	switch sidesNum {
	case SidesCircle:
		result = math.Pow(sideLen, 2) * math.Pi
	case SidesTriangle:
		result = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case SidesSquare:
		result = math.Pow(sideLen, 2)
	default:
		result = 0
	}
	return result
}

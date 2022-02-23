package square

import "testing"

func TestCalcSquareCircle(t *testing.T) {

	result := CalcSquare(2, SidesCircle)

	if result != 12.566370614359172 {
		t.Error("Expected:4, actual was:", result)
	}
}

func TestCalcSquareTriangle(t *testing.T) {

	result := CalcSquare(11.2, SidesTriangle)

	if result != 54.31711332535998 {
		t.Error("Expected:4, actual was:", result)
	}
}

func TestCalcSquareSquare(t *testing.T) {

	result := CalcSquare(2, SidesSquare)

	if result != 4 {
		t.Error("Expected:4, actual was:", result)
	}
}
func TestCalcSquareUndefined(t *testing.T) {

	result := CalcSquare(2, SidesUndefined)

	if result != 0 {
		t.Error("Expected:4, actual was:", result)
	}
}

package lattice

import (
	"testing"
)

func TestLatticeDimention(t *testing.T) {
	var expectedDim int = 2

	lattice, err := New(20, 10)

	if err != nil {
		t.Error("The lattice's creation fail.")
	}

	if lattice.Dimention != expectedDim {
		t.Error("The lattice is nil afeter criation.")
	}
}

func TestLatticeSetCellValue(t *testing.T) {
	X := 7
	Y := 5
	var expectedValue float32 = 2

	lattice, err := New(X, Y)

	if err != nil {
		t.Error(err)
	}

	for y := 0; y < Y; y += 1 {
		for x := 0; x < X; x += 1 {
			if lattice.At(x, y) != 0 {
				t.Errorf("Cell (%d, %d) value should be not nil and equal to 0.", x, y)
			}
		}
	}

	lattice.Set(expectedValue, 3, 2)

	retrieved := lattice.At(3, 2)

	if retrieved != expectedValue {
		t.Errorf("Expected value is %f, but %f was retrieved.", expectedValue, retrieved)
	}
}

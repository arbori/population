package lattice

import (
	"math/rand"
	"testing"

	"github.com/arbori/population.git/population/space"
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

	for y := 0; y < lattice.Limits[1]; y += 1 {
		for x := 0; x < lattice.Limits[0]; x += 1 {
			if lattice.At(x, y) != nil {
				t.Errorf("The value at (%d, %d) suppose to be nil.", x, y)
			}
		}
	}

	expectedDim = 1
	lattice, err = New(42)

	if err != nil {
		t.Error("The lattice's creation fail.")
	}

	if lattice.Dimention != expectedDim {
		t.Error("The lattice is nil afeter criation.")
	}

	for x := 0; x < lattice.Limits[0]; x += 1 {
		if lattice.At(x) != nil {
			t.Errorf("The value at (%d) suppose to be nil.", x)
		}
	}
}

func TestLatticeSetCellValue(t *testing.T) {
	X := 7
	Y := 5

	lattice, err := New(X, Y)

	if err != nil {
		t.Error(err)
	}

	var expectedValue float32
	var retrieved float32

	for y := 0; y < Y; y += 1 {
		for x := 0; x < X; x += 1 {
			expectedValue = 10 * rand.Float32()

			lattice.Set(expectedValue, x, y)
			retrieved = lattice.At(x, y).(float32)

			if retrieved != expectedValue {
				t.Errorf("Expected value is %f, but %f was retrieved.", expectedValue, retrieved)
			}
		}
	}
}

func TestArrayOutOfBoundsLessThanInterval(t *testing.T) {
	var X int = 0
	var Y int = 0

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("An array out of bounds was not catch with y == %d.", Y)
		}
	}()

	X = 13
	Y = 7
	var valueTest float32 = 42.0

	lattice, err := New(X, Y)

	if err != nil {
		t.Error(err)
	}

	X_out := 0
	Y_out := -1
	lattice.Set(valueTest, X_out, Y_out)
}

func TestArrayOutOfBoundsGreaterThanInterval(t *testing.T) {
	var X int = 0
	var Y int = 0

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("An array out of bounds was not catch with y == %d.", Y)
		}
	}()

	X = 13
	Y = 7
	var valueTest float32 = 42.0

	lattice, err := New(X, Y)

	if err != nil {
		t.Error(err)
	}

	X_out := 0
	Y_out := 7
	lattice.Set(valueTest, X_out, Y_out)
}

func TestDimentionOutOfBounds(t *testing.T) {
	var X int = 0
	var Y int = 0
	var Z int = 0

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("A dimention out of bounds was not catch with only (%d, %d), not (%d, %d, %d).", X, Y, X, Y, Z)
		}
	}()

	X = 13
	Y = 7
	Z = 2
	var valueTest float32 = 42.0

	lattice, err := New(X, Y, Z)

	if err != nil {
		t.Error(err)
	}

	X_out := 0
	Y_out := 7
	lattice.Set(valueTest, X_out, Y_out)
}

func TestMotion(t *testing.T) {
	X := space.Point([]int{5, 5})

	env, err := New(X...)

	if err != nil {
		t.Error(err)
	}

	x := space.Point([]int{4, 4})
	expected := space.Point([]int{x[0] % X[0], x[1] % X[1]})
	point := env.Enclose(x)

	if !point.Equals(expected) {
		t.Errorf("The expected point is (%d, %d), after motion to (%d, %d). But (%d, %d) was obtained.",
			expected[0], expected[1], x[0], x[1], point[0], point[1])
	}

	x = space.Point([]int{23, 42})
	expected = space.Point([]int{x[0] % X[0], x[1] % X[1]})
	point = env.Enclose(x)

	if !point.Equals(expected) {
		t.Errorf("The expected point is (%d, %d), after motion to (%d, %d). But (%d, %d) was obtained.",
			expected[0], expected[1], x[0], x[1], point[0], point[1])
	}

	x = space.Point([]int{-51, -16})
	expected = space.Point([]int{(-x[0] + X[0]) % X[0], (-x[1] + X[1]) % X[1]})

	point = env.Enclose(x)

	if !point.Equals(expected) {
		t.Errorf("The expected point is (%d, %d), after motion to (%d, %d). But (%d, %d) was obtained.",
			expected[0], expected[1], x[0], x[1], point[0], point[1])
	}
}

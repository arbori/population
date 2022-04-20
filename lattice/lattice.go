package lattice

import (
	"errors"
	"fmt"
)

type Lattice struct {
	Dimention int
	Limits    []int
	lines     []interface{}
}

func New(dim ...int) (Lattice, error) {
	if len(dim) <= 0 {
		return Lattice{}, errors.New("Wrong dimention")
	}

	dimention := len(dim)

	result := Lattice{
		Dimention: dimention,
		Limits:    dim,
		lines:     makeLine(dim...),
	}

	return result, nil
}

func (l Lattice) At(x ...int) float32 {
	if len(x) != l.Dimention {
		panic(fmt.Sprintf("Dimention out of bounds: The X's dimention %d is diferent than lattice's dimention %d.", len(x), l.Dimention))
	}

	line := get(x, l.lines)

	return (*line).(float32)
}

func (l Lattice) Set(value float32, x ...int) {
	if len(x) != l.Dimention {
		panic(fmt.Sprintf("Dimention out of bounds: The X's dimention %d is diferent than lattice's dimention %d.", len(x), l.Dimention))
	}

	line := get(x, l.lines)

	(*line) = value
}

var defaultCellValue float32 = 0

func makeLine(dim ...int) []interface{} {
	var rows []interface{} = make([]interface{}, dim[0])

	for i := 0; i < dim[0]; i += 1 {
		fillLine(&rows[i], dim[1:])
	}

	return rows
}

func fillLine(cell *interface{}, dim []int) {
	if len(dim) < 1 {
		(*cell) = defaultCellValue
		return
	}

	*cell = make([]interface{}, dim[0])

	size := dim[0]

	if len(dim) == 1 {
		for i := 0; i < size; i += 1 {
			(*cell).([]interface{})[i] = defaultCellValue
		}
	} else {
		for i := 0; i < size; i += 1 {
			fillLine(&(*cell).([]interface{})[i], dim[1:])
		}
	}
}

func get(x []int, row []interface{}) *interface{} {
	if x[0] < 0 || x[0] >= len(row) {
		panic(fmt.Sprintf("Index out of bounds: %d is out of [0, %d).", x[0], len(row)))
	}

	if len(x) == 1 {
		return &row[x[0]]
	}

	return get(x[1:], row[x[0]].([]interface{}))
}

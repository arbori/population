package lattice

import (
	"errors"
)

type Lattice struct {
	Dimention int
	lines     []interface{}
}

func New(dim ...int) (Lattice, error) {
	if len(dim) < 1 {
		return Lattice{}, errors.New("Wrong dimention")
	}

	dimention := len(dim)

	result := Lattice{
		Dimention: dimention,
		lines:     makeLine(dim...),
	}

	return result, nil
}

func (l Lattice) At(x ...int) float32 {
	return getCellValue(x, l.lines)
}

func (l Lattice) Set(value float32, x ...int) {
	setCellValue(value, x, l.lines)
}

func makeLine(dim ...int) []interface{} {
	var rows []interface{} = make([]interface{}, dim[0])

	for i := 0; i < dim[0]; i += 1 {
		fillLine(&rows[i], dim[1:])
	}

	return rows
}

var defaultCellValue float32 = 0

func fillLine(cell *interface{}, dim []int) {
	if len(dim) < 1 {
		return
	}

	*cell = make([]interface{}, dim[0])

	size := dim[0]

	if len(dim) == 1 {
		for i := 0; i < size; i += 1 {
			(*cell).([]interface{})[i] = defaultCellValue
		}
	} else {
		dim = dim[1:]

		for i := 0; i < size; i += 1 {
			fillLine(&(*cell).([]interface{})[i], dim)
		}
	}
}

func getCellValue(x []int, row []interface{}) float32 {
	if len(x) == 1 {
		return row[x[0]].(float32)
	}

	return getCellValue(x[1:], row[x[0]].([]interface{}))
}

func setCellValue(value float32, x []int, row []interface{}) {
	if len(x) == 1 {
		row[x[0]] = value
		return
	}

	setCellValue(value, x[1:], row[x[0]].([]interface{}))
}

package space

type PointError struct {
	msg string
}

func (pe PointError) Error() string {
	return pe.msg
}

type Point []int

func (p *Point) Assign(point *Point) error {
	if len(*p) != len(*point) {
		return PointError{msg: "The length of these two points are different."}
	}

	for i := 0; i < len(*point); i += 1 {
		(*p)[i] = (*point)[i]
	}

	return nil
}

func (p *Point) Add(point *Point) error {
	if len(*p) != len(*point) {
		return PointError{msg: "The length of these two points are different."}
	}

	for i := 0; i < len(*point); i += 1 {
		(*p)[i] += (*point)[i]
	}

	return nil
}

func (p *Point) Equals(point Point) bool {
	if point == nil || len(*p) != len(point) {
		return false
	}

	for i := 0; i < len(point); i += 1 {
		if (*p)[i] != (point)[i] {
			return false
		}
	}

	return true
}

type NeighborhoodMotion struct {
	Size       int
	Directions []Point
}

func MakeNeighborhoodMotion(size int, dimention int) NeighborhoodMotion {
	result := NeighborhoodMotion{
		Size:       size,
		Directions: make([]Point, size),
	}

	for s := 0; s < size; s += 1 {
		result.Directions[s] = make([]int, dimention)
	}

	return result
}

type Cell struct {
	Value   float32
	Content interface{}
}

type Environment struct {
	X       int
	Y       int
	Cells   [][]Cell
	mirror  [][]Cell
	Motion  NeighborhoodMotion
	Inertia float32
}

func MakeEnvironment(X int, Y int, neighborhoodMotion *NeighborhoodMotion, inertia float32) Environment {
	environment := Environment{
		X:       X,
		Y:       Y,
		Cells:   make([][]Cell, X),
		mirror:  make([][]Cell, X),
		Motion:  *neighborhoodMotion,
		Inertia: inertia,
	}

	for x := 0; x < environment.X; x += 1 {
		environment.Cells[x] = make([]Cell, Y)
		environment.mirror[x] = make([]Cell, Y)

		for y := 0; y < environment.Y; y += 1 {
			environment.Cells[x][y] = Cell{Value: 0.0, Content: nil}
			environment.mirror[x][y] = Cell{Value: 0.0, Content: nil}
		}
	}

	return environment
}

func (e *Environment) NeighborhoodValues(x int, y int) []float32 {
	neighborhood := make([]float32, e.Motion.Size)

	if len(e.Motion.Directions) == 0 || len(e.Motion.Directions) != e.Motion.Size {
		return neighborhood
	}

	var i int
	var j int

	for index := 0; index < e.Motion.Size; index += 1 {
		j = e.Motion.Directions[index][0] + x
		i = e.Motion.Directions[index][1] + y

		if j < 0 {
			j = e.X + j
		} else if j >= e.X {
			j = j - e.X
		}

		if i < 0 {
			i = e.Y + i
		} else if i >= e.Y {
			i = i - e.Y
		}

		neighborhood[index] = e.Cells[j][i].Value
	}

	return neighborhood
}

func (e *Environment) GetNewPosition(position *Point, directionChoosed int) Point {
	result := []int{
		e.Motion.Directions[directionChoosed][0] + (*position)[0],
		e.Motion.Directions[directionChoosed][1] + (*position)[1]}

	if result[0] < 0 {
		result[0] = e.X + result[0]
	} else if result[0] >= e.X {
		result[0] = result[0] - e.X
	}

	if result[1] < 0 {
		result[1] = e.Y + result[1]
	} else if result[1] >= e.Y {
		result[1] = result[1] - e.Y
	}

	return result
}

/*
func (e *Environment) ApplyRule(r rule.Rule) {
	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			e.mirror[x][y].Value = r.Transition(e.NeighborhoodValues(x, y))
		}
	}

	for y := 0; y < e.Y; y += 1 {
		for x := 0; x < e.X; x += 1 {
			e.Cells[x][y].Value = e.mirror[x][y].Value
		}
	}
}
*/

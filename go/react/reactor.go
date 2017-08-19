package react

type DefaultReactor struct {
}

func New() Reactor {
	return &DefaultReactor{}
}

func (r *DefaultReactor) CreateInput(val int) InputCell {
	inputCell := &DefaultInputCell{
		Val:       val,
		Observers: []CellObserver{},
	}
	return inputCell
}

func (r *DefaultReactor) CreateCompute1(cell Cell, f func(int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Cell1:         cell.Value(),
		Type:          COMPUTE1_TYPE,
		Observers:     []CellObserver{},
		Compute1:      f,
		Callbacks:     []func(int){},
		ChangeHistory: map[string]int{},
	}
	switch v := cell.(type) {
	case Observable:
		v.AddObserver(computeCell, 0)
	}
	return computeCell
}

func (r *DefaultReactor) CreateCompute2(cell1 Cell, cell2 Cell, f func(int, int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Cell1:         cell1.Value(),
		Cell2:         cell2.Value(),
		Type:          COMPUTE2_TYPE,
		Observers:     []CellObserver{},
		Compute2:      f,
		Callbacks:     []func(int){},
		ChangeHistory: map[string]int{},
		EmptyInput:    checkForEmptyInput(f, cell1, cell2),
	}

	switch v := cell1.(type) {
	case Observable:
		v.AddObserver(computeCell, 0)
	}
	switch v := cell2.(type) {
	case Observable:
		v.AddObserver(computeCell, 1)
	}
	return computeCell
}

// checks if the previous inputs might produces the same output all the time
func checkForEmptyInput(compute2 func(int, int) int, cell1, cell2 Cell) bool {
	var aCompute1, bCompute1 func(int) int
	var aCompute2, bCompute2 func(int, int) int
	switch v := cell1.(type) {
	case *DefaultComputeCell:
		aCompute1 = v.Compute1
		aCompute2 = v.Compute2
	}
	switch v := cell2.(type) {
	case *DefaultComputeCell:
		bCompute1 = v.Compute1
		bCompute2 = v.Compute2
	}
	if aCompute1 != nil && bCompute1 != nil {
		return compute2(aCompute1(10), bCompute1(10)) == compute2(aCompute1(100), bCompute1(100))
	}
	if aCompute2 != nil && bCompute2 != nil {
		return compute2(aCompute2(10, 10), bCompute2(10, 10)) == compute2(aCompute2(100, 100), bCompute2(100, 100))
	}
	return false
}

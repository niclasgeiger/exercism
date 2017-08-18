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
		Cell1:     cell.Value(),
		Type:      COMPUTE1_TYPE,
		Observers: []CellObserver{},
		Compute1:  f,
	}
	return computeCell
}

func (r *DefaultReactor) CreateCompute2(cell1 Cell, cell2 Cell, f func(int, int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Cell1:     cell1.Value(),
		Cell2:     cell2.Value(),
		Type:      COMPUTE2_TYPE,
		Observers: []CellObserver{},
		Compute2:  f,
	}
	return computeCell
}

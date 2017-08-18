package react

type DefaultReactor struct {
}

func New() Reactor {
	return &DefaultReactor{}
}

func (r *DefaultReactor) CreateInput(val int) InputCell {
	return NewInputCell(val)
}

func (r *DefaultReactor) CreateCompute1(cell Cell, f func(int) int) ComputeCell {
	return NewCompute1Cell(cell, f)
}

func (r *DefaultReactor) CreateCompute2(cell1 Cell, cell2 Cell, f func(int, int) int) ComputeCell {
	return NewCompute2Cell(cell1, cell2, f)
}

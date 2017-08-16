package react

const testVersion = 5

type DefaultReactor struct {
	ActionMap map[Cell][]ComputeCell
}

type DefaultCell struct {
	Val int
	Reactor *DefaultReactor
}
type DefaultInputCell struct {
	Val int
	Reactor *DefaultReactor
}
type DefaultComputeCell struct {
	Val int
	Reactor *DefaultReactor
	Callback func(int)
	Compute1 func(int) int
	Compute2 func(int, int) int
}

type DefaultCanceler struct {

}

func New() Reactor {
	return DefaultReactor{
		ActionMap: make(map[Cell][]ComputeCell),
	}
}


// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (r *DefaultReactor) CreateInput(val int) InputCell {
	cell := &DefaultInputCell{
		Val:val,
		Reactor:r,
	}
	r.ActionMap[cell] = []ComputeCell{}
	return cell
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r *DefaultReactor) CreateCompute1(cell Cell, f func(int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Val:      f(cell.Value()),
		Compute1: f,
		Reactor:  r,
	}
	r.ActionMap[cell] = append(r.ActionMap[cell], computeCell)
	return computeCell
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r *DefaultReactor) CreateCompute2(cell1 Cell, cell2 Cell, f func(int, int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Val: f(cell1.Value(), cell2.Value()),
		Compute2:f,
	}
	return computeCell
}

func (d *DefaultCell) Value() int {
	return d.Val
}

func (d *DefaultInputCell) Value() int {
	return d.Val
}

func (d *DefaultInputCell) SetValue(i int) {
	d.Val = i
}

func (d *DefaultComputeCell) Value() int {
	return d.Val
}

func (d *DefaultComputeCell) AddCallback(f func(int)) Canceler {
	d.Callback = f
	return new(DefaultCanceler)
}

func (d *DefaultCanceler) Cancel() {

}
package react

const testVersion = 5

const(
	ONE ComputeCellType = iota
	TWO
)
type ComputeCellType int
type DefaultReactor struct {
	ActionMap map[Cell][]func(int)
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
	Type ComputeCellType
	Cell1 int
	Cell2 int
	Reactor *DefaultReactor
	Callback func(int)
	Compute1 func(int) int
	Compute2 func(int, int) int
	Canceler *DefaultCanceler
}

type DefaultCanceler struct {
	ComputeCell *DefaultComputeCell
}

func New() Reactor {
	return &DefaultReactor{
		ActionMap: make(map[Cell][]func(int)),
	}
}


// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (r *DefaultReactor) CreateInput(val int) InputCell {
	cell := &DefaultInputCell{
		Val:val,
		Reactor:r,
	}
	r.ActionMap[cell] = []func(int){}
	return cell
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (r *DefaultReactor) CreateCompute1(cell Cell, f func(int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Type:ONE,
		Cell1:      cell.Value(),
		Compute1: f,
		Reactor:  r,
	}
	callback := func(n int) {
		computeCell.Cell1 = n
	}
	computeCell.Canceler = computeCell.AddCallback(callback).(*DefaultCanceler)
	r.ActionMap[cell] = append(r.ActionMap[cell], callback)
	return computeCell
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (r *DefaultReactor) CreateCompute2(cell1 Cell, cell2 Cell, f func(int, int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		Type:TWO,
		Cell1: cell1.Value(),
		Cell2: cell2.Value(),
		Compute2:f,
	}
	callback1 := func(n int){
		computeCell.Cell1 = n
	}
	callback2 := func(n int){
		computeCell.Cell2 = n
	}
	//TODO: Propagation for compute cells
	computeCell.Canceler = computeCell.AddCallback(callback1).(*DefaultCanceler)
	computeCell.Canceler = 	computeCell.AddCallback(callback2).(*DefaultCanceler)
	r.ActionMap[cell1] = append(r.ActionMap[cell1], callback1)
	r.ActionMap[cell2] = append(r.ActionMap[cell2], callback2)
	return computeCell
}

func (d *DefaultCell) Value() int {
	return d.Val
}

func (d *DefaultInputCell) Value() int {
	return d.Val
}

func (d *DefaultInputCell) SetValue(i int) {
	for _, action := range d.Reactor.ActionMap[d]{
		action(i)
	}
	d.Val = i
}

func (d *DefaultComputeCell) Value() int {
	switch d.Type {
	case ONE:
		return d.Compute1(d.Cell1)
	case TWO:
		return d.Compute2(d.Cell1, d.Cell2)
	}
	return -1
}

func (d *DefaultComputeCell) AddCallback(f func(int)) Canceler {
	if d.Canceler == nil {
		d.Canceler = &DefaultCanceler{
			ComputeCell:d,
		}
	}
	d.Callback = f
	return d.Canceler
}

func (d *DefaultCanceler) Cancel() {
	d.ComputeCell.Callback = func(n int){}
}
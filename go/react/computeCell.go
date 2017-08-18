package react

type CellType int

const (
	COMPUTE1_TYPE = iota
	COMPUTE2_TYPE
)

type DefaultComputeCell struct {
	cell1     Cell
	cell2     Cell
	oldValue  int
	compute1  func(int) int
	compute2  func(int, int) int
	Type      CellType
	observers []Observer
	callbacks []func(int)
}

func NewCompute1Cell(cell Cell, f func(int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		cell1:     cell,
		Type:      COMPUTE1_TYPE,
		observers: []Observer{},
		compute1:  f,
		callbacks: []func(int){},
	}
	switch v := cell.(type) {
	case Observable:
		{
			v.AddObserver(computeCell)
		}
	}
	return computeCell
}

func NewCompute2Cell(cell1, cell2 Cell, f func(int, int) int) ComputeCell {
	computeCell := &DefaultComputeCell{
		cell1:     cell1,
		cell2:     cell2,
		Type:      COMPUTE2_TYPE,
		observers: []Observer{},
		compute2:  f,
		callbacks: []func(int){},
	}
	switch v := cell1.(type) {
	case Observable:
		{
			v.AddObserver(computeCell)
		}
	}
	switch v := cell2.(type) {
	case Observable:
		{
			v.AddObserver(computeCell)
		}
	}
	return computeCell
}
func (c *DefaultComputeCell) currentVal() int {
	val := 0
	switch c.Type {
	case COMPUTE1_TYPE:
		val = c.compute1(c.cell1.Value())
	case COMPUTE2_TYPE:
		val = c.compute2(c.cell1.Value(), c.cell2.Value())
	}
	return val
}

func (c *DefaultComputeCell) Value() int {
	val := c.currentVal()
	c.oldValue = val
	return val
}

func (c *DefaultComputeCell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, callback)
	return &DefaultCanceler{} //TODO
}

func (c *DefaultComputeCell) AddObserver(observer Observer, cellIndex int) {
	c.observers = append(c.observers, observer)
}

func (c *DefaultComputeCell) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update()
	}
}

func (c *DefaultComputeCell) Update() {
	if c.currentVal() != c.oldValue {
		c.NotifyObservers()
		for _, callback := range c.callbacks {
			callback(c.Value())
		}
	}
}

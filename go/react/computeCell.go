package react

type CellType int

const (
	COMPUTE1_TYPE = iota
	COMPUTE2_TYPE
)

type DefaultComputeCell struct {
	Cell1     int // TODO replace with cell instance?
	Cell2     int // TODO replace with cell instance?
	Compute1  func(int) int
	Compute2  func(int, int) int
	Type      CellType
	Observers []CellObserver
}

func (c *DefaultComputeCell) Value() int {
	val := 0
	switch c.Type {
	case COMPUTE1_TYPE:
		val = c.Compute1(c.Cell1)
	case COMPUTE2_TYPE:
		val = c.Compute2(c.Cell1, c.Cell2)
	}
	c.NotifyObservers(val)
	return val
}

func (c *DefaultComputeCell) AddCallback(func(int)) Canceler {
	// TODO
	return &DefaultCanceler{}
}

func (c *DefaultComputeCell) AddObserver(observer Observer, cellIndex int) {
	c.Observers = append(c.Observers, CellObserver{cellIndex, observer})
}

func (c *DefaultComputeCell) NotifyObservers(val int) {
	for _, observer := range c.Observers {
		observer.Update(val, observer.Index)
	}
}

func (c *DefaultComputeCell) Update(val int, index int) {
	switch index {
	case 0:
		c.Cell1 = val
	case 1:
		c.Cell2 = val
	}
	c.NotifyObservers(c.Value())
}

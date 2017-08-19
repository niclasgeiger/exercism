package react

type CellType int

const (
	COMPUTE1_TYPE = iota
	COMPUTE2_TYPE
)

type DefaultComputeCell struct {
	Cell1         int
	Cell2         int
	Compute1      func(int) int
	Compute2      func(int, int) int
	Type          CellType
	Observers     []CellObserver
	Callbacks     []func(int)
	ChangeHistory map[string]int
	EmptyInput    bool
}

func (c *DefaultComputeCell) compute(a, b int) int {
	switch c.Type {
	case COMPUTE1_TYPE:
		return c.Compute1(a)
	case COMPUTE2_TYPE:
		return c.Compute2(a, b)
	}
	return 0
}

func (c *DefaultComputeCell) Value() int {
	return c.compute(c.Cell1, c.Cell2)
}

func (c *DefaultComputeCell) AddCallback(callback func(int)) Canceler {
	canceler := &DefaultCanceler{
		Active: true,
	}
	c.Callbacks = append(c.Callbacks, func(val int) {
		if canceler.Active {
			callback(val)
		}
	})
	return canceler
}

func (c *DefaultComputeCell) AddObserver(observer Observer, cellIndex int) {
	c.Observers = append(c.Observers, CellObserver{cellIndex, observer})
}

func (c *DefaultComputeCell) NotifyObservers(val int, correlationId string) {
	for _, observer := range c.Observers {
		observer.Update(val, observer.Index, correlationId)
	}
	if c.ChangeHistory[correlationId] < 1 {
		for _, callback := range c.Callbacks {
			callback(val)
		}
	}
	c.ChangeHistory[correlationId] = 1
}

func (c *DefaultComputeCell) Update(val int, index int, correlationId string) {
	oldVal := c.compute(c.Cell1, c.Cell2)
	changed := false
	switch index {
	case 0:
		if val != c.Cell1 {
			c.Cell1 = val
			changed = true
		}
	case 1:
		if val != c.Cell2 {
			c.Cell2 = val
			changed = true
		}
	}
	if !c.EmptyInput && changed && c.compute(c.Cell1, c.Cell2) != oldVal {
		c.NotifyObservers(c.Value(), correlationId)
	}
}

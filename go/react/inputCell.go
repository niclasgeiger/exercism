package react

type DefaultInputCell struct {
	Val       int
	Observers []CellObserver
}

func (c *DefaultInputCell) Value() int {
	return c.Val
}

func (c *DefaultInputCell) SetValue(val int) {
	c.Val = val
}

func (c *DefaultInputCell) AddObserver(cell Observer, cellIndex int) {
	c.Observers = append(c.Observers, CellObserver{cellIndex, cell})
}

func (c *DefaultInputCell) NotifyObservers(val int) {
	for _, observer := range c.Observers {
		observer.Update(val, observer.Index)
	}
}

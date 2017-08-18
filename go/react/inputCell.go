package react

type DefaultInputCell struct {
	Val       int
	Observers []Observer
}

func NewInputCell(val int) InputCell {
	return &DefaultInputCell{
		Val:       val,
		Observers: []Observer{},
	}
}

func (c *DefaultInputCell) Value() int {
	return c.Val
}

func (c *DefaultInputCell) SetValue(val int) {

	if val != c.Val {
		c.Val = val
		c.NotifyObservers()
	}
}

func (c *DefaultInputCell) AddObserver(observer Observer) {
	c.Observers = append(c.Observers, observer)
}

func (c *DefaultInputCell) NotifyObservers() {
	for _, observer := range c.Observers {
		observer.Update()
	}
}

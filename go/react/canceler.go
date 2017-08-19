package react

type DefaultCanceler struct {
	Active bool
}

func (c *DefaultCanceler) Cancel() {
	c.Active = false
}

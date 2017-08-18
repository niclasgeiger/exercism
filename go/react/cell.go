package react

type DefaultCell struct {
	Val int
}

func (c *DefaultCell) Value() int {
	return c.Val
}

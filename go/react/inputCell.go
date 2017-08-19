package react

import (
	"math/rand"
	"time"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type DefaultInputCell struct {
	Val       int
	Observers []CellObserver
}

func (c *DefaultInputCell) Value() int {
	return c.Val
}

func (c *DefaultInputCell) SetValue(val int) {
	c.Val = val
	c.NotifyObservers(val, GenerateCorrelationId())
}

func (c *DefaultInputCell) AddObserver(cell Observer, cellIndex int) {
	c.Observers = append(c.Observers, CellObserver{cellIndex, cell})
}

func (c *DefaultInputCell) NotifyObservers(val int, correlationId string) {
	for _, observer := range c.Observers {
		observer.Update(val, observer.Index, correlationId)
	}
}

func GenerateCorrelationId() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 16)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

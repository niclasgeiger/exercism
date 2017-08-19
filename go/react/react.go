package react

const testVersion = 5

type Observable interface {
	AddObserver(Observer, int)
	NotifyObservers(int, string)
}

type Observer interface {
	Update(val int, index int, correlationId string)
}

type CellObserver struct {
	Index int
	Observer
}

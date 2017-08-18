package react

const testVersion = 5

type Observable interface {
	AddObserver(Observer)
	NotifyObservers()
}

type Observer interface {
	Update()
}

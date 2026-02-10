package domain

type Event[T any] interface {
	GetData() T
	GetName() string
}

type EventDispatcher[T any] interface {
	Dispatch(e Event[T])
}

type EventListener[T any] interface {
	Listen(name string, listener func(e Event[T]))
}

type EventManager[T any] interface {
	EventDispatcher[T]
	EventListener[T]
}

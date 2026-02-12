package domain

type Event[T any] interface {
	GetData() T
	GetName() string
}

type Listener func(e Event[any])
type Unlistener func()

type EventDispatcher[T any] interface {
	Dispatch(e Event[T])
}

type EventListener[T any] interface {
	Listen(name string, listener Listener) Unlistener
}

type EventManager[T any] interface {
	EventDispatcher[T]
	EventListener[T]
}

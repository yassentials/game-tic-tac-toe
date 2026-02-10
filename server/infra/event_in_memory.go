package infra

import (
	"sync"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
)

type InMemoryEventManager struct {
	mu               sync.RWMutex
	eventSubscribers map[string][]func(e domain.Event[any])
}

func NewInMemoryEventManager() *InMemoryEventManager {
	return &InMemoryEventManager{
		eventSubscribers: map[string][]func(e domain.Event[any]){},
	}
}

func (em *InMemoryEventManager) Dispatch(e domain.Event[any]) {
	em.mu.RLock()
	listeners, exists := em.eventSubscribers[e.GetName()]

	if !exists {
		em.mu.RUnlock()
		return
	}

	// make sure data didn't change
	temp := make([]func(e domain.Event[any]), len(listeners))
	copy(temp, listeners)

	em.mu.RUnlock()

	for _, l := range temp {
		go l(e)
	}
}

func (em *InMemoryEventManager) Listen(name string, listener func(e domain.Event[any])) {
	em.eventSubscribers[name] = append(em.eventSubscribers[name], listener)
}

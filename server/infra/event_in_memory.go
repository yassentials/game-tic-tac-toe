package infra

import (
	"maps"
	"sync"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
)

type InMemoryEventManager struct {
	mu               sync.RWMutex
	eventSubscribers map[string]map[int]domain.Listener
	nextId           int
}

func NewInMemoryEventManager() *InMemoryEventManager {
	return &InMemoryEventManager{
		eventSubscribers: map[string]map[int]domain.Listener{},
		nextId:           0,
	}
}

func (em *InMemoryEventManager) Dispatch(e domain.Event[any]) {
	em.mu.RLock()
	subs, exists := em.eventSubscribers[e.GetName()]

	if !exists {
		em.mu.RUnlock()
		return
	}

	// make sure if original map change, doesn't effect the process
	temp := maps.Clone(subs)

	em.mu.RUnlock()

	for _, l := range temp {
		go l(e)
	}
}

func (em *InMemoryEventManager) Listen(name string, listener domain.Listener) domain.Unlistener {
	em.mu.Lock()
	defer em.mu.Unlock()

	id := em.nextId
	em.nextId++

	subs := em.eventSubscribers[name]
	subs[id] = listener

	return func() {
		delete(subs, id)
	}
}

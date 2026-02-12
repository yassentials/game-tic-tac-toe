package infra

import (
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
)

type InMemoryLobby struct {
	games map[string]domain.Game
	mu    sync.RWMutex
}

func NewInMemoryLobby() *InMemoryLobby {
	return &InMemoryLobby{
		games: map[string]domain.Game{},
	}
}

func (l *InMemoryLobby) AddGame(g domain.Game) error {
	code := g.GetCode()

	if l.games[code] != nil {
		return fmt.Errorf("failed to create game, already exists")
	}

	l.games[code] = g

	return nil
}

func (l *InMemoryLobby) FindGameByCode(code string) (domain.Game, error) {
	game := l.games[code]

	if game == nil {
		return game, fmt.Errorf("game with code %s not found.", code)
	}

	if game.IsFull() {
		return game, fmt.Errorf("game with code %s is full.", code)
	}

	return game, nil
}

func (l *InMemoryLobby) DeleteGameByCode(code string) {
	delete(l.games, code)
}

func (l *InMemoryLobby) ListAvailableGames() []domain.Game {
	l.mu.RLock()
	defer l.mu.RUnlock()

	games := []domain.Game{}

	for _, game := range l.games {
		if !game.IsFull() && game.IsPublic() {
			games = append(games, game)
		}
	}

	return games
}

func (l *InMemoryLobby) ListPublicGames() []domain.Game {
	l.mu.RLock()
	defer l.mu.RUnlock()

	games := []domain.Game{}

	for _, game := range l.games {
		if game.IsPublic() {
			games = append(games, game)
		}
	}

	return games
}

func (l *InMemoryLobby) FindRandomGameAvailable() (domain.Game, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	keys := make([]string, 0, len(l.games))

	for key, game := range l.games {
		if !game.IsFull() && game.IsPublic() {
			keys = append(keys, key)
		}
	}

	if len(keys) <= 0 {
		return nil, fmt.Errorf("[Find Random Game] unvailable.\n")
	}

	randKey := rand.IntN(len(keys))

	game := l.games[keys[randKey]]

	return game, nil
}

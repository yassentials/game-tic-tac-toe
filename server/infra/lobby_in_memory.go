package infra

import (
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
)

type InMemoryLobby struct {
	games map[string]domain.Game
}

func NewInMemoryLobby() *InMemoryLobby {
	return &InMemoryLobby{}
}

func (l InMemoryLobby) AddGame(g domain.Game) error {
	code := g.GetCode()

	if l.games[code] != nil {
		return fmt.Errorf("failed to create game, already exists")
	}

	l.games[code] = g

	return nil
}

func (l InMemoryLobby) FindGameByCode(code string) (domain.Game, error) {
	game := l.games[code]

	if game == nil {
		return game, fmt.Errorf("game with code %s not found.", code)
	}

	return game, nil
}

func (l InMemoryLobby) DeleteGameByCode(code string) {
	delete(l.games, code)
}

func (l InMemoryLobby) ListAvailableGames() []domain.Game {
}
func (l InMemoryLobby) FindRandomGameAvailable() domain.Game {
}

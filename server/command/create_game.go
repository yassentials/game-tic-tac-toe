package command

import (
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
	"github.com/yassentials/game-tic-tac-toe/server/event"
)

type CreateGameHandler struct {
	lobby   domain.Lobby
	codeGen func() string
}

func NewCreateGame(lobby domain.Lobby, codeGen func() string) CreateGameHandler {
	return CreateGameHandler{lobby, codeGen}
}

type CreateGameCommand struct {
	Player       domain.Player
	Type         domain.GameType
	EventManager domain.EventManager[any]
}

func (h CreateGameHandler) Handle(cmd CreateGameCommand) (domain.Game, error) {
	const CAPACITY = 2

	game := domain.NewBaseGame(CAPACITY, cmd.Type, cmd.EventManager, h.codeGen)

	game.GetEventManager().Listen(event.EVENT_ROOM_FULL, func(e domain.Event[any]) {
		game.Restart()
	})

	// join self
	if err := game.Join(cmd.Player); err != nil {
		return nil, fmt.Errorf("[Join Game] Failed: %w", err)
	}

	h.lobby.AddGame(game)

	return game, nil
}

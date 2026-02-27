package command

import (
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
	"github.com/yassentials/game-tic-tac-toe/server/infra"
	"github.com/yassentials/game-tic-tac-toe/server/shared/event"
)

type CreateGameHandler struct {
	lobby        domain.Lobby
	codeGen      func() string
	eventManager domain.EventManager[any]
}

func NewCreateGameHandler(lobby domain.Lobby, codeGen func() string) CreateGameHandler {
	return CreateGameHandler{
		lobby:        lobby,
		codeGen:      codeGen,
		eventManager: infra.NewInMemoryEventManager(),
	}
}

type CreateGameCommand struct {
	PlayerName       string
	PlayerCharacater domain.Character
	Type             domain.GameType
}

func (h CreateGameHandler) Handle(cmd CreateGameCommand) (domain.Game, domain.Player, error) {
	const CAPACITY = 2

	game := domain.NewBaseGame(CAPACITY, cmd.Type, h.eventManager, h.codeGen)

	game.GetEventManager().Listen(event.EVENT_ROOM_FULL, func(e domain.Event[any]) {
		game.Restart()
	})

	player := domain.NewGamePlayer(cmd.PlayerName, cmd.PlayerCharacater)

	// join self
	if err := game.Join(player); err != nil {
		return nil, nil, fmt.Errorf("[Join Game] Failed: %w", err)
	}

	h.lobby.AddGame(game)

	return game, player, nil
}

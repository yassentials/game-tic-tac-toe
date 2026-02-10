package command

import (
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
	"github.com/yassentials/game-tic-tac-toe/server/event"
)

type JoinRandomGameHandler struct {
	lobby domain.Lobby
}

func NewJoinRandomGameHandler(lobby domain.Lobby) *JoinRandomGameHandler {
	return &JoinRandomGameHandler{
		lobby,
	}
}

type JoinRandomGameCommand struct {
	Player domain.Player
}

func (h *JoinRandomGameHandler) Handle(cmd JoinRandomGameCommand) (domain.Game, error) {
	game := h.lobby.FindRandomGameAvailable()

	if err := game.Join(cmd.Player); err != nil {
		return nil, fmt.Errorf("[Join Game] Failed: %w", err)
	}

	game.GetEventManager().Dispatch(event.NewPlayerJoinedEvent(event.PlayerJoinedEventData{
		Name:      cmd.Player.GetName(),
		Character: cmd.Player.GetCharacter(),
	}))

	if game.IsFull() {
		game.GetEventManager().Dispatch(event.NewRoomFullEvent())
	}

	return game, nil
}

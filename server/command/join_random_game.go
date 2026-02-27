package command

import (
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
	"github.com/yassentials/game-tic-tac-toe/server/shared/event"
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
	PlayerName       string
	PlayerCharacater domain.Character
}

func (h *JoinRandomGameHandler) Handle(cmd JoinRandomGameCommand) (domain.Game, domain.Player, error) {
	game, err := h.lobby.FindRandomGameAvailable()

	if err != nil {
		return nil, nil, fmt.Errorf("[Join Random Game] Failed: %w", err)
	}

	player := domain.NewGamePlayer(cmd.PlayerName, cmd.PlayerCharacater)

	if err := game.Join(player); err != nil {
		return nil, nil, fmt.Errorf("[Join Random Game] Failed: %w", err)
	}

	game.GetEventManager().Dispatch(event.NewPlayerJoinedEvent(event.PlayerJoinedEventData{
		Name:      player.GetName(),
		Character: player.GetCharacter(),
	}))

	if game.IsFull() {
		game.GetEventManager().Dispatch(event.NewRoomFullEvent())
	}

	return game, player, nil
}

package command

import (
	"fmt"

	"github.com/up9t/game-tic-tac-toe/backend/domain"
	"github.com/up9t/game-tic-tac-toe/backend/event"
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
	PlayerName string
	// PlayerIP         string
	PlayerCharacater domain.Character
}

func (h *JoinRandomGameHandler) Handle(cmd JoinRandomGameCommand) (domain.Game, domain.Player, error) {
	// add a validation if the connnection has alraedy joined to a game.
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

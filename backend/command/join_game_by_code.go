package command

import (
	"fmt"

	"github.com/up9t/game-tic-tac-toe/backend/domain"
	"github.com/up9t/game-tic-tac-toe/backend/event"
)

type JoinGameByCodeHandler struct {
	lobby      domain.Lobby
	codeLength int
}

func NewJoinGameByCodeHandler(lobby domain.Lobby, codeLength int) *JoinGameByCodeHandler {
	return &JoinGameByCodeHandler{
		lobby,
		codeLength,
	}
}

type JoinGameByCodeCommand struct {
	PlayerName       string
	PlayerCharacater domain.Character
	Code             string
}

func (h *JoinGameByCodeHandler) Handle(cmd JoinGameByCodeCommand) (domain.Game, domain.Player, error) {
	code := cmd.Code

	if len(code) != h.codeLength {
		return nil, nil, fmt.Errorf("[Join Code] code must be equal to %d digit\n", h.codeLength)
	}

	game, err := h.lobby.FindGameByCode(code)
	if err != nil {
		return nil, nil, fmt.Errorf("[Join Code] Failed: %w", err)
	}

	player := domain.NewGamePlayer(cmd.PlayerName, cmd.PlayerCharacater)

	if err := game.Join(player); err != nil {
		return nil, nil, fmt.Errorf("[Join Code] Failed: %w", err)
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

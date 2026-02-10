package command

import (
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
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
	Player domain.Player
	Code   string
}

func (h *JoinGameByCodeHandler) Handle(cmd JoinGameByCodeCommand) (domain.Game, error) {
	code := cmd.Code

	if len(code) != h.codeLength {
		return nil, fmt.Errorf("code must be equal to %d digit\n", h.codeLength)
	}

	game, err := h.lobby.FindGameByCode(code)

	if err != nil {
		return nil, err
	}

	if err := game.Join(cmd.Player); err != nil {
		return nil, fmt.Errorf("[Join Game] Failed: %w", err)
	}

	return game, nil
}

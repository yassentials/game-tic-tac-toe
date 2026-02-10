package command

import (
	"errors"
	"fmt"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
	"github.com/yassentials/game-tic-tac-toe/server/event"
)

type TakePositionHandler struct {
	lobby domain.Lobby
}

func NewTakePositionHandler(lobby domain.Lobby) *TakePositionHandler {
	return &TakePositionHandler{lobby}
}

type TakePositionCommand struct {
	Player domain.Player
	Index  int
	Game   domain.Game
}

func (h *TakePositionHandler) Handle(cmd TakePositionCommand) error {
	if cmd.Player == nil {
		return errors.New("player hasn't been initialized")
	}

	if cmd.Game == nil {
		return errors.New("game hasn't been initialized")
	}

	game := cmd.Game

	if err := game.TakePosition(cmd.Player, cmd.Index); err != nil {
		return fmt.Errorf("[Failed Take]: %w", err)
	}

	if game.GetState() == domain.GAME_STATE_RESULT {

		result := game.GetResult()
		switch result {
		case domain.GAME_RESULT_DRAW:
		case domain.GAME_RESULT_EITHER_WIN:
			winner := game.GetWinner()
			game.GetEventManager().Dispatch(event.NewResultAnnouncedEvent(result, winner))
		case domain.GAME_RESULT_NONE:
		default:
			return fmt.Errorf("[Result None or Unknown]\n")
		}
	}

	return nil
}

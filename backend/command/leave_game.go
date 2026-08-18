package command

import "github.com/up9t/game-tic-tac-toe/backend/domain"

type LeaveGameHandler struct {
}

func NewLeaveGameHandler() *LeaveGameHandler {
	return &LeaveGameHandler{}
}

type LeaveGameCommand struct {
	Player domain.Player
	Game   domain.Game
}

func (h *LeaveGameHandler) Handle(cmd LeaveGameCommand) {
	cmd.Game.Leave(cmd.Player)
}

package event

import "github.com/up9t/game-tic-tac-toe/backend/domain"

const EVENT_TAKE_POSITION_SUCCEED = "take-position-Succeed"

type TakePositionSucceedEventData struct {
	Index     int
	Character domain.Character
}

type TakePositionSucceedEvent struct {
	Name string
	Data TakePositionSucceedEventData
}

func NewTakePositionSucceedEvent(data TakePositionSucceedEventData) *TakePositionSucceedEvent {
	return &TakePositionSucceedEvent{
		Name: EVENT_TAKE_POSITION_SUCCEED,
		Data: data,
	}
}

func (e *TakePositionSucceedEvent) GetData() any {
	return e.Data
}

func (e *TakePositionSucceedEvent) GetName() string {
	return e.Name
}

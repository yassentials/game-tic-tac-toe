package event

import "github.com/yassentials/game-tic-tac-toe/server/domain"

const EVENT_TAKE_POSITION_SUCCEED = "take-position-Succeed"

type TakePositionSucceedEventData struct {
	Index     int              `json:"ind"`
	Character domain.Character `json:"cha"`
}

type TakePositionSucceedEvent struct {
	Name string                       `json:"name"`
	Data TakePositionSucceedEventData `json:"data"`
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

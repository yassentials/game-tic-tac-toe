package event

import "github.com/yassentials/game-tic-tac-toe/server/domain"

const EVENT_TAKE_POSITION_SUCCEED = "take-position-Succeed"

type TakePositionSucceedEventData struct {
	Index     int              `json:"ind"`
	Character domain.Character `json:"cha"`
}

type TakePositionSucceedEvent struct {
	data TakePositionSucceedEventData
	name string
}

func NewTakePositionSucceedEvent(data TakePositionSucceedEventData) *TakePositionSucceedEvent {
	return &TakePositionSucceedEvent{
		name: EVENT_TAKE_POSITION_SUCCEED,
		data: data,
	}
}

func (e *TakePositionSucceedEvent) GetData() any {
	return e.data
}

func (e *TakePositionSucceedEvent) GetName() string {
	return e.name
}

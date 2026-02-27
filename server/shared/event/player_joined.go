package event

import "github.com/yassentials/game-tic-tac-toe/server/domain"

const EVENT_PLAYER_JOINED = "player-joined"

type PlayerJoinedEventData struct {
	Name      string           `json:"nam"`
	Character domain.Character `json:"cha"`
}

type PlayerJoinedEvent struct {
	Name string                `json:"name"`
	Data PlayerJoinedEventData `json:"data"`
}

func NewPlayerJoinedEvent(data PlayerJoinedEventData) *PlayerJoinedEvent {
	return &PlayerJoinedEvent{
		Name: EVENT_ROOM_FULL,
		Data: data,
	}
}

func (e *PlayerJoinedEvent) GetData() any {
	return e.Data
}

func (e *PlayerJoinedEvent) GetName() string {
	return e.Name
}

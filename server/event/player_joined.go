package event

import "github.com/yassentials/game-tic-tac-toe/server/domain"

const EVENT_PLAYER_JOINED = "player-joined"

type PlayerJoinedEventData struct {
	Name      string           `json:"nam"`
	Character domain.Character `json:"cha"`
}

type PlayerJoinedEvent struct {
	name string
	data PlayerJoinedEventData
}

func NewPlayerJoinedEvent(data PlayerJoinedEventData) *PlayerJoinedEvent {
	return &PlayerJoinedEvent{
		name: EVENT_ROOM_FULL,
		data: data,
	}
}

func (e *PlayerJoinedEvent) GetData() any {
	return e.data
}

func (e *PlayerJoinedEvent) GetName() string {
	return e.name
}

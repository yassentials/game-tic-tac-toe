package event

import "github.com/up9t/game-tic-tac-toe/backend/domain"

const EVENT_PLAYER_JOINED = "player-joined"

type PlayerJoinedEventData struct {
	Name      string
	Character domain.Character
}

type PlayerJoinedEvent struct {
	Name string
	Data PlayerJoinedEventData
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

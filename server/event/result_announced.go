package event

import "github.com/yassentials/game-tic-tac-toe/server/domain"

const EVENT_RESULT_ANNOUNCED = "result-announced"

type ResultAnnouncedEventData struct {
	Result domain.GameResult
	Player domain.Player
}

type ResultAnnouncedEvent struct {
	name string
	data ResultAnnouncedEventData
}

func NewResultAnnouncedEvent(result domain.GameResult, player domain.Player) *ResultAnnouncedEvent {
	return &ResultAnnouncedEvent{
		name: EVENT_RESULT_ANNOUNCED,
		data: ResultAnnouncedEventData{
			Result: result,
			Player: player,
		},
	}
}

func (e *ResultAnnouncedEvent) GetData() any {
	return e.data
}

func (e *ResultAnnouncedEvent) GetName() string {
	return e.name
}

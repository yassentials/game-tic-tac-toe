package event

import "github.com/yassentials/game-tic-tac-toe/server/domain"

const EVENT_RESULT_ANNOUNCED = "result-announced"

type ResultAnnouncedEventData struct {
	Result domain.GameResult
	Player domain.Player
}

type ResultAnnouncedEvent struct {
	Name string                   `json:"name"`
	Data ResultAnnouncedEventData `json:"data"`
}

func NewResultAnnouncedEvent(result domain.GameResult, player domain.Player) *ResultAnnouncedEvent {
	return &ResultAnnouncedEvent{
		Name: EVENT_RESULT_ANNOUNCED,
		Data: ResultAnnouncedEventData{
			Result: result,
			Player: player,
		},
	}
}

func (e *ResultAnnouncedEvent) GetData() any {
	return e.Data
}

func (e *ResultAnnouncedEvent) GetName() string {
	return e.Name
}

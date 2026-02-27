package event

const EVENT_ROOM_FULL = "room-full"

type RoomFullEvent struct {
	Name string   `json:"name"`
	Data struct{} `json:"data"`
}

func NewRoomFullEvent() *RoomFullEvent {
	return &RoomFullEvent{
		Name: EVENT_ROOM_FULL,
		Data: struct{}{},
	}
}

func (e *RoomFullEvent) GetData() any {
	return e.Data
}

func (e *RoomFullEvent) GetName() string {
	return e.Name
}

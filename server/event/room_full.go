package event

const EVENT_ROOM_FULL = "room-full"

type RoomFullEvent struct {
	name string
}

func NewRoomFullEvent() *RoomFullEvent {
	return &RoomFullEvent{
		name: EVENT_ROOM_FULL,
	}
}

func (e *RoomFullEvent) GetData() any {
	return struct{}{}
}

func (e *RoomFullEvent) GetName() string {
	return e.name
}

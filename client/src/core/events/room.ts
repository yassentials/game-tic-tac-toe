export class RoomFullEvent extends CustomEvent<void> {
  public static readonly EVENT_NAME = "room-full";

  public constructor() {
    super(RoomFullEvent.EVENT_NAME);
  }
}

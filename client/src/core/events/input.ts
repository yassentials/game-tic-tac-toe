export type MouseClickEventData = {
  event: Event;
  index: number;
};

export class MouseClickEvent extends CustomEvent<MouseClickEventData> {
  public static readonly EVENT_NAME = "tile-click";

  public constructor(data: MouseClickEventData) {
    super(MouseClickEvent.EVENT_NAME, { detail: data });
  }
}

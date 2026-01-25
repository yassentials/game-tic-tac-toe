export interface INotifyTurnEventData {
  availablePositions: number[];
}

export class NotifyTurnEvent extends CustomEvent<INotifyTurnEventData> {
  public static readonly EVENT_NAME = "notify-turn";
  public constructor(data: INotifyTurnEventData) {
    super(NotifyTurnEvent.EVENT_NAME, {
      detail: data,
    });
  }
}

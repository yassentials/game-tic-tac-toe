import type { IPlayerCharacter } from "../players/interface";

export interface IUIUpdatePositionTakeEventData {
  index: number;
  character: IPlayerCharacter;
}

export interface IUIActionPositionTakeEventData {
  index: number;
  successfulCallback: () => Promise<void>;
}

export class UIUpdatePositionTakeEvent extends CustomEvent<IUIUpdatePositionTakeEventData> {
  public static readonly EVENT_NAME = "ui-update-position-take";

  public constructor(data: IUIUpdatePositionTakeEventData) {
    super(UIUpdatePositionTakeEvent.EVENT_NAME, {
      detail: data,
    });
  }
}

export class UIActionPositionTakeEvent extends CustomEvent<IUIActionPositionTakeEventData> {
  public static readonly EVENT_NAME = "ui-action-position-take";

  public constructor(data: IUIActionPositionTakeEventData) {
    super(UIActionPositionTakeEvent.EVENT_NAME, {
      detail: data,
    });
  }
}

import type { IPlayer } from "../players/interface.js";

export interface IGameResultEvent {
  player: IPlayer | null;
  result:
    | typeof GameResultEvent.RESULT_DRAW
    | typeof GameResultEvent.RESULT_WIN;
}

export class GameResultEvent extends CustomEvent<IGameResultEvent> {
  public static readonly EVENT_NAME = "result";
  public static readonly RESULT_DRAW = 0;
  public static readonly RESULT_WIN = 1;

  public constructor(data: IGameResultEvent) {
    super(GameResultEvent.EVENT_NAME, {
      detail: data,
    });
  }
}

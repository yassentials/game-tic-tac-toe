import type { IPlayer, IPlayerCharacter } from "../../players/interface.js";

export interface IGameMode extends EventTarget {
  title: string;
  registerSelf(player: IPlayer): Promise<void>;
  getPositions(): Promise<(IPlayerCharacter | undefined)[]>;
}

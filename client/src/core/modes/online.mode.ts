import type { IPlayer, IPlayerCharacter } from "../players/interface.js";
import type { IGameMode } from "./common/interfaces.js";

/**
 * There should be three options: join room with code, host room, and quick play randomly choose room
 */
export default class OnlineMode extends EventTarget implements IGameMode {
  public readonly title: string = "Online Mode";
  public async getPositions(): Promise<(IPlayerCharacter | undefined)[]> {
    return [undefined];
  }

  public async registerSelf(_player: IPlayer): Promise<void> {}
}

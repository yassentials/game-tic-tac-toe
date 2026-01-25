import type { IPlayer } from "../players/interface";

export interface IRoom extends EventTarget {
  addPlayer(player: IPlayer): void;
  getPlayers(): Promise<IPlayer[]>;
  getPlayerCount(): Promise<number>;
}

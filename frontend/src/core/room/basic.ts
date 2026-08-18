import { RoomFullEvent } from "../events/room.js";
import type { IPlayer } from "../players/interface.js";
import type { IRoom } from "./interface.js";

export default class BasicRoom extends EventTarget implements IRoom {
  private room;
  private MAX_ROOM_PLAYER = 2;

  public constructor(players: Iterable<IPlayer> = []) {
    super();
    this.room = Array.from<IPlayer>(players);
  }

  public getPlayers() {
    return Promise.resolve(this.room);
  }

  public addPlayer(player: IPlayer): void {
    if (this.room.length >= this.MAX_ROOM_PLAYER) {
      this.dispatchEvent(new RoomFullEvent());

      return;
    }

    this.room.push(player);

    if (this.room.length >= this.MAX_ROOM_PLAYER) {
      this.dispatchEvent(new RoomFullEvent());
    }
  }

  public getPlayerCount(): Promise<number> {
    return Promise.resolve(this.room.length);
  }
}

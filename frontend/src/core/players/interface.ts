import type setting from "../../config/settings";

export type IPlayerCharacter = (typeof setting.PLAYER_CHARACTERS)[number];

export interface IPlayer extends EventTarget {
  name: Promise<string>;
  set character(c: IPlayerCharacter);
  get character(): Promise<IPlayerCharacter>;
}

export type ICapability = (instance: IPlayer) => {
  takePosition(index: number): Promise<void>;
};

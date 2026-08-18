import type { IPlayer, IPlayerCharacter } from "./interface.js";

export default abstract class BasePlayer
  extends EventTarget
  implements IPlayer
{
  private __character: IPlayerCharacter | undefined;

  public constructor(private readonly __name: string) {
    super();
  }

  public get name(): Promise<string> {
    return Promise.resolve(this.__name);
  }

  public set character(char: IPlayerCharacter) {
    this.__character = char;
  }

  public get character(): Promise<IPlayerCharacter> {
    if (!this.__character) {
      throw new Error("character undefined, please set the character");
    }

    return Promise.resolve(this.__character as IPlayerCharacter);
  }
}

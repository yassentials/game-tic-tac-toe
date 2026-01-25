import { getRandomIntBetween, shuffleArray } from "../../common/utils.js";
import { NotifyTurnEvent } from "../events/player.js";
import BasePlayer from "./base.js";
import type { ICapability, IPlayer } from "./interface.js";

export default class Bot extends BasePlayer implements IPlayer {
  public constructor(
    name: string,
    private readonly capability: ICapability,
  ) {
    super(name);

    this.addEventListener(NotifyTurnEvent.EVENT_NAME, async (event: Event) => {
      const evt = event as NotifyTurnEvent;
      const { availablePositions } = evt.detail;

      this.simulateThinking(async () => {
        const index = shuffleArray(availablePositions)[0] as number;
        await this.capability(this).takePosition(index);
      });
    });
  }

  private async simulateThinking<T>(
    callback: () => T | Promise<T>,
    thinkingAverageMilliseconds: number = 1000,
  ): Promise<T> {
    const offset = thinkingAverageMilliseconds / 5;
    const minDelay = thinkingAverageMilliseconds - offset;
    const maxDelay = thinkingAverageMilliseconds + offset;
    const delay = getRandomIntBetween(minDelay, maxDelay);

    return new Promise((resolve) => {
      setTimeout(async () => {
        const result = await callback();
        resolve(result);
      }, delay);
    });
  }
}

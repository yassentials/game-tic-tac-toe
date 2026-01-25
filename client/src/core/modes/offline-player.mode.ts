import { getRandomIntBetween } from "../../common/utils.js";
import settings from "../../config/settings.js";
import system from "../../config/system.js";
import { GameResultEvent } from "../events/game.js";
import { NotifyTurnEvent } from "../events/player.js";
import { RoomFullEvent } from "../events/room.js";
import {
  UIActionPositionTakeEvent,
  UIUpdatePositionTakeEvent,
} from "../events/ui.js";
import type { IPlayer, IPlayerCharacter } from "../players/interface.js";
import Player from "../players/offline-player.js";
import BasicRoom from "../room/basic.js";
import type { IRoom } from "../room/interface.js";
import { GameOverStatus } from "./common/enums.js";
import type { IGameMode } from "./common/interfaces.js";

export default class OfflinePlayerGameMode
  extends EventTarget
  implements IGameMode
{
  public readonly title: string = "2 Player Mode";
  private readonly room: IRoom;
  private readonly positions: (IPlayerCharacter | undefined)[] = [];
  private turn: number = 0;
  private ready = false;
  private over = false;
  private gameOverStatus:
    | (typeof GameOverStatus)[keyof typeof GameOverStatus]
    | undefined;

  public constructor() {
    super();
    this.room = new BasicRoom();

    this.positions = Array.from(
      { length: settings.TILE_SIZE ** 2 },
      () => undefined,
    );

    /**
     * Register events after initialization
     */
    this.registerEvents();
  }

  private registerEvents() {
    this.room.addEventListener(RoomFullEvent.EVENT_NAME, () => {
      this.ready = true;
      this.start();
    });
    this.addEventListener(
      UIActionPositionTakeEvent.EVENT_NAME,
      async (e: Event) => {
        const evt = e as UIActionPositionTakeEvent;
        const { index, successfulCallback } = evt.detail;
        const player = await this.getCurrentPlayer();

        if (!(player instanceof Player)) {
          return;
        }

        const ok = await this.takePosition(index, await player.character);

        if (!ok) {
          return;
        }
        await successfulCallback();

        await this.acquireTurn();

        if (this.over) {
          return;
        }

        const nextPlayer = await this.getCurrentPlayer();
        nextPlayer.dispatchEvent(
          new NotifyTurnEvent({
            availablePositions: await this.getAvailablePositions(),
          }),
        );
      },
    );
  }

  private async shuffleTurn() {
    this.turn = getRandomIntBetween(0, (await this.room.getPlayerCount()) - 1);
  }

  private async assignEachPlayerRandomCharacter() {
    const availableChars = [...settings.PLAYER_CHARACTERS];

    for (const player of await this.room.getPlayers()) {
      const index = getRandomIntBetween(0, availableChars.length - 1);
      const char = availableChars.splice(index, 1)[0];

      if (typeof char === "undefined") {
        throw new TypeError("char is undefined");
      }

      player.character = char;
    }
  }

  public async getPositions() {
    return this.positions;
  }

  private async getAvailablePositions(): Promise<number[]> {
    const positions = await this.getPositions();
    const availablePos = (
      await Promise.all(
        positions.map(async (character, index) => {
          return typeof character === "undefined" ? index : null;
        }),
      )
    ).filter((v) => v !== null);

    return availablePos;
  }

  public registerSelf(player: IPlayer): Promise<void> {
    return Promise.resolve(this.room.addPlayer(player));
  }

  private async start(): Promise<void> {
    await this.shuffleTurn();
    await this.assignEachPlayerRandomCharacter();

    const player = await this.getCurrentPlayer();

    player.dispatchEvent(
      new NotifyTurnEvent({
        availablePositions: await this.getAvailablePositions(),
      }),
    );
  }

  private async getCurrentPlayer(): Promise<IPlayer> {
    return (await this.room.getPlayers())[this.turn] as IPlayer;
  }

  private async takePosition(
    position: number,
    character: IPlayerCharacter,
  ): Promise<boolean> {
    if (!this.ready || this.over) return false;

    if (typeof this.positions[position] === "undefined") {
      this.positions[position] = character;
      this.dispatchEvent(
        new UIUpdatePositionTakeEvent({
          index: position,
          character,
        }),
      );

      return true;
    }

    return false;
  }

  private async acquireTurn(): Promise<void> {
    const status = await this.getGameOverStatus();

    switch (status) {
      case GameOverStatus.DRAW: {
        this.over = true;
        this.dispatchEvent(
          new GameResultEvent({
            result: GameResultEvent.RESULT_DRAW,
            player: null,
          }),
        );
        return;
      }
      case GameOverStatus.WIN: {
        this.over = true;
        this.dispatchEvent(
          new GameResultEvent({
            result: GameResultEvent.RESULT_WIN,
            player: await this.getCurrentPlayer(),
          }),
        );
        return;
      }
      default:
        break;
    }

    this.turn = await this.getNextTurn();
  }

  private async getGameOverStatus(): Promise<
    (typeof GameOverStatus)[keyof typeof GameOverStatus]
  > {
    if (typeof this.gameOverStatus !== "undefined") {
      return this.gameOverStatus;
    }

    const calculateResult = async () => {
      const positions = await this.getPositions();

      const size = settings.TILE_SIZE;
      const winPositionOffset = system.WIN_POSITION_OFFSET;

      for (let row = 0; row < size - 2; row++) {
        for (let col = 0; col < size - 2; col++) {
          const position = [
            positions[row * size + col],
            positions[row * size + col + 1],
            positions[row * size + col + 2],
            positions[(row + 1) * size + col],
            positions[(row + 1) * size + col + 1],
            positions[(row + 1) * size + col + 2],
            positions[(row + 2) * size + col],
            positions[(row + 2) * size + col + 1],
            positions[(row + 2) * size + col + 2],
          ];

          for (let i = 0; i < system.WIN_POSITIONS.length; i++) {
            const winPos = system.WIN_POSITIONS[i] as [number, number, number];
            const blockOne = position[winPos[0] + winPositionOffset];
            const blockTwo = position[winPos[1] + winPositionOffset];
            const blockThree = position[winPos[2] + winPositionOffset];

            if (
              [blockOne, blockTwo, blockThree].some(
                (block) => typeof block === "undefined",
              )
            ) {
              continue;
            }

            if (blockOne === blockTwo && blockTwo === blockThree) {
              return GameOverStatus.WIN;
            }
          }
        }
      }

      if (positions.every((pos) => typeof pos !== "undefined")) {
        return GameOverStatus.DRAW;
      }

      return GameOverStatus.NONE;
    };

    const result = await calculateResult();

    if (result === GameOverStatus.NONE) {
      return result;
    }

    this.gameOverStatus = result;

    return result;
  }

  private async getNextTurn(): Promise<number> {
    const turn = this.turn + 1;

    return turn % (await this.room.getPlayerCount());
  }
}

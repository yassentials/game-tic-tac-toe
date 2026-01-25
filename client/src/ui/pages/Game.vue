<script setup lang="ts">
import { onMounted, ref, shallowRef } from "vue";
import { useRoute, useRouter } from "vue-router";
import { GameResultEvent } from "../../core/events/game";
import { NotifyTurnEvent } from "../../core/events/player";
import {
  UIActionPositionTakeEvent,
  UIUpdatePositionTakeEvent,
} from "../../core/events/ui";
import { Mode } from "../../core/modes/common/enums";
import type { IGameMode } from "../../core/modes/common/interfaces";
import BotGameMode from "../../core/modes/offline-bot.mode";
import OfflinePlayerGameMode from "../../core/modes/offline-player.mode";
import type { IPlayerCharacter } from "../../core/players/interface";
import Player from "../../core/players/offline-player";

const route = useRoute();
const router = useRouter();
const name = "yassentials";
const isOver = ref<boolean>(false);
const infoText = ref<string>("");

const positions = ref<(IPlayerCharacter | undefined)[]>([]);
const gameRef = shallowRef<IGameMode | null>(null);

function handleClick(index: number) {
  gameRef.value?.dispatchEvent(
    new UIActionPositionTakeEvent({
      index,
      async successfulCallback() {
        infoText.value = "Their turn.";
      },
    }),
  );
}

function goToMenu() {
  router.push({
    name: "menu",
  });
}

async function ConstructGameMode(mode: Mode): Promise<IGameMode> {
  let game: IGameMode | undefined;
  switch (mode) {
    case Mode.OFFLINE_BOT: {
      const player = new Player(name);

      player.addEventListener(NotifyTurnEvent.EVENT_NAME, () => {
        infoText.value = "Your Turn!";
      });

      game = new BotGameMode();

      await game.registerSelf(player);

      break;
    }
    case Mode.OFFLINE_PLAYER: {
      const p1 = new Player("P1");
      const p2 = new Player("P2");

      game = new OfflinePlayerGameMode();

      for (const p of [p1, p2]) {
        p.addEventListener(NotifyTurnEvent.EVENT_NAME, async () => {
          infoText.value = `${await p.character} Turn`;
        });

        await game.registerSelf(p);
      }

      break;
    }
    default:
      break;
  }

  if (typeof game === "undefined") {
    throw new Error("Failed to construct game instance");
  }

  return game;
}

function resetStateToDefault() {
  isOver.value = false;
  infoText.value = "";

  positions.value = [];
  gameRef.value = null;
}

async function StartNewGame() {
  resetStateToDefault();

  const game = await ConstructGameMode(Number(route.params.mode) as Mode);

  game.addEventListener(UIUpdatePositionTakeEvent.EVENT_NAME, (e: Event) => {
    const evt = e as UIUpdatePositionTakeEvent;
    const { character, index } = evt.detail;
    positions.value[index] = character;
  });

  game.addEventListener(GameResultEvent.EVENT_NAME, async (e: Event) => {
    const evt = e as GameResultEvent;
    const { result, player: p } = evt.detail;

    switch (result) {
      case GameResultEvent.RESULT_DRAW:
        infoText.value = "Draw";
        break;
      case GameResultEvent.RESULT_WIN:
        if (game instanceof BotGameMode) {
          if (p instanceof Player) {
            infoText.value = "You Win!";
            break;
          }
          infoText.value = "You Lose!";
          break;
        }

        if (game instanceof OfflinePlayerGameMode) {
          infoText.value = `${await p?.character} Win!`;
          break;
        }

        break;
      default:
        break;
    }

    setTimeout(() => {
      isOver.value = true;
    }, 3 * 1000);
  });

  gameRef.value = game;
  positions.value = [...(await game.getPositions())];
}

onMounted(async () => {
  await StartNewGame();
});
</script>

<template>
    <section class="min-h-screen flex items-center bg-slate-800">
        <div class="mx-auto max-w-7xl w-100 px-6">
            <div class="text-center mb-6 text-xl font-bold">
                <h1>{{ gameRef?.title }}</h1>
            </div>
            <div
                class="grid grid-cols-3 gap-4 rounded-2xl overflow-hidden text-slate-900 p-3 border-4 border-slate-300">
                <button
                    class="aspect-square rounded-lg bg-slate-100 hover:bg-slate-300 text-2xl font-bold focus:outline-slate-500 focus:outline-6"
                    type="button" v-for="(char, i) in positions" @click="() => handleClick(i)">{{ char }}</button>
            </div>
        </div>

        <div v-if="infoText !== ''" class="fixed inset-x-0 top-0 mt-3 me-3 flex justify-end">
            <div class="bg-slate-200/80  rounded-full backdrop-blur-md px-6 py-2">
                <p class="text-slate-800 font-medium">{{ infoText }}</p>
            </div>
        </div>

        <div v-if="isOver" class="absolute inset-0 bg-slate-800/80 flex justify-center items-center p-4">
            <div class="bg-white shadow-xl w-full max-w-[350px] text-slate-800 text-xl rounded-lg p-6">
                <p class="text-2xl font-bold mb-6 bg-slate-600 text-slate-200 text-center rounded-md">{{ infoText }}</p>
                <p class="text-xl font-semibold mb-6">Do you want to continue?</p>
                <div class="ml-auto flex justify-end gap-x-3">
                    <button class="bg-gray-200 hover:bg-gray-300 px-4 rounded-lg" @click="goToMenu">Main Menu</button>
                    <button @click="StartNewGame"
                        class="bg-blue-600 hover:bg-blue-700 px-4 rounded-lg text-white">Yes</button>
                </div>
            </div>
        </div>
    </section>
</template>
<script setup lang="ts">
import { useRouter } from "vue-router";
import { Mode } from "../../core/modes/common/enums";

const router = useRouter();

const modes = [
  {
    name: "Bot",
    action: () => {
      router.push({
        name: "game",
        params: {
          mode: Mode.OFFLINE_BOT,
        },
      });
    },
  },
  {
    name: "2 Player",
    action: () => {
      router.push({
        name: "game",
        params: {
          mode: Mode.OFFLINE_PLAYER,
        },
      });
    },
  },
  {
    name: "Online",
    disable: true,
  },
];
</script>

<template>
  <section class="min-h-screen flex items-center bg-slate-600">
    <div class="px-4 mx-auto max-w-7xl w-full">
      <div class="flex flex-col-reverse md:flex-row justify-between items-center">
        <div class="flex flex-col gap-y-4 min-w-44 w-full max-w-84">
          <h1 class="text-2xl md:text-4xl font-black mb-6">Select Mode</h1>
          <button v-for="mode in modes" @click="mode.action"
            class="bg-slate-700 hover:bg-slate-400 hover:text-slate-700 py-2 rounded-xl">{{ mode.name }}</button>
        </div>

        <div class="hidden md:block shrink-0">
          <img class="rounded-2xl shadow-2xl border-4 border-slate-400" src="https://picsum.photos/300/300.jpg" alt="">
        </div>
      </div>
    </div>
  </section>
</template>
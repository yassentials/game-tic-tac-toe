<script setup lang="ts">
import { RouterLink } from "vue-router";
import NavigationBar from "../components/NavigationBar.vue";
</script>

<template>
    <NavigationBar />
    <section class="bg-slate-500 text-white min-h-screen flex items-center">
        <div class="mx-auto px-4 pb-6 sm:px-6 w-full max-w-7xl">
            <div class="flex flex-col-reverse md:flex-row gap-12 justify-between items-center">
                <div class="flex flex-col items-center md:items-start text-center md:text-start">
                    <h1 class="text-6xl my-2 font-bold">Tic Tac Toe</h1>
                    <p class="mb-8 opacity-90">Simple game made by me</p>
                    <RouterLink to="/menu" class="bg-slate-600 hover:bg-slate-700 transition-colors px-6 py-4 rounded-xl border-2">I'm ready to start!</RouterLink>
                </div>
                <div class="shrink-0">
                    <img 
                    class="rounded-2xl shadow-2xl border-4 border-slate-400"
                    src="https://picsum.photos/300/300.jpg" alt="">
                </div>
            </div>
        </div>
    </section>

    <section id="about" class="min-h-[50vh] bg-slate-300 text-slate-600 flex items-center">
        <div class="mx-auto px-4 text-center">
            <h1 class="text-4xl font-bold mb-8">About</h1>
            <p class="leading-relaxed text-lg">
                This is just a personal project, I wanted to learn about Typescript and Frontend Development. Using various technologies such as Vue.js, Typescript, Tailwind, and many others.
            </p>
        </div>
    </section>
</template>
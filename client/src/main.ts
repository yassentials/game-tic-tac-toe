import { createApp } from "vue";
import "./ui/style.css";
import { createRouter, createWebHistory } from "vue-router";
import App from "./ui/App.vue";
import Game from "./ui/pages/Game.vue";
import Landing from "./ui/pages/Landing.vue";
import MainMenu from "./ui/pages/MainMenu.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "root",
      component: Landing,
    },
    {
      path: "/menu",
      name: "menu",
      component: MainMenu,
    },
    {
      path: "/game/:mode",
      name: "game",
      component: Game,
    },
  ],
});

createApp(App).use(router).mount("#app");

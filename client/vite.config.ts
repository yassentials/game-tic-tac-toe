import tailwindcss from "@tailwindcss/vite";
import vue from "@vitejs/plugin-vue";
import { defineConfig, type ConfigEnv } from "vite";

// https://vite.dev/config/
export default defineConfig((env: ConfigEnv) => ({
  plugins: [vue(), tailwindcss()],
  base: env.mode === "github" ? "/game-tic-tac-toe/" : "/",
}));

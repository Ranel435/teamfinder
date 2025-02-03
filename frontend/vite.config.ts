import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig(({ mode }) => ({
  plugins: [sveltekit()],
  server: {
    port: 3000,
    proxy: mode === 'development'
      ? { '/api': 'http://localhost:8090' }
      : undefined
  }
}));

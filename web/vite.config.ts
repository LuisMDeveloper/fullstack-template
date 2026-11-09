import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
  ],
  server: {
    // Forward API calls to the Go server during `npm run dev`
    proxy: { '/api': 'http://localhost:8080' },
  },
  build: {
    // The Go binary embeds this directory (see api/internal/webui)
    outDir: fileURLToPath(new URL('../api/internal/webui/dist', import.meta.url)),
    emptyOutDir: true,
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})

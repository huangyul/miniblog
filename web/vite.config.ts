import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'
import UnoCSS from 'unocss/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [UnoCSS(), react()],
  server: {
    port: 7777,
  },
  resolve: {
    alias: {
      "@": path.join(__dirname, "src")
    }
  }
})

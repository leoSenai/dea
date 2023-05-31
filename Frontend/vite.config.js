import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'
import eslint from 'vite-plugin-eslint'

export default defineConfig({
  plugins: [
    vue({
      template: { transformAssetUrls }
    }),
    eslint(),
    quasar()
  ]
})

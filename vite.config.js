import { defineConfig } from 'vite'
import vuePlugin from '@vitejs/plugin-vue'
import path from 'node:path'

export default defineConfig({
    plugins: [vuePlugin()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './web'),
        },
    },
    build: {
        manifest: true,
        rollupOptions: {
            input: '/web/main.js',
        },
    }
})

import { defineConfig, loadEnv } from 'vite'
import vuePlugin from '@vitejs/plugin-vue'
import path from 'node:path'

export default defineConfig(({ command, mode }) => {
    const env = loadEnv(mode, process.cwd(), '')

    return {
        plugins: [vuePlugin()],
        server: {
            host: env.VITE_HOST,
            port: env.VITE_PORT,
        },
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
    }
})
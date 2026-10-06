import { fileURLToPath } from 'node:url';
import { defineConfig } from 'vite';

export default defineConfig({
  base: '/go-gpui/',
  build: {
    outDir: fileURLToPath(new URL('../docs', import.meta.url)),
    emptyOutDir: true,
  },
});

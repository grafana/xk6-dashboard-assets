import { defineConfig } from "vite"
import react from "@vitejs/plugin-react-swc"
import svgr from "vite-plugin-svgr"
import { visualizer } from "rollup-plugin-visualizer"
import tsconfigPaths from "vite-tsconfig-paths"
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin"

export default defineConfig({
  plugins: [
    svgr({
      include: "**/*.svg",
      svgrOptions: {
        exportType: "named",
        namedExport: "ReactComponent"
      }
    }),
    react(),
    tsconfigPaths(),
    vanillaExtractPlugin(),
    visualizer({
      filename: "bundle-stats.html"
    })
  ],
  build: { chunkSizeWarningLimit: 512 /*, outDir: '../../ui'*/ },
  base: "",
  server: {
    watch: {
      ignored: ["**/node_modules/**", "**/packages/*/node_modules/**", "**/dist/**", "**/vendor/**"]
    }
  }
})

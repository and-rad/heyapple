import { fileURLToPath, URL } from "url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const path = require("path");

// https://vitejs.dev/config/
export default defineConfig({
	build: {
		outDir: "../../../internal/web/static",
		rollupOptions: {
			output: {
				entryFileNames: `app/[name].js`,
				chunkFileNames: `app/[name].js`,
				assetFileNames: `app/[name].[ext]`,
			},
		},
	},
	plugins: [vue()],
	resolve: {
		alias: {
			"@": fileURLToPath(new URL("./src", import.meta.url)),
		},
	},
});

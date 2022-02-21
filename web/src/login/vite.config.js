import { fileURLToPath, URL } from "url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const path = require("path");

// https://vitejs.dev/config/
export default defineConfig({
	build: {
		outDir: "../../",
		rollupOptions: {
			output: {
				entryFileNames: `static/login/[name].js`,
				chunkFileNames: `static/login/[name].js`,
				assetFileNames: `static/login/[name].[ext]`,
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

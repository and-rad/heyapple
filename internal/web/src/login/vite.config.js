import { fileURLToPath, URL } from "url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const path = require("path");

// https://vitejs.dev/config/
export default defineConfig({
	build: {
		outDir: "../../static",
		rollupOptions: {
			output: {
				entryFileNames: `login/[name].js`,
				chunkFileNames: `login/[name].js`,
				assetFileNames: `login/[name].[ext]`,
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

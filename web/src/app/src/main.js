import { createApp } from "vue";
import { createI18n } from "vue-i18n";
import App from "./App.vue";
import router from "./router";

let lang = document.documentElement.lang || navigator.language;
if (!lang && navigator.languages != undefined) {
	lang = navigator.languages[0];
}

const csrfMeta = document.querySelector("meta[name='_csrf']");
const csrfToken = csrfMeta ? csrfMeta.content : "";

fetch("/app/l10n.json")
	.then((response) => response.json())
	.then((messages) => {
		const i18n = createI18n({
			locale: lang.split("-")[0],
			fallbackLocale: "en",
			messages,
		});

		const app = createApp(App);
		app.provide("csrfToken", csrfToken);
		app.use(router);
		app.use(i18n);

		app.mount("#app");
	});

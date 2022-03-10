import { createApp } from "vue";
import { createI18n } from "vue-i18n";
import App from "./App.vue";
import router from "./router";
import Message from "./components/Message.vue";
import HeaderImage from "./components/images/ImageHeader.vue";

const csrfMeta = document.querySelector("meta[name='_csrf']");
const csrfToken = csrfMeta ? csrfMeta.content : "";

fetch("/login/l10n.json")
	.then((response) => response.json())
	.then((messages) => {
		let lang = navigator.language;
		if (navigator.languages != undefined) {
			lang = navigator.languages[0];
		}

		const i18n = createI18n({
			legacy: false,
			locale: lang.split("-")[0],
			fallbackLocale: "en",
			messages,
		});

		const app = createApp(App);
		app.component("Message", Message);
		app.component("HeaderImage", HeaderImage);
		app.provide("csrfToken", csrfToken);
		app.use(router);
		app.use(i18n);

		app.mount("#app");
	});

import { createApp } from "vue";
import { createI18n } from "vue-i18n";
import App from "./App.vue";
import router from "./router";
import Message from "./components/Message.vue";

fetch("/login/l10n.json")
	.then((response) => response.json())
	.then((messages) => {
		const i18n = createI18n({
			locale: "de",
			fallbackLocale: "en",
			messages,
		});

		const app = createApp(App);
		app.component("Message", Message);
		app.use(router);
		app.use(i18n);

		app.mount("#app");
	});

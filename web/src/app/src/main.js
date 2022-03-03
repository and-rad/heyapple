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

const perms = (function () {
	const _current = document.documentElement.dataset.perm || 1;
	const _createFood = 0x00010000;
	const _editFood = 0x00020000;

	function _check(perm) {
		return (_current & perm) == perm;
	}

	return {
		canCreateFood: _check(_createFood),
		canEditFood: _check(_editFood),
	};
})();

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
		app.provide("perms", perms);
		app.use(router);
		app.use(i18n);

		app.mount("#app");
	});

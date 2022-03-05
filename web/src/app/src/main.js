import { createApp } from "vue";
import { createI18n } from "vue-i18n";
import App from "./App.vue";
import router from "./router";

// preferred UI language
let lang = document.documentElement.lang || navigator.language;
if (!lang && navigator.languages != undefined) {
	lang = navigator.languages[0];
}

// CSRF protection
const csrfMeta = document.querySelector("meta[name='_csrf']");
const csrfToken = csrfMeta ? csrfMeta.content : "";

// local user permissions
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

// remote data
const app = createApp(App);
let locale = undefined;
let food = undefined;

function initLocale(messages) {
	locale = createI18n({
		locale: lang.split("-")[0],
		fallbackLocale: "en",
		messages,
	});

	if (food) {
		mountApp();
	}
}

function initFoods(data) {
	food = data;
	if (locale) {
		mountApp();
	}
}

function mountApp() {
	food.forEach((f) => {
		f.name = locale.global.t(f.id.toString());
	});

	app.provide("csrfToken", csrfToken);
	app.provide("perms", perms);
	app.provide("food", food);
	app.use(router);
	app.use(locale);
	app.mount("#app");
}

fetch("/app/l10n.json")
	.then((response) => response.json())
	.then(initLocale);

fetch("api/v1/foods")
	.then((response) => response.json())
	.then(initFoods);

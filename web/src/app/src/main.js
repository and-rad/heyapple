import { createApp, ref } from "vue";
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

// message bus
const log = (function () {
	function _getMessage(obj) {
		if (typeof obj === "string") {
			return obj;
		}
		if ("message" in obj) {
			return obj.message;
		}
		return locale.global.t("err.err");
	}

	return {
		msg: function (obj) {
			let payload = { msg: _getMessage(obj), timeout: 3000 };
			window.dispatchEvent(new CustomEvent("message", { detail: payload }));
		},
		warn: function (obj) {
			let payload = { msg: _getMessage(obj), timeout: 4000 };
			window.dispatchEvent(new CustomEvent("warning", { detail: payload }));
		},
		err: function (obj) {
			let payload = { msg: _getMessage(obj), timeout: 5000 };
			window.dispatchEvent(new CustomEvent("error", { detail: payload }));
		},
	};
})();

// remote data
const app = createApp(App);
let locale = undefined;
let food = undefined;
let recipes = undefined;

function initLocale(messages) {
	locale = createI18n({
		locale: lang.split("-")[0],
		fallbackLocale: "en",
		messages,
	});

	if (food && recipes) {
		mountApp();
	}
}

function initFoods(data) {
	food = data;
	if (locale && recipes) {
		mountApp();
	}
}

function initRecipes(data) {
	recipes = data;
	console.log(recipes);
	if (locale && food) {
		mountApp();
	}
}

function mountApp() {
	food.forEach((f) => {
		f.name = locale.global.t(f.id.toString());
	});

	recipes.forEach((r) => {
		r.isOwner = true;
	});

	app.provide("csrfToken", csrfToken);
	app.provide("perms", perms);
	app.provide("food", ref(food));
	app.provide("recipes", ref(recipes));
	app.provide("log", log);
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

fetch("api/v1/recipes")
	.then((response) => response.json())
	.then(initRecipes);

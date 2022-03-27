import { createApp, ref } from "vue";
import { createI18n } from "vue-i18n";
import { Settings } from "luxon";
import App from "./App.vue";
import router from "./router";

// preferred UI language
let lang = document.documentElement.lang || navigator.language;
if (!lang && navigator.languages != undefined) {
	lang = navigator.languages[0];
}
Settings.defaultLocale = lang;

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
let diary = undefined;
let prefs = undefined;

function initLocale(messages) {
	locale = createI18n({
		legacy: false,
		locale: lang,
		fallbackLocale: "en",
		messages,
	});

	if (food && recipes && diary && prefs) {
		mountApp();
	}
}

function initFoods(data) {
	food = data;
	if (locale && recipes && diary && prefs) {
		mountApp();
	}
}

function initRecipes(data) {
	recipes = data;
	if (locale && food && diary && prefs) {
		mountApp();
	}
}

function initDiary(data) {
	diary = {};
	data.forEach((d) => (diary[d.date] = d));
	if (locale && food && recipes && prefs) {
		mountApp();
	}
}

function initPrefs(data) {
	prefs = data;
	if (locale && food && recipes && diary) {
		mountApp();
	}
}

function mountApp() {
	food.forEach((f) => {
		f.name = locale.global.t(f.id.toString());
	});

	app.provide("csrfToken", csrfToken);
	app.provide("perms", perms);
	app.provide("food", ref(food));
	app.provide("recipes", ref(recipes));
	app.provide("diary", ref(diary));
	app.provide("prefs", ref(prefs));
	app.provide("log", log);
	app.use(router);
	app.use(locale);
	app.mount("#app");
}

fetch(`/app/l10n/${lang}.json`)
	.then((response) => response.json())
	.then(initLocale)
	.catch(() => {
		fetch("/app/l10n/en.json")
			.then((response) => response.json())
			.then(initLocale);
	});

fetch("/api/v1/foods")
	.then((response) => response.json())
	.then(initFoods);

fetch("/api/v1/recipes")
	.then((response) => response.json())
	.then(initRecipes);

fetch("/api/v1/diary")
	.then((response) => response.json())
	.then(initDiary);

// TODO this needs an actual implementation
fetch("/api/v1/prefs")
	.then((response) => response.json())
	.then(initPrefs)
	.catch(() => {
		initPrefs({
			rdi: {
				kcal: 2000,
				fat: 60,
				carb: 270,
				prot: 80,
			},
			ui: {
				neutralCharts: false,
			},
		});
	});

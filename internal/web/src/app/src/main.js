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
		f.cat = locale.global.t("cat." + f.cat);
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
				fib: 32,
				salt: 5.8,
				fatsat: 22,
				fato3: 1.6,
				fato6: 3.2,
				vita: 0.9,
				vitb1: 1.2,
				vitb2: 1.3,
				vitb3: 16,
				vitb5: 5,
				vitb6: 1.7,
				vitb7: 0.03,
				vitb9: 0.4,
				vitb12: 0.003,
				vitc: 90,
				vitd: 0.015,
				vite: 15,
				vitk: 0.12,
				pot: 3400,
				chl: 2300,
				sod: 2300,
				calc: 1000,
				phos: 700,
				mag: 400,
				iron: 8,
				zinc: 11,
				mang: 2.3,
				cop: 0.9,
				iod: 0.15,
				chr: 0.035,
				mol: 0.045,
				sel: 0.055,
			},
			ui: {
				neutralCharts: false,
			},
		});
	});

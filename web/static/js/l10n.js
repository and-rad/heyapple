var HA = HA || {};

HA.L10n = (function () {
	var _locmap;

	return {
		init: function () {
			let lang = document.documentElement.lang || navigator.language;
			if (!lang && navigator.languages != undefined) {
				lang = navigator.languages[0];
			}

			fetch("/api/v1/l10n/" + lang)
				.then((response) => response.json())
				.then((data) => (_locmap = data));
		},

		t: function (key) {
			return _locmap[key] || key;
		},
	};
})();

HA.L10n.init();

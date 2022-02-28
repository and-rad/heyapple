var HA = HA || {};

HA.Auth = (function () {
	function _initSubmitButtons() {
		document
			.querySelectorAll("form.reset input[type='submit']")
			.forEach((s) => s.addEventListener("click", _onResetButtonClicked));
	}

	function _showError(msg) {
		let field = document.querySelector(".msg");
		field.textContent = msg;
		field.classList.remove("hidden");
		field.classList.add("err");
	}

	function _showMessage(msg) {
		let field = document.querySelector(".msg");
		field.textContent = msg;
		field.classList.remove("hidden", "err");
	}

	var _onResetButtonClicked = function (evt) {
		evt.preventDefault();
		let form = evt.target.closest("form");
		let formData = new FormData(form);
		if (!formData.getAll("pass").every((v, i, a) => v === a[0])) {
			_showError(HA.L10n.t("reset.nomatch"));
			return;
		}

		fetch("/auth/reset", {
			method: "PUT",
			headers: {
				"Content-Type": "application/x-www-form-urlencoded",
				"X-CSRF-Token": document.querySelector("meta[name='_csrf']").content,
			},
			body: new URLSearchParams(formData),
		}).then((response) => {
			let msg = HA.L10n.t("reset." + response.status);
			response.ok ? _showMessage(msg) : _showError(msg);
		});
	};

	return {
		initLocal: function () {
			_initSubmitButtons();
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	HA.Auth.initLocal();
});

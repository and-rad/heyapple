var HA = HA || {};

HA.Auth = (function () {
	var _scoreColor = ["#e53935", "#FB8C00", "#FDD835", "#7CB342", "#00897B"];

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

	function _onResetButtonClicked(evt) {
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
	}

	function _onPasswordInput(evt) {
		let pass = evt.target.value;
		let score = zxcvbn(pass).score;
		let bar = evt.target.parentNode.querySelector(".password-strength-bar");
		bar.style.width = pass.length > 0 ? 20 + score * 20 + "%" : "0%";
		bar.style.background = _scoreColor[score];
	}

	return {
		init: function () {
			document
				.querySelectorAll("form.reset input[type='submit']")
				.forEach((s) => s.addEventListener("click", _onResetButtonClicked));

			document.querySelectorAll(".password-field > input[type='password']").forEach((pw) => {
				pw.addEventListener("input", _onPasswordInput);
			});
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	HA.Auth.init();
});

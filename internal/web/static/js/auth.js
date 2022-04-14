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

		fetch("/auth/reset", {
			method: "PUT",
			headers: {
				"Content-Type": "application/x-www-form-urlencoded",
				"X-CSRF-Token": document.querySelector("meta[name='_csrf']").content,
			},
			body: new URLSearchParams(new FormData(form)),
		}).then((response) => {
			let msg = HA.L10n.t("reset." + response.status);
			response.ok ? _showMessage(msg) : _showError(msg);
		});
	}

	function _onPasswordInput(evt) {
		let pass = evt.target.value;
		let score = zxcvbn(pass).score;
		let bar = evt.target.parentNode.querySelector(".password .strength-bar");
		bar.style.width = pass.length > 0 ? 20 + score * 20 + "%" : "0%";
		bar.style.background = _scoreColor[score];
	}

	function _onPasswordVisibilityToggle(evt) {
		evt.target.closest(".password").classList.toggle("visible");
	}

	return {
		init: function () {
			document
				.querySelectorAll("form.reset input[type='submit']")
				.forEach((s) => s.addEventListener("click", _onResetButtonClicked));

			document.querySelectorAll(".password > input[type='password']").forEach((pw) => {
				pw.addEventListener("input", _onPasswordInput);
			});

			document.querySelectorAll(".password > button.visibility").forEach((btn) => {
				btn.addEventListener("click", _onPasswordVisibilityToggle);
			});
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	HA.Auth.init();
});

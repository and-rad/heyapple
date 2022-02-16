var HA = HA || {};

HA.Auth = (function () {
	var _initSubmitButtons = function () {
		document
			.querySelectorAll("form.auth[method='post'] input[type='submit']")
			.forEach((s) => s.addEventListener("click", _onLoginButtonClicked));
		document
			.querySelectorAll("form.auth[method='delete'] input[type='submit']")
			.forEach((s) => s.addEventListener("click", _onLogoutButtonClicked));
	};

	var _onLoginButtonClicked = function (evt) {
		evt.preventDefault();
		let form = new FormData(evt.target.closest("form"));
		fetch("/auth/local", {
			method: "POST",
			headers: { "Content-Type": "application/x-www-form-urlencoded" },
			body: new URLSearchParams(form),
		}).then((response) => {
			if (response.ok) {
				window.location.reload();
			} else {
				window.dispatchEvent(
					new CustomEvent("loginfail", { detail: { code: response.status } })
				);
			}
		});
	};

	var _onLogoutButtonClicked = function (evt) {
		evt.preventDefault();
		fetch("/auth/local", {
			method: "DELETE",
		}).then((response) => {
			if (response.ok) {
				window.location.reload();
			} else {
				window.dispatchEvent(
					new CustomEvent("logoutfail", { detail: { code: response.status } })
				);
			}
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

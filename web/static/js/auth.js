var HA = HA || {};

HA.Auth = (function () {
	var _initSubmitButtons = function () {
		document
			.querySelectorAll("form.auth input[type='submit']")
			.forEach((s) => s.addEventListener("click", _onSubmitButtonClicked));
	};

	var _onSubmitButtonClicked = function (evt) {
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

	return {
		initLocal: function () {
			_initSubmitButtons();
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	HA.Auth.initLocal();
});

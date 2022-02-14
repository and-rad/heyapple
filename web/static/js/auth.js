var HA = HA || {};

HA.Auth = (function () {
	var _initSubmitButtons = function () {
		document
			.querySelectorAll("form.auth input[type='submit']")
			.forEach((s) => {
				s.addEventListener("click", function (evt) {
					evt.preventDefault();
				});
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

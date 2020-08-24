if (!OCA.HeyApple) {
	OCA.HeyApple = {};
}

OCA.HeyApple.Core = (function(){
	var _data = {};

	return {
		init: function() {
			OCA.HeyApple.Backend.getLists(function(obj) {
				_data = obj.data;
				OCA.HeyApple.UI.init();
			});
		},

		listNames: function() {
			return Object.keys(_data);
		}
	};
})();

OCA.HeyApple.UI = (function(){
	var _refreshLists = function() {
		let frag = document.createDocumentFragment();

		OCA.HeyApple.Core.listNames().forEach(function(name) {
			let item = document.createElement("li");
			item.textContent = name;
			item.dataset.name = name;
			item.addEventListener("click", function(evt) {
				_showList(evt.target);
			});

			frag.appendChild(item);
		});

		let list = document.querySelector("#list-category");
		list.textContent = "";
		list.appendChild(frag);

		if (list.firstElementChild) {
			_showList(list.firstElementChild);
		}
	};

	var _showList = function(elem) {
		for (let i = 0, e; e = elem.parentNode.children[i]; i++) {
			e == elem ? e.classList.add("active") : e.classList.remove("active");
		}
		console.log(elem.dataset.name);
	};

	return {
		init: function() {
			let picker = new Pikaday({
				field: document.getElementById("calendar"),
				bound: false,
				container: document.getElementById("calendar"),
				showDaysInNextAndPreviousMonths: true,
				firstDay: 1,
				i18n: {
					previousMonth : "←",
					nextMonth     : "→",
					months        : ["January","February","March","April","May","June","July","August","September","October","November","December"],
					weekdays      : ["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"],
					weekdaysShort : ["Su","Mo","Tu","We","Th","Fr","Sa"]
				}
			});

			_refreshLists();
		},
	};
})();

OCA.HeyApple.Backend = (function() {
	return {
		get: function(uri, callback) {
			let xhr = new XMLHttpRequest();
			xhr.addEventListener("load", callback);
			xhr.open("GET", uri);
			xhr.setRequestHeader("requesttoken", OC.requestToken);
			xhr.send();
		},

		post: function(uri, data, callback) {
			let xhr = new XMLHttpRequest();
			xhr.addEventListener("load", callback);
			xhr.open("POST", uri);
			xhr.setRequestHeader("requesttoken", OC.requestToken);
			xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
			xhr.send(data);
		},

		getLists: function(callback) {
			this.get(OC.generateUrl("apps/heyapple/api/0.1/lists"), function() {
				callback(JSON.parse(this.response));
			});
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	OCA.HeyApple.Core.init();
});
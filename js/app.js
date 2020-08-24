if (!OCA.HeyApple) {
	OCA.HeyApple = {};
}

OCA.HeyApple.Core = (function(){
	var _data = {};

	var _add = function(amount, name) {
		return amount + 1;
	}

	var _amount = function(name) {
		return 1;
	}

	var _selected = function(listId, itemId) {
		return false;
	}

	var _trim = function(name) {
		return name;
	}

	return {
		init: function() {
			OCA.HeyApple.Backend.getLists(function(obj) {
				_data = obj.data;
				OCA.HeyApple.UI.init();
			});
		},

		listNames: function() {
			return Object.keys(_data);
		},

		list: function(name) {
			let out = {};

			if (_data[name]) {
				_data[name].forEach(function(elem) {
					let name = elem[1];
					let id = elem[2];
					let amount = _amount(name);

					if (out[id]) {
						out[id].amount = _add(out[id].amount, amount);
					} else {
						out[id] = {
							id: id,
							selected: _selected(name, id),
							name: _trim(name),
							amount: amount,
							aisle: "None"
						};
					}
				});
			}

			return Object.values(out);
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

		let frag = document.createDocumentFragment();
		let tmpl = document.createElement("tr");
		tmpl.innerHTML = document.querySelector("#template-item").innerHTML;

		let list = OCA.HeyApple.Core.list(elem.dataset.name);
		for (let i = 0, item; item = list[i]; i++) {
			let row = tmpl.cloneNode(true);
			row.dataset.id = item.id;
			row.addEventListener("click", _onItemClicked);
			if (item.selected) {
				row.classList.add("selected");
			}

			let fields = row.querySelectorAll("td span");
			fields[0].textContent = item.amount;
			fields[1].textContent = item.name;
			fields[2].textContent = item.aisle;

			frag.appendChild(row);
		}

		let table = document.querySelector("#app-content tbody");
		table.textContent = "";
		table.appendChild(frag);
	};

	var _onItemClicked = function(evt) {
		let item = evt.target.closest("tr");
		item.classList.toggle("selected");
		console.log(item.dataset.id);
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
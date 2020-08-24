if (!OCA.HeyApple) {
	OCA.HeyApple = {};
}

OCA.HeyApple.Core = (function(){
	var _data = {};
	var _progress = {};
	var _rxTrim = /\d+\s?(ml|l|g|kg)\s/;
	var _rxAmount = /\d+\s?(ml|l|g|kg)/;

	var _add = function(amount, name) {
		let a1 = amount.split(" ");
		let a2 = _amount(name).split(" ");
		let n1 = parseInt(a1[0]);
		let n2 = parseInt(a2[0]);
		let u1 = a1[1];
		let u2 = a2[1];

		if (u1 == undefined && u2 == undefined) {
			return n1 + n2;
		}

		if (["g","ml"].indexOf(u1) >= 0 && ["kg","l"].indexOf(u2) >= 0) {
			n2 *= 1000;
		}else if (["kg","l"].indexOf(u1) >= 0 && ["g","ml"].indexOf(u2) >= 0) {
			n2 /= 1000;
		}

		return `${n1+n2} ${u1}`;
	}

	var _amount = function(name) {
		let out = name.match(_rxAmount);
		return out ? out[0] : "1";
	}

	var _bought = function(listId, itemId) {
		let listInfo = _progress[listId];
		if (listInfo) {
			return listInfo[itemId] || false;
		}
		return false;
	}

	var _trim = function(name) {
		return name.replace(_rxTrim, "");
	}

	return {
		init: function() {
			OCA.HeyApple.Backend.getConfig(function(obj) {
				document.querySelector("#path-settings").value = obj.directory;
			});

			OCA.HeyApple.Backend.getLists(function(obj) {
				_data = obj.success ? obj.data : {};
				OCA.HeyApple.UI.init();
			});
		},

		listNames: function() {
			return Object.keys(_data);
		},

		list: function(listId) {
			let out = {};

			if (_data[listId]) {
				_data[listId].forEach(function(elem) {
					let name = elem[1];
					let id = elem[2];
					let amount = _amount(name);

					if (out[id]) {
						out[id].amount = _add(out[id].amount, amount);
					} else {
						out[id] = {
							id: id,
							bought: _bought(listId, id),
							name: _trim(name),
							amount: amount,
							aisle: "None"
						};
					}
				});
			}

			return Object.values(out);
		},

		toggleBought: function(listId, itemId) {
			if (!_progress[listId]) {
				_progress[listId] = {};
			}
			let bought = _progress[listId][itemId] || false;
			_progress[listId][itemId] = !bought;
		}
	};
})();

OCA.HeyApple.UI = (function(){
	var _refreshCalendar = function() {
		let month = document.querySelector("#calendar2 select.month").value;
		let year = document.querySelector("#calendar2 select.year").value;
		let date = new Date(year, month);
		date.setDate(1 - date.getDay());

		let today = new Date();
		let d = today.getDate();
		let m = today.getMonth();
		let y = today.getFullYear();

		let cells = document.querySelectorAll("#calendar2 tbody td");
		for (let i = 0, cell; cell = cells[i]; i++) {
			let day = date.getDate();

			cell.firstElementChild.textContent = day;
			cell.classList.remove("selected");

			if (date.getMonth() != month) {
				cell.classList.add("outside");
			} else {
				cell.classList.remove("outside");
			}

			if (day == d && month == m && year == y) {
				cell.classList.add("today");
			} else {
				cell.classList.remove("today");
			}

			date.setDate(day+1);
		}
	};

	var _refreshLists = function() {
		let activeItem = document.querySelector("#list-category .active");
		let frag = document.createDocumentFragment();

		OCA.HeyApple.Core.listNames().forEach(function(name) {
			let item = document.createElement("li");
			item.textContent = name;
			item.dataset.name = name;
			item.addEventListener("click", function(evt) {
				_showList(evt.target);
			});

			if (!activeItem || activeItem.dataset.name == name) {
				activeItem = item;
			}

			frag.appendChild(item);
		});

		let list = document.querySelector("#list-category");
		list.textContent = "";
		list.appendChild(frag);

		if (activeItem) {
			_showList(activeItem);
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
			if (item.bought) {
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

		let list = document.querySelector("#list-category li.active");
		OCA.HeyApple.Core.toggleBought(list.dataset.name, item.dataset.id);
	};

	var _onDateClicked = function(evt) {
		let btn = evt.target.closest("td");
		btn.classList.toggle("selected");
	};

	return {
		init: function() {
			document.querySelector("#settings-item-scan").addEventListener("click", function() {
				OCA.HeyApple.Backend.scan(document.querySelector("#path-settings").value, function(obj) {
					if (obj.success) {
						window.location.reload();
					}
				});
			});

			let now = new Date();
			let months = document.querySelector("#calendar2 select.month");
			months.value = now.getMonth();
			months.addEventListener("change", _refreshCalendar);

			let years = document.querySelector("#calendar2 select.year");
			years.value = now.getFullYear();
			years.addEventListener("change", _refreshCalendar);

			let btns = document.querySelectorAll("#calendar2 > div > button");
			btns[0].addEventListener("click", function() {
				months.selectedIndex = (months.selectedIndex + 11) % 12;
				if (months.selectedIndex == 11) {
					years.selectedIndex = (years.selectedIndex + 2) % 3;
				}
				_refreshCalendar();
			});
			btns[1].addEventListener("click", function() {
				months.selectedIndex = (months.selectedIndex + 1) % 12;
				if (months.selectedIndex == 0) {
					years.selectedIndex = (years.selectedIndex + 1) % 3;
				}
				_refreshCalendar();
			});

			let days = document.querySelectorAll("#calendar2 tbody td button");
			for (let i = 0, day; day = days[i]; i++) {
				day.addEventListener("click", _onDateClicked);
			}

			_refreshLists();
			_refreshCalendar();
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

		getConfig: function(callback) {
			this.get(OC.generateUrl("apps/heyapple/api/0.1/config"), function() {
				callback(JSON.parse(this.response));
			});
		},

		getLists: function(callback) {
			this.get(OC.generateUrl("apps/heyapple/api/0.1/lists"), function() {
				callback(JSON.parse(this.response));
			});
		},

		scan: function(dir, callback) {
			let data = `dir=${dir}`;
			this.post(OC.generateUrl("apps/heyapple/api/0.1/scan"), data, function() {
				callback(JSON.parse(this.response));
			});
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	OCA.HeyApple.Core.init();
});
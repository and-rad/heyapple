if (!OCA.HeyApple) {
	OCA.HeyApple = {};
}

OCA.HeyApple.Core = (function () {
	var _data = {};
	var _progress = {};
	var _progressTimeout = undefined;
	var _rxTrim = /\d+\s?(ml|l|g|kg)\s/;
	var _rxAmount = /\d+\s?(ml|l|g|kg)/;

	var _add = function (amount, name) {
		let a1 = amount.split(" ");
		let a2 = _amount(name).split(" ");
		let n1 = parseInt(a1[0]);
		let n2 = parseInt(a2[0]);
		let u1 = a1[1];
		let u2 = a2[1];

		if (u1 == undefined && u2 == undefined) {
			return n1 + n2;
		}

		if (["g", "ml"].indexOf(u1) >= 0 && ["kg", "l"].indexOf(u2) >= 0) {
			n2 *= 1000;
		} else if (
			["kg", "l"].indexOf(u1) >= 0 &&
			["g", "ml"].indexOf(u2) >= 0
		) {
			n2 /= 1000;
		}

		return `${n1 + n2} ${u1}`;
	};

	var _amount = function (name) {
		let out = name.match(_rxAmount);
		return out ? out[0] : "1";
	};

	var _bought = function (listId, itemId) {
		let listInfo = _progress[listId];
		if (listInfo) {
			return listInfo.indexOf(itemId) != -1;
		}
		return false;
	};

	var _formatDateStrings = function () {
		Object.keys(_data).forEach(function (prop) {
			let list = _data[prop];
			for (let i = 0, item; (item = list[i]); i++) {
				let [d, m, y] = item[0].split(" ")[0].split(".");
				let date = new Date(y, m - 1, d);
				if (!isNaN(date)) {
					item[0] = date.toISOString().split("T")[0];
				}
			}
		});
	};

	var _trim = function (name) {
		return name.replace(_rxTrim, "");
	};

	return {
		init: function () {
			OCA.HeyApple.Backend.getConfig(function (obj) {
				document.querySelector("#path-settings").value = obj.directory;
			});

			OCA.HeyApple.Backend.getLists(function (obj) {
				_data = obj.success ? obj.data.lists : {};
				_progress = obj.success ? obj.data.completed || {} : {};
				_formatDateStrings();
				OCA.HeyApple.UI.init();
			});
		},

		listDates: function () {
			let dates = {};
			for (let key in _data) {
				_data[key].forEach((item) => (dates[item[0]] = true));
			}
			return Object.keys(dates);
		},

		listNames: function () {
			return Object.keys(_data);
		},

		list: function (listId, days) {
			let out = {};

			if (_data[listId]) {
				_data[listId].forEach(function (elem) {
					let date = elem[0];
					if (date.length != 0 && days.indexOf(date) == -1) {
						return;
					}

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
							aisle: elem[3],
						};
					}
				});
			}

			return Object.values(out);
		},

		toggleBought: function (listId, itemId) {
			clearTimeout(_progressTimeout);

			if (!_progress[listId]) {
				_progress[listId] = [];
			}

			let listInfo = _progress[listId];
			let idx = listInfo.indexOf(itemId);
			if (idx > -1) {
				listInfo.splice(idx, 1);
			} else {
				listInfo.push(itemId);
			}

			_progressTimeout = setTimeout(function () {
				OCA.HeyApple.Backend.setCompleted(_progress, function () {});
			}, 1000);
		},
	};
})();

OCA.HeyApple.UI = (function () {
	var _selection = {};
	var _sortBy = "name";
	var _sortAsc = true;
	var _hideSelected = false;

	var _refreshCalendar = function () {
		let month = document.querySelector("#calendar2 select.month").value;
		let year = document.querySelector("#calendar2 select.year").value;
		let date = new Date(year, month);
		date.setDate(1 - date.getDay());

		let today = new Date();
		let d = today.getDate();
		let m = today.getMonth();
		let y = today.getFullYear();

		let dates = OCA.HeyApple.Core.listDates();
		let cells = document.querySelectorAll("#calendar2 tbody td");
		for (let i = 0, cell; (cell = cells[i]); i++) {
			let day = date.getDate();
			let mon = date.getMonth();
			let iso = date.toISOString().split("T")[0];

			cell.firstElementChild.textContent = day;
			cell.firstElementChild.dataset.date = iso;

			if (_selection[iso]) {
				cell.classList.add("selected");
			} else {
				cell.classList.remove("selected");
			}

			if (mon != month) {
				cell.classList.add("outside");
			} else {
				cell.classList.remove("outside");
			}

			if (day == d && mon == m && year == y) {
				cell.classList.add("today");
			} else {
				cell.classList.remove("today");
			}

			if (dates.indexOf(iso) != -1) {
				cell.classList.add("has-entries");
			} else {
				cell.classList.remove("has-entries");
			}

			date.setDate(day + 1);
		}
	};

	var _refreshLists = function () {
		let activeItem = document.querySelector("#list-category .active");
		let frag = document.createDocumentFragment();
		let tmpl = document.createElement("li");
		tmpl.innerHTML = document.querySelector(
			"#template-menu-item"
		).innerHTML;

		OCA.HeyApple.Core.listNames().forEach(function (name) {
			let item = tmpl.cloneNode(true);
			item.dataset.name = name;
			item.firstElementChild.textContent = name;
			item.firstElementChild.addEventListener("click", function (evt) {
				_showList(evt.target.parentNode);
				evt.preventDefault();
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

	var _refreshHeadToggle = function () {
		let x = document.querySelectorAll("#app-content tbody tr").length;
		let y = document.querySelectorAll(
			"#app-content tbody tr.selected"
		).length;
		let head = document.querySelector("#app-content th.selection");
		x == y
			? head.classList.add("selected")
			: head.classList.remove("selected");
	};

	var _showList = function (elem) {
		for (let i = 0, e; (e = elem.parentNode.children[i]); i++) {
			e == elem
				? e.classList.add("active")
				: e.classList.remove("active");
		}

		let frag = document.createDocumentFragment();
		let tmpl = document.createElement("tr");
		tmpl.innerHTML = document.querySelector("#template-item").innerHTML;

		let list = OCA.HeyApple.Core.list(
			elem.dataset.name,
			Object.keys(_selection)
		);
		for (let i = 0, item; (item = list[i]); i++) {
			let row = tmpl.cloneNode(true);
			row.dataset.id = item.id;
			row.firstElementChild.addEventListener("click", _onItemClicked);
			if (item.bought) {
				row.classList.add("selected");
				if (_hideSelected) {
					row.style.display = "none";
				}
			}

			let fields = row.querySelectorAll("td span");
			fields[0].textContent = item.amount;
			fields[1].textContent = item.name;
			fields[2].textContent = item.aisle;

			frag.appendChild(row);
		}

		let title = document.querySelector("#controls .list-title");
		title.textContent = elem.dataset.name;

		let table = document.querySelector("#app-content tbody");
		table.textContent = "";
		table.appendChild(frag);

		_sortTable(_sortBy);
		_refreshHeadToggle();
	};

	var _sortTable = function (cat, toggle) {
		_sortBy = cat;

		if (toggle) {
			if (
				document.querySelector(
					`#app-content th.${_sortBy} > span:not(.hidden)`
				)
			) {
				_sortAsc = !_sortAsc;
			}
		}

		let heads = document.querySelectorAll("#app-content th.sort");
		for (let i = 0, head; (head = heads[i]); i++) {
			if (head.classList.contains(_sortBy)) {
				head.firstElementChild.classList.remove("hidden");
			} else {
				head.firstElementChild.classList.add("hidden");
			}

			if (_sortAsc) {
				head.firstElementChild.classList.remove("icon-triangle-s");
				head.firstElementChild.classList.add("icon-triangle-n");
			} else {
				head.firstElementChild.classList.remove("icon-triangle-n");
				head.firstElementChild.classList.add("icon-triangle-s");
			}
		}

		let body = document.querySelector("#app-content tbody");
		let tr = Array.from(body.querySelectorAll("tr"));
		tr.sort(function (a, b) {
			let text1 = a.querySelector(`.${_sortBy} > div > span`).textContent;
			let text2 = b.querySelector(`.${_sortBy} > div > span`).textContent;
			let locale = document.documentElement.dataset.locale || "en";
			let out = text1.localeCompare(text2, locale, { numeric: true });
			if (!_sortAsc) out *= -1;
			return out;
		});

		tr.forEach((t) => {
			body.appendChild(t);
		});
	};

	var _onItemClicked = function (evt) {
		let item = evt.target.closest("tr");
		item.classList.toggle("selected");
		_refreshHeadToggle();

		let list = document.querySelector("#list-category li.active");
		OCA.HeyApple.Core.toggleBought(list.dataset.name, item.dataset.id);

		if (_hideSelected && item.classList.contains("selected")) {
			item.style.display = "none";
		}
	};

	var _onHeadClicked = function (evt) {
		let box = evt.target.closest("th");
		box.classList.toggle("selected");

		let on = box.classList.contains("selected");
		let list = document.querySelector("#list-category li.active");
		let items = document.querySelectorAll("#app-content tbody tr");

		for (let i = 0, item; (item = items[i]); i++) {
			if (item.classList.contains("selected") != on) {
				item.classList.toggle("selected");
				item.style.display =
					_hideSelected && item.classList.contains("selected")
						? "none"
						: "table-row";
				OCA.HeyApple.Core.toggleBought(
					list.dataset.name,
					item.dataset.id
				);
			}
		}
	};

	var _onDateClicked = function (evt) {
		evt.target.closest("td").classList.toggle("selected");

		let date = evt.target.dataset.date;
		if (_selection[date]) {
			delete _selection[date];
		} else {
			_selection[date] = true;
		}

		_refreshLists();
	};

	var _onHideClicked = function (evt) {
		evt.target.classList.toggle("selected");
		_hideSelected = evt.target.classList.contains("selected");

		let rows = document.querySelectorAll("#app-content tbody tr.selected");
		for (let i = 0, row; (row = rows[i]); i++) {
			row.style.display = _hideSelected ? "none" : "table-row";
		}
	};

	return {
		init: function () {
			document
				.querySelector("#settings-item-scan")
				.addEventListener("click", function () {
					OCA.HeyApple.Backend.scan(
						document.querySelector("#path-settings").value,
						function (obj) {
							if (obj.success) {
								window.location.reload();
							}
						}
					);
				});

			document
				.querySelector("#app-content th.selection")
				.addEventListener("click", _onHeadClicked);
			document
				.querySelector("#controls .selection > div")
				.addEventListener("click", _onHideClicked);

			let cols = document.querySelectorAll("th.sort");
			for (let i = 0, col; (col = cols[i]); i++) {
				col.addEventListener("click", function (evt) {
					_sortTable(evt.target.dataset.sort, true);
				});
			}

			let now = new Date();
			let months = document.querySelector("#calendar2 select.month");
			months.value = now.getMonth();
			months.addEventListener("change", _refreshCalendar);

			let years = document.querySelector("#calendar2 select.year");
			years.value = now.getFullYear();
			years.addEventListener("change", _refreshCalendar);

			let btns = document.querySelectorAll("#calendar2 > div > button");
			btns[0].addEventListener("click", function () {
				months.selectedIndex = (months.selectedIndex + 11) % 12;
				if (months.selectedIndex == 11) {
					years.selectedIndex = (years.selectedIndex + 2) % 3;
				}
				_refreshCalendar();
			});
			btns[1].addEventListener("click", function () {
				months.selectedIndex = (months.selectedIndex + 1) % 12;
				if (months.selectedIndex == 0) {
					years.selectedIndex = (years.selectedIndex + 1) % 3;
				}
				_refreshCalendar();
			});

			let days = document.querySelectorAll("#calendar2 tbody td button");
			for (let i = 0, day; (day = days[i]); i++) {
				day.addEventListener("click", _onDateClicked);
			}

			let date = new Date(
				now.getFullYear(),
				now.getMonth(),
				now.getDate()
			);
			let dates = OCA.HeyApple.Core.listDates();
			for (let i = 0; i < 7; i++) {
				let iso = date.toISOString().split("T")[0];
				if (dates.indexOf(iso) != -1) {
					_selection[iso] = true;
				}
				date.setDate(date.getDate() + 1);
			}

			_refreshLists();
			_refreshCalendar();
		},
	};
})();

OCA.HeyApple.Backend = (function () {
	return {
		get: function (uri, callback) {
			let xhr = new XMLHttpRequest();
			xhr.addEventListener("load", callback);
			xhr.open("GET", uri);
			xhr.setRequestHeader("requesttoken", OC.requestToken);
			xhr.send();
		},

		post: function (uri, data, callback) {
			let xhr = new XMLHttpRequest();
			xhr.addEventListener("load", callback);
			xhr.open("POST", uri);
			xhr.setRequestHeader("requesttoken", OC.requestToken);
			xhr.setRequestHeader(
				"Content-Type",
				"application/x-www-form-urlencoded"
			);
			xhr.send(data);
		},

		postJSON: function (uri, data, callback) {
			let xhr = new XMLHttpRequest();
			xhr.addEventListener("load", callback);
			xhr.open("POST", uri);
			xhr.setRequestHeader("requesttoken", OC.requestToken);
			xhr.setRequestHeader(
				"Content-Type",
				"application/json; charset=UTF-8"
			);
			xhr.send(JSON.stringify(data));
		},

		getConfig: function (callback) {
			this.get(
				OC.generateUrl("apps/heyapple/api/0.1/config"),
				function () {
					callback(JSON.parse(this.response));
				}
			);
		},

		getLists: function (callback) {
			this.get(
				OC.generateUrl("apps/heyapple/api/0.1/lists"),
				function () {
					callback(JSON.parse(this.response));
				}
			);
		},

		setCompleted: function (completed, callback) {
			this.postJSON(
				OC.generateUrl("apps/heyapple/api/0.1/complete"),
				completed,
				function () {
					callback(JSON.parse(this.response));
				}
			);
		},

		scan: function (dir, callback) {
			let data = `dir=${dir}`;
			this.post(
				OC.generateUrl("apps/heyapple/api/0.1/scan"),
				data,
				function () {
					callback(JSON.parse(this.response));
				}
			);
		},
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	//OCA.HeyApple.Core.init();
	/*
	let callback = function () {
		console.log(this.response);
	};
	let xhr = new XMLHttpRequest();
	xhr.addEventListener("load", callback);
	xhr.open("GET", "http://localhost:8080");
	xhr.setRequestHeader("requesttoken", OC.requestToken);
	xhr.send();
	*/
	fetch("http://localhost:8080")
		.then((response) => response.json())
		.then((data) => console.log(data));
});

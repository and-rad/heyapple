if (!OCA.HeyApple) {
	OCA.HeyApple = {};
}

OCA.HeyApple.Core = (function(){
	return {
		init: function() {
			OCA.HeyApple.UI.init();
			OCA.HeyApple.Diary.init();
		}
	};
})();

OCA.HeyApple.Diary = (function(){
	var _refreshCalendar = function() {

	};

	return {
		init: function() {
			let picker = new Pikaday({
				field: document.getElementById('calendar'),
				bound: false,
				container: document.getElementById('calendar'),
				showDaysInNextAndPreviousMonths: true,
				firstDay: 1,
				onDraw: _refreshCalendar,
				i18n: {
					previousMonth : '←',
					nextMonth     : '→',
					months        : ['January','February','March','April','May','June','July','August','September','October','November','December'],
					weekdays      : ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'],
					weekdaysShort : ['Su','Mo','Tu','We','Th','Fr','Sa']
				}
			});
		}
	};
})();

OCA.HeyApple.UI = (function(){
	var _showPage = function(idx) {
		let nav = document.querySelectorAll('#list-category > li');
		for (let i = 0, entry; entry = nav[i]; i++) {
			i == idx ? entry.classList.add("active") : entry.classList.remove("active");
		}

		let cat = document.querySelectorAll('#category > div');
		for (let i = 0, entry; entry = cat[i]; i++) {
			i == idx ? entry.style.display = "block" : entry.style.display = "none";
		}
	}

	return {
		init: function() {
			let nav = document.querySelectorAll('#list-category > li');
			nav[0].addEventListener("click", function() { _showPage(0); });
			nav[1].addEventListener("click", function() { _showPage(1); });
			nav[2].addEventListener("click", function() { _showPage(2); });
			nav[3].addEventListener("click", function() { _showPage(3); });
			_showPage(0);
		}
	};
})();

document.addEventListener("DOMContentLoaded", function () {
	OCA.HeyApple.Core.init();
});
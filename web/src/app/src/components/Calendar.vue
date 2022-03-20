<script setup>
import ArrowImage from "./images/ImageRightArrow.vue";
import { ref, computed, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t } = useI18n();
const prop = defineProps(["mode", "items"]);
const emit = defineEmits(["selection"]);
const month = ref(DateTime.now().month);
const year = ref(DateTime.now().year);

const calendar = ref(null);

const years = computed(() => {
	let dates = prop.items.map((i) => parseInt(i.split("-")[0])).filter((i, idx, self) => self.indexOf(i) === idx);
	return [Math.min(...dates) - 1, ...dates, Math.max(...dates) + 1];
});

const hasPrev = computed(() => {
	return year.value != years.value[0] || month.value > 1;
});

const hasNext = computed(() => {
	return year.value != years.value[years.value.length - 1] || month.value < 12;
});

let selection = [];

function onCalendarChanged() {
	let today = DateTime.now().toISODate();
	let date = DateTime.local(year.value, month.value);
	date = date.minus({ days: date.weekday });

	if (!selection.length && prop.mode != "toggle") {
		selection = [today];
	}

	calendar.value.querySelectorAll("td").forEach((cell) => {
		let iso = date.toISODate();

		cell.firstElementChild.textContent = date.day;
		cell.firstElementChild.dataset.date = iso;

		if (date.month != month.value) {
			cell.classList.add("outside");
		} else {
			cell.classList.remove("outside");
		}

		if (iso == today) {
			cell.classList.add("today");
		} else {
			cell.classList.remove("today");
		}

		if (prop.items.indexOf(iso) != -1) {
			cell.classList.add("has-entries");
		} else {
			cell.classList.remove("has-entries");
		}

		if (selection.indexOf(iso) != -1) {
			cell.classList.add("active");
		} else {
			cell.classList.remove("active");
		}

		date = date.plus({ days: 1 });
	});
}

function onPrev() {
	if (--month.value < 1) {
		month.value = 12;
		--year.value;
	}
	onCalendarChanged();
}

function onNext() {
	if (++month.value > 12) {
		month.value = 1;
		++year.value;
	}
	onCalendarChanged();
}

function onDay(evt) {
	let iso = evt.target.dataset.date;

	if (prop.mode == "toggle") {
		let idx = selection.indexOf(iso);
		if (idx == -1) {
			selection.push(iso);
		} else {
			selection.splice(idx, 1);
		}
	} else {
		selection = [iso];
	}

	calendar.value.querySelectorAll("td>button").forEach((btn) => {
		if (selection.indexOf(btn.dataset.date) != -1) {
			btn.parentNode.classList.add("active");
		} else {
			btn.parentNode.classList.remove("active");
		}
	});

	emit("selection", selection);
}

onMounted(() => {
	onCalendarChanged();
	emit("selection", selection);
});
</script>

<template>
	<div class="calendar" :class="[mode]">
		<div>
			<button class="prev icon" @click="onPrev" :disabled="!hasPrev"><ArrowImage /></button>
			<div>
				<select class="month" v-model.number="month" @change="onCalendarChanged">
					<option value="1">{{ t("month.1") }}</option>
					<option value="2">{{ t("month.2") }}</option>
					<option value="3">{{ t("month.3") }}</option>
					<option value="4">{{ t("month.4") }}</option>
					<option value="5">{{ t("month.5") }}</option>
					<option value="6">{{ t("month.6") }}</option>
					<option value="7">{{ t("month.7") }}</option>
					<option value="8">{{ t("month.8") }}</option>
					<option value="9">{{ t("month.9") }}</option>
					<option value="10">{{ t("month.10") }}</option>
					<option value="11">{{ t("month.11") }}</option>
					<option value="12">{{ t("month.12") }}</option>
				</select>
				<select class="year" v-model.number="year" @change="onCalendarChanged">
					<option v-for="y in years" :value="y">{{ y }}</option>
				</select>
			</div>
			<button class="next icon" @click="onNext" :disabled="!hasNext"><ArrowImage /></button>
		</div>
		<table ref="calendar">
			<thead>
				<tr>
					<th>{{ t("day.cal7") }}</th>
					<th>{{ t("day.cal1") }}</th>
					<th>{{ t("day.cal2") }}</th>
					<th>{{ t("day.cal3") }}</th>
					<th>{{ t("day.cal4") }}</th>
					<th>{{ t("day.cal5") }}</th>
					<th>{{ t("day.cal6") }}</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
				</tr>
				<tr>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
				</tr>
				<tr>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
				</tr>
				<tr>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
				</tr>
				<tr>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
				</tr>
				<tr>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
					<td><button @click="onDay"></button></td>
				</tr>
			</tbody>
		</table>
	</div>
</template>

<style>
.calendar {
	white-space: nowrap;
}

.calendar > div {
	display: flex;
	justify-content: space-between;
}

.calendar > div button[disabled] {
	background: none;
}

.calendar > div button[disabled] svg {
	fill: #e0e0e0;
}

.calendar > div > div {
	display: flex;
	margin: 0 0.25em;
}

.calendar > div > div select {
	display: inline-block;
	flex-basis: 40%;
}

.calendar > div > div .month {
	margin-right: 0.25em;
	flex-basis: 60%;
}

.calendar .prev svg {
	transform: rotate(180deg);
}

.calendar table {
	border-collapse: collapse;
	width: 100%;
	margin-top: 0.5em;
}

.calendar th,
.calendar td {
	width: calc(100% / 7);
	padding: 1px;
}

.calendar th {
	font-weight: 700;
	color: var(--color-primary);
}

.calendar td button {
	width: 38px;
	height: 38px;
	background: none;
	/*border-radius: 50%;*/
	color: var(--color-text);
	line-height: 1;
	position: relative;
}

.calendar td.active button {
	border: 1px solid var(--color-primary);
}

.calendar td.today button {
	border: 1px solid var(--color-secondary);
}

@media (hover: hover) {
	.calendar td.active button:hover {
		box-shadow: inset 0 0 100px var(--color-primary-light);
	}
	.calendar td.today button:hover {
		box-shadow: inset 0 0 100px var(--color-secondary-light);
	}
}

.calendar td.today.active button:before {
	content: "";
	border: 1px solid var(--color-primary);
	position: absolute;
	top: -4px;
	right: -4px;
	bottom: -4px;
	left: -4px;
	border-radius: 5px;
}

.calendar.toggle td.active button {
	border: none;
	box-shadow: inset 0 0 100px var(--color-primary);
	color: #fff;
}

.calendar.toggle td.today.active button {
	box-shadow: inset 0 0 100px var(--color-secondary);
	color: var(--color-text);
}

.calendar.toggle td.today.active button:before {
	display: none;
}

.calendar td.outside button {
	opacity: 0.3;
	font-weight: 300;
}

.calendar td.has-entries button {
	font-weight: 700;
}

.calendar td.outside.has-entries button {
	font-weight: 400;
}
</style>

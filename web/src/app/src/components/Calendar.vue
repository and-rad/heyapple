<script setup>
import ArrowImage from "./images/ImageRightArrow.vue";
import { ref, computed, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t } = useI18n();
const years = ref([2020, 2021, 2022, 2023]);
const dates = ref(["2022-03-12", "2022-03-14", "2022-03-15", "2022-03-16"]);
const month = ref(DateTime.now().month);
const year = ref(DateTime.now().year);

const calendar = ref(null);

const hasPrev = computed(() => {
	return year.value != years.value[0] || month.value > 1;
});

const hasNext = computed(() => {
	return year.value != years.value[years.value.length - 1] || month.value < 12;
});

function onCalendarChanged() {
	let today = DateTime.now().toISODate();
	let date = DateTime.local(year.value, month.value);

	date = date.minus({ days: date.weekday });
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

		if (dates.value.indexOf(iso) != -1) {
			cell.classList.add("has-entries");
		} else {
			cell.classList.remove("has-entries");
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
	console.log(evt.target.dataset.date);
}

onMounted(() => onCalendarChanged());
</script>

<template>
	<div class="calendar">
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

.calendar > div > div select {
	width: auto;
	display: inline-block;
}

.calendar > div > div .month {
	margin-right: 0.25em;
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
	padding: 0;
}

.calendar th {
	font-weight: 700;
	color: var(--color-primary);
}

.calendar td button {
	width: 100%;
	height: 40px;
	background: none;
	border-radius: 0;
	color: var(--color-text);
	line-height: 1;
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

.calendar td.today button {
	border: 1px solid var(--color-primary);
}

.calendar td.today button:hover {
	box-shadow: inset 0 0 100px var(--color-primary-light);
}
</style>

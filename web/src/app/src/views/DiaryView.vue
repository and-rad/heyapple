<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import EntryList from "../components/DiaryEntryList.vue";
import PieChart from "../components/PieChart.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, computed, inject } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime, Duration } from "luxon";

const { t } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");
const diary = inject("diary");
const prefs = inject("prefs");

const current = ref(null);
const currentDate = ref(DateTime.now());
const currentNutrient = ref("kcal");
const disableSave = ref(false);
const editMode = ref(false);

const calendar = ref(null);
const entries = ref(null);
const main = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

const currentWeek = computed(() => {
	let result = [];

	let isoNow = DateTime.now().toISODate();
	let isoCurrent = currentDate.value.toISODate();
	let date = currentDate.value.minus({ days: currentDate.value.weekday - 1 });
	for (let i = 0; i < 7; ++i) {
		let iso = date.toISODate();
		result.push({
			weekday: t("day.cal" + date.weekday),
			calday: date.day,
			date: iso,
			active: iso == isoCurrent,
			today: iso == isoNow,
			value: diary.value[iso] ? diary.value[iso].kcal : 0,
		});
		date = date.plus({ days: 1 });
	}

	return result;
});

let hasTabDrag = false;

function onTabSlide(evt) {
	evt.stopPropagation();
	evt.preventDefault();
	moveTabBar(evt.target.closest("ul"), evt.movementX);
	hasTabDrag = true;
}

function onTabWheel(evt) {
	evt.stopPropagation();
	evt.preventDefault();
	let delta = Math.max(-16, Math.min(-evt.deltaY, 16));
	moveTabBar(evt.target.closest("ul"), delta);
}

function moveTabBar(elem, delta) {
	let offset = parseInt(elem.style.getPropertyValue("--offset")) || 0;
	let final = Math.max(elem.clientWidth - elem.scrollWidth, Math.min(offset + delta, 0));
	elem.style.setProperty("--offset", final + "px");
}

function onTabsPress(evt) {
	let handle = evt.target.closest("ul");
	handle.addEventListener("pointermove", onTabSlide);
	handle.addEventListener("mouseup", onTabsRelease);
	handle.addEventListener("touchend", onTabsRelease);
	handle.addEventListener("mouseleave", onTabsRelease);
	handle.addEventListener("touchcancel", onTabsRelease);
	hasTabDrag = false;
}

function onTabsRelease(evt) {
	let handle = evt.target.closest("ul");
	handle.removeEventListener("pointermove", onTabSlide);
	handle.removeEventListener("mouseup", onTabsRelease);
	handle.removeEventListener("touchend", onTabsRelease);
	handle.removeEventListener("mouseleave", onTabsRelease);
	handle.removeEventListener("touchcancel", onTabsRelease);
}

function onTabClick(evt) {
	if (!hasTabDrag) {
		currentNutrient.value = evt.target.dataset.name;
	}
}

function onEditMode() {
	if (current.value && current.value.entries) {
		editMode.value ? saveEntries() : (editMode.value = true);
	}
}

function saveEntries() {
	let items = entries.value.getDiff();
	if (items.length == 0) {
		editMode.value = false;
		return;
	}

	disableSave.value = true;

	let date = current.value.date;
	let params = new URLSearchParams();
	items.forEach((i) => {
		params.append("id", i.id);
		params.append("amount", i.amount);
		params.append("time", i.time);
		params.append("recipe", i.recipe);
	});

	fetch(`/api/v1/diary/${date}`, {
		method: "PUT",
		body: params,
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savediary.err" + response.status);
			}
			editMode.value = false;
			return fetch("/api/v1/diary/" + date.replaceAll("-", "/"));
		})
		.then((response) => response.json())
		.then((data) => {
			if (data[0] && data[0].kcal) {
				diary.value[date] = data[0];
			} else {
				delete diary.value[date];
			}
			onDateSelected([date]);
			log.msg(t("savediary.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				disableSave.value = false;
			}, 500);
		});
}

function onDateSelected(dates) {
	currentDate.value = DateTime.fromISO(dates[0]);

	let date = currentDate.value.toISODate();
	if (!diary.value[date]) {
		current.value = { date: date };
		return;
	}
	if (diary.value[date].entries) {
		current.value = diary.value[date];
		return;
	}

	fetch(`/api/v1/diary/${date.replaceAll("-", "/")}/entries`)
		.then((response) => response.json())
		.then((data) => {
			data.forEach((d) => (d.time = d.date.split("T")[1].slice(0, 5)));
			data.sort((a, b) => {
				if (a.time < b.time) return -1;
				if (b.time < a.time) return 1;
				return 0;
			});
			diary.value[date].entries = data;
		})
		.catch((err) => log.err(err))
		.finally(() => (current.value = diary.value[date]));
}
</script>

<template>
	<Main ref="main" class="diary" :class="{ 'edit-mode': editMode }" @detailVisibility="editMode = false">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar ref="calendar" :items="daysWithEntries" @selection="onDateSelected" />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main>
			<section id="charts-week">
				<button
					v-for="day in currentWeek"
					:data-date="day.date"
					:class="{ today: day.today, active: day.active }"
					@click="calendar.onDay"
				>
					<PieChart range="360" :value="day.value" :max="prefs.rdi.kcal">
						<template #details>
							<span>{{ day.weekday }}</span>
							<span>{{ day.calday }}</span>
						</template>
					</PieChart>
				</button>
			</section>
			<section id="charts-macro">
				<PieChart
					class="kcal"
					start="225"
					range="270"
					frac="0"
					:label="t('food.energy')"
					:unit="t('unit.cal')"
					:value="current ? current.kcal : 0"
					:max="prefs.rdi.kcal"
				/>
				<PieChart
					class="fat"
					start="225"
					range="270"
					frac="0"
					:label="t('food.fat')"
					:unit="t('unit.g')"
					:value="current ? current.fat : 0"
					:max="prefs.rdi.fat"
				/>
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.carbs2')"
					:unit="t('unit.g')"
					:value="current ? current.carb : 0"
					:max="prefs.rdi.carb"
				/>
				<PieChart
					class="prot"
					start="225"
					range="270"
					frac="0"
					:label="t('food.protein')"
					:unit="t('unit.g')"
					:value="current ? current.prot : 0"
					:max="prefs.rdi.prot"
				/>
			</section>
		</template>

		<template #head-details>
			<h2 class="no-edit-mode">{{ currentDate.weekdayLong }}</h2>
		</template>

		<template #details>
			<section class="subtitle no-edit-mode">
				{{ currentDate.toLocaleString(DateTime.DATE_FULL) }}
			</section>
			<section class="tabs">
				<div>
					<ul @mousedown="onTabsPress" @touchstart="onTabsPress" @wheel="onTabWheel">
						<li :class="{ active: currentNutrient == 'kcal' }">
							<button data-name="kcal" @click="onTabClick">{{ t("food.energy") }}</button>
						</li>
						<li :class="{ active: currentNutrient == 'fat' }">
							<button data-name="fat" @click="onTabClick">{{ t("food.fat") }}</button>
						</li>
						<li :class="{ active: currentNutrient == 'carb' }">
							<button data-name="carb" @click="onTabClick">{{ t("food.carbs2") }}</button>
						</li>
						<li :class="{ active: currentNutrient == 'prot' }">
							<button data-name="prot" @click="onTabClick">{{ t("food.protein") }}</button>
						</li>
					</ul>
				</div>
				<button class="icon async" :disabled="disableSave" @click="onEditMode">
					<EditImage v-if="!editMode" />
					<SaveImage v-if="editMode" />
				</button>
			</section>
			<hr />
			<section v-if="current">
				<h2>{{ t("aria.headlog") }}</h2>
				<p class="msg-noitems" v-if="!current.entries" v-html="t('diary.noitems')"></p>
				<EntryList
					ref="entries"
					v-if="current.entries"
					:entries="current.entries"
					:nutrient="currentNutrient"
					:disabled="!editMode"
				/>
			</section>
		</template>
	</Main>
</template>

<style>
:root {
	--color-kcal: var(--color-primary);
	--color-kcal-light: var(--color-primary-light);
	--color-fat: #03a9f4;
	--color-fat-light: #e1f5fe;
	--color-carb: #ffa726;
	--color-carb-light: #ffe0b2;
	--color-prot: #ab47bc;
	--color-prot-light: #f3e5f5;
}

#charts-week {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 0 1em 1em;
}

#charts-week button {
	background: none;
	box-shadow: none !important;
	border-radius: 50%;
	margin: 0;
	padding: 0;
	color: var(--color-text);
	flex-basis: 14%;
	max-width: 4em;
	max-height: 4em;
	font-size: 12px;
	line-height: normal;
}

#charts-week button.active {
	color: var(--color-primary-dark);
}

#charts-week button.today.active {
	color: var(--color-secondary-dark);
}

#charts-week button div span:first-child {
	font-weight: 700;
}

#charts-week .pie-chart {
	width: auto;
	height: auto;
	user-select: none;
	pointer-events: none;
}

#charts-week .pie-chart circle.base {
	stroke: var(--color-background);
	fill: var(--color-primary-lighter);
	stroke-width: 16;
}

#charts-week .pie-chart circle.good {
	stroke-width: 8;
}

#charts-week .pie-chart circle.bad {
	stroke-width: 24;
}

#charts-week button.today .pie-chart circle.base {
	fill: var(--color-secondary-lighter);
}

#charts-week button.today .pie-chart circle.good {
	stroke: var(--color-secondary);
}

#charts-week button.active .pie-chart circle.base {
	fill: var(--color-primary-light);
}

#charts-week button.today.active .pie-chart circle.base {
	fill: var(--color-secondary-light);
}

@media (hover: hover) {
	#charts-week button:hover .pie-chart circle.base {
		fill: var(--color-primary-light);
	}

	#charts-week button.today:hover .pie-chart circle.base {
		fill: var(--color-secondary-light);
	}
}

#charts-macro {
	display: flex;
	flex-wrap: wrap;
	justify-content: space-between;
	align-items: center;
	padding: 1em;
}

#charts-macro .pie-chart {
	flex-basis: 30%;
	margin: 1em 0;
	width: auto;
	height: auto;
}

#charts-macro .pie-chart > div span:last-child {
	color: var(--color-text-light);
}

#charts-macro .pie-chart.kcal {
	flex-basis: 100%;
	margin: 1em 15%;
}

#charts-macro .pie-chart.kcal figcaption {
	bottom: 10%;
}

#charts-macro .pie-chart.kcal circle.base {
	stroke: var(--color-kcal-light);
}

#charts-macro .pie-chart.kcal circle.good {
	stroke: var(--color-kcal);
}

#charts-macro .pie-chart.fat circle.base {
	stroke: var(--color-fat-light);
}

#charts-macro .pie-chart.fat circle.good {
	stroke: var(--color-fat);
}

#charts-macro .pie-chart.carb circle.base {
	stroke: var(--color-carb-light);
}

#charts-macro .pie-chart.carb circle.good {
	stroke: var(--color-carb);
}

#charts-macro .pie-chart.prot circle.base {
	stroke: var(--color-prot-light);
}

#charts-macro .pie-chart.prot circle.good {
	stroke: var(--color-prot);
}

@media only screen and (min-width: 800px) {
	#charts-week button {
		font-size: unset;
	}

	#charts-macro {
		flex-wrap: nowrap;
	}

	#charts-macro .pie-chart {
		flex-basis: 20%;
	}

	#charts-macro .pie-chart.kcal {
		flex-basis: 30%;
		margin: 1em 0;
	}

	#charts-macro .pie-chart figcaption {
		bottom: 10%;
	}
}

#details section.tabs {
	padding: 0 3em 0 0;
	position: relative;
}

#details section.tabs:before,
#details section.tabs:after {
	position: absolute;
	pointer-events: none;
	top: 0;
	bottom: 2px;
	width: 2em;
	content: "";
	z-index: 1;
}

#details section.tabs:before {
	left: 0;
	background: linear-gradient(to right, #fff 10%, transparent);
}

#details section.tabs:after {
	width: 5em;
	right: 0em;
	background: linear-gradient(to left, #fff 64%, transparent);
}

#details section.tabs button.icon {
	position: absolute;
	right: 0.5em;
	bottom: 0.5em;
	z-index: 2;
}

#details section.tabs > div {
	overflow: hidden;
	margin-bottom: -1px;
}

#details section.tabs ul {
	--offset: 0px;
	list-style: none;
	padding: 0;
	white-space: nowrap;
	touch-action: none;
	transform: translateX(var(--offset));
}

#details section.tabs li {
	display: inline-block;
	min-width: 8em;
}

#details section.tabs li button {
	background: none;
	box-shadow: none !important;
	color: var(--color-placeholder);
	border-radius: 0;
	padding: 0.5em 0.5em 0.35em;
	border-bottom: 2px solid transparent;
	transition: color var(--transition-style), border-color var(--transition-style);
}

@media (hover: hover) {
	#details section.tabs li button:hover {
		border-color: var(--color-text-light);
		box-shadow: none;
		color: var(--color-text-light);
	}
}

#details section.tabs li.active button {
	border-color: var(--color-secondary);
	color: var(--color-text);
}

#details .diary-entry-list label {
	color: var(--color-text);
}
</style>

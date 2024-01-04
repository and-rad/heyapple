<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import EntryList from "../components/DiaryEntryList.vue";
import PieChart from "../components/PieChart.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, computed, inject, onMounted } from "vue";
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
const select = ref(null);

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

function onTabSelect(evt) {
	currentNutrient.value = evt.target.value;
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

onMounted(() => {
	onDateSelected(calendar.value.selection);
});
</script>

<template>
	<Main ref="main" class="diary" :class="{ 'edit-mode': editMode }" @detailVisibility="editMode = false">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar ref="calendar" storage="caldiary" :items="daysWithEntries" @selection="onDateSelected" />
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
					@click="calendar.onDay">
					<PieChart range="360" :value="day.value" :max="prefs.rdi.kcal">
						<template #details>
							<span>{{ day.weekday }}</span>
							<span>{{ day.calday }}</span>
						</template>
					</PieChart>
				</button>
			</section>
			<section class="charts-nutrient" id="charts-macro">
				<PieChart
					class="kcal"
					start="225"
					range="270"
					frac="0"
					:label="t('food.energy')"
					:unit="t('unit.cal')"
					:value="current ? current.kcal : 0"
					:max="prefs.rdi.kcal" />
				<PieChart
					class="fat"
					start="225"
					range="270"
					frac="0"
					:label="t('food.fat')"
					:unit="t('unit.g')"
					:value="current ? current.fat : 0"
					:max="prefs.rdi.fat" />
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.carbs2')"
					:unit="t('unit.g')"
					:value="current ? current.carb : 0"
					:max="prefs.rdi.carb" />
				<PieChart
					class="prot"
					start="225"
					range="270"
					frac="0"
					:label="t('food.protein')"
					:unit="t('unit.g')"
					:value="current ? current.prot : 0"
					:max="prefs.rdi.prot" />
			</section>
			<h2>{{ t("aria.headcarbcomp") }}</h2>
			<section class="charts-nutrient">
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.sugar')"
					:unit="t('unit.g')"
					:value="current ? current.sug : 0"
					:max="current ? current.carb : 0" />
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.fiber')"
					:unit="t('unit.g')"
					:value="current ? current.fib : 0"
					:max="current ? current.carb : 0"
					start2="180"
					range2="360"
					:label2="t('food.rdi')"
					:value2="current ? current.fib : 0"
					:max2="prefs.rdi.fib" />
			</section>
			<h3>{{ t("aria.headsugcomp") }}</h3>
			<section class="charts-nutrient">
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.fruc')"
					:unit="t('unit.g')"
					:value="current ? current.fruc : 0"
					:max="current ? current.sug : 0" />
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.gluc')"
					:unit="t('unit.g')"
					:value="current ? current.gluc : 0"
					:max="current ? current.sug : 0" />
				<PieChart
					class="carb"
					start="225"
					range="270"
					frac="0"
					:label="t('food.suc')"
					:unit="t('unit.g')"
					:value="current ? current.suc : 0"
					:max="current ? current.sug : 0" />
			</section>
			<h2>{{ t("aria.headfatcomp") }}</h2>
			<section class="charts-nutrient">
				<PieChart
					class="fat"
					start="225"
					range="270"
					frac="1"
					:label="t('food.fatsat2')"
					:unit="t('unit.g')"
					:value="current ? current.fatsat : 0"
					:max="current ? current.fat : 0"
					start2="180"
					range2="360"
					:label2="t('food.rdm')"
					:value2="current ? current.fatsat : 0"
					:max2="prefs.rdi.fatsat" />
				<PieChart
					class="fat"
					start="225"
					range="270"
					frac="1"
					:label="t('food.fato3')"
					:unit="t('unit.g')"
					:value="current ? current.fato3 : 0"
					:max="current ? current.fat : 0"
					start2="180"
					range2="360"
					:label2="t('food.rdi')"
					:value2="current ? current.fato3 : 0"
					:max2="prefs.rdi.fato3" />
				<PieChart
					class="fat"
					start="225"
					range="270"
					frac="1"
					:label="t('food.fato6')"
					:unit="t('unit.g')"
					:value="current ? current.fato6 : 0"
					:max="current ? current.fat : 0"
					start2="180"
					range2="360"
					:label2="t('food.rdi')"
					:value2="current ? current.fato6 : 0"
					:max2="prefs.rdi.fato6" />
			</section>
			<h2>{{ t("aria.headvits") }}</h2>
			<section class="charts-nutrient">
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="1"
					:label="t('food.vitashort')"
					:unit="t('unit.mg')"
					:value="current ? current.vita : 0"
					:max="prefs.rdi.vita" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="1"
					:label="t('food.vitb1short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb1 : 0"
					:max="prefs.rdi.vitb1" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="1"
					:label="t('food.vitb2short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb2 : 0"
					:max="prefs.rdi.vitb2" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="0"
					:label="t('food.vitb3short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb3 : 0"
					:max="prefs.rdi.vitb3" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="1"
					:label="t('food.vitb5short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb5 : 0"
					:max="prefs.rdi.vitb5" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="1"
					:label="t('food.vitb6short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb6 : 0"
					:max="prefs.rdi.vitb6" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="2"
					:label="t('food.vitb7short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb7 : 0"
					:max="prefs.rdi.vitb7" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="2"
					:label="t('food.vitb9short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb9 : 0"
					:max="prefs.rdi.vitb9" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="3"
					:label="t('food.vitb12short')"
					:unit="t('unit.mg')"
					:value="current ? current.vitb12 : 0"
					:max="prefs.rdi.vitb12" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="0"
					:label="t('food.vitcshort')"
					:unit="t('unit.mg')"
					:value="current ? current.vitc : 0"
					:max="prefs.rdi.vitc" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="3"
					:label="t('food.vitdshort')"
					:unit="t('unit.mg')"
					:value="current ? current.vitd : 0"
					:max="prefs.rdi.vitd" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="0"
					:label="t('food.viteshort')"
					:unit="t('unit.mg')"
					:value="current ? current.vite : 0"
					:max="prefs.rdi.vite" />
				<PieChart
					class="vit"
					start="225"
					range="270"
					frac="2"
					:label="t('food.vitkshort')"
					:unit="t('unit.mg')"
					:value="current ? current.vitk : 0"
					:max="prefs.rdi.vitk" />
			</section>
			<h2>{{ t("aria.headminerals") }}</h2>
			<section class="charts-nutrient">
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.calc')"
					:unit="t('unit.mg')"
					:value="current ? current.calc : 0"
					:max="prefs.rdi.calc" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.pot')"
					:unit="t('unit.mg')"
					:value="current ? current.pot : 0"
					:max="prefs.rdi.pot" />
				<PieChart
					v-if="prefs.ui.trackSaltAsSodium"
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.sod')"
					:unit="t('unit.mg')"
					:value="current ? current.sod : 0"
					:max="prefs.rdi.sod" />
				<PieChart
					v-if="!prefs.ui.trackSaltAsSodium"
					class="mins"
					start="225"
					range="270"
					frac="1"
					:label="t('food.salt')"
					:unit="t('unit.g')"
					:value="current ? current.salt : 0"
					:max="prefs.rdi.salt" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.mag')"
					:unit="t('unit.mg')"
					:value="current ? current.mag : 0"
					:max="prefs.rdi.mag" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="1"
					:label="t('food.iron')"
					:unit="t('unit.mg')"
					:value="current ? current.iron : 0"
					:max="prefs.rdi.iron" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.zinc')"
					:unit="t('unit.mg')"
					:value="current ? current.zinc : 0"
					:max="prefs.rdi.zinc" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.chl')"
					:unit="t('unit.mg')"
					:value="current ? current.chl : 0"
					:max="prefs.rdi.chl" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="0"
					:label="t('food.phos')"
					:unit="t('unit.mg')"
					:value="current ? current.phos : 0"
					:max="prefs.rdi.phos" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="1"
					:label="t('food.mang')"
					:unit="t('unit.mg')"
					:value="current ? current.mang : 0"
					:max="prefs.rdi.mang" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="2"
					:label="t('food.cop')"
					:unit="t('unit.mg')"
					:value="current ? current.cop : 0"
					:max="prefs.rdi.cop" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="3"
					:label="t('food.iod')"
					:unit="t('unit.mg')"
					:value="current ? current.iod : 0"
					:max="prefs.rdi.iod" />
				<!--<PieChart
					class="mins"
					start="225"
					range="270"
					frac="3"
					:label="t('food.chr')"
					:unit="t('unit.mg')"
					:value="current ? current.chr : 0"
					:max="prefs.rdi.chr" />
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="3"
					:label="t('food.mol')"
					:unit="t('unit.mg')"
					:value="current ? current.mol : 0"
					:max="prefs.rdi.mol" />-->
				<PieChart
					class="mins"
					start="225"
					range="270"
					frac="3"
					:label="t('food.sel')"
					:unit="t('unit.mg')"
					:value="current ? current.sel : 0"
					:max="prefs.rdi.sel" />
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
						<li :class="{ active: currentNutrient == (select ? select.value : '') }">
							<select ref="select" @change="onTabSelect">
								<option value="" selected disabled>{{ t("food.hintnut") }}</option>
								<optgroup :label="t('aria.headmacro3')">
									<option value="fib">{{ t("food.fiber") }}</option>
									<option value="sug">{{ t("food.sugar") }}</option>
									<option value="fruc">{{ t("food.fruc") }}</option>
									<option value="gluc">{{ t("food.gluc") }}</option>
									<option value="suc">{{ t("food.suc") }}</option>
									<option value="fatsat">{{ t("food.fatsat") }}</option>
									<option value="fato3">{{ t("food.fato3") }}</option>
									<option value="fato6">{{ t("food.fato6") }}</option>
								</optgroup>
								<optgroup :label="t('aria.headvits')">
									<option value="vita">{{ t("food.vita") }}</option>
									<option value="vitb1">{{ t("food.vitb1") }}</option>
									<option value="vitb2">{{ t("food.vitb2") }}</option>
									<option value="vitb3">{{ t("food.vitb3") }}</option>
									<option value="vitb5">{{ t("food.vitb5") }}</option>
									<option value="vitb6">{{ t("food.vitb6") }}</option>
									<option value="vitb7">{{ t("food.vitb7") }}</option>
									<option value="vitb9">{{ t("food.vitb9") }}</option>
									<option value="vitb12">{{ t("food.vitb12") }}</option>
									<option value="vitc">{{ t("food.vitc") }}</option>
									<option value="vitd">{{ t("food.vitd") }}</option>
									<option value="vite">{{ t("food.vite") }}</option>
									<option value="vitk">{{ t("food.vitk") }}</option>
								</optgroup>
								<optgroup :label="t('aria.headminerals')">
									<option value="calc">{{ t("food.calc") }}</option>
									<option value="pot">{{ t("food.pot") }}</option>
									<option v-if="prefs.ui.trackSaltAsSodium" value="sod">
										{{ t("food.sod") }}
									</option>
									<option v-if="!prefs.ui.trackSaltAsSodium" value="salt">
										{{ t("food.salt") }}
									</option>
									<option value="mag">{{ t("food.mag") }}</option>
									<option value="iron">{{ t("food.iron") }}</option>
									<option value="zinc">{{ t("food.zinc") }}</option>
									<option value="chl">{{ t("food.chl") }}</option>
									<option value="phos">{{ t("food.phos") }}</option>
									<option value="mang">{{ t("food.mang") }}</option>
									<option value="cop">{{ t("food.cop") }}</option>
									<option value="iod">{{ t("food.iod") }}</option>
									<option value="sel">{{ t("food.sel") }}</option>
								</optgroup>
							</select>
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
					:disabled="!editMode" />
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
	--color-vit: #ff4fdc;
	--color-vit-light: #ffc8f4;
	--color-mins: #b77400;
	--color-mins-light: #e6c58f;
}

main.diary .content > h2,
main.diary .content > h3 {
	text-align: center;
	margin-top: 1em;
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

main.neutral-charts #charts-week button.today .pie-chart circle.bad {
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

.charts-nutrient {
	display: flex;
	flex-wrap: wrap;
	justify-content: space-between;
	align-items: center;
	padding: 1em;
}

.charts-nutrient .pie-chart {
	flex-basis: 30%;
	margin: 0.5em 0;
	width: auto;
	height: auto;
}

.charts-nutrient .pie-chart figcaption {
	bottom: -2px;
}

.charts-nutrient .pie-chart > div span:last-child {
	color: var(--color-text-light);
}

#charts-macro .pie-chart.kcal {
	flex-basis: 100%;
	margin: 0 10%;
}

#charts-macro .pie-chart.kcal figcaption {
	bottom: 10%;
	font-size: 1.5rem;
}

#charts-macro .pie-chart.kcal circle.base {
	stroke: var(--color-kcal-light);
}

#charts-macro .pie-chart.kcal circle.good {
	stroke: var(--color-kcal);
}

.charts-nutrient .pie-chart.fat circle.base {
	stroke: var(--color-fat-light);
}

.charts-nutrient .pie-chart.fat circle.good {
	stroke: var(--color-fat);
}

main.neutral-charts .charts-nutrient .pie-chart.fat circle.bad {
	stroke: var(--color-fat);
}

.charts-nutrient .pie-chart.carb circle.base {
	stroke: var(--color-carb-light);
}

.charts-nutrient .pie-chart.carb circle.good {
	stroke: var(--color-carb);
}

main.neutral-charts .charts-nutrient .pie-chart.carb circle.bad {
	stroke: var(--color-carb);
}

.charts-nutrient .pie-chart.vit circle.base {
	stroke: var(--color-vit-light);
}

.charts-nutrient .pie-chart.vit circle.good {
	stroke: var(--color-vit);
}

main.neutral-charts .charts-nutrient .pie-chart.vit circle.bad {
	stroke: var(--color-vit);
}

.charts-nutrient .pie-chart.mins circle.base {
	stroke: var(--color-mins-light);
}

.charts-nutrient .pie-chart.mins circle.good {
	stroke: var(--color-mins);
}

main.neutral-charts .charts-nutrient .pie-chart.mins circle.bad {
	stroke: var(--color-mins);
}

#charts-macro .pie-chart.prot circle.base {
	stroke: var(--color-prot-light);
}

#charts-macro .pie-chart.prot circle.good {
	stroke: var(--color-prot);
}

main.neutral-charts #charts-macro .pie-chart.prot circle.bad {
	stroke: var(--color-prot);
}

@media only screen and (min-width: 800px) {
	#charts-week button {
		font-size: unset;
	}

	.charts-nutrient {
		justify-content: center;
	}

	.charts-nutrient .pie-chart {
		flex-basis: 20%;
		margin-left: 1em;
		margin-right: 1em;
	}

	.charts-nutrient .pie-chart figcaption {
		bottom: 0;
	}

	#charts-macro {
		flex-wrap: nowrap;
		justify-content: space-between;
	}

	#charts-macro .pie-chart {
		margin-left: 0;
		margin-right: 0;
	}

	#charts-macro .pie-chart.kcal {
		flex-basis: 30%;
		margin: 0;
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
	vertical-align: bottom;
	min-width: 8em;
}

#details section.tabs li:last-child {
	padding-right: 1rem;
}

#details section.tabs li button,
#details section.tabs li select {
	background: none;
	box-shadow: none !important;
	color: var(--color-placeholder);
	border-radius: 0;
	padding: 0.5em 0.5em 0.35em;
	border: none;
	border-bottom: 2px solid transparent;
	transition: color var(--transition-style), border-color var(--transition-style);
}

@media (hover: hover) {
	#details section.tabs li button:hover,
	#details section.tabs li select:hover {
		border-color: var(--color-text-light);
		box-shadow: none;
		color: var(--color-text-light);
	}
}

#details section.tabs li.active button,
#details section.tabs li.active select {
	border-color: var(--color-secondary);
	color: var(--color-text);
}

#details .diary-entry-list label {
	color: var(--color-text);
}
</style>

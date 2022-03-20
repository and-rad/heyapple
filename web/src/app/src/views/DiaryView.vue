<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import EntryList from "../components/DiaryEntryList.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, computed, inject } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime, Duration } from "luxon";

const { t } = useI18n();
const log = inject("log");
const diary = inject("diary");
const current = ref(null);
const currentDate = ref(DateTime.now());
const currentNutrient = ref("kcal");
const disableSave = ref(false);
const editMode = ref(false);
const main = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

let hasTabDrag = false;

function onTabSlide(evt) {
	evt.stopPropagation();
	moveTabBar(evt.target.closest("ul"), evt.movementX);
	hasTabDrag = true;
}

function onTabWheel(evt) {
	evt.stopPropagation();
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

	fetch("/api/v1/diary/" + date.replaceAll("-", "/"))
		.then((response) => response.json())
		.then((data) => {
			data.forEach((d) => (d.time = d.date.split("T")[1].slice(0, 5)));
			diary.value[date].entries = data;
		})
		.catch((err) => log.err(err))
		.finally(() => (current.value = diary.value[date]));
}
</script>

<template>
	<Main ref="main" class="diary">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar :items="daysWithEntries" @selection="onDateSelected" />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main> Diary </template>

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
				<EntryList ref="entries" :entries="current.entries" :nutrient="currentNutrient" :disabled="!editMode" />
			</section>
		</template>
	</Main>
</template>

<style>
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
	transition: color 0.2s, border-color 0.2s;
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
</style>

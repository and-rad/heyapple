<script setup>
import MoreImage from "./images/ImageMore.vue";
import { computed, ref, inject } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t } = useI18n();
const foods = inject("food");
const prefs = inject("prefs");
const prop = defineProps(["entries", "nutrient", "disabled"]);

/**
 * Controls how entries are grouped.
 * @values hour, day, meal, custom
 */
const groupMode = ref("hour");

/**
 * Controls whether nutrients are displayed in metric
 * numbers or fraction of the daily recommended dose.
 * @values metric, relative
 */
const nutrientMode = ref("relative");

const form = ref(null);

const groupedEntries = computed(() => {
	switch (groupMode.value) {
		case "hour":
			return groupedByHour();
		case "day":
			return groupedByDay();
		default:
			return [];
	}
});

const nutrientUnit = computed(() => {
	if (nutrientMode.value == "relative") {
		return "%";
	}
	switch (prop.nutrient) {
		case "kcal":
			return " " + t("unit.cal");
		case "fat":
		case "carb":
		case "prot":
			return " " + t("unit.g");
		default:
			return " " + t("unit.mg");
	}
});

/**
 * Returns the entries property grouped by time of day
 * in 2-hour increments, starting with even numbers.
 */
function groupedByHour() {
	let result = [];

	if (prop.entries) {
		let groups = {};
		prop.entries.forEach((entry) => {
			let hour = parseInt(entry.time);
			let start = DateTime.fromObject({ hours: hour - (hour % 2) });
			let name = start.toFormat("t") + " - " + start.plus({ hours: 2 }).toFormat("t");
			if (!(name in groups)) {
				groups[name] = { name: name, entries: [] };
			}

			let next = {
				id: entry.food.id,
				name: t(entry.food.id.toString()),
				amount: entry.food.amount,
				nutrient: getNutrient(entry.food),
				recipe: entry.recipe,
				time: entry.time,
				isrec: false,
			};

			if (entry.recipe) {
				let exists = false;
				groups[name].entries.every((e) => {
					if (e.isrec && e.name == entry.recipe) {
						e.entries.push(next);
						e.amount += next.amount;
						e.nutrient += next.nutrient;
						exists = true;
						return false;
					}
					return true;
				});
				if (!exists) {
					groups[name].entries.push({
						name: entry.recipe,
						isrec: true,
						amount: next.amount,
						nutrient: next.nutrient,
						entries: [next],
					});
				}
			} else {
				groups[name].entries.push(next);
			}
		});

		result = Object.values(groups);
		result.sort((a, b) => a.name < b.name);
	}

	return result;
}

function groupedByDay() {
	return [
		{ name: "Morning", entries: [{ type: "recipe" }, { type: "single" }] },
		{ name: "Noon", entries: [{ type: "recipe" }, { type: "single" }] },
	];
}

function getNutrient(food) {
	let data = foods.value.filter((f) => f.id == food.id)[0];
	let amount = data[prop.nutrient] * food.amount * 0.01;

	if (nutrientMode.value == "metric") {
		return Math.round(amount);
	}

	let rdi = prefs.value.rdi[prop.nutrient];
	return Math.round((amount * 100) / rdi);
}

function toggleNutrientMode() {
	if (nutrientMode.value == "relative") {
		nutrientMode.value = "metric";
	} else if (nutrientMode.value == "metric") {
		nutrientMode.value = "relative";
	}
}

function onRecipeDetails(evt) {
	let parent = evt.target.closest("div");
	let fields = parent.querySelector("fieldset");
	fields.style.setProperty("--max-height", fields.childElementCount * 41 + "px");
	parent.classList.toggle("open");
}

function onInput(evt) {
	evt.target.blur();
	let val = parseFloat(evt.target.value);
	if (isNaN(val) || val < 0) {
		evt.target.value = 0;
	}
}

function getDiff() {
	let data = new FormData(form.value);
	let ids = data.getAll("id");
	let amounts = data.getAll("amount");
	let recipes = data.getAll("recipe");
	let times = data.getAll("time");

	let result = [];
	prop.entries.forEach((entry) => {
		for (let i = 0; i < ids.length; ++i) {
			if (ids[i] != entry.food.id) {
				continue;
			}
			if (recipes[i] != entry.recipe) {
				continue;
			}
			if (times[i] != entry.time) {
				continue;
			}
			if (amounts[i] != entry.food.amount) {
				result.push({ id: ids[i], amount: amounts[i], recipe: recipes[i], time: times[i] });
			}
			break;
		}
	});
	return result;
}

defineExpose({ getDiff });
</script>

<template>
	<form class="diary-entry-list" ref="form">
		<fieldset v-for="group in groupedEntries" :key="group.name">
			<legend>{{ group.name }}</legend>
			<div v-for="entry in group.entries" :class="{ recipe: entry.isrec }">
				<template v-if="entry.isrec">
					<label>
						{{ entry.name }}
					</label>
					<button class="icon" type="button" @click="onRecipeDetails"><MoreImage /></button>
					<span>{{ entry.amount }}</span>
					<span class="unit">{{ t("unit.g") }}</span>
					<span class="nutrient" :class="[nutrient, nutrientMode]">
						{{ entry.nutrient }}{{ nutrientUnit }}
					</span>
					<fieldset :disabled="disabled">
						<div v-for="food in entry.entries" :key="food.id">
							<label>{{ food.name }}</label>
							<input type="number" name="amount" :value="food.amount" @change="onInput" />
							<span class="unit">{{ t("unit.g") }}</span>
							<span class="nutrient" :class="[nutrient, nutrientMode]">
								{{ food.nutrient }}{{ nutrientUnit }}
							</span>
							<input type="hidden" name="id" :value="food.id" />
							<input type="hidden" name="recipe" :value="food.recipe" />
							<input type="hidden" name="time" :value="food.time" />
						</div>
					</fieldset>
				</template>
				<template v-else>
					<label>{{ entry.name }}</label>
					<input type="number" name="amount" :value="entry.amount" :disabled="disabled" @change="onInput" />
					<span class="unit">{{ t("unit.g") }}</span>
					<span class="nutrient" :class="[nutrient, nutrientMode]">
						{{ entry.nutrient }}{{ nutrientUnit }}
					</span>
					<input type="hidden" name="id" :value="entry.id" />
					<input type="hidden" name="recipe" :value="entry.recipe" />
					<input type="hidden" name="time" :value="entry.time" />
				</template>
			</div>
		</fieldset>
		<button type="button" class="nutrient-mode-switch icon" @click="toggleNutrientMode">
			<span v-if="nutrientMode == 'relative'">%</span>
			<span v-if="nutrientMode == 'metric'">g</span>
		</button>
	</form>
</template>

<style>
.diary-entry-list {
	position: relative;
}

.diary-entry-list > fieldset {
	margin-top: 1em;
	white-space: nowrap;
}

.diary-entry-list fieldset legend {
	margin-bottom: 0 !important;
}

.diary-entry-list fieldset > div {
	display: flex;
	align-items: baseline;
	padding: 0.5em 0;
}

.diary-entry-list fieldset > div.recipe {
	flex-wrap: wrap;
}

.diary-entry-list .recipe label {
	flex-grow: 1;
	flex-basis: 40% !important;
}

.diary-entry-list .recipe label + button {
	align-self: center;
	padding: 0 0.25em;
	height: 24px;
	width: 26px;
}

.diary-entry-list .recipe label + button > svg {
	transform: rotate(90deg);
	fill: var(--color-text-light);
}

.diary-entry-list .recipe label + button + span {
	flex-grow: 0 !important;
}

.diary-entry-list .recipe fieldset {
	--max-height: 42px;
	flex-grow: 1;
	overflow: hidden;
	max-height: 0;
	padding-left: 1em !important;
	border-color: transparent !important;
	transition: 0.15s ease-in;
}

.diary-entry-list .recipe.open fieldset {
	margin: 0.5em 0;
	border-top: var(--border-light) !important;
	max-height: var(--max-height);
}

.diary-entry-list .nutrient {
	color: var(--color-text-light);
	/*color: var(--color-primary);*/
	margin-left: 1em;
	min-width: 3em;
	text-align: right;
}

.diary-entry-list .nutrient.metric {
	min-width: 4em;
}

button.nutrient-mode-switch {
	position: absolute;
	top: -2.75em;
	right: -0.5em;
	color: var(--color-primary);
	width: auto;
	height: auto;
	box-shadow: none !important;
}

button.nutrient-mode-switch:hover {
	color: var(--color-primary-dark);
}

button.nutrient-mode-switch span {
	font-weight: 700 !important;
}
</style>

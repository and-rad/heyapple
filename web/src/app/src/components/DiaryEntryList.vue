<script setup>
import MoreImage from "./images/ImageMore.vue";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t, locale } = useI18n();
const prop = defineProps(["entries", "disabled"]);

/*
 * Controls how entries are grouped.
 * @values hour, day, meal, custom
 */
const groupMode = ref("hour");

const form = ref(null);

const collator = new Intl.Collator(locale.value, { numeric: true });

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

/*
 * Returns the entries property grouped by time of day
 * in 2-hour increments, starting with even numbers.
 */
function groupedByHour() {
	let result = [];

	if (prop.entries) {
		let groups = {};
		prop.entries.forEach((entry) => {
			console.log(entry);
			let hour = parseInt(entry.time);
			let start = DateTime.fromObject({ hours: hour - (hour % 2) });
			let name = start.toFormat("t") + " - " + start.plus({ hours: 2 }).toFormat("t");
			if (!(name in groups)) {
				groups[name] = { name: name, entries: [] };
			}

			let next = {
				name: t(entry.food.id.toString()),
				amount: entry.food.amount,
				isrec: false,
				id: entry.food.id,
			};

			if (entry.recipe) {
				let exists = false;
				groups[name].entries.every((e) => {
					if (e.name == entry.recipe) {
						e.entries.push(next);
						e.amount += next.amount;
						exists = true;
						return false;
					}
				});
				if (!exists) {
					groups[name].entries.push({
						name: entry.recipe,
						isrec: true,
						amount: next.amount,
						entries: [next],
					});
				}
			} else {
				groups[name].entries.push(next);
			}
		});

		result = Object.values(groups);
		result.sort((a, b) => a.name < b.name);
		console.log(result);
	}

	return result;
}

function groupedByDay() {
	return [
		{ name: "Morning", entries: [{ type: "recipe" }, { type: "single" }] },
		{ name: "Noon", entries: [{ type: "recipe" }, { type: "single" }] },
	];
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

	let result = [];
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
					<fieldset :disabled="disabled">
						<div v-for="food in entry.entries" :key="food.id">
							<label>{{ food.name }}</label>
							<input type="number" name="amount" :value="food.amount" @change="onInput" />
							<span class="unit">{{ t("unit.g") }}</span>
							<input type="hidden" name="id" :value="food.id" />
						</div>
					</fieldset>
				</template>
				<template v-else>
					<label>{{ entry.name }}</label>
					<input type="number" name="amount" :value="entry.amount" :disabled="disabled" @change="onInput" />
					<span class="unit">{{ t("unit.g") }}</span>
					<input type="hidden" name="id" :value="entry.id" />
				</template>
			</div>
		</fieldset>
	</form>
</template>

<style>
.diary-entry-list > fieldset {
	margin-top: 1em;
}

.diary-entry-list fieldset legend {
	margin-bottom: 0 !important;
}

.diary-entry-list fieldset > div {
	display: flex;
	flex-wrap: wrap;
	align-items: baseline;
	padding: 0.5em 0;
}

.diary-entry-list .recipe label {
	flex-grow: 1;
}

.diary-entry-list .recipe label + button {
	align-self: center;
	padding: 0 0.25em;
	height: 24px;
	width: 28px;
}

.diary-entry-list .recipe label + button > svg {
	transform: rotate(90deg);
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
</style>

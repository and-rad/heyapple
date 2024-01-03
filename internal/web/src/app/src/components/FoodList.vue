<script setup>
import Arrow from "./images/ImageSortArrow.vue";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();
const prop = defineProps(["items"]);
const emit = defineEmits("selected");
const sortBy = ref("name");
const sortDir = ref("asc");

const collator = new Intl.Collator(locale.value, { numeric: true });

const categories = [
	{ cat: "name", name: t("food.name") },
	{ cat: "kcal", name: t("food.energy") },
	{ cat: "fat", name: t("food.fat") },
	{ cat: "carb", name: t("food.carbs") },
	{ cat: "prot", name: t("food.protein") },
];

const sortedItems = computed(() => {
	const factor = sortDir.value == "asc" ? 1 : -1;
	const val = sortBy.value;
	if (val == "name") {
		return [...prop.items].sort((a, b) => {
			return collator.compare(a[val], b[val]) * factor;
		});
	}

	return [...prop.items].sort((a, b) => {
		let amountA = a[val] / (a.size || 1);
		let amountB = b[val] / (b.size || 1);
		return collator.compare(amountA, amountB) * factor;
	});
});

function perServing(val, size, frac = 1) {
	size = size || 1;
	return +parseFloat(val / size).toFixed(frac);
}

function setSortCategory(cat, dir) {
	if (cat && dir) {
		sortBy.value = cat;
		sortDir.value = dir;
	} else if (sortBy.value == cat) {
		sortDir.value = sortDir.value == "asc" ? "desc" : "asc";
	} else {
		sortBy.value = cat;
	}
}

defineExpose({ setSortCategory, categories, sortBy, sortDir });
</script>

<template>
	<table>
		<thead>
			<tr :class="sortDir">
				<th class="name sort" :class="{ active: sortBy == 'name' }" @click="setSortCategory('name')">
					{{ t("food.name") }} <Arrow />
				</th>
				<th class="num sort" :class="{ active: sortBy == 'kcal' }" @click="setSortCategory('kcal')">
					<Arrow /> {{ t("food.energy") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'fat' }" @click="setSortCategory('fat')">
					<Arrow /> {{ t("food.fat") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'carb' }" @click="setSortCategory('carb')">
					<Arrow /> {{ t("food.carbs2") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'prot' }" @click="setSortCategory('prot')">
					<Arrow /> {{ t("food.protein") }}
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item in sortedItems" :key="item.id" @click="$emit('selected', item.id)">
				<td class="name">
					{{ item.name }}
					<div class="subtitle">
						{{ perServing(item.fat, item.size, 0) }}{{ t("unit.g") }} {{ t("food.fat") }},
						{{ perServing(item.carb, item.size, 0) }}{{ t("unit.g") }} {{ t("food.carbs2") }},
						{{ perServing(item.prot, item.size, 0) }}{{ t("unit.g") }} {{ t("food.protein") }}
					</div>
				</td>
				<td class="num">
					{{ perServing(item.kcal, item.size, 0) }} <span class="unit">{{ t("unit.cal") }}</span>
				</td>
				<td class="m num">
					{{ perServing(item.fat, item.size) }} <span class="unit">{{ t("unit.g") }}</span>
				</td>
				<td class="m num">
					{{ perServing(item.carb, item.size) }} <span class="unit">{{ t("unit.g") }}</span>
				</td>
				<td class="m num">
					{{ perServing(item.prot, item.size) }} <span class="unit">{{ t("unit.g") }}</span>
				</td>
			</tr>
		</tbody>
	</table>
</template>

<style></style>

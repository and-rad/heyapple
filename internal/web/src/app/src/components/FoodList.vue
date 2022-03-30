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

const sortedItems = computed(() => {
	if (sortDir.value == "asc") {
		return [...prop.items].sort((a, b) => {
			return collator.compare(a[sortBy.value], b[sortBy.value]);
		});
	} else {
		return [...prop.items].sort((a, b) => {
			return -collator.compare(a[sortBy.value], b[sortBy.value]);
		});
	}
});

function perServing(val, size, frac = 1) {
	if (size) {
		return +parseFloat(val / size).toFixed(frac);
	}
	return val;
}

function setActive(evt) {
	let cat = evt.target.dataset.sort;
	if (sortBy.value == cat) {
		sortDir.value = sortDir.value == "asc" ? "desc" : "asc";
	} else {
		sortBy.value = cat;
	}
}
</script>

<template>
	<table>
		<thead>
			<tr :class="sortDir">
				<th class="name sort" :class="{ active: sortBy == 'name' }" @click="setActive" data-sort="name">
					{{ t("food.name") }} <Arrow />
				</th>
				<th class="num sort" :class="{ active: sortBy == 'kcal' }" @click="setActive" data-sort="kcal">
					<Arrow /> {{ t("food.energy") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'fat' }" @click="setActive" data-sort="fat">
					<Arrow /> {{ t("food.fat") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'carb' }" @click="setActive" data-sort="carb">
					<Arrow /> {{ t("food.carbs2") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'prot' }" @click="setActive" data-sort="prot">
					<Arrow /> {{ t("food.protein") }}
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item in sortedItems" :key="item.id" @click="$emit('selected', item.id)">
				<td class="name">
					{{ item.name }}
				</td>
				<td class="num">
					{{ perServing(item.kcal, item.size) }} <span class="unit">{{ t("unit.cal") }}</span>
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

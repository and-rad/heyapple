<script setup>
import Arrow from "./images/ImageSortArrow.vue";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();
const prop = defineProps(["items"]);
const emit = defineEmits("selected");
const sortBy = ref("amount");
const sortDir = ref("desc");

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

function formattedAmount(item) {
	if (item.amount > 999) {
		var amount = +parseFloat(item.amount * 0.001).toFixed(1);
		var unit = t("unit.kg");
	} else {
		var amount = item.amount;
		var unit = t("unit.g");
	}
	return `${amount} <span class="unit">${unit}</span>`;
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
				<th class="num sort" :class="{ active: sortBy == 'amount' }" @click="setActive" data-sort="amount">
					<Arrow /> {{ t("food.amount") }}
				</th>
				<th class="name sort" :class="{ active: sortBy == 'name' }" @click="setActive" data-sort="name">
					{{ t("food.name") }} <Arrow />
				</th>
				<th class="m sort" :class="{ active: sortBy == 'aisle' }" @click="setActive" data-sort="aisle">
					<Arrow /> {{ t("food.aisle") }}
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item in sortedItems" :key="item.id">
				<td class="num" v-html="formattedAmount(item)"></td>
				<td class="name" @click="$emit('selected', item.id)">{{ item.name }}</td>
				<td class="m">{{ t("aisle." + item.aisle) }}</td>
			</tr>
		</tbody>
	</table>
</template>

<style></style>

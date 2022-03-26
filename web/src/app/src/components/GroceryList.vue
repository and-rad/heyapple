<script setup>
import Arrow from "./images/ImageSortArrow.vue";
import Checkbox from "./Checkbox.vue";
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

const allChecked = computed(() => prop.items.length && prop.items.length == prop.items.filter((i) => i.done).length);

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

function onCheckedAll(evt) {
	let val = evt.target.checked;
	prop.items.forEach((i) => (i.done = val));
}

function onChecked(evt) {
	let id = evt.target.closest("label").dataset.id;
	let val = evt.target.checked;
	prop.items.filter((i) => i.id == id)[0].done = val;
}
</script>

<template>
	<table>
		<thead>
			<tr :class="sortDir">
				<th class="select">
					<Checkbox :checked="allChecked" @click="onCheckedAll" />
				</th>
				<th class="num sort" :class="{ active: sortBy == 'amount' }" @click="setActive" data-sort="amount">
					<Arrow /> {{ t("food.amount") }}
				</th>
				<th class="name sort" :class="{ active: sortBy == 'name' }" @click="setActive" data-sort="name">
					{{ t("food.name") }} <Arrow />
				</th>
				<th class="s sort" :class="{ active: sortBy == 'aisle' }" @click="setActive" data-sort="aisle">
					<Arrow /> {{ t("food.aisle") }}
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item in sortedItems" :key="item.id" :class="{ done: item.done }">
				<td class="select">
					<Checkbox :data-id="item.id" :checked="item.done" @change="onChecked" />
				</td>
				<td class="num" v-html="formattedAmount(item)"></td>
				<td class="name" @click="$emit('selected', item.id)">{{ item.name }}</td>
				<td class="s">{{ t("aisle." + item.aisle) }}</td>
			</tr>
		</tbody>
	</table>
</template>

<style></style>

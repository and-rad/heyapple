<script setup>
import Arrow from "./images/ImageSortArrow.vue";
import { computed, ref } from "vue";

const prop = defineProps(["items"]);
const sortBy = ref("name");
const sortDir = ref("asc");

const sortedItems = computed(() => {
	if (sortDir.value == "asc") {
		return [...prop.items].sort((a, b) => {
			if (a[sortBy.value] < b[sortBy.value]) return -1;
			if (a[sortBy.value] > b[sortBy.value]) return 1;
			return 0;
		});
	} else {
		return [...prop.items].sort((a, b) => {
			if (a[sortBy.value] > b[sortBy.value]) return -1;
			if (a[sortBy.value] < b[sortBy.value]) return 1;
			return 0;
		});
	}
});

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
					{{ $t("food.name") }} <Arrow />
				</th>
				<th class="num sort" :class="{ active: sortBy == 'kcal' }" @click="setActive" data-sort="kcal">
					<Arrow /> {{ $t("food.energy") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'fat' }" @click="setActive" data-sort="fat">
					<Arrow /> {{ $t("food.fat") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'carb' }" @click="setActive" data-sort="carb">
					<Arrow /> {{ $t("food.carbs2") }}
				</th>
				<th class="m num sort" :class="{ active: sortBy == 'prot' }" @click="setActive" data-sort="prot">
					<Arrow /> {{ $t("food.protein") }}
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item in sortedItems" :key="item.id">
				<td class="name">
					{{ item.name }}
				</td>
				<td class="num">
					{{ item.kcal }} <span class="unit">{{ $t("unit.cal") }}</span>
				</td>
				<td class="m num">
					{{ item.fat }} <span class="unit">{{ $t("unit.g") }}</span>
				</td>
				<td class="m num">
					{{ item.carb }} <span class="unit">{{ $t("unit.g") }}</span>
				</td>
				<td class="m num">
					{{ item.prot }} <span class="unit">{{ $t("unit.g") }}</span>
				</td>
			</tr>
		</tbody>
	</table>
</template>

<style></style>

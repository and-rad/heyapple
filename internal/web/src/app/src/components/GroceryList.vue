<script setup>
import Arrow from "./images/ImageSortArrow.vue";
import Checkbox from "./Checkbox.vue";
import { computed, ref, inject, watch, onMounted } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");

const prop = defineProps(["items", "offline"]);
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

const allChecked = computed(() => {
	if (prop.items && prop.items.length) {
		return prop.items.length == prop.items.filter((i) => i.done).length;
	}
	return false;
});

watch(() => prop.offline, resync);

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

function resync() {
	if (!prop.offline) {
		let params = new URLSearchParams();
		prop.items.forEach((i) => {
			if (i.local) {
				params.append("id", i.id);
				params.append("done", i.done);
			}
		});

		if (params.has("id")) {
			fetch("/api/v1/list/diary/done", {
				method: "PUT",
				body: params,
				headers: {
					"Content-Type": "application/x-www-form-urlencoded",
					"X-CSRF-Token": csrf,
				},
			})
				.then((response) => {
					if (response.ok) {
						prop.items.forEach((i) => (i.local = false));
					}
				})
				.catch(() => {});
		}
	}
}

function onCheckedAll(evt) {
	let val = evt.target.checked;
	let params = new URLSearchParams();
	prop.items.forEach((i) => {
		i.done = val;
		i.local = prop.offline;
		params.append("id", i.id);
		params.append("done", i.done);
	});

	if (!prop.offline) {
		fetch("/api/v1/list/diary/done", {
			method: "PUT",
			body: params,
			headers: {
				"Content-Type": "application/x-www-form-urlencoded",
				"X-CSRF-Token": csrf,
			},
		}).catch(() => {});
	}
}

function onChecked(evt) {
	let id = evt.target.closest("label").dataset.id;
	let val = evt.target.checked;
	let item = prop.items.filter((i) => i.id == id)[0];
	item.done = val;
	item.local = prop.offline;

	if (!prop.offline) {
		fetch("/api/v1/list/diary/done", {
			method: "PUT",
			body: new URLSearchParams({ id: id, done: val }),
			headers: {
				"Content-Type": "application/x-www-form-urlencoded",
				"X-CSRF-Token": csrf,
			},
		}).catch(() => {});
	}
}

defineExpose({ setSortCategory });
onMounted(resync);
</script>

<template>
	<table>
		<thead>
			<tr :class="sortDir">
				<th class="select">
					<Checkbox :checked="allChecked" @click="onCheckedAll" />
				</th>
				<th class="num sort" :class="{ active: sortBy == 'amount' }" @click="setSortCategory('amount')">
					<Arrow /> {{ t("food.amount") }}
				</th>
				<th class="name sort" :class="{ active: sortBy == 'name' }" @click="setSortCategory('name')">
					{{ t("food.name") }} <Arrow />
				</th>
				<th class="s sort" :class="{ active: sortBy == 'aisle' }" @click="setSortCategory('aisle')">
					{{ t("food.aisle") }} <Arrow />
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item in sortedItems" :key="item.id" :class="{ done: item.done }">
				<td class="select">
					<Checkbox :data-id="item.id" :checked="item.done" @change="onChecked" />
				</td>
				<td class="num" v-html="formattedAmount(item)"></td>
				<td class="name" @click="$emit('selected', item.id)">
					<span>{{ item.name }}</span>
					<div class="subtitle">{{ t("aisle." + item.aisle) }}</div>
				</td>
				<td class="s">{{ t("aisle." + item.aisle) }}</td>
			</tr>
		</tbody>
	</table>
</template>

<style></style>

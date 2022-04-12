<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import FoodList from "../components/GroceryList.vue";
import { ref, computed, inject, watch, onBeforeMount, onMounted, onUnmounted } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const log = inject("log");
const diary = inject("diary");

const filtered = ref([]);
const offline = ref(false);

const main = ref(null);
const calendar = ref(null);
const list = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

let pingIntervalHandle = undefined;

watch(filtered, (val) => saveItemsLocal(), { deep: true });

function onDateSelected(dates) {
	if (dates.length == 0) {
		filtered.value = [];
		return;
	}

	if (offline.value) {
		return;
	}

	let params = new URLSearchParams(dates.map((d) => ["date", d]));
	fetch("/api/v1/list/diary?" + params)
		.then((response) => {
			if (!response.ok) {
				throw t("getlist.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			data.forEach((d) => {
				d.name = t(d.id.toString());
				let old = filtered.value.filter((i) => i.id == d.id)[0];
				if (old && old.local) {
					d.local = true;
					d.done = old.done;
				}
			});
			filtered.value = data;
		})
		.catch((err) => log.err(err));
}

function onOffline() {
	if (!offline.value) {
		offline.value = true;
		log.warn(t("conn.off"));
	}
}

function onOnline() {
	if (offline.value) {
		offline.value = false;
		log.msg(t("conn.on"));
	}
}

function onSort(evt) {
	let [cat, dir] = evt.target.value.split(" ");
	list.value.setSortCategory(cat, dir);
	evt.target.selectedIndex = 0;
}

function showDetails(id) {
	main.value.showDetails();
}

function saveItemsLocal() {
	let str = JSON.stringify(filtered.value);
	window.localStorage.setItem("listdiary", str);
}

function loadItemsLocal() {
	let str = window.localStorage.getItem("listdiary") || "[]";
	filtered.value = JSON.parse(str);
}

function ping() {
	return fetch("/ping", { method: "HEAD" }).then(onOnline, onOffline);
}

onBeforeMount(() => {
	loadItemsLocal();
});

onMounted(() => {
	pingIntervalHandle = setInterval(ping, 15000);
	ping().then(() => onDateSelected(calendar.value.selection));
});

onUnmounted(() => {
	clearInterval(pingIntervalHandle);
});
</script>

<template>
	<Main ref="main">
		<template #filter>
			<section :class="{ offline: offline }">
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar
					ref="calendar"
					mode="toggle"
					storage="calshop"
					:items="daysWithEntries"
					@selection="onDateSelected"
				/>
			</section>
			<hr />
			<section></section>
		</template>

		<template #controls>
			<select class="sort s" @change="onSort">
				<option value="" disabled selected hidden>{{ t("sort.hint") }}</option>
				<option value="aisle asc">{{ t("food.aisle") }} {{ t("sort.asc") }}</option>
				<option value="aisle desc">{{ t("food.aisle") }} {{ t("sort.desc") }}</option>
			</select>
			<span class="spacer"></span>
		</template>

		<template #main>
			<FoodList ref="list" :items="filtered" :offline="offline" @selected="showDetails" />
		</template>
	</Main>
</template>

<style>
section.offline {
	pointer-events: none;
	position: relative;
	opacity: 0.3;
}

section.offline:after {
	content: "";
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	position: absolute;
	background-image: var(--icon-wifi-off);
	background-size: 75%;
	background-position: center;
	background-repeat: no-repeat;
	opacity: 0.2;
}
</style>

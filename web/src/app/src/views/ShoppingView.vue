<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import FoodList from "../components/GroceryList.vue";
import { ref, computed, inject, onMounted, onUnmounted } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const log = inject("log");
const diary = inject("diary");

const filtered = ref([]);
const offline = ref(false);

const main = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

let pingIntervalHandle = undefined;

function onDateSelected(dates) {
	if (dates.length == 0) {
		filtered.value = [];
		saveItemsLocal();
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
			data.forEach((d) => (d.name = t(d.id.toString())));
			filtered.value = data;
			saveItemsLocal();
		})
		.catch((err) => {
			if (err instanceof TypeError) {
				onOffline();
			} else {
				log.err(err);
			}
		});
}

function onOffline() {
	if (!offline.value) {
		offline.value = true;
		log.warn(t("conn.off"));
		loadItemsLocal();
	}
}

function onOnline() {
	if (offline.value) {
		offline.value = false;
		log.msg(t("conn.on"));
	}
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

onMounted(() => {
	pingIntervalHandle = setInterval(() => {
		fetch("/ping", { method: "HEAD" }).then(onOnline, onOffline);
	}, 15000);
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
				<Calendar mode="toggle" storage="calshop" :items="daysWithEntries" @selection="onDateSelected" />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main>
			<FoodList :items="filtered" @selected="showDetails" />
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

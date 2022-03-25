<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import FoodList from "../components/GroceryList.vue";
import { ref, computed, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const log = inject("log");
const diary = inject("diary");

const filtered = ref([]);

const main = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

function onDateSelected(dates) {
	if (dates.length == 0) {
		filtered.value = [];
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
		})
		.catch((err) => log.err(err));
}

function showDetails(id) {
	main.value.showDetails();
}
</script>

<template>
	<Main ref="main">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar mode="toggle" :items="daysWithEntries" @selection="onDateSelected" />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main>
			<FoodList :items="filtered" @selected="showDetails" />
		</template>
	</Main>
</template>

<style></style>

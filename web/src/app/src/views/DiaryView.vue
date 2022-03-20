<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import EntryList from "../components/DiaryEntryList.vue";
import { ref, computed, inject } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime, Duration } from "luxon";

const { t } = useI18n();
const log = inject("log");
const diary = inject("diary");
const current = ref(null);
const currentDate = ref(DateTime.now());
const editMode = ref(false);
const main = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

function onDateSelected(dates) {
	currentDate.value = DateTime.fromISO(dates[0]);

	let date = currentDate.value.toISODate();
	if (!diary.value[date]) {
		current.value = { date: date };
		return;
	}
	if (diary.value[date].entries) {
		current.value = diary.value[date];
		return;
	}

	fetch("/api/v1/diary/" + date.replaceAll("-", "/"))
		.then((response) => response.json())
		.then((data) => {
			data.forEach((d) => (d.time = d.date.split("T")[1].slice(0, 5)));
			diary.value[date].entries = data;
		})
		.catch((err) => log.err(err))
		.finally(() => (current.value = diary.value[date]));
}
</script>

<template>
	<Main ref="main" class="diary">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar :items="daysWithEntries" @selection="onDateSelected" />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main> Diary </template>

		<template #head-details>
			<h2 class="no-edit-mode">{{ currentDate.weekdayLong }}</h2>
		</template>

		<template #details>
			<section class="subtitle no-edit-mode">
				{{ currentDate.toLocaleString(DateTime.DATE_FULL) }}
			</section>
			<hr />
			<section v-if="current">
				<h2>{{ t("aria.headlog") }}</h2>
				<p class="msg-noitems" v-if="!current.entries" v-html="t('diary.noitems')"></p>
				<EntryList ref="entries" :entries="current.entries" :disabled="!editMode" />
			</section>
		</template>
	</Main>
</template>

<style>
main.diary #details section.subtitle {
	padding-bottom: 0.5em;
}
</style>

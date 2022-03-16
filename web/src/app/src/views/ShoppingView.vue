<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import { ref, computed, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const diary = inject("diary");
const main = ref(null);

const daysWithEntries = computed(() => Object.keys(diary.value));

function onDateSelected(dates) {
	console.log(dates);
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

		<template #main> Shopping Lists </template>
	</Main>
</template>

<style></style>

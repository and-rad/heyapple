<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const main = ref(null);

function onDateSelected(dates) {
	console.log(dates);
}
</script>

<template>
	<Main ref="main">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar @selection="onDateSelected" />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main> Diary </template>
	</Main>
</template>

<style></style>

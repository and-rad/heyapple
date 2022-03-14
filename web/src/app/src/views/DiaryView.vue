<script setup>
import Main from "../components/Main.vue";
import Calendar from "../components/Calendar.vue";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const main = ref(null);
</script>

<template>
	<Main ref="main">
		<template #filter>
			<section>
				<h2>{{ t("aria.headcal") }}</h2>
				<Calendar />
			</section>
			<hr />
			<section></section>
		</template>

		<template #main> Diary </template>
	</Main>
</template>

<style></style>

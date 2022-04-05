<script setup>
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();
const prop = defineProps(["item"]);

const collator = new Intl.Collator(locale.value, { numeric: true });
const tagMax = 4096;

const sortedTags = computed(() => {
	let result = [];
	for (let i = 1; i <= tagMax; i *= 2) {
		if ((prop.item.flags & i) == i) {
			result.push({ id: i, name: t("tag." + i) });
		}
	}
	return result.sort((a, b) => collator.compare(a.name, b.name));
});
</script>

<template>
	<span class="tag" v-for="tag in sortedTags" :key="tag.id">{{ tag.name }}</span>
</template>

<style></style>

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
	<span class="tag" v-for="tag in sortedTags" :key="tag.id" :class="'tag-' + tag.id">{{ tag.name }}</span>
</template>

<style>
.tags .tag {
	display: inline-block;
	padding: 2px 4px;
	min-width: 5em;
	margin: 0 0.25em;
	font-size: 12px;
	font-weight: bold;
	text-align: center;
	border-radius: 4px;
	border: var(--border);
}

.tags .tag-1,
.tags .tag-2 {
	border-color: #388e3c;
	color: #388e3c;
}

.tags .tag-4,
.tags .tag-8 {
	border-color: #303f9f;
	color: #303f9f;
}

span.tag-16,
span.tag-256 {
	border-color: #0288d1;
	color: #0288d1;
}

.tags .tag-32,
.tags .tag-64,
.tags .tag-128 {
	border-color: #7b1fa2;
	color: #7b1fa2;
}

.tags .tag-512,
.tags .tag-1024 {
	border-color: #6d4c41;
	color: #6d4c41;
}

.tags .tag-2048 {
	border-color: #f57c00;
	color: #f57c00;
}

.tags .tag-4096 {
	border-color: #d32f2f;
	color: #d32f2f;
}
</style>

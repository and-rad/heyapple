<script setup>
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";

const { locale } = useI18n();
const prop = defineProps(["label", "placeholder", "items"]);
const emit = defineEmits(["confirm"]);
const value = ref(0);

const collator = new Intl.Collator(locale.value, { numeric: true });

const sortedItems = computed(() => {
	return [...prop.items].sort((a, b) => collator.compare(a.name, b.name));
});

function confirm(evt) {
	evt.preventDefault();
	if (value.value) {
		emit("confirm", value.value);
		value.value = 0;
	}
}
</script>

<template>
	<form class="clickable-select">
		<select v-model="value">
			<option value="0" selected>{{ placeholder }}</option>
			<option v-for="item in sortedItems" :key="item.id" :value="item.id">{{ item.name }}</option>
		</select>
		<input type="submit" @click="confirm" :value="label" />
	</form>
</template>

<style>
.clickable-select {
	display: flex;
}

.clickable-select select {
	flex-grow: 1;
	flex-basis: 50%;
	border-top-right-radius: 0;
	border-bottom-right-radius: 0;
}

.clickable-select input[type="submit"] {
	width: auto;
	border-top-left-radius: 0;
	border-bottom-left-radius: 0;
	margin-left: -1px;
	background-color: var(--color-primary);
}
</style>

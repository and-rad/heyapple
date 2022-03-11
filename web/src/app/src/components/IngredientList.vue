<script setup>
import { computed, ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();
const prop = defineProps(["items", "disabled"]);
const foods = inject("food");
const sortBy = ref("name");

const collator = new Intl.Collator(locale.value, { numeric: true });

const sortedItems = computed(() => {
	let items = prop.items.map((i) => ({ id: i.id, amount: i.amount, name: t(i.id.toString()) }));
	return items.sort((a, b) => collator.compare(a[sortBy.value], b[sortBy.value]));
});

function onInput(evt) {
	evt.target.blur();
	let val = parseFloat(evt.target.value);
	if (isNaN(val) || val < 0) {
		evt.target.value = 0;
	}
}
</script>

<template>
	<form class="ingredients">
		<fieldset :disabled="disabled">
			<div v-for="item in sortedItems" :key="item.id">
				<label>{{ item.name }}</label>
				<input type="number" name="amount" :value="item.amount" @change="onInput" />
				<span class="unit">{{ t("unit.g") }}</span>
				<input type="hidden" name="id" :value="item.id" />
			</div>
		</fieldset>
	</form>
</template>

<style>
.ingredients fieldset > div {
	display: flex;
	align-items: baseline;
	padding: 0.5em 0;
}
</style>

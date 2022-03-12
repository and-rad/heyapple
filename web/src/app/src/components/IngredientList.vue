<script setup>
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();
const prop = defineProps(["items", "disabled"]);
const sortBy = ref("name");
const form = ref(null);

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

function getDiff() {
	let data = new FormData(form.value);
	let ids = data.getAll("id");
	let amounts = data.getAll("amount");

	let result = [];
	prop.items.forEach((item) => {
		let idx = parseInt(ids.indexOf(item.id.toString()));
		let amount = parseFloat(amounts[idx]);
		if (amount != item.amount) {
			result.push({ id: item.id, amount: amount });
		}
	});
	return result;
}

defineExpose({ getDiff });
</script>

<template>
	<form class="ingredients" ref="form">
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

<script setup>
import { computed, ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const prop = defineProps(["items"]);
const foods = inject("food");
const sortBy = ref("name");

const collator = new Intl.Collator(useI18n().locale.value, { numeric: true });

const sortedItems = computed(() => {
	let ids = prop.items.map((i) => i.id);
	let current = foods.value.filter((f) => ids.includes(f.id));
	return current.sort((a, b) => {
		return collator.compare(a[sortBy.value], b[sortBy.value]);
	});
});
</script>

<template>
	<form class="ingredients">
		<fieldset disabled>
			<div v-for="item in sortedItems" :key="item.id">
				<label>{{ item.name }}</label>
				<input type="text" name="amount" :value="item.amount" />
				<span class="unit">{{ $t("unit.g") }}</span>
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

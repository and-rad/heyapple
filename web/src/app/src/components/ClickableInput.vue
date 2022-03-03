<script setup>
import { ref } from "vue";

const prop = defineProps(["label"]);
const emit = defineEmits(["confirm"]);
const value = ref("");

function confirm(evt) {
	evt.preventDefault();
	emit("confirm", value.value);
	value.value = "";
}
</script>

<template>
	<form class="clickable-input">
		<input type="text" v-model="value" :placeholder="$t('filter.namehint')" />
		<input type="submit" @click="confirm" :value="label" />
	</form>
</template>

<style>
.clickable-input {
	display: flex;
}

.clickable-input input[type="text"] {
	flex-grow: 1;
	flex-basis: 50%;
	border-top-right-radius: 0;
	border-bottom-right-radius: 0;
}

.clickable-input input[type="submit"] {
	width: auto;
	border-top-left-radius: 0;
	border-bottom-left-radius: 0;
	margin-left: -1px;
}

</style>

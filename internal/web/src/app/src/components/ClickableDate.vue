<script setup>
const prop = defineProps(["label", "date", "time", "disabled"]);
const emit = defineEmits(["confirm"]);

function confirm(evt) {
	evt.preventDefault();
	let data = new FormData(evt.target.closest("form"));
	emit("confirm", data.get("date"), data.get("time"));
}
</script>

<template>
	<form class="clickable-date">
		<input type="date" name="date" :value="date" />
		<input type="time" name="time" :value="time" />
		<button type="submit" class="async" @click="confirm" :disabled="disabled">{{ label }}</button>
	</form>
</template>

<style>
.clickable-date {
	display: flex;
}

.clickable-date input[type="date"],
.clickable-date input[type="time"] {
	flex-grow: 1;
	flex-basis: 50%;
	border-top-right-radius: 0;
	border-bottom-right-radius: 0;
}

.clickable-date input[type="time"] {
	border-radius: 0;
	border-left: none;
	border-right: none;
	flex-basis: 35%;
}

.clickable-date button {
	width: auto;
	border-top-left-radius: 0;
	border-bottom-left-radius: 0;
	background-color: var(--color-primary);
}
</style>

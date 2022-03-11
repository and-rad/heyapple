<script setup>
import { ref, computed } from "vue";

const prop = defineProps(["label", "date", "time"]);
const emit = defineEmits(["confirm"]);

function confirm(evt) {
	evt.preventDefault();
	let data = new FormData(evt.target.closest("form"));
	emit("confirm", data.get("date") + "T" + data.get("time"));
}
</script>

<template>
	<form class="clickable-date">
		<input type="date" name="date" :value="date" />
		<input type="time" name="time" :value="time" />
		<input type="submit" @click="confirm" :value="label" />
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

.clickable-date input[type="submit"] {
	width: auto;
	border-top-left-radius: 0;
	border-bottom-left-radius: 0;
}
</style>

<script setup>
const prop = defineProps(["placeholder"]);
const emit = defineEmits(["result"]);

var timeoutHandle = undefined;

function confirm(evt) {
	evt.preventDefault();
	clearTimeout(timeoutHandle);
	timeoutHandle = setTimeout(function () {
		fetchData(evt.target.closest("form"));
	}, 500);
}

function fetchData(form) {
	let params = new URLSearchParams(new FormData(form));
	fetch(form.action + "?" + params).then((response) => {
		response.json().then((data) => emit("result", response.status, data));
	});
}
</script>

<template>
	<form>
		<input type="text" name="name" autocomplete="off" :placeholder="placeholder" @input="confirm" />
		<slot :confirm="confirm">Additional filters</slot>
	</form>
</template>

<style></style>

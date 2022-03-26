<script setup>
const prop = defineProps(["checked"]);
const emit = defineEmits(["change", "click"]);

function onChange(evt) {
	emit("change", evt);
}

function onClick(evt) {
	emit("click", evt);
}
</script>

<template>
	<label>
		<input type="checkbox" :checked="checked" @change="onChange" @click="onClick" />
		<div></div>
	</label>
</template>

<style>
label > input[type="checkbox"] {
	display: none;
}

label > input[type="checkbox"] + div {
	width: 1.1em;
	height: 1.1em;
	border: var(--border);
	border-radius: 2px;
	background-repeat: no-repeat;
	background-size: 1em;
	background-position: center;
	transition: var(--transition-style);
}

@media (hover: hover) {
	label:hover > input[type="checkbox"] + div {
		border-color: var(--color-primary);
	}
}

label > input:checked + div {
	border-color: var(--color-primary);
	background-color: var(--color-primary);
	background-image: var(--icon-check);
}
</style>

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
	<label class="checkbox">
		<input type="checkbox" :checked="checked" @change="onChange" @click="onClick" />
		<div></div>
		<slot></slot>
	</label>
</template>

<style>
label.checkbox {
	display: flex;
	align-items: center;
}

label.checkbox > input {
	display: none;
}

label.checkbox > input + div {
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
	label.checkbox:hover > input + div {
		border-color: var(--color-primary);
	}
}

label.checkbox > input:checked + div {
	border-color: var(--color-primary);
	background-color: var(--color-primary);
	background-image: var(--icon-check);
}
</style>

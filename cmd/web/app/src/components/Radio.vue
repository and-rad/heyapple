<script setup>
const prop = defineProps(["checked", "name", "value"]);
const emit = defineEmits(["change", "click"]);

function onChange(evt) {
	emit("change", evt);
}

function onClick(evt) {
	emit("click", evt);
}
</script>

<template>
	<label class="radio">
		<input type="radio" :checked="checked" :name="name" :value="value" @change="onChange" @click="onClick" />
		<div></div>
		<slot></slot>
	</label>
</template>

<style>
label.radio {
	display: flex;
	align-items: center;
}

label.radio > input {
	display: none;
}

label.radio > input + div {
	width: 1.1em;
	height: 1.1em;
	border: var(--border);
	border-radius: 50%;
	background-repeat: no-repeat;
	background-size: 0.9em;
	background-position: center;
}

@media (hover: hover) {
	label.radio:hover > input + div {
		border-color: var(--color-primary);
	}
}

label.radio > input:checked + div {
	border-color: var(--color-primary);
	background-color: var(--color-primary);
	background-image: var(--icon-check);
}
</style>

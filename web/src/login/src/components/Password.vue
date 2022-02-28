<script setup>
import { reactive } from "vue";

const prop = defineProps(["modelValue", "name", "withBar"]);
const emit = defineEmits(["update:modelValue"]);
const scoreColor = ["#e53935", "#FB8C00", "#FDD835", "#7CB342", "#00897B"];
const barStyle = reactive({ width: "0", background: "transparent" });

function onInput(evt) {
	const pw = evt.target.value;
	emit("update:modelValue", pw);
	updatePasswordStrength(pw);
}

function updatePasswordStrength(password) {
	let score = zxcvbn(password).score;
	barStyle.width = password.length > 0 ? 20 + score * 20 + "%" : "0%";
	barStyle.background = scoreColor[score];
}
</script>

<template>
	<span>
		<input type="password" :name="name" :value="modelValue" @input="onInput" />
		<div class="password-strength-bar" v-if="withBar" :style="barStyle"></div>
	</span>
</template>

<style>
.password-strength-bar {
	background: green;
	height: 4px;
	width: 100%;
	margin-top: 4px;
	transition: 0.2s ease-in-out;
}
</style>

<script setup>
import VisibleImage from "./images/ImageVisible.vue";
import HiddenImage from "./images/ImageHidden.vue";
import { reactive, ref } from "vue";

const prop = defineProps(["modelValue", "name", "withBar"]);
const emit = defineEmits(["update:modelValue"]);
const scoreColor = ["#e53935", "#FB8C00", "#FDD835", "#7CB342", "#00897B"];
const barStyle = reactive({ width: "0", background: "transparent" });
const fieldType = ref("password");

function onInput(evt) {
	prop.modelValue = evt.target.value;
	emit("update:modelValue", prop.modelValue);
	updatePasswordStrength(prop.modelValue);
}

function updatePasswordStrength(password) {
	let score = zxcvbn(password).score;
	barStyle.width = password.length > 0 ? 20 + score * 20 + "%" : "0%";
	barStyle.background = scoreColor[score];
}

function toggle() {
	fieldType.value = fieldType.value == "password" ? "text" : "password";
}

function hide() {
	fieldType.value = "password";
}

function show() {
	fieldType.value = "text";
}

defineExpose({ hide, show });
</script>

<template>
	<span class="password">
		<input :type="fieldType" :name="name" :value="modelValue" @input="onInput" />
		<div class="strength-bar" v-if="withBar" :style="barStyle"></div>
		<button class="visibility icon" type="button" @click="toggle">
			<VisibleImage v-if="fieldType == 'password'" />
			<HiddenImage v-if="fieldType == 'text'" />
		</button>
	</span>
</template>

<style>
.password {
	position: relative;
	display: inline-block;
	width: 100%;
}

.password .strength-bar {
	height: 4px;
	width: 0;
	margin-top: 4px;
	background: transparent;
	transition: var(--transition-move);
}

.password button.visibility {
	position: absolute;
	top: 0;
	right: 0;
}

.password button.visibility svg {
	fill: var(--color-text-light);
	transition: var(--transition-style);
}

@media (hover: hover) {
	.password button.visibility:hover {
		box-shadow: none;
	}

	.password button.visibility:hover svg {
		fill: var(--color-text);
	}
}
</style>

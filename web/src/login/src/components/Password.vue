<script setup>
const prop = defineProps(['modelValue']);
const emit = defineEmits(['update:modelValue']);
const scoreColor = ["#e53935", "#FB8C00", "#FDD835", "#7CB342", "#00897B"];

function onInput(evt) {
    const pw = evt.target.value;
    emit('update:modelValue', pw);
    updatePasswordStrength(pw);
}

function updatePasswordStrength(password) {
	let score = zxcvbn(password).score;
	console.log(score);
}
</script>

<template>
    <input type="password" :value="modelValue" @input="onInput" >
</template>

<style>

</style>
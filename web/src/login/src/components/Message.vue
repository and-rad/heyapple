<script setup>
import { ref, computed } from "vue";

const props = defineProps({ msg: String, level: Number });

const hasMessage = computed(() => {
	return props.msg;
});
</script>

<template>
	<div class="message" :class="[{ hidden: !hasMessage }, level]">{{ msg }}</div>
</template>

<style>
.message {
	padding: 0.5em;
	background-color: var(--color-good-light);
	border: 1px solid var(--color-good);
	border-radius: 4px;
	margin: 1em 0;
}

.message.hidden {
	display: none;
}

.message.err {
	background-color: var(--color-bad-light);
	border-color: var(--color-bad);
}

.message.warn {
	background-color: var(--color-warn-light);
	border-color: var(--color-warn);
}
</style>

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
	--color-bad: #db5343;
	--color-bad-light: rgba(219, 83, 67, 0.3);
	--color-good: #21ab53;
	--color-good-light: rgba(33, 171, 83, 0.1);
	--color-warn: #ffc107;
	--color-warn-light: rgba(255, 193, 7, 0.1);

	padding: 0.5em;
	background-color: var(--color-good-light);
	border: 1px solid var(--color-good);
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

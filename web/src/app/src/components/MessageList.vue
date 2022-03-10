<script setup>
import Message from "./Message.vue";
import { onMounted, ref } from "vue";

const messages = ref([]);
let nextId = 0;

function onMessage(evt) {
	messages.value.push({
		id: nextId++,
		type: evt.type,
		msg: evt.detail.msg,
		time: evt.detail.timeout,
	});
}

function onTimeout(id) {
	messages.value = messages.value.filter((m) => m.id != id);
}

onMounted(() => {
	window.addEventListener("message", onMessage);
	window.addEventListener("warning", onMessage);
	window.addEventListener("error", onMessage);
});
</script>

<template>
	<div id="messages">
		<Message v-for="msg in messages" :key="msg.id" :msg="msg" @timeout="onTimeout" />
	</div>
</template>

<style>
#messages {
	position: fixed;
	z-index: 1000;
	width: 100%;
	max-width: 480px;
	right: 0;
	bottom: 0.5em;
}
</style>

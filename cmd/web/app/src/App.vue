<script setup>
import Header from "./components/Header.vue";
import MessageList from "./components/MessageList.vue";
import { RouterView } from "vue-router";
import { onMounted } from "vue";

onMounted(() => {
	document.querySelector("body > .spinner-container").remove();
});
</script>

<template>
	<Header />
	<RouterView />
	<MessageList />
</template>

<style>
@import "@/assets/base.css";

html,
body {
	overflow-y: hidden;
	height: 100%;
}
</style>

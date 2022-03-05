<script setup>
import { RouterView } from "vue-router";
import Header from "./components/Header.vue";
import { onMounted } from "vue";

onMounted(() => {
	document.querySelector("body > .spinner-container").remove();
});
</script>

<template>
	<Header />
	<RouterView />
</template>

<style>
@import "@/assets/base.css";

html,
body {
	overflow-y: hidden;
	height: 100%;
}
</style>

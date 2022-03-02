<script setup>
import { RouterView } from "vue-router";
import Header from "./components/Header.vue";
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

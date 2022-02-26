<script setup>
import { RouterView } from "vue-router";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

</script>

<template>
	<Header />
	<RouterView />
	<Footer />
</template>

<style>
@import "@/assets/base.css";

html {
	scroll-behavior: smooth;
}

#app {
	display: flex;
	flex-direction: column;
	min-height: 100vh;
}

main {
	flex-grow: 1;
	margin-top: 4em;
}
</style>

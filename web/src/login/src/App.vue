<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
	<RouterView />
</template>

<style>
@import "@/assets/base.css";

#app {
	width: 1080px;
	max-width: 1080px;
	height: 720px;
	max-height: 720px;
	margin: 0 auto;
	transform: translateY(-50%) translateX(-50%);
	background-color: #fff;
	border-radius: 4px;
	box-shadow: 0 0 24px #ccc;
	overflow-y: hidden;
	display: flex;
	justify-content: space-between;
	position: absolute;
	top: 50%;
	left: 50%;
}

#app > * {
	flex-basis: 50%;
	position: relative;
}

#app > .image svg {
	position: absolute;
	bottom: 0;
}
</style>

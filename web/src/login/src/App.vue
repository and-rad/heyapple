<script setup>
import { RouterView } from "vue-router";
</script>

<template>
	<RouterView />
</template>

<style>
@import "../../app/src/assets/base.css";

body {
	display: flex;
	align-items: center;
}

#app {
	padding: 1em;
	width: 100%;
	max-width: 640px;
	margin: 0 auto;
}

h1 {
	margin-bottom: 1em;
}

label {
	margin-top: 1em;
	display: flex;
	justify-content: space-between;
}

input[type="submit"] {
	margin-top: 2em;
}

form footer {
	margin-top: 0.5em;
	display: flex;
	justify-content: space-between;
}

header {
	display: flex;
	align-items: center;
	margin-bottom: 1em;
}

header #logo {
	height: 3em;
	width: 3em;
}

header #app-name {
	margin-left: 0.5em;
	font-size: 2em;
	text-transform: uppercase;
	flex-grow: 1;
	color: var(--color-primary);
}

header #app-name span:last-child {
	font-weight: 300;
}
</style>

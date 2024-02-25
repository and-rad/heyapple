<script setup>
import Footer from "./components/Footer.vue";
import { RouterView } from "vue-router";
</script>

<template>
	<div>
		<main>
			<RouterView />
		</main>
	</div>
	<Footer />
</template>

<style>
@import "../../app/src/assets/base.css";

#app {
	width: 100%;
	min-height: 100vh;
	margin: 0 auto;
	display: flex;
	flex-direction: column;
	hyphens: auto;
}

#app > div {
	flex-grow: 1;
	display: flex;
	align-items: center;
}

main {
	width: 100%;
	max-width: 560px;
	margin: 0 auto;
	padding: 1em;
}

h1 {
	margin-bottom: 1em;
}

label {
	margin-top: 1em;
	display: flex;
	justify-content: space-between;
}

input[type="submit"],
button[type="submit"] {
	margin: 2em 0 1em;
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

<script setup>
import { RouterView } from "vue-router";
</script>

<template>
	<RouterView />
</template>

<style>
@import "../../app/src/assets/base.css";

#app > *:first-child {
	padding: 3em 1em 1em;
}

#app > .image {
	display: none;
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

@media screen and (min-width: 720px) {
	body {
		background: var(--color-background);
	}

	#app {
		width: calc(100% - 48px);
		height: calc(100% - 48px);
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

	#app > *:first-child {
		flex-basis: 60%;
		padding: 3em;
	}

	#app > *:nth-child(2) {
		flex-basis: 40%;
	}

	#app > .image {
		display: block;
	}

	#app > .image svg {
		position: absolute;
		bottom: 0;
	}
}

@media screen and (min-width: 1000px) {
	#app {
		width: 960px;
		max-width: calc(100% - 48px);
		height: 640px;
		max-height: calc(100% - 48px);
	}
}
</style>

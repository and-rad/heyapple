<script setup>
import MenuImage from "./images/ImageMenu.vue";
import { ref } from "vue";

const filter = ref("");
const details = ref("");

function toggleFilter() {
	if (filter.value == "") {
		filter.value = "open-filter";
		details.value = "";
	} else {
		filter.value = "";
	}
}

function toggleDetails() {
	if (details.value == "") {
		details.value = "open-details";
		filter.value = "";
	} else {
		details.value = "";
	}
}
</script>

<template>
	<main :class="[filter, details]">
		<div id="filter">
			<section>This is the "Create New" widget</section>
			<section>
				<slot name="filter"> This is the main search & filter area</slot>
			</section>
		</div>

		<div id="main">
			<div class="controls">
				<button @click="toggleFilter" class="open-close"><MenuImage /></button>
			</div>
			<slot name="main">This is the main area</slot>
		</div>

		<div id="details">
			<slot name="details">This is the details area</slot>
			<section></section>
		</div>
	</main>
</template>

<style>
#filter,
#details,
#main {
	position: fixed;
	top: 4em;
	bottom: 0;
	overflow-y: auto;
	background-color: var(--color-background);
}

#filter {
	width: 300px;
	left: -300px;
	border-right: var(--border-light);
	white-space: nowrap;
	user-select: none;
	transition: left 0.25s;
}

main.open-filter #filter {
	left: 0;
}

#details {
	right: -500px;
	width: 480px;
	max-width: 100%;
	box-shadow: var(--shadow-menu);
	transition: right 0.25s;
}

main.open-details #details {
	right: 0;
}

#main {
	left: 0;
	right: 0;
	transition: 0.25s;
}

main.open-filter #main {
	transform: translateX(300px);
}

#main .controls {
	display: flex;
	align-items: center;
	justify-content: space-between;
	padding: 0.25em;
	user-select: none;
}

#main .controls button {
	width: auto;
}

button.open-close {
	height: 2.5em;
	width: 2.5em;
	background: none;
}

button.open-close > svg {
	width: 1.5em;
	height: 1.5em;
	fill: var(--color-primary);
}

/*
button.open-close > div {
	width: 1.5em;
	height: 2px;
	background-color: #aaa;
	transform-origin: center;
	transition: 0.5s ease-in-out;
}

button.open-close > div:nth-child(2) {
	margin: 6px auto;
}

main.open-filter button.open-close div:first-child {
	background-color: var(--color-secondary);
	transform: translateY(10px) rotate(45deg);
}

main.open-filter button.open-close div:nth-child(2) {
	opacity: 0;
}

main.open-filter button.open-close div:last-child {
	background-color: var(--color-secondary);
	transform: translateY(-9px) rotate(-45deg);
}
*/

/* screen size medium */
@media only screen and (min-width: 800px) {
	#filter {
		left: 0;
	}

	#main {
		left: 300px;
		transform: none !important;
	}

	#main .controls button.open-close {
		display: none;
	}
}

/* screen size large */
@media only screen and (min-width: 1280px) {
	#main {
		right: 480px;
	}

	#details {
		right: 0;
		border-left: var(--border-light);
		box-shadow: none;
	}
}
</style>

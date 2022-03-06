<script setup>
import MenuImage from "./images/ImageMenu.vue";
import DetailsImage from "./images/ImageHeaderMono.vue";
import MoreImage from "./images/ImageMore.vue";
import BackArrow from "./images/ImageRightArrow.vue";
import { ref, watch } from "vue";

const prop = defineProps(["current"]);
const filter = ref("");
const details = ref("");

watch(
	() => prop.current,
	(newValue) => {
		if (newValue) {
			showDetails();
		}
	}
);

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

function showDetails() {
	if (details.value == "") {
		details.value = "open-details";
		filter.value = "";
	}
}
</script>

<template>
	<main :class="[filter, details]">
		<div id="filter">
			<slot name="filter"> This is the main search & filter area</slot>
		</div>

		<div id="main">
			<div class="controls">
				<button @click="toggleFilter" class="open-filter icon"><MenuImage /></button>
				<span class="spacer"></span>
				<button @click="toggleDetails" class="open-details icon"><MoreImage /></button>
			</div>
			<div class="content">
				<slot name="main">This is the main area</slot>
			</div>
		</div>

		<div id="details">
			<div class="controls">
				<h2 v-if="current">{{ current.name }}</h2>
				<span class="spacer"></span>
				<button @click="toggleDetails" class="open-details icon"><BackArrow /></button>
			</div>
			<slot name="details">
				<div class="placeholder">
					<DetailsImage />
					<p>{{ $t("details.noitem") }}</p>
				</div>
			</slot>
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
	overflow-x: hidden;
	user-select: none;
	background-color: var(--color-background);
}

#details > section,
#filter > section {
	padding: 0.5em;
	border-bottom: var(--border-light);
}

#details fieldset legend,
#filter fieldset legend,
#details > section > h2,
#filter > section > h2 {
	margin-bottom: 1rem;
	font-size: 12px;
	font-weight: 700;
	color: var(--color-text);
	text-transform: uppercase;
}

#details fieldset legend,
#filter fieldset legend {
	color: var(--color-text-light);
}

#details > section:last-child,
#filter > section:last-child {
	border: none;
}

#filter {
	width: 300px;
	left: -300px;
	padding-bottom: 2em;
	border-right: var(--border-light);
	white-space: nowrap;
	transition: left 0.25s;
}

main.open-filter #filter {
	left: 0;
}

#filter > section.new-item {
	height: 105px;
}

#filter .slider {
	margin-top: 2em;
}

#filter fieldset {
	margin-top: 3em;
}

#filter fieldset .slider:first-of-type {
	margin-top: 0;
}

#details {
	right: -500px;
	width: 480px;
	max-width: 100%;
	box-shadow: var(--shadow-menu);
	transition: right 0.25s;
	display: flex;
	flex-direction: column;
}

main.open-details #details {
	right: 0;
}

#details .placeholder {
	width: 100%;
	height: 100%;
	position: relative;
	font-size: 2em;
	color: var(--color-placeholder);
	text-align: center;
}

#details .placeholder > * {
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translateX(-50%) translateY(-50%);
	width: 66%;
}

#details .placeholder svg {
	fill: var(--color-placeholder);
	opacity: 0.2;
}

#details .controls {
	align-items: unset;
}

#details .controls h2 {
	font-size: 2em;
	text-overflow: ellipsis;
	overflow-x: hidden;
	margin-right: 1em;
}

#details .tag {
	display: inline-block;
	padding: 2px 4px;
	min-width: 5em;
	margin: 0 0.25em;
	font-size: 12px;
	font-weight: bold;
	text-align: center;
	background-color: var(--color-secondary);
	color: #fff;
	border-radius: 4px;
}

#details label,
#details .unit {
	color: var(--color-text-light);
}

#details input:disabled,
#details input:read-only {
	border-color: transparent !important;
	background: none !important;
}

#details input:disabled + .unit,
#details input:read-only + .unit {
	border-color: transparent !important;
}

#filter fieldset,
#details fieldset {
	padding: 0;
	border: none;
}

#details fieldset > div {
	display: flex;
	align-items: baseline;
	padding: 0.5em 0;
}

#details fieldset label {
	flex-basis: 60%;
	margin-right: 0.5em;
}

#details fieldset input[type="text"] {
	flex-basis: 2.5em;
	flex-grow: 1;
	padding: 0 0.25em 0 0;
	border: none;
	border-radius: 0;
	border-bottom: var(--border);
	text-align: right;
}

#details fieldset .unit {
	border-bottom: var(--border);
}

#details fieldset input:active + .unit,
#details fieldset input:focus + .unit {
	border-color: var(--color-primary);
}

#main {
	left: 0;
	right: 0;
	display: flex;
	flex-direction: column;
	transition: 0.25s;
}

main.open-filter #main {
	transform: translateX(300px);
}

main .controls {
	display: flex;
	align-items: center;
	justify-content: space-between;
	white-space: nowrap;
	padding: 0.5em;
	min-height: 3.5em;
	user-select: none;
}

#main .controls button.open-details > svg {
	transform: rotate(90deg);
}

main .controls button {
	width: auto;
}

main .controls button.icon {
	height: 2.5em;
	width: 2.5em;
	background: none;
}

main .controls button > svg {
	width: 1.5em;
	height: 1.5em;
	fill: var(--color-primary);
}

#main .content {
	flex-grow: 1;
	overflow: auto;
}

#main table {
	width: 100%;
	border-collapse: separate;
	border-spacing: 0;
	white-space: nowrap;
}

#main td,
#main th {
	padding: 1em 0.5em;
	text-align: left;
	cursor: default;
	overflow: hidden;
	max-width: 6em;
}

#main td.name {
	text-overflow: ellipsis;
}

#main th {
	border-bottom: var(--border-light);
	color: var(--color-text-light);
	padding-top: 0.5em;
}

#main th.sort {
	cursor: pointer;
}

#main th.sort .icon {
	opacity: 0;
	transition: opacity 0.2s;
}

#main th.sort.active .icon {
	opacity: 0.4;
}

#main thead th {
	background-color: var(--color-background);
	position: sticky;
	z-index: 1;
	top: 0;
}

#main table .num {
	text-align: right;
	width: 6em;
}

#main table .num .unit {
	color: var(--color-text-light);
}

#main table .l,
#main table .m,
#main table .s {
	display: none;
}

@media (hover: hover) {
	#main tbody tr:hover {
		background-color: #f8f8f8;
	}

	#main th.sort:hover .icon {
		opacity: 0.5;
	}
}

/* screen size small */
@media only screen and (min-width: 400px) {
	#main table .s {
		display: table-cell;
	}
}

/* screen size smallish */
@media only screen and (min-width: 640px) {
	#main table .m {
		display: table-cell;
	}
}

/* screen size medium */
@media only screen and (min-width: 800px) {
	#filter {
		left: 0;
	}

	#main {
		left: 300px;
		transform: none !important;
	}

	main .controls button.open-filter {
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

	main .controls button.open-details {
		display: none;
	}
}
</style>

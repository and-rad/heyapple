<script setup>
import MenuImage from "./images/ImageMenu.vue";
import DetailsImage from "./images/ImageHeaderMono.vue";
import MoreImage from "./images/ImageMore.vue";
import BackImage from "./images/ImageRightArrow.vue";
import { ref, inject, computed } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const prefs = inject("prefs");
const emit = defineEmits(["detailVisibility"]);
const filter = ref(false);
const details = ref(false);

const mainClass = computed(() => ({
	"open-filter": filter.value,
	"open-details": details.value,
	"neutral-charts": prefs.value.ui.neutralCharts,
}));

function toggleFilter() {
	if (!filter.value) {
		filter.value = true;
		details.value = false;
		emit("detailVisibility");
	} else {
		filter.value = false;
	}
}

function showDetails() {
	emit("detailVisibility");
	if (!details.value) {
		details.value = true;
		filter.value = false;
	}
}

function hideDetails() {
	emit("detailVisibility");
	if (details.value) {
		details.value = false;
	}
}

defineExpose({ showDetails, hideDetails });
</script>

<template>
	<main :class="mainClass">
		<div id="filter">
			<slot name="filter"> This is the main search & filter area</slot>
		</div>

		<div id="main">
			<div class="controls">
				<button @click="toggleFilter" class="open-filter icon"><MenuImage /></button>
				<slot name="controls"><span class="spacer"></span></slot>
				<button @click="showDetails" class="open-details icon"><MoreImage /></button>
			</div>
			<div class="content">
				<slot name="main">This is the main area</slot>
			</div>
		</div>

		<div id="details">
			<slot name="details">
				<div class="controls">
					<span class="spacer"></span>
					<button @click="hideDetails" class="open-details icon"><BackImage /></button>
				</div>
				<div class="placeholder">
					<DetailsImage />
					<p>{{ t("details.noitem") }}</p>
				</div>
			</slot>
		</div>
	</main>
</template>

<style>
:root {
	--icon-sort: url(data:image/svg+xml;base64,PHN2ZyBpZD0ic3ZnODc1IiB2ZXJzaW9uPSIxLjEiIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KIDxwYXRoIGlkPSJwYXRoMTA0NyIgZD0ibTAgM3YyaDI0di0yaC0yNHptMCA4djJoMTZ2LTJoLTE2em0wIDh2Mmg4di0yaC04eiIgZmlsbD0iIzEzYWQ3MyIvPgo8L3N2Zz4=);
}

main.no-dt #details,
main.no-dt button.open-details {
	display: none;
}

main.no-ctrl #main > .controls {
	display: none;
}

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
	transition: opacity var(--transition-style);
}

#details fieldset legend,
#filter fieldset legend,
#details > section > h2,
#filter > section > h2 {
	margin-bottom: 1rem;
	padding: 0;
	font-size: 12px;
	font-weight: 700;
	color: var(--color-text);
	text-transform: uppercase;
}

#details fieldset legend,
#filter fieldset legend {
	color: var(--color-text-light);
}

#filter {
	width: 300px;
	left: -300px;
	padding-bottom: 2em;
	border-right: var(--border-light);
	white-space: nowrap;
	transition: left var(--transition-move);
}

main.open-filter #filter {
	left: 0;
}

#filter > section.new-item {
	height: 104px;
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
	transition: right var(--transition-move);
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
	height: auto;
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

#details label,
#details .unit {
	color: var(--color-text-light);
}

#details label + span,
#details label + button + span,
#details input:disabled,
#details input[type="text"]:read-only,
#details input[type="number"]:read-only {
	border-color: transparent !important;
	background: none !important;
}

#details input:disabled + .unit,
#details input[type="text"]:read-only + .unit,
#details input[type="number"]:read-only + .unit {
	border-color: transparent !important;
}

#filter fieldset,
#details fieldset {
	padding: 0;
	border: none;
}

#details .col50 > div {
	display: flex;
	align-items: baseline;
	padding: 0.5em 0;
}

#details label {
	/*flex-basis: 60%;
	width: 5em;*/
	margin-right: 0.5em;
	white-space: nowrap;
	text-overflow: ellipsis;
	overflow-x: hidden;
}

#details label + span,
#details label + button + span,
#details input[type="text"],
#details input[type="number"] {
	flex-basis: 2.5em;
	flex-grow: 1;
	padding: 0 0.25em 0 0;
	border: none;
	border-radius: 0;
	border-bottom: var(--border);
	text-align: right;
}

#details label + span:last-child,
#details label + button + span:last-child,
#details input[type="text"]:last-child,
#details input[type="number"]:last-child {
	padding: 0;
}

#details input + .unit {
	border-bottom: var(--border);
}

#details input:active + .unit,
#details input:focus + .unit {
	border-color: var(--color-primary);
}

#details .controls form {
	margin-right: 1em;
}

#details .controls input[type="text"] {
	text-align: left;
	text-overflow: ellipsis;
	font-size: 2em;
}

#main {
	left: 0;
	right: 0;
	display: flex;
	flex-direction: column;
	transition: var(--transition-move);
}

main.open-filter #main {
	transform: translateX(300px);
}

main .no-edit-mode {
	transition: opacity var(--transition-style);
}

main.edit-mode .no-edit-mode {
	pointer-events: none;
	opacity: 0.2;
}

button.cancel-edit-mode > svg {
	transition: fill 0.25s ease-in;
}

main.edit-mode button.cancel-edit-mode > svg {
	fill: var(--color-bad);
}

main .controls {
	display: flex;
	align-items: center;
	white-space: nowrap;
	padding: 0.5em;
	min-height: 3.5em;
	user-select: none;
}

main .controls .sort-menu .options {
	top: 3em;
}

#main .controls .sort-menu.open .options {
	left: -2.5em;
}

#main .controls button.open-details > svg {
	transform: rotate(90deg);
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

#main table.s thead,
#main table.m thead,
#main table.l thead {
	display: none;
}

#main tr {
	transition: var(--transition-style);
}

#main tr.done {
	opacity: 0.3;
	background-color: #f0f0f0;
	color: var(--color-text-light);
}

#main td,
#main th {
	padding: 1em 0.5em;
	text-align: left;
	cursor: default;
	overflow: hidden;
	text-overflow: ellipsis;
	vertical-align: top;
	max-width: 6em;
}

#main th.select,
#main td.select {
	width: 3em;
	padding-left: 1em;
	padding-right: 1em;
	vertical-align: middle;
	text-align: center;
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
	transition: opacity var(--transition-style);
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

#main table .subtitle {
	color: var(--color-text-light);
	font-size: 0.8em;
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
	#main table.s thead {
		display: table-header-group;
	}

	#main table .s {
		display: table-cell;
	}

	#main table.s .subtitle {
		display: none;
	}

	main .controls .sort-menu.s {
		display: none;
	}
}

/* screen size smallish */
@media only screen and (min-width: 640px) {
	#main table.m thead {
		display: table-header-group;
	}

	#main table .m {
		display: table-cell;
	}

	#main table.m .subtitle {
		display: none;
	}

	main .controls .sort-menu.m {
		display: none;
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

	main .controls .sort-menu.l {
		display: none;
	}

	#main .controls .sort-menu.open .options {
		left: 0;
	}

	#main table.l .subtitle {
		display: none;
	}

	#main table.l thead {
		display: table-header-group;
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

	main.edit-mode .controls button.cancel-edit-mode {
		display: block;
	}
}
</style>

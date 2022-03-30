<script setup>
import { computed, ref } from "vue";

const radius = 112;

const prop = defineProps(["start", "range", "value", "max", "label", "unit", "frac"]);

const formattedValue = computed(() => +(parseFloat(prop.value) || 0).toFixed(prop.frac || 0));
const formattedMax = computed(() => +(parseFloat(prop.max) || 0).toFixed(prop.frac || 0));

const transform = computed(() => {
	let rot = -90 + (parseFloat(prop.start) || 0);
	return `rotate(${rot} 128 128)`;
});

const fullArc = computed(() => {
	let ratio = (parseFloat(prop.range) || 0) / 360;
	return Math.PI * 2 * radius * ratio + " 710";
});

const valueArc = computed(() => {
	let normalized = (parseFloat(prop.value) || 0) / (parseFloat(prop.max) || 1);
	let range = parseFloat(prop.range) || 0;
	let ratio = (Math.min(normalized, 1) * range) / 360;
	return Math.PI * 2 * radius * ratio + " 710";
});

const overArc = computed(() => {
	let normalized = (parseFloat(prop.value) || 0) / (parseFloat(prop.max) || 1);
	let range = parseFloat(prop.range) || 0;
	let ratio = (Math.min(Math.max(0, normalized - 1), 1) * range) / 360;
	return Math.PI * 2 * radius * ratio + " 710";
});
</script>

<template>
	<figure class="pie-chart">
		<svg viewBox="0 0 256 256">
			<circle class="base" r="112" cx="128" cy="128" :transform="transform" :stroke-dasharray="fullArc" />
			<circle class="good" r="112" cx="128" cy="128" :stroke-dasharray="valueArc" :transform="transform" />
			<circle class="bad" r="112" cx="128" cy="128" :stroke-dasharray="overArc" :transform="transform" />
		</svg>
		<div>
			<slot name="details">
				<span>{{ formattedValue }} {{ unit }}</span>
				<hr />
				<span>{{ formattedMax }} {{ unit }}</span>
			</slot>
		</div>
		<figcaption>{{ label }}</figcaption>
	</figure>
</template>

<style>
.pie-chart {
	white-space: nowrap;
	position: relative;
	width: 4em;
	height: 4em;
}

.pie-chart > svg {
	display: block;
}

.pie-chart hr {
	border-bottom: var(--border);
}

.pie-chart circle {
	stroke-width: 12;
	fill: none;
	transition: all 0.5s ease-in-out, fill 0.25s;
}

.pie-chart circle.base {
	stroke: var(--color-primary-light);
}

.pie-chart circle.good {
	stroke: var(--color-primary);
}

.pie-chart circle.bad {
	stroke: var(--color-bad);
	stroke-width: 13;
	transition: 0.5s 0.5s ease-in-out;
}

main.neutral-charts .pie-chart circle.bad {
	stroke: var(--color-primary);
	stroke-width: 32;
}

.pie-chart circle.soft {
	stroke: #ccc;
}

.pie-chart figcaption {
	position: absolute;
	left: 0;
	right: 0;
	bottom: 0;
	text-align: center;
}

.pie-chart > div {
	position: absolute;
	left: 50%;
	top: 50%;
	transform: translateX(-50%) translateY(-50%);
	min-width: 50%;
	max-width: 75%;
	text-align: center;
}
.pie-chart > div > * {
	display: block;
}
</style>

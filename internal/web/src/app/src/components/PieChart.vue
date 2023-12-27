<script setup>
import { computed, ref } from "vue";

const radius = 112;

const prop = defineProps([
	"start",
	"range",
	"value",
	"max",
	"label",
	"unit",
	"frac",
	"start2",
	"range2",
	"value2",
	"max2",
	"label2",
]);

const formattedValue = computed(() => +(parseFloat(prop.value) || 0).toFixed(prop.frac || 0));
const formattedMax = computed(() => +(parseFloat(prop.max) || 0).toFixed(prop.frac || 0));

const transform = computed(() => {
	return computeTransform(prop.start);
});

const transform2 = computed(() => {
	return computeTransform(prop.start2);
});

const fullArc = computed(() => {
	return computeFullArc(prop.range);
});

const fullArc2 = computed(() => {
	return computeFullArc(prop.range2);
});

const valueArc = computed(() => {
	return computeValueArc(prop.value, prop.max, prop.range);
});

const valueArc2 = computed(() => {
	return computeValueArc(prop.value2, prop.max2, prop.range2);
});

const overArc = computed(() => {
	return computeOverArc(prop.value, prop.max, prop.range);
});

const overArc2 = computed(() => {
	return computeOverArc(prop.value2, prop.max2, prop.range2);
});

function computeTransform(startProp) {
	let rot = -90 + (parseFloat(startProp) || 0);
	return `rotate(${rot} 128 128)`;
}

function computeFullArc(rangeProp) {
	let ratio = (parseFloat(rangeProp) || 0) / 360;
	return Math.PI * 2 * radius * ratio + " 710";
}

function computeValueArc(valProp, maxProp, rangeProp) {
	let normalized = (parseFloat(valProp) || 0) / (parseFloat(maxProp) || 1);
	let range = parseFloat(rangeProp) || 0;
	let ratio = (Math.min(normalized, 1) * range) / 360;
	return Math.PI * 2 * radius * ratio + " 710";
}

function computeOverArc(valProp, maxProp, rangeProp) {
	let normalized = (parseFloat(valProp) || 0) / (parseFloat(maxProp) || 1);
	let range = parseFloat(rangeProp) || 0;
	let ratio = (Math.min(Math.max(0, normalized - 1), 1) * range) / 360;
	return Math.PI * 2 * radius * ratio + " 710";
}
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

		<figure v-if="prop.label2" class="pie-chart secondary">
			<svg viewBox="0 0 256 256">
				<circle class="base" r="112" cx="128" cy="128" :transform="transform2" :stroke-dasharray="fullArc2" />
				<circle class="good" r="112" cx="128" cy="128" :stroke-dasharray="valueArc2" :transform="transform2" />
				<circle class="bad" r="112" cx="128" cy="128" :stroke-dasharray="overArc2" :transform="transform2" />
			</svg>
			<div>{{ label2 }}</div>
		</figure>
	</figure>
</template>

<style>
.pie-chart {
	white-space: nowrap;
	position: relative;
	width: 4em;
	height: 4em;
}

.pie-chart.secondary {
	position: absolute;
	margin: 0 !important;
	top: 0 !important;
	right: 0 !important;
	width: 35% !important;
	height: 35% !important;
	background: #fff;
	border-radius: 50%;
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

.pie-chart.secondary circle {
	stroke-width: 18;
}

.pie-chart circle.base {
	stroke: var(--color-primary-light);
}

.pie-chart circle.good {
	stroke: var(--color-primary);
}

.pie-chart circle.bad {
	stroke: var(--color-bad);
	stroke-width: 16;
	transition: 0.5s 0.5s ease-in-out;
}

.pie-chart.pie-chart.secondary circle.bad {
	stroke-width: 22;
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

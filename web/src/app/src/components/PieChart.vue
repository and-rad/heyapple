<script setup></script>

<template>
	<figure class="pie-chart">
		<svg viewBox="0 0 256 256">
			<circle r="112" cx="128" cy="128" />
			<circle class="good" r="112" cx="128" cy="128" stroke-dasharray="216 710" transform="rotate(16 128 128)" />
		</svg>
		<figcaption>123</figcaption>
	</figure>
</template>

<style>
.pie-chart {
	position: relative;
    width: 4em;
    height: 4em;
}

.pie-chart circle {
	stroke: var(--color-good-light);
	stroke-width: 12;
	fill: none !important;
}

.pie-chart circle.good {
	stroke: var(--color-good);
}

.pie-chart circle.bad {
	stroke: var(--color-bad);
}

.pie-chart circle.soft {
	stroke: #ccc;
}

.pie-chart figcaption {
	position: absolute;
	left: 0;
	right: 0;
	top: 50%;
	transform: translateY(-50%);
	text-align: center;
}
</style>

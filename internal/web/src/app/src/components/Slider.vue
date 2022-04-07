<script setup>
import { ref, watch } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const prop = defineProps(["label", "name", "unit", "min", "max", "frac"]);
const emit = defineEmits(["input"]);

const minVal = ref(parseFloat(prop.min).toFixed(prop.frac));
const maxVal = ref(parseFloat(prop.max).toFixed(prop.frac));
const minPercent = ref(0);
const maxPercent = ref(100);

watch(
	() => prop.min,
	(val, old) => setMin(val)
);

watch(
	() => prop.max,
	(val, old) => setMax(val)
);

function setMin(value) {
	let min = parseFloat(prop.min) || 0;
	let val = parseFloat(value) || min;
	minVal.value = val;
	maxVal.value = Math.max(minVal.value, maxVal.value);

	let max = parseFloat(prop.max) || 0;
	minPercent.value = ((val - min) * 100) / (max - min);
	maxPercent.value = Math.max(minPercent.value, maxPercent.value);
}

function setMax(value) {
	let max = parseFloat(prop.max) || 0;
	let val = parseFloat(value) || max;
	maxVal.value = val;
	minVal.value = Math.min(minVal.value, maxVal.value);

	let min = parseFloat(prop.min) || 0;
	maxPercent.value = ((val - min) * 100) / (max - min);
	minPercent.value = Math.min(minPercent.value, maxPercent.value);
}

function onMin(evt) {
	evt.target.blur();
	setMin(evt.target.value);
	emit("input", evt);
}

function onMax(evt) {
	evt.target.blur();
	setMax(evt.target.value);
	emit("input", evt);
}

function onSlide(evt) {
	let bar = evt.target.closest(".slide");
	let rect = bar.getBoundingClientRect();
	let pos = evt.pageX !== undefined ? evt.pageX : evt.changedTouches[0].pageX;
	pos = Math.min(Math.max(pos - rect.left, 0), rect.width);

	let percent = (pos * 100) / rect.width;
	let min = parseFloat(prop.min) || 0;
	let max = parseFloat(prop.max) || 0;
	let val = (percent / 100) * (max - min) + min;
	if (evt.target.closest("button").classList.contains("min")) {
		minVal.value = val.toFixed(prop.frac);
		maxVal.value = Math.max(minVal.value, maxVal.value);
		minPercent.value = percent;
		maxPercent.value = Math.max(minPercent.value, maxPercent.value);
	} else {
		maxVal.value = val.toFixed(prop.frac);
		minVal.value = Math.min(minVal.value, maxVal.value);
		maxPercent.value = percent;
		minPercent.value = Math.min(minPercent.value, maxPercent.value);
	}

	emit("input", evt);
}

function onPress(evt) {
	let handle = evt.target.closest("button");
	handle.addEventListener("mousemove", onSlide);
	handle.addEventListener("touchmove", onSlide);
	handle.addEventListener("mouseup", onRelease);
	handle.addEventListener("touchend", onRelease);
	handle.addEventListener("mouseleave", onRelease);
	handle.addEventListener("touchcancel", onRelease);
	handle.classList.add("active");
}

function onRelease(evt) {
	let handle = evt.target.closest("button");
	handle.removeEventListener("mousemove", onSlide);
	handle.removeEventListener("touchmove", onSlide);
	handle.removeEventListener("mouseup", onRelease);
	handle.removeEventListener("touchend", onRelease);
	handle.removeEventListener("mouseleave", onRelease);
	handle.removeEventListener("touchcancel", onRelease);
	handle.classList.remove("active");
}
</script>

<template>
	<div class="slider">
		<label>
			<span>{{ label }} ({{ t("unit." + unit) }})</span>
			<input type="number" :name="name" :value="minVal" @change="onMin" />
			&nbsp;&ndash;&nbsp;
			<input type="number" :name="name" :value="maxVal" @change="onMax" />
		</label>

		<div class="bar">
			<div class="slide">
				<div class="overlay" :style="{ left: minPercent + '%', right: 100 - maxPercent + '%' }"></div>
				<button
					type="button"
					class="handle min"
					:style="{ left: minPercent + '%' }"
					@mousedown="onPress"
					@touchstart="onPress"
				>
					<div class="interact-area"></div>
				</button>
				<button
					type="button"
					class="handle max"
					:style="{ left: maxPercent + '%' }"
					@mousedown="onPress"
					@touchstart="onPress"
				>
					<div class="interact-area"></div>
				</button>
			</div>
		</div>
	</div>
</template>

<style>
.slider label {
	display: flex;
	align-items: baseline;
}

.slider span {
	flex-grow: 1;
	margin-right: 1em;
}

.slider input[type="number"] {
	flex-basis: 3em;
	padding: 0;
	line-height: 1;
	border: none;
	border-radius: 0;
	border-bottom: var(--border);
	text-align: center;
	font-weight: 300;
}

.slider .bar {
	position: relative;
	height: 4px;
	background-color: var(--color-placeholder);
	border-radius: 2px;
	margin: 16px 0;
}

.slider .bar .slide {
	position: absolute;
	top: 0;
	bottom: 0;
	left: 20px;
	right: 20px;
}

.slider .bar .overlay {
	background-color: var(--color-primary);
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
}

.slider .bar .handle {
	position: absolute;
	width: 20px;
	height: 20px;
	padding: 0;
	background-color: #fff;
	border-radius: 10px;
	border: var(--border);
	top: -8px;
	cursor: pointer;
	transition: box-shadow var(--transition-style);
}

.slider .bar .handle.min {
	transform: translateX(-20px);
}

.slider .bar .handle.max {
	left: 100%;
}

.slider .bar .handle .interact-area {
	position: absolute;
	width: 40px;
	height: 40px;
	left: -11px;
	top: -11px;
	border-radius: 20px;
}

.slider .bar .handle.active ~ .handle {
	pointer-events: none;
}

.slider .bar .handle.active .interact-area {
	transform: scale(5);
	z-index: 10;
}

@media (hover: hover) {
	.slider .bar .handle:hover {
		box-shadow: var(--shadow-button);
	}
}
</style>

<script setup>
import Radio from "./Radio.vue";
import SortImage from "./images/ImageSort.vue";
import { ref, computed, onMounted, onUnmounted } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const prop = defineProps(["list"]);

const isOpen = ref(false);

const cat = computed(() => (prop.list ? prop.list.sortBy : ""));
const dir = computed(() => (prop.list ? prop.list.sortDir : ""));
const options = computed(() => (prop.list ? prop.list.categories : []));

function onChangeDir(evt) {
	prop.list.setSortCategory(cat.value, evt.target.value);
}

function onChangeCat(evt) {
	prop.list.setSortCategory(evt.target.value, dir.value);
}

function toggleOpen(evt) {
	evt.stopPropagation();
	isOpen.value = !isOpen.value;
}

function close(evt) {
	if (isOpen.value && !evt.target.closest(".sort-menu .options")) {
		isOpen.value = false;
	}
}

onMounted(() => document.addEventListener("click", close));
onUnmounted(() => document.removeEventListener("click", close));
</script>

<template>
	<div class="sort-menu" :class="{ open: isOpen }">
		<button class="icon" @click="toggleOpen">
			<SortImage />
		</button>
		<div class="options">
			<Radio v-for="opt in options" name="cat" :value="opt.cat" :checked="cat == opt.cat" @change="onChangeCat">
				{{ opt.name }}
			</Radio>
			<hr />
			<Radio name="dir" value="asc" :checked="dir == 'asc'" @change="onChangeDir">{{ t("sort.asc") }}</Radio>
			<Radio name="dir" value="desc" :checked="dir == 'desc'" @change="onChangeDir">{{ t("sort.desc") }}</Radio>
		</div>
	</div>
</template>

<style>
.sort-menu {
	position: relative;
}

.sort-menu .options {
	position: absolute;
	z-index: 2;
	min-width: 200px;
	max-width: 66vw;
	box-shadow: var(--shadow-menu);
	border-radius: 8px;
	background-color: #fff;
	padding: 0.5em 0;
	opacity: 0;
	left: -2000px;
	transition: opacity var(--transition-style);
}

.sort-menu.open .options {
	left: 0;
	opacity: 1;
}

.sort-menu label.radio {
	padding: 0.5em;
	transition: color var(--transition-style);
}

.sort-menu label.radio > input + div {
	margin-right: 0.5em;
}

@media (hover: hover) {
	.sort-menu label.radio:hover {
		color: var(--color-primary);
	}
}
</style>

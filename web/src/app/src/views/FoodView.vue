<script setup>
import Main from "../components/Main.vue";
import Search from "../components/LocalSearch.vue";
import Slider from "../components/Slider.vue";
import NewFood from "../components/ClickableInput.vue";
import FoodList from "../components/FoodList.vue";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const perms = inject("perms");
const foods = inject("food");
const filteredFood = ref([]);
const current = ref(null);

function newFood(name) {
	// TODO create new food
	console.log(name);
}

function updateList(items) {
	filteredFood.value = items;

	// TODO not sure yet about this part
	if (current.value) {
		if (filteredFood.value.filter(f => f.id == current.value.id).length == 0) {
			current.value = null;
		}
	}
}

function showDetails(id) {
	current.value = filteredFood.value.filter(f => f.id == id)[0];
}
</script>

<template>
	<Main :current="current">
		<template #filter>
			<section v-if="perms.canCreateFood">
				<h2>{{ $t("aria.headnew") }}</h2>
				<NewFood :label="$t('btn.new')" :placeholder="$t('food.hintnew')" @confirm="newFood" />
			</section>
			<section>
				<h2>{{ $t("aria.headsearch") }}</h2>
				<Search :data="foods" v-slot="slotProps" :placeholder="$t('food.hintsearch')" @result="updateList">
					<Slider :label="$t('food.energy')" @input="slotProps.confirm" name="kcal" unit="cal" min="0" max="900" frac="0" />
					<Slider :label="$t('food.fat')" @input="slotProps.confirm" name="fat" unit="g" min="0" max="100" frac="0" />
					<Slider :label="$t('food.carbs')" @input="slotProps.confirm" name="carb" unit="g" min="0" max="100" frac="0" />
					<Slider :label="$t('food.protein')" @input="slotProps.confirm" name="prot" unit="g" min="0" max="100" frac="0" />
				</Search>
			</section>
		</template>
		<template #main>
			<FoodList :items="filteredFood" @selected="showDetails" />
		</template>
		<template #details v-if="current">
			<section>
				{{ current.name }}
			</section>
		</template>
	</Main>
</template>

<style></style>

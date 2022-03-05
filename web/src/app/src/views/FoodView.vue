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
		if (filteredFood.value.filter((f) => f.id == current.value.id).length == 0) {
			current.value = null;
		}
	}
}

function showDetails(id) {
	current.value = filteredFood.value.filter((f) => f.id == id)[0];
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
					<Slider
						:label="$t('food.energy')"
						@input="slotProps.confirm"
						name="kcal"
						unit="cal"
						min="0"
						max="900"
						frac="0"
					/>
					<Slider
						:label="$t('food.fat')"
						@input="slotProps.confirm"
						name="fat"
						unit="g"
						min="0"
						max="100"
						frac="0"
					/>
					<Slider
						:label="$t('food.carbs')"
						@input="slotProps.confirm"
						name="carb"
						unit="g"
						min="0"
						max="100"
						frac="0"
					/>
					<Slider
						:label="$t('food.protein')"
						@input="slotProps.confirm"
						name="prot"
						unit="g"
						min="0"
						max="100"
						frac="0"
					/>
				</Search>
			</section>
		</template>
		<template #main>
			<FoodList :items="filteredFood" @selected="showDetails" />
		</template>
		<template #details v-if="current">
			<section class="subtitle">Some food category</section>
			<section class="tags">
				<span class="tag">Tag 1</span>
				<span class="tag">Tag 2</span>
				<span class="tag">Tag 3</span>
			</section>
			<section class="nutrients">
				<fieldset :disabled="!(perms.canCreateFood || perms.canEditFood)">
					<div>
						<label>{{ $t("food.energy") }}</label>
						<input type="text" :value="current.kcal" name="kcal" />
						<span class="unit">{{ $t("unit.cal") }}</span>
					</div>
					<div>
						<label>{{ $t("food.fat") }}</label>
						<input type="text" :value="current.fat" name="fat" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.carbs2") }}</label>
						<input type="text" :value="current.carb" name="carb" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.protein") }}</label>
						<input type="text" :value="current.prot" name="prot" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.fiber") }}</label>
						<input type="text" :value="current.fib" name="fib" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
				</fieldset>
				<fieldset>
					<div>
						<label>{{ $t("food.fatsat") }}</label>
						<input type="text" :value="current.fatsat" name="fatsat" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.fato3") }}</label>
						<input type="text" :value="current.fato3" name="fato3" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.fato6") }}</label>
						<input type="text" :value="current.fato6" name="fato6" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.sugar") }}</label>
						<input type="text" :value="current.sug" name="sug" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
					<div>
						<label>{{ $t("food.salt") }}</label>
						<input type="text" :value="current.sod" name="sod" />
						<span class="unit">{{ $t("unit.g") }}</span>
					</div>
				</fieldset>
			</section>
		</template>
	</Main>
</template>

<style>
#details .controls {
	padding-bottom: 0;
}

#details section.subtitle {
	padding-top: 0;
	padding-bottom: 0;
	border: none;
}

#details section.tags {
	padding: 0 0.25em 0.5em;
	display: block;
}

#details section.nutrients:not(:last-child) {
	padding-bottom: 2em;
	border: none;
}
</style>

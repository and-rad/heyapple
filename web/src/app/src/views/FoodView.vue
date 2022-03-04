<script setup>
import Main from "../components/Main.vue";
import NewFood from "../components/ClickableInput.vue";
import FoodSearch from "../components/Search.vue";
import { inject } from "vue";

const perms = inject("perms");

function newFood(name) {
	// TODO create new food
	console.log(name);
}

function updateList(status, items) {
	// TODO update food list
	console.log(status, items);
}
</script>

<template>
	<Main>
		<template #filter>
			<section v-if="perms.canCreateFood">
				<h2>{{ $t("aria.headnew") }}</h2>
				<NewFood :label="$t('btn.new')" :placeholder="$t('food.hintnew')" @confirm="newFood" />
			</section>
			<section>
				<h2>{{ $t("aria.headsearch") }}</h2>
				<FoodSearch
					action="/api/v1/foods"
					v-slot="slotProps"
					:placeholder="$t('food.hintsearch')"
					@result="updateList"
				></FoodSearch>
			</section>
		</template>
		<template #main> Food </template>
	</Main>
</template>

<style></style>

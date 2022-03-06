<script setup>
import Main from "../components/Main.vue";
import NewRecipe from "../components/ClickableInput.vue";

function newRecipe(name) {
	// TODO create new food
	console.log(name);
}
</script>

<template>
	<Main>
		<template #filter>
			<section class="new-item">
				<h2>{{ $t("aria.headnew") }}</h2>
				<NewRecipe :label="$t('btn.new')" :placeholder="$t('recipe.hintnew')" @confirm="newRecipe" />
			</section>
			<section>
				<h2>{{ $t("aria.headsearch") }}</h2>
			</section>
		</template>
		<template #main> Recipes </template>
	</Main>
</template>

<style></style>

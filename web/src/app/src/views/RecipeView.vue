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
			<section>
				<NewRecipe :label="$t('filter.new')" @confirm="newRecipe" />
			</section>
		</template>
		<template #main> Recipes </template>
	</Main>
</template>

<style></style>

<script setup>
import Main from "../components/Main.vue";
import NewFood from "../components/ClickableInput.vue";

function newFood(name) {
	// TODO create new food
	console.log(name);
}
</script>

<template>
	<Main>
		<template #filter>
			<section>
				<h2>{{ $t("aria.headnew") }}</h2>
				<NewFood :label="$t('filter.new')" @confirm="newFood" />
			</section>
		</template>
		<template #main> Food </template>
	</Main>
</template>

<style></style>

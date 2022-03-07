<script setup>
import Main from "../components/Main.vue";
import NewRecipe from "../components/ClickableInput.vue";
import { inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");

function newRecipe(name) {
	fetch("/api/v1/recipe", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams({ name: name }),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("createrec.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			console.log(data);
			log.msg(t("createrec.ok"));
		})
		.catch((err) => log.err(err));
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

<script setup>
import Main from "../components/Main.vue";
import NewRecipe from "../components/ClickableInput.vue";
import FoodList from "../components/FoodList.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, inject, onMounted } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");
const recipes = inject("recipes");

const filtered = ref([]);
const current = ref(null);
const editMode = ref(false);
const isSaving = ref(false);

const main = ref(null);

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

function saveRecipe() {
	isSaving.value = true;
	console.log("TODO save recipe");
}

function showDetails(id) {
	current.value = filtered.value.filter((f) => f.id == id)[0];
	main.value.showDetails();
}

function onEditMode() {
	editMode.value ? saveRecipe() : (editMode.value = true);
}

onMounted(() => {
	filtered.value = recipes.value;
})
</script>

<template>
	<Main ref="main" @detailVisibility="editMode = false">
		<template #filter>
			<section class="new-item">
				<h2>{{ $t("aria.headnew") }}</h2>
				<NewRecipe :label="$t('btn.new')" :placeholder="$t('recipe.hintnew')" @confirm="newRecipe" />
			</section>
			<section>
				<h2>{{ $t("aria.headsearch") }}</h2>
			</section>
		</template>

		<template #main>
			<FoodList :items="filtered" @selected="showDetails" />
		</template>

		<template #head-details v-if="current">
			<h2>{{ current.name }}</h2>
		</template>

		<template #details v-if="current">
			<section class="tags">
				<span class="tag">Tag 1</span>
				<span class="tag">Tag 2</span>
				<span class="tag">Tag 3</span>
				<button
					class="icon async"
					:disabled="isSaving"
					@click="onEditMode"
				>
					<EditImage v-if="!editMode" />
					<SaveImage v-if="editMode" />
				</button>
			</section>
		</template>
	</Main>
</template>

<style></style>

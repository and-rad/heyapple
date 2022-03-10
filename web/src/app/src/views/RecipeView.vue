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
const perms = inject("perms");
const recipes = inject("recipes");

const filtered = ref([]);
const current = ref(null);
const editMode = ref(false);
const isSaving = ref(false);
const ownerInfo = ref("&nbsp;");

const main = ref(null);
const form = ref(null);

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

	let id = current.value.id;
	let owner = current.value.owner;
	let isOwner = current.value.isOwner;

	fetch("/api/v1/recipe/" + id, {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form.value)),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("saverec.err" + response.status);
			}
			editMode.value = false;
			return fetch("/api/v1/recipe/" + id);
		})
		.then((response) => response.json())
		.then((data) => {
			data.owner = owner;
			data.isOwner = isOwner;
			recipes.value = recipes.value.map((r) => (data.id == r.id ? data : r));
			filtered.value = filtered.value.map((r) => (data.id == r.id ? data : r));
			current.value = current.value.id == data.id ? data : current.value;
			log.msg(t("saverec.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				isSaving.value = false;
			}, 500);
		});
}

function showDetails(id) {
	current.value = filtered.value.filter((r) => r.id == id)[0];
	main.value.showDetails();

	if ("isOwner" in current.value) {
		updateOwnerInfo();
	} else {
		fetchOwnerInfo();
	}
}

function fetchOwnerInfo() {
	console.log("TODO get owner info");
}

function updateOwnerInfo() {
	if (current.value.isOwner) {
		ownerInfo.value = t("recipe.isowner");
	} else if (current.value.owner) {
		ownerInfo.value = t("recipe.owner", { name: current.value.owner });
	} else {
		ownerInfo.value = t("recipe.ispublic");
	}
}

function onEditMode() {
	editMode.value ? saveRecipe() : (editMode.value = true);
}

onMounted(() => (filtered.value = recipes.value));
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
			<form ref="form" autocomplete="off">
				<fieldset :disabled="!editMode">
					<input type="text" name="name" :value="current.name" />
				</fieldset>
			</form>
		</template>

		<template #details v-if="current">
			<section class="subtitle" v-html="ownerInfo"></section>
			<section class="tags">
				<span class="tag">Tag 1</span>
				<span class="tag">Tag 2</span>
				<span class="tag">Tag 3</span>
				<button class="icon async" :disabled="isSaving" @click="onEditMode" v-if="current.isOwner">
					<EditImage v-if="!editMode" />
					<SaveImage v-if="editMode" />
				</button>
			</section>
			<section>
				<h2>{{ $t("aria.headtrack") }}</h2>
				Add to diary here
			</section>
			<section>
				<h2>{{ $t("aria.headingred") }}</h2>
				Ingredients go here
			</section>
			<section>
				<h2>{{ $t("aria.headnutrients") }}</h2>
				Nutrients go here
			</section>
			<section>
				<h2>{{ $t("aria.headprep") }}</h2>
				Cooking instructions go here
			</section>
		</template>
	</Main>
</template>

<style></style>

<script setup>
import Main from "../components/Main.vue";
import NewRecipe from "../components/ClickableInput.vue";
import FoodList from "../components/FoodList.vue";
import IngredientList from "../components/IngredientList.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import ListImage from "../components/images/ImageList.vue";
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
			data.isowner = true;
			recipes.value.push(data);
			filtered.value.push(data);
			log.msg(t("createrec.ok"));
			showDetails(data.id);
		})
		.catch((err) => log.err(err));
}

function saveRecipe() {
	isSaving.value = true;

	let id = current.value.id;
	let owner = current.value.owner;
	let isowner = current.value.isowner;

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
			data.isowner = isowner;
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

	if ("isowner" in current.value) {
		updateOwnerInfo();
	} else {
		fetchOwnerInfo();
	}
}

function fetchOwnerInfo() {
	fetch(`/api/v1/recipe/${current.value.id}/owner`)
		.then((response) => {
			if (!response.ok) {
				throw response;
			}
			return response.json();
		})
		.then((data) => {
			current.value.isowner = data.isowner;
			current.value.owner = data.owner;
			updateOwnerInfo();
		})
		.catch(() => {
			log.err(t("recowner.err"));
			ownerInfo.value = "&nbsp;";
		});
}

function updateOwnerInfo() {
	if (current.value.isowner) {
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

function onInput(evt) {
	evt.target.blur();
	if (isNaN(parseFloat(evt.target.value))) {
		evt.target.value = current.value[evt.target.name];
	}
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
			<form ref="form" autocomplete="off" id="form-recipe">
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
				<button class="icon async" :disabled="isSaving" @click="onEditMode" v-if="current.isowner">
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
				<IngredientList :items="current.items" />
			</section>
			<section>
				<h2>{{ $t("aria.headnutrients") }}</h2>
				<div class="nutrient-block">
					<div class="col50">
						<div>
							<label>{{ $t("food.energy") }}</label>
							<span>{{ current.kcal }}</span>
							<span class="unit">{{ $t("unit.cal") }}</span>
						</div>
						<div>
							<label>{{ $t("food.fat") }}</label>
							<span>{{ current.fat }}</span>
							<span class="unit">{{ $t("unit.g") }}</span>
						</div>
					</div>
					<div class="col50">
						<div>
							<label>{{ $t("food.carbs2") }}</label>
							<span>{{ current.carb }}</span>
							<span class="unit">{{ $t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ $t("food.protein") }}</label>
							<span>{{ current.prot }}</span>
							<span class="unit">{{ $t("unit.g") }}</span>
						</div>
					</div>
				</div>
			</section>
			<section class="prep">
				<h2>{{ $t("aria.headprep") }}</h2>
				<div>
					<fieldset :disabled="!editMode" class="col50">
						<div>
							<label>{{ $t("recipe.size") }}</label>
							<input type="text" name="size" form="form-recipe" :value="current.size" @change="onInput" />
						</div>
						<div>
							<label>{{ $t("recipe.time") }}</label>
							<input
								type="text"
								disabled
								:value="current.preptime + current.cooktime + current.misctime"
							/>
							<span class="unit">{{ $t("unit.min") }}</span>
						</div>
					</fieldset>
					<fieldset :disabled="!editMode" class="col50">
						<div>
							<label>{{ $t("recipe.preptime") }}</label>
							<input
								type="text"
								name="preptime"
								form="form-recipe"
								:value="current.preptime"
								@change="onInput"
							/>
							<span class="unit">{{ $t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ $t("recipe.cooktime") }}</label>
							<input
								type="text"
								name="cooktime"
								form="form-recipe"
								:value="current.cooktime"
								@change="onInput"
							/>
							<span class="unit">{{ $t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ $t("recipe.misctime") }}</label>
							<input
								type="text"
								name="misctime"
								form="form-recipe"
								:value="current.misctime"
								@change="onInput"
							/>
							<span class="unit">{{ $t("unit.min") }}</span>
						</div>
					</fieldset>
				</div>
				<div class="placeholder">
					<ListImage />
					<p>{{ $t("todo.instructions") }}</p>
				</div>
			</section>
		</template>
	</Main>
</template>

<style>
#details section.prep .placeholder {
	height: 66vw;
}

@media only screen and (min-width: 400px) {
	#details section.prep > div {
		display: flex;
		justify-content: space-between;
	}
}

@media only screen and (min-width: 480px) {
	#details section.prep .placeholder {
		height: 320px;
	}
}
</style>

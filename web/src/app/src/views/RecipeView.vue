<script setup>
import Main from "../components/Main.vue";
import Search from "../components/LocalSearch.vue";
import Slider from "../components/Slider.vue";
import NewRecipe from "../components/ClickableInput.vue";
import FoodList from "../components/FoodList.vue";
import DiarySelect from "../components/ClickableDate.vue";
import IngredientList from "../components/IngredientList.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import ListImage from "../components/images/ImageList.vue";
import { ref, inject, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");
const perms = inject("perms");
const recipes = inject("recipes");
const diary = inject("diary");

const filtered = ref([]);
const current = ref(null);
const editMode = ref(false);
const isSaving = ref(false);
const disableToDiary = ref(false);
const amount = ref(1);
const ownerInfo = ref("&nbsp;");
const today = ref(DateTime.now().toISODate());
const now = ref(DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE));

const main = ref(null);
const form = ref(null);
const ingredients = ref(null);

function perServing(val, frac = 2) {
	return +parseFloat(val / (current.value.size || 1)).toFixed(frac);
}

function minSearchAttr(attr) {
	return Math.min.apply(
		Math,
		recipes.value.map((r) => Math.floor(r[attr] / r.size))
	);
}

function maxSearchAttr(attr) {
	return Math.max.apply(
		Math,
		recipes.value.map((r) => Math.ceil(r[attr] / r.size))
	);
}

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
			return saveIngredients(id);
		})
		.then(() => fetch("/api/v1/recipe/" + id))
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

function saveIngredients(id) {
	let items = ingredients.value.getDiff();
	if (items.length == 0) {
		return new Promise((resolve) => {
			editMode.value = false;
			resolve();
		});
	}

	return new Promise((resolve, reject) => {
		let count = 0;
		let error = undefined;
		items.forEach((item) => {
			fetch(`/api/v1/recipe/${id}/ingredient/${item.id}`, {
				method: "PUT",
				headers: {
					"Content-Type": "application/x-www-form-urlencoded",
					"X-CSRF-Token": csrf,
				},
				body: new URLSearchParams({ amount: item.amount }),
			})
				.then((response) => {
					++count;
					if (!response.ok) {
						throw t("saverec.err" + response.status);
					}
				})
				.catch((err) => (error = err))
				.finally(() => {
					if (count == items.length) {
						if (error) {
							reject(error);
						} else {
							editMode.value = false;
							resolve();
						}
					}
				});
		});
	});
}

function addToDiary(date, time) {
	disableToDiary.value = true;

	let params = new URLSearchParams();
	current.value.items.forEach((i) => {
		params.append("id", i.id);
		params.append("amount", perServing(i.amount) * amount.value);
		params.append("time", time);
		params.append("recipe", current.value.name);
	});

	fetch(`/api/v1/diary/${date}`, {
		method: "POST",
		body: params,
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savediary.err" + response.status);
			}
			return fetch("/api/v1/diary/" + date.replaceAll("-", "/"));
		})
		.then((response) => response.json())
		.then((data) => {
			diary.value[date] = data[0];
			amount.value = 1;
			log.msg(t("savediary.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				disableToDiary.value = false;
			}, 500);
		});
}

function showDetails(id) {
	current.value = filtered.value.filter((r) => r.id == id)[0];
	today.value = DateTime.now().toISODate();
	now.value = DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE);
	amount.value = 1;
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

function updateList(items) {
	filtered.value = items;
	if (current.value) {
		if (filtered.value.filter((r) => r.id == current.value.id).length == 0) {
			current.value = null;
		}
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

onMounted(() => (filtered.value = [...recipes.value]));
</script>

<template>
	<Main ref="main" @detailVisibility="editMode = false" :class="{ 'edit-mode': editMode }">
		<template #filter>
			<section class="new-item">
				<h2>{{ t("aria.headnewrec") }}</h2>
				<NewRecipe :label="t('btn.new')" :placeholder="t('recipe.hintnew')" @confirm="newRecipe" />
			</section>
			<hr />
			<section>
				<h2>{{ t("aria.headsearch") }}</h2>
				<Search :data="recipes" v-slot="slotProps" :placeholder="t('recipe.hintsearch')" @result="updateList">
					<fieldset>
						<Slider
							:label="t('food.energy')"
							:min="minSearchAttr('kcal')"
							:max="maxSearchAttr('kcal')"
							@input="slotProps.confirm"
							name="kcal"
							unit="cal"
							frac="0"
						/>
						<Slider
							:label="t('food.fat')"
							:min="minSearchAttr('fat')"
							:max="maxSearchAttr('fat')"
							@input="slotProps.confirm"
							name="fat"
							unit="g"
							frac="0"
						/>
						<Slider
							:label="t('food.carbs')"
							:min="minSearchAttr('carb')"
							:max="maxSearchAttr('carb')"
							@input="slotProps.confirm"
							name="carb"
							unit="g"
							frac="0"
						/>
						<Slider
							:label="t('food.protein')"
							:min="minSearchAttr('prot')"
							:max="maxSearchAttr('prot')"
							@input="slotProps.confirm"
							name="prot"
							unit="g"
							frac="0"
						/>
					</fieldset>
				</Search>
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
			<section class="subtitle no-edit-mode" v-html="ownerInfo"></section>
			<section class="tags">
				<span class="tag no-edit-mode">Tag 1</span>
				<span class="tag no-edit-mode">Tag 2</span>
				<span class="tag no-edit-mode">Tag 3</span>
				<button class="icon async" :disabled="isSaving" @click="onEditMode" v-if="current.isowner">
					<EditImage v-if="!editMode" />
					<SaveImage v-if="editMode" />
				</button>
			</section>
			<hr />
			<section v-if="current.items.length" class="tracking no-edit-mode">
				<h2>{{ t("aria.headtrack") }}</h2>
				<fieldset class="tracking-amount">
					<div>
						<label>{{ t("food.amount") }}</label>
						<input type="number" v-model="amount" name="amount" />
						<span class="unit">{{ t("recipe.size", amount) }}</span>
					</div>
				</fieldset>
				<fieldset>
					<label>Add to diary</label>
					<DiarySelect
						:label="t('btn.add')"
						:time="now"
						:date="today"
						:disabled="disableToDiary"
						@confirm="addToDiary"
					/>
				</fieldset>
			</section>
			<hr />
			<section>
				<h2>{{ t("aria.headingred") }}</h2>
				<p class="msg-noitems" v-if="!current.items.length" v-html="t('recipe.noitems')"></p>
				<IngredientList ref="ingredients" :items="current.items" :disabled="!editMode" />
			</section>
			<hr />
			<section class="no-edit-mode">
				<h2>{{ t("aria.headnutrients") }}</h2>
				<div class="nutrient-block">
					<div class="col50">
						<div>
							<label>{{ t("food.energy") }}</label>
							<span>{{ perServing(current.kcal, 1) }}</span>
							<span class="unit">{{ t("unit.cal") }}</span>
						</div>
						<div>
							<label>{{ t("food.fat") }}</label>
							<span>{{ perServing(current.fat, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
					</div>
					<div class="col50">
						<div>
							<label>{{ t("food.carbs2") }}</label>
							<span>{{ perServing(current.carb, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.protein") }}</label>
							<span>{{ perServing(current.prot, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
					</div>
				</div>
			</section>
			<hr />
			<section class="prep">
				<h2>{{ t("aria.headprep") }}</h2>
				<div>
					<fieldset :disabled="!editMode" class="col50">
						<div class="prep-size">
							<label>{{ t("recipe.size", 2) }}</label>
							<input type="text" name="size" form="form-recipe" :value="current.size" @change="onInput" />
							<label>{{ t("recipe.size", current.size) }}</label>
						</div>
					</fieldset>
					<fieldset :disabled="!editMode" class="col50">
						<div>
							<label>{{ t("recipe.time") }}</label>
							<input
								type="text"
								disabled
								:value="current.preptime + current.cooktime + current.misctime"
							/>
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ t("recipe.preptime") }}</label>
							<input
								type="text"
								name="preptime"
								form="form-recipe"
								:value="current.preptime"
								@change="onInput"
							/>
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ t("recipe.cooktime") }}</label>
							<input
								type="text"
								name="cooktime"
								form="form-recipe"
								:value="current.cooktime"
								@change="onInput"
							/>
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ t("recipe.misctime") }}</label>
							<input
								type="text"
								name="misctime"
								form="form-recipe"
								:value="current.misctime"
								@change="onInput"
							/>
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
					</fieldset>
				</div>
				<div class="placeholder">
					<ListImage />
					<p>{{ t("todo.instructions") }}</p>
				</div>
			</section>
		</template>
	</Main>
</template>

<style>
#details section.prep .placeholder {
	height: 66vw;
}

#details section.prep .prep-size label:last-child {
	display: none;
}

#details .msg-noitems {
	padding: 0.5em 0;
	color: var(--color-text-light);
	text-align: center;
}

#details .ingredients label {
	color: var(--color-text);
}

@media only screen and (min-width: 400px) {
	#details section.prep > div {
		display: flex;
		justify-content: space-between;
	}

	#details section.prep .prep-size input {
		font-size: 5em;
		flex-basis: 50%;
		padding: 0;
		flex-grow: 1;
		margin-top: 1rem;
	}

	#details section.prep .prep-size label:first-child {
		display: none;
	}

	#details section.prep .prep-size label:last-child {
		flex-basis: 0;
		display: block;
		overflow-x: visible;
	}
}

@media only screen and (min-width: 480px) {
	#details section.prep .placeholder {
		height: 320px;
	}
}
</style>

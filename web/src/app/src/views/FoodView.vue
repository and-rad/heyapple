<script setup>
import Main from "../components/Main.vue";
import Search from "../components/LocalSearch.vue";
import Slider from "../components/Slider.vue";
import NewFood from "../components/ClickableInput.vue";
import RecipeSelect from "../components/ClickableSelect.vue";
import DiarySelect from "../components/ClickableDate.vue";
import FoodList from "../components/FoodList.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");
const perms = inject("perms");
const foods = inject("food");
const recipes = inject("recipes");
const diary = inject("diary");

const filtered = ref([]);
const current = ref(null);
const editMode = ref(false);
const disableSave = ref(false);
const disableToRecipe = ref(false);
const disableToDiary = ref(false);
const amount = ref(100);
const today = ref(DateTime.now().toISODate());
const now = ref(DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE));

const main = ref(null);
const form = ref(null);

function newFood(name) {
	fetch("/api/v1/food", {
		method: "POST",
		headers: { "X-CSRF-Token": csrf },
	})
		.then((response) => {
			if (!response.ok) {
				throw t("createfood.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			data.name = name;
			foods.value.push(data);
			filtered.value.push(data);
			log.msg(t("createfood.ok"));
			showDetails(data.id);
		})
		.catch((err) => log.err(err));
}

function saveFood() {
	disableSave.value = true;
	let id = current.value.id;
	fetch("/api/v1/food/" + id, {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form.value)),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savefood.err" + response.status);
			}
			editMode.value = false;
			return fetch("/api/v1/food/" + id);
		})
		.then((response) => response.json())
		.then((data) => {
			data.name = t(data.id.toString());
			foods.value = foods.value.map((f) => (data.id == f.id ? data : f));
			filtered.value = filtered.value.map((f) => (data.id == f.id ? data : f));
			current.value = current.value.id == data.id ? data : current.value;
			log.msg(t("savefood.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				disableSave.value = false;
			}, 500);
		});
}

function addToRecipe(id) {
	disableToRecipe.value = true;
	fetch(`/api/v1/recipe/${id}/ingredient/${current.value.id}`, {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams({ amount: amount.value }),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("addfood.err" + response.status);
			}
			return fetch("/api/v1/recipe/" + id);
		})
		.then((response) => response.json())
		.then((data) => {
			recipes.value = recipes.value.map((r) => (data.id == r.id ? data : r));
			amount.value = 100;
			log.msg(t("addfood.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				disableToRecipe.value = false;
			}, 500);
		});
}

function addToDiary(date, time) {
	disableToDiary.value = true;
	fetch(`/api/v1/diary/${date}`, {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams({
			id: current.value.id,
			amount: amount.value,
			time: time,
		}),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("adddiary.err" + response.status);
			}
			return fetch("/api/v1/diary/" + date.replaceAll("-", "/"));
		})
		.then((response) => response.json())
		.then((data) => {
			diary.value[date] = data[0];
			amount.value = 100;
			log.msg(t("adddiary.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				disableToDiary.value = false;
			}, 500);
		});
}

function updateList(items) {
	filtered.value = items;
	if (current.value) {
		if (filtered.value.filter((f) => f.id == current.value.id).length == 0) {
			current.value = null;
		}
	}
}

function showDetails(id) {
	current.value = filtered.value.filter((f) => f.id == id)[0];
	today.value = DateTime.now().toISODate();
	now.value = DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE);
	amount.value = 100;
	main.value.showDetails();
}

function onEditMode() {
	editMode.value ? saveFood() : (editMode.value = true);
}

function onInput(evt) {
	evt.target.blur();
	if (isNaN(parseFloat(evt.target.value))) {
		evt.target.value = current.value[evt.target.name];
	}
}
</script>

<template>
	<Main ref="main" @detailVisibility="editMode = false" :class="{ 'edit-mode': editMode }">
		<template #filter>
			<section v-if="perms.canCreateFood" class="new-item">
				<h2>{{ t("aria.headnew") }}</h2>
				<NewFood :label="t('btn.new')" :placeholder="t('food.hintnew')" @confirm="newFood" />
			</section>
			<hr />
			<section>
				<h2>{{ t("aria.headsearch") }}</h2>
				<Search :data="foods" v-slot="slotProps" :placeholder="t('food.hintsearch')" @result="updateList">
					<fieldset>
						<legend>{{ t("aria.headmacro1") }}</legend>
						<Slider
							:label="t('food.energy')"
							@input="slotProps.confirm"
							name="kcal"
							unit="cal"
							min="0"
							max="900"
							frac="0"
						/>
						<Slider
							:label="t('food.fat')"
							@input="slotProps.confirm"
							name="fat"
							unit="g"
							min="0"
							max="100"
							frac="0"
						/>
						<Slider
							:label="t('food.carbs')"
							@input="slotProps.confirm"
							name="carb"
							unit="g"
							min="0"
							max="100"
							frac="0"
						/>
						<Slider
							:label="t('food.protein')"
							@input="slotProps.confirm"
							name="prot"
							unit="g"
							min="0"
							max="89"
							frac="0"
						/>
						<Slider
							:label="t('food.fiber')"
							@input="slotProps.confirm"
							name="fib"
							unit="g"
							min="0"
							max="71"
							frac="0"
						/>
					</fieldset>
					<fieldset>
						<legend>{{ t("aria.headmacro2") }}</legend>
						<Slider
							:label="t('food.fatsat')"
							@input="slotProps.confirm"
							name="fatsat"
							unit="g"
							min="0"
							max="83"
							frac="0"
						/>
						<Slider
							:label="t('food.fato3')"
							@input="slotProps.confirm"
							name="fato3"
							unit="g"
							min="0"
							max="54"
							frac="0"
						/>
						<Slider
							:label="t('food.fato6')"
							@input="slotProps.confirm"
							name="fato6"
							unit="g"
							min="0"
							max="70"
							frac="0"
						/>
						<Slider
							:label="t('food.sugar')"
							@input="slotProps.confirm"
							name="sug"
							unit="g"
							min="0"
							max="100"
							frac="0"
						/>
						<Slider
							:label="t('food.salt')"
							@input="slotProps.confirm"
							name="salt"
							unit="g"
							min="0"
							max="100"
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
			<h2>{{ current.name }}</h2>
		</template>

		<template #details v-if="current">
			<section class="subtitle">Some food category</section>
			<section class="tags">
				<span class="tag">Tag 1</span>
				<span class="tag">Tag 2</span>
				<span class="tag">Tag 3</span>
				<button
					class="icon async"
					:disabled="disableSave"
					@click="onEditMode"
					v-if="perms.canCreateFood || perms.canEditFood"
				>
					<EditImage v-if="!editMode" />
					<SaveImage v-if="editMode" />
				</button>
			</section>
			<hr />
			<section class="tracking no-edit-mode">
				<h2>{{ t("aria.headtrack") }}</h2>
				<fieldset class="tracking-amount">
					<div>
						<label>{{ t("food.amount") }}</label>
						<input type="number" v-model="amount" name="amount" />
						<span class="unit">{{ t("unit.g") }}</span>
					</div>
				</fieldset>
				<fieldset>
					<label>Add to recipe</label>
					<RecipeSelect
						:label="t('btn.add')"
						:placeholder="t('food.hintrec')"
						:items="recipes"
						:disabled="disableToRecipe"
						@confirm="addToRecipe"
					/>
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
				<h2>{{ t("aria.headnutrients") }}</h2>
				<form ref="form">
					<div class="nutrient-block">
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.energy") }}</label>
								<input type="text" :value="current.kcal" name="kcal" @change="onInput" />
								<span class="unit">{{ t("unit.cal") }}</span>
							</div>
							<div>
								<label>{{ t("food.fat") }}</label>
								<input type="text" :value="current.fat" name="fat" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.carbs2") }}</label>
								<input type="text" :value="current.carb" name="carb" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.protein") }}</label>
								<input type="text" :value="current.prot" name="prot" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fiber") }}</label>
								<input type="text" :value="current.fib" name="fib" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
						</fieldset>
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.fatsat") }}</label>
								<input type="text" :value="current.fatsat" name="fatsat" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fato3") }}</label>
								<input type="text" :value="current.fato3" name="fato3" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fato6") }}</label>
								<input type="text" :value="current.fato6" name="fato6" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.sugar") }}</label>
								<input type="text" :value="current.sug" name="sug" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.salt") }}</label>
								<input type="text" :value="current.salt" name="salt" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
						</fieldset>
					</div>
				</form>
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
}

#details section.subtitle strong {
	font-weight: normal;
	color: var(--color-secondary);
}

#details section.tags {
	padding: 0.5em 3em 0.5em 0.25em;
	position: relative;
	display: flex;
	flex-wrap: wrap;
}

#details section.tags button.icon {
	position: absolute;
	right: 0.5em;
	bottom: 0.5em;
}

#details section.tracking > fieldset {
	margin-bottom: 1em;
}

#details .tracking-amount > div {
	display: flex;
	flex-wrap: wrap;
	align-items: baseline;
	padding: 0.5em 0;
}

#details .tracking-amount label {
	flex-basis: 100%;
}

#details .nutrient-block:not(:first-of-type) {
	margin-top: 3em;
}

@media only screen and (min-width: 400px) {
	#details .nutrient-block {
		display: flex;
		justify-content: space-between;
	}
}
</style>

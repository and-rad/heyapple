<script setup>
import Main from "../components/Main.vue";
import Search from "../components/LocalSearch.vue";
import Slider from "../components/Slider.vue";
import NewFood from "../components/ClickableInput.vue";
import RecipeSelect from "../components/ClickableSelect.vue";
import DiarySelect from "../components/ClickableDate.vue";
import FoodList from "../components/FoodList.vue";
import TagList from "../components/TagList.vue";
import SortMenu from "../components/SortMenu.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, inject, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { DateTime } from "luxon";

const { t } = useI18n();
const log = inject("log");
const csrf = inject("csrfToken");
const perms = inject("perms");
const foods = inject("food");
const recipes = inject("recipes");
const diary = inject("diary");
const prefs = inject("prefs");

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
const list = ref(null);

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
			data.cat = "";
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
			data.cat = t("cat." + data.cat);
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
	today.value = date;
	now.value = time;

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
			recipe: "",
		}),
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
			amount.value = 100;
			log.msg(t("savediary.ok"));
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

function onSort(evt) {
	let [cat, dir] = evt.target.value.split(" ");
	list.value.setSortCategory(cat, dir);
	evt.target.selectedIndex = 0;
}

onMounted(() => (filtered.value = [...foods.value]));
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
							frac="0" />
						<Slider
							:label="t('food.fat')"
							@input="slotProps.confirm"
							name="fat"
							unit="g"
							min="0"
							max="100"
							frac="0" />
						<Slider
							:label="t('food.carbs')"
							@input="slotProps.confirm"
							name="carb"
							unit="g"
							min="0"
							max="100"
							frac="0" />
						<Slider
							:label="t('food.protein')"
							@input="slotProps.confirm"
							name="prot"
							unit="g"
							min="0"
							max="89"
							frac="0" />
						<Slider
							:label="t('food.fiber')"
							@input="slotProps.confirm"
							name="fib"
							unit="g"
							min="0"
							max="71"
							frac="0" />
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
							frac="0" />
						<Slider
							:label="t('food.fato3')"
							@input="slotProps.confirm"
							name="fato3"
							unit="g"
							min="0"
							max="54"
							frac="0" />
						<Slider
							:label="t('food.fato6')"
							@input="slotProps.confirm"
							name="fato6"
							unit="g"
							min="0"
							max="70"
							frac="0" />
						<Slider
							:label="t('food.sugar')"
							@input="slotProps.confirm"
							name="sug"
							unit="g"
							min="0"
							max="100"
							frac="0" />
						<Slider
							:label="t('food.fruc')"
							@input="slotProps.confirm"
							name="fruc"
							unit="g"
							min="0"
							max="56"
							frac="0" />
						<Slider
							:label="t('food.gluc')"
							@input="slotProps.confirm"
							name="gluc"
							unit="g"
							min="0"
							max="36"
							frac="0" />
						<Slider
							:label="t('food.suc')"
							@input="slotProps.confirm"
							name="suc"
							unit="g"
							min="0"
							max="100"
							frac="0" />
					</fieldset>
					<fieldset>
						<legend>{{ t("aria.headvits") }}</legend>
						<Slider
							:label="t('food.vita')"
							@input="slotProps.confirm"
							name="vita"
							unit="mg"
							min="0"
							max="30"
							frac="1" />
						<Slider
							:label="t('food.vitb1')"
							@input="slotProps.confirm"
							name="vitb1"
							unit="mg"
							min="0"
							max="24"
							frac="1" />
						<Slider
							:label="t('food.vitb2')"
							@input="slotProps.confirm"
							name="vitb2"
							unit="mg"
							min="0"
							max="18"
							frac="1" />
						<Slider
							:label="t('food.vitb3')"
							@input="slotProps.confirm"
							name="vitb3"
							unit="mg"
							min="0"
							max="86"
							frac="0" />
						<Slider
							:label="t('food.vitb5')"
							@input="slotProps.confirm"
							name="vitb5"
							unit="mg"
							min="0"
							max="30"
							frac="1" />
						<Slider
							:label="t('food.vitb6')"
							@input="slotProps.confirm"
							name="vitb6"
							unit="mg"
							min="0"
							max="8"
							frac="1" />
						<Slider
							:label="t('food.vitb7')"
							@input="slotProps.confirm"
							name="vitb7"
							unit="mg"
							min="0"
							max="0.2"
							frac="2" />
						<Slider
							:label="t('food.vitb9')"
							@input="slotProps.confirm"
							name="vitb9"
							unit="mg"
							min="0"
							max="0.6"
							frac="2" />
						<Slider
							:label="t('food.vitb12')"
							@input="slotProps.confirm"
							name="vitb12"
							unit="mg"
							min="0"
							max="0.1"
							frac="2" />
						<Slider
							:label="t('food.vitc')"
							@input="slotProps.confirm"
							name="vitc"
							unit="mg"
							min="0"
							max="210"
							frac="0" />
						<Slider
							:label="t('food.vitd')"
							@input="slotProps.confirm"
							name="vitd"
							unit="mg"
							min="0"
							max="0.04"
							frac="3" />
						<Slider
							:label="t('food.vite')"
							@input="slotProps.confirm"
							name="vite"
							unit="mg"
							min="0"
							max="45"
							frac="0" />
						<Slider
							:label="t('food.vitk')"
							@input="slotProps.confirm"
							name="vitk"
							unit="mg"
							min="0"
							max="0.9"
							frac="2" />
					</fieldset>
					<fieldset>
						<legend>{{ t("aria.headminerals") }}</legend>
						<Slider
							:label="t('food.calc')"
							@input="slotProps.confirm"
							name="calc"
							unit="mg"
							min="0"
							max="1200"
							frac="0" />
						<Slider
							:label="t('food.pot')"
							@input="slotProps.confirm"
							name="pot"
							unit="mg"
							min="0"
							max="1800"
							frac="0" />
						<Slider
							:label="t('food.sod')"
							@input="slotProps.confirm"
							v-if="prefs.ui.trackSaltAsSodium"
							name="sod"
							unit="mg"
							min="0"
							max="5000"
							frac="0" />
						<Slider
							:label="t('food.salt')"
							@input="slotProps.confirm"
							v-if="!prefs.ui.trackSaltAsSodium"
							name="salt"
							unit="g"
							min="0"
							max="100"
							frac="0" />
						<Slider
							:label="t('food.mag')"
							@input="slotProps.confirm"
							name="mag"
							unit="mg"
							min="0"
							max="550"
							frac="0" />
						<Slider
							:label="t('food.iron')"
							@input="slotProps.confirm"
							name="iron"
							unit="mg"
							min="0"
							max="35"
							frac="1" />
						<Slider
							:label="t('food.zinc')"
							@input="slotProps.confirm"
							name="zinc"
							unit="mg"
							min="0"
							max="100"
							frac="0" />
						<Slider
							:label="t('food.chl')"
							@input="slotProps.confirm"
							name="chl"
							unit="mg"
							min="0"
							max="850"
							frac="0" />
						<Slider
							:label="t('food.phos')"
							@input="slotProps.confirm"
							name="phos"
							unit="mg"
							min="0"
							max="1200"
							frac="0" />
						<Slider
							:label="t('food.mang')"
							@input="slotProps.confirm"
							name="mang"
							unit="mg"
							min="0"
							max="60"
							frac="1" />
						<Slider
							:label="t('food.cop')"
							@input="slotProps.confirm"
							name="cop"
							unit="mg"
							min="0"
							max="16"
							frac="2" />
						<Slider
							:label="t('food.iod')"
							@input="slotProps.confirm"
							name="iod"
							unit="mg"
							min="0"
							max="3"
							frac="3" />
						<!--<Slider
							:label="t('food.chr')"
							@input="slotProps.confirm"
							name="chr"
							unit="mg"
							min="0"
							max="0.5"
							frac="3"
						/>
						<Slider
							:label="t('food.mol')"
							@input="slotProps.confirm"
							name="mol"
							unit="mg"
							min="0"
							max="0.9"
							frac="3"
						/>-->
						<Slider
							:label="t('food.sel')"
							@input="slotProps.confirm"
							name="sel"
							unit="mg"
							min="0"
							max="0.3"
							frac="3" />
					</fieldset>
				</Search>
			</section>
		</template>

		<template #controls>
			<SortMenu class="m" :list="list" />
			<span class="spacer"></span>
		</template>

		<template #main>
			<FoodList class="m" ref="list" :items="filtered" @selected="showDetails" />
		</template>

		<template #head-details v-if="current">
			<h2>{{ current.name }}</h2>
		</template>

		<template #details v-if="current">
			<section class="subtitle">{{ current.cat }}</section>
			<section class="tags">
				<TagList :item="current" />
				<button
					class="icon async"
					:disabled="disableSave"
					@click="onEditMode"
					v-if="perms.canCreateFood || perms.canEditFood">
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
					<label>{{ t("food.torec") }}</label>
					<RecipeSelect
						:label="t('btn.add')"
						:placeholder="t('food.hintrec')"
						:items="recipes"
						:disabled="disableToRecipe"
						@confirm="addToRecipe" />
				</fieldset>
				<fieldset>
					<label>{{ t("food.todiary") }}</label>
					<DiarySelect
						:label="t('btn.add')"
						:time="now"
						:date="today"
						:disabled="disableToDiary"
						@confirm="addToDiary" />
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
								<input type="number" :value="current.kcal" name="kcal" @change="onInput" />
								<span class="unit">{{ t("unit.cal") }}</span>
							</div>
							<div>
								<label>{{ t("food.fat") }}</label>
								<input type="number" :value="current.fat" name="fat" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.carbs2") }}</label>
								<input type="number" :value="current.carb" name="carb" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.sugar") }}</label>
								<input type="number" :value="current.sug" name="sug" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.protein") }}</label>
								<input type="number" :value="current.prot" name="prot" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fiber") }}</label>
								<input type="number" :value="current.fib" name="fib" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
						</fieldset>
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.fatsat") }}</label>
								<input type="number" :value="current.fatsat" name="fatsat" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fato3") }}</label>
								<input type="number" :value="current.fato3" name="fato3" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fato6") }}</label>
								<input type="number" :value="current.fato6" name="fato6" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.fruc") }}</label>
								<input type="number" :value="current.fruc" name="fruc" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.gluc") }}</label>
								<input type="number" :value="current.gluc" name="gluc" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.suc") }}</label>
								<input type="number" :value="current.suc" name="suc" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
						</fieldset>
					</div>
					<div class="nutrient-block">
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.vita") }}</label>
								<input type="number" :value="current.vita" name="vita" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb1") }}</label>
								<input type="number" :value="current.vitb1" name="vitb1" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb2") }}</label>
								<input type="number" :value="current.vitb2" name="vitb2" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb3") }}</label>
								<input type="number" :value="current.vitb3" name="vitb3" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb5") }}</label>
								<input type="number" :value="current.vitb5" name="vitb5" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb6") }}</label>
								<input type="number" :value="current.vitb6" name="vitb6" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb7") }}</label>
								<input type="number" :value="current.vitb7" name="vitb7" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
						</fieldset>
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.vitb9") }}</label>
								<input type="number" :value="current.vitb9" name="vitb9" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitb12") }}</label>
								<input type="number" :value="current.vitb12" name="vitb12" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitc") }}</label>
								<input type="number" :value="current.vitc" name="vitc" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitd") }}</label>
								<input type="number" :value="current.vitd" name="vitd" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vite") }}</label>
								<input type="number" :value="current.vite" name="vite" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.vitk") }}</label>
								<input type="number" :value="current.vitk" name="vitk" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
						</fieldset>
					</div>
					<div class="nutrient-block">
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.calc") }}</label>
								<input type="number" :value="current.calc" name="calc" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.pot") }}</label>
								<input type="number" :value="current.pot" name="pot" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div v-if="prefs.ui.trackSaltAsSodium">
								<label>{{ t("food.sod") }}</label>
								<input type="number" :value="current.sod" name="sod" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div v-if="!prefs.ui.trackSaltAsSodium">
								<label>{{ t("food.salt") }}</label>
								<input type="number" :value="current.salt" name="salt" @change="onInput" />
								<span class="unit">{{ t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ t("food.mag") }}</label>
								<input type="number" :value="current.mag" name="mag" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.iron") }}</label>
								<input type="number" :value="current.iron" name="iron" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.zinc") }}</label>
								<input type="number" :value="current.zinc" name="zinc" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
						</fieldset>
						<fieldset :disabled="!editMode" class="col50">
							<div>
								<label>{{ t("food.chl") }}</label>
								<input type="number" :value="current.chl" name="chl" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.phos") }}</label>
								<input type="number" :value="current.phos" name="phos" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.mang") }}</label>
								<input type="number" :value="current.mang" name="mang" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.cop") }}</label>
								<input type="number" :value="current.cop" name="cop" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.iod") }}</label>
								<input type="number" :value="current.iod" name="iod" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
							</div>
							<div>
								<label>{{ t("food.sel") }}</label>
								<input type="number" :value="current.sel" name="sel" @change="onInput" />
								<span class="unit">{{ t("unit.mg") }}</span>
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
	padding: 0.25em 3em 0.25em 0.25em;
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

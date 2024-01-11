<script setup>
import Main from "../components/Main.vue";
import Search from "../components/LocalSearch.vue";
import Slider from "../components/Slider.vue";
import NewRecipe from "../components/ClickableInput.vue";
import FoodList from "../components/FoodList.vue";
import DiarySelect from "../components/ClickableDate.vue";
import IngredientList from "../components/IngredientList.vue";
import TagList from "../components/TagList.vue";
import SortMenu from "../components/SortMenu.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import BackImage from "../components/images/ImageRightArrow.vue";
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
const prefs = inject("prefs");

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
const list = ref(null);

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
	today.value = date;
	now.value = time;

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

function onBack() {
	main.value.hideDetails();
}

function onDetailVisibility() {
	editMode.value = false;
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

onMounted(() => (filtered.value = [...recipes.value]));
</script>

<template>
	<Main ref="main" @detailVisibility="onDetailVisibility" :class="{ 'edit-mode': editMode }">
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
						<legend>{{ t("aria.headmacro1") }}</legend>
						<Slider
							:label="t('food.energy')"
							:min="minSearchAttr('kcal')"
							:max="maxSearchAttr('kcal')"
							@input="slotProps.confirm"
							name="kcal"
							unit="cal"
							frac="0" />
						<Slider
							:label="t('food.fat')"
							:min="minSearchAttr('fat')"
							:max="maxSearchAttr('fat')"
							@input="slotProps.confirm"
							name="fat"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.carbs')"
							:min="minSearchAttr('carb')"
							:max="maxSearchAttr('carb')"
							@input="slotProps.confirm"
							name="carb"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.protein')"
							:min="minSearchAttr('prot')"
							:max="maxSearchAttr('prot')"
							@input="slotProps.confirm"
							name="prot"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.fiber')"
							:min="minSearchAttr('fib')"
							:max="maxSearchAttr('fib')"
							@input="slotProps.confirm"
							name="fib"
							unit="g"
							frac="0" />
					</fieldset>
					<fieldset>
						<legend>{{ t("aria.headmacro2") }}</legend>
						<Slider
							:label="t('food.fatsat')"
							:min="minSearchAttr('fatsat')"
							:max="maxSearchAttr('fatsat')"
							@input="slotProps.confirm"
							name="fatsat"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.fato3')"
							:min="minSearchAttr('fato3')"
							:max="maxSearchAttr('fato3')"
							@input="slotProps.confirm"
							name="fato3"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.fato6')"
							:min="minSearchAttr('fato6')"
							:max="maxSearchAttr('fato6')"
							@input="slotProps.confirm"
							name="fato6"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.sugar')"
							:min="minSearchAttr('sug')"
							:max="maxSearchAttr('sug')"
							@input="slotProps.confirm"
							name="sug"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.fruc')"
							:min="minSearchAttr('fruc')"
							:max="maxSearchAttr('fruc')"
							@input="slotProps.confirm"
							name="fruc"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.gluc')"
							:min="minSearchAttr('gluc')"
							:max="maxSearchAttr('gluc')"
							@input="slotProps.confirm"
							name="gluc"
							unit="g"
							frac="0" />
						<Slider
							:label="t('food.suc')"
							:min="minSearchAttr('suc')"
							:max="maxSearchAttr('suc')"
							@input="slotProps.confirm"
							name="suc"
							unit="g"
							frac="0" />
					</fieldset>
					<fieldset>
						<legend>{{ t("aria.headvits") }}</legend>
						<Slider
							:label="t('food.vita')"
							:min="minSearchAttr('vita')"
							:max="maxSearchAttr('vita')"
							@input="slotProps.confirm"
							name="vita"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.vitb1')"
							:min="minSearchAttr('vitb1')"
							:max="maxSearchAttr('vitb1')"
							@input="slotProps.confirm"
							name="vitb1"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.vitb2')"
							:min="minSearchAttr('vitb2')"
							:max="maxSearchAttr('vitb2')"
							@input="slotProps.confirm"
							name="vitb2"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.vitb3')"
							:min="minSearchAttr('vitb3')"
							:max="maxSearchAttr('vitb3')"
							@input="slotProps.confirm"
							name="vitb3"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.vitb5')"
							:min="minSearchAttr('vitb5')"
							:max="maxSearchAttr('vitb5')"
							@input="slotProps.confirm"
							name="vitb5"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.vitb6')"
							:min="minSearchAttr('vitb6')"
							:max="maxSearchAttr('vitb6')"
							@input="slotProps.confirm"
							name="vitb6"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.vitb7')"
							:min="minSearchAttr('vitb7')"
							:max="maxSearchAttr('vitb7')"
							@input="slotProps.confirm"
							name="vitb7"
							unit="mg"
							frac="2" />
						<Slider
							:label="t('food.vitb9')"
							:min="minSearchAttr('vitb9')"
							:max="maxSearchAttr('vitb9')"
							@input="slotProps.confirm"
							name="vitb9"
							unit="mg"
							frac="2" />
						<Slider
							:label="t('food.vitb12')"
							:min="minSearchAttr('vitb12')"
							:max="maxSearchAttr('vitb12')"
							@input="slotProps.confirm"
							name="vitb12"
							unit="mg"
							frac="2" />
						<Slider
							:label="t('food.vitc')"
							:min="minSearchAttr('vitc')"
							:max="maxSearchAttr('vitc')"
							@input="slotProps.confirm"
							name="vitc"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.vitd')"
							:min="minSearchAttr('vitd')"
							:max="maxSearchAttr('vitd')"
							@input="slotProps.confirm"
							name="vitd"
							unit="mg"
							frac="3" />
						<Slider
							:label="t('food.vite')"
							:min="minSearchAttr('vite')"
							:max="maxSearchAttr('vite')"
							@input="slotProps.confirm"
							name="vite"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.vitk')"
							:min="minSearchAttr('vitk')"
							:max="maxSearchAttr('vitk')"
							@input="slotProps.confirm"
							name="vitk"
							unit="mg"
							frac="2" />
					</fieldset>
					<fieldset>
						<legend>{{ t("aria.headminerals") }}</legend>
						<Slider
							:label="t('food.calc')"
							:min="minSearchAttr('calc')"
							:max="maxSearchAttr('calc')"
							@input="slotProps.confirm"
							name="calc"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.pot')"
							:min="minSearchAttr('pot')"
							:max="maxSearchAttr('pot')"
							@input="slotProps.confirm"
							name="pot"
							unit="mg"
							frac="0" />
						<Slider
							v-if="prefs.ui.trackSaltAsSodium"
							:label="t('food.sod')"
							:min="minSearchAttr('sod')"
							:max="maxSearchAttr('sod')"
							@input="slotProps.confirm"
							name="sod"
							unit="mg"
							frac="0" />
						<Slider
							v-if="!prefs.ui.trackSaltAsSodium"
							:label="t('food.salt')"
							:min="minSearchAttr('salt')"
							:max="maxSearchAttr('salt')"
							@input="slotProps.confirm"
							name="salt"
							unit="g"
							frac="1" />
						<Slider
							:label="t('food.mag')"
							:min="minSearchAttr('mag')"
							:max="maxSearchAttr('mag')"
							@input="slotProps.confirm"
							name="mag"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.iron')"
							:min="minSearchAttr('iron')"
							:max="maxSearchAttr('iron')"
							@input="slotProps.confirm"
							name="iron"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.zinc')"
							:min="minSearchAttr('zinc')"
							:max="maxSearchAttr('zinc')"
							@input="slotProps.confirm"
							name="zinc"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.chl')"
							:min="minSearchAttr('chl')"
							:max="maxSearchAttr('chl')"
							@input="slotProps.confirm"
							name="chl"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.phos')"
							:min="minSearchAttr('phos')"
							:max="maxSearchAttr('phos')"
							@input="slotProps.confirm"
							name="phos"
							unit="mg"
							frac="0" />
						<Slider
							:label="t('food.mang')"
							:min="minSearchAttr('mang')"
							:max="maxSearchAttr('mang')"
							@input="slotProps.confirm"
							name="mang"
							unit="mg"
							frac="1" />
						<Slider
							:label="t('food.cop')"
							:min="minSearchAttr('cop')"
							:max="maxSearchAttr('cop')"
							@input="slotProps.confirm"
							name="cop"
							unit="mg"
							frac="2" />
						<Slider
							:label="t('food.iod')"
							:min="minSearchAttr('iod')"
							:max="maxSearchAttr('iod')"
							@input="slotProps.confirm"
							name="iod"
							unit="mg"
							frac="3" />
						<!--<Slider
							:label="t('food.chr')"
							:min="minSearchAttr('chr')"
							:max="maxSearchAttr('chr')"
							@input="slotProps.confirm"
							name="chr"
							unit="mg"
							frac="3"
						/>
						<Slider
							:label="t('food.mol')"
							:min="minSearchAttr('mol')"
							:max="maxSearchAttr('mol')"
							@input="slotProps.confirm"
							name="mol"
							unit="mg"
							frac="3"
						/>-->
						<Slider
							:label="t('food.sel')"
							:min="minSearchAttr('sel')"
							:max="maxSearchAttr('sel')"
							@input="slotProps.confirm"
							name="sel"
							unit="mg"
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

		<template #details v-if="current">
			<div class="controls">
				<form ref="form" autocomplete="off" id="form-recipe">
					<fieldset :disabled="!editMode">
						<input type="text" name="name" :value="current.name" />
					</fieldset>
				</form>
				<span class="spacer"></span>
				<button @click="onBack" class="open-details icon cancel-edit-mode"><BackImage /></button>
			</div>
			<section class="subtitle no-edit-mode" v-html="ownerInfo"></section>
			<section class="tags">
				<TagList :item="current" />
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
						<div>
							<label>{{ t("food.carbs2") }}</label>
							<span>{{ perServing(current.carb, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.sugar") }}</label>
							<span>{{ perServing(current.sug, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.protein") }}</label>
							<span>{{ perServing(current.prot, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.fiber") }}</label>
							<span>{{ perServing(current.fib, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
					</div>
					<div class="col50">
						<div>
							<label>{{ t("food.fatsat") }}</label>
							<span>{{ perServing(current.fatsat, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.fato3") }}</label>
							<span>{{ perServing(current.fato3, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.fato6") }}</label>
							<span>{{ perServing(current.fato6, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.fruc") }}</label>
							<span>{{ perServing(current.fruc, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.gluc") }}</label>
							<span>{{ perServing(current.gluc, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.suc") }}</label>
							<span>{{ perServing(current.suc, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
					</div>
				</div>
				<div class="nutrient-block">
					<div class="col50">
						<div>
							<label>{{ t("food.vita") }}</label>
							<span>{{ perServing(current.vita, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb1") }}</label>
							<span>{{ perServing(current.vitb1, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb2") }}</label>
							<span>{{ perServing(current.vitb2, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb3") }}</label>
							<span>{{ perServing(current.vitb3, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb5") }}</label>
							<span>{{ perServing(current.vitb5, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb6") }}</label>
							<span>{{ perServing(current.vitb6, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb7") }}</label>
							<span>{{ perServing(current.vitb7, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
					</div>
					<div class="col50">
						<div>
							<label>{{ t("food.vitb9") }}</label>
							<span>{{ perServing(current.vitb9, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitb12") }}</label>
							<span>{{ perServing(current.vitb12, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitc") }}</label>
							<span>{{ perServing(current.vitc, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitd") }}</label>
							<span>{{ perServing(current.vitd, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vite") }}</label>
							<span>{{ perServing(current.vite, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.vitk") }}</label>
							<span>{{ perServing(current.vitk, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
					</div>
				</div>
				<div class="nutrient-block">
					<div class="col50">
						<div>
							<label>{{ t("food.calc") }}</label>
							<span>{{ perServing(current.calc, 1) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.pot") }}</label>
							<span>{{ perServing(current.pot, 1) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div v-if="prefs.ui.trackSaltAsSodium">
							<label>{{ t("food.sod") }}</label>
							<span>{{ perServing(current.sod, 1) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div v-if="!prefs.ui.trackSaltAsSodium">
							<label>{{ t("food.salt") }}</label>
							<span>{{ perServing(current.salt, 1) }}</span>
							<span class="unit">{{ t("unit.g") }}</span>
						</div>
						<div>
							<label>{{ t("food.mag") }}</label>
							<span>{{ perServing(current.mag, 1) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.iron") }}</label>
							<span>{{ perServing(current.iron, 2) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.zinc") }}</label>
							<span>{{ perServing(current.zinc, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
					</div>
					<div class="col50">
						<div>
							<label>{{ t("food.chl") }}</label>
							<span>{{ perServing(current.chl, 1) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.phos") }}</label>
							<span>{{ perServing(current.phos, 1) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.mang") }}</label>
							<span>{{ perServing(current.mang, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.cop") }}</label>
							<span>{{ perServing(current.cop, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.iod") }}</label>
							<span>{{ perServing(current.iod, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
						</div>
						<div>
							<label>{{ t("food.sel") }}</label>
							<span>{{ perServing(current.sel, 3) }}</span>
							<span class="unit">{{ t("unit.mg") }}</span>
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
							<input
								type="number"
								name="size"
								form="form-recipe"
								:value="current.size"
								@change="onInput" />
							<label>{{ t("recipe.size", current.size) }}</label>
						</div>
					</fieldset>
					<fieldset :disabled="!editMode" class="col50">
						<div>
							<label>{{ t("recipe.time") }}</label>
							<input
								type="number"
								disabled
								:value="current.preptime + current.cooktime + current.misctime" />
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ t("recipe.preptime") }}</label>
							<input
								type="number"
								name="preptime"
								form="form-recipe"
								:value="current.preptime"
								@change="onInput" />
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ t("recipe.cooktime") }}</label>
							<input
								type="number"
								name="cooktime"
								form="form-recipe"
								:value="current.cooktime"
								@change="onInput" />
							<span class="unit">{{ t("unit.min") }}</span>
						</div>
						<div>
							<label>{{ t("recipe.misctime") }}</label>
							<input
								type="number"
								name="misctime"
								form="form-recipe"
								:value="current.misctime"
								@change="onInput" />
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

<script setup>
import Main from "../components/Main.vue";
import Search from "../components/LocalSearch.vue";
import Slider from "../components/Slider.vue";
import NewFood from "../components/ClickableInput.vue";
import FoodList from "../components/FoodList.vue";
import EditImage from "../components/images/ImageEdit.vue";
import SaveImage from "../components/images/ImageSave.vue";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const csrf = inject("csrfToken");
const perms = inject("perms");
const foods = inject("food");

const filtered = ref([]);
const current = ref(null);
const editMode = ref(false);
const isSaving = ref(false);

const main = ref(null);
const form = ref(null);

function newFood(name) {
	// TODO create new food
	console.log(name);
}

function saveFood() {
	isSaving.value = true;
	let id = current.value.id;
	let data = new FormData(form.value);
	fetch("/api/v1/food/" + id, {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(data),
	}).then((response) => {
		if (response.ok) {
			editMode.value = false;
			refreshFood(id);
		} else {
			isSaving.value = false;
		}
	});
}

function refreshFood(id) {
	fetch("/api/v1/food/" + id)
		.then((response) => response.json())
		.then((data) => {
			data.name = t(data.id.toString());
			foods.value = foods.value.map((f) => (data.id == f.id ? data : f));
			filtered.value = filtered.value.map((f) => (data.id == f.id ? data : f));
			current.value = current.value.id == data.id ? data : current.value;
			setTimeout(function () {
				isSaving.value = false;
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
</script>

<template>
	<Main ref="main" @detailVisibility="editMode = false">
		<template #filter>
			<section v-if="perms.canCreateFood" class="new-item">
				<h2>{{ $t("aria.headnew") }}</h2>
				<NewFood :label="$t('btn.new')" :placeholder="$t('food.hintnew')" @confirm="newFood" />
			</section>
			<section>
				<h2>{{ $t("aria.headsearch") }}</h2>
				<Search :data="foods" v-slot="slotProps" :placeholder="$t('food.hintsearch')" @result="updateList">
					<fieldset>
						<legend>Primary Macronutrients</legend>
						<Slider
							:label="$t('food.energy')"
							@input="slotProps.confirm"
							name="kcal"
							unit="cal"
							min="0"
							max="900"
							frac="0"
						/>
						<Slider
							:label="$t('food.fat')"
							@input="slotProps.confirm"
							name="fat"
							unit="g"
							min="0"
							max="100"
							frac="0"
						/>
						<Slider
							:label="$t('food.carbs')"
							@input="slotProps.confirm"
							name="carb"
							unit="g"
							min="0"
							max="100"
							frac="0"
						/>
						<Slider
							:label="$t('food.protein')"
							@input="slotProps.confirm"
							name="prot"
							unit="g"
							min="0"
							max="89"
							frac="0"
						/>
						<Slider
							:label="$t('food.fiber')"
							@input="slotProps.confirm"
							name="fib"
							unit="g"
							min="0"
							max="71"
							frac="0"
						/>
					</fieldset>
					<fieldset>
						<legend>Secondary Macronutrients</legend>
						<Slider
							:label="$t('food.fatsat')"
							@input="slotProps.confirm"
							name="fatsat"
							unit="g"
							min="0"
							max="83"
							frac="0"
						/>
						<Slider
							:label="$t('food.fato3')"
							@input="slotProps.confirm"
							name="fato3"
							unit="g"
							min="0"
							max="54"
							frac="0"
						/>
						<Slider
							:label="$t('food.fato6')"
							@input="slotProps.confirm"
							name="fato6"
							unit="g"
							min="0"
							max="70"
							frac="0"
						/>
						<Slider
							:label="$t('food.sugar')"
							@input="slotProps.confirm"
							name="sug"
							unit="g"
							min="0"
							max="100"
							frac="0"
						/>
						<Slider
							:label="$t('food.salt')"
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
					:disabled="isSaving"
					@click="onEditMode"
					v-if="perms.canCreateFood || perms.canEditFood"
				>
					<EditImage v-if="!editMode" />
					<SaveImage v-if="editMode" />
				</button>
			</section>
			<section class="nutrients">
				<h2>Nutrients</h2>
				<form ref="form">
					<div>
						<fieldset :disabled="!editMode">
							<div>
								<label>{{ $t("food.energy") }}</label>
								<input type="text" :value="current.kcal" name="kcal" @change="onInput" />
								<span class="unit">{{ $t("unit.cal") }}</span>
							</div>
							<div>
								<label>{{ $t("food.fat") }}</label>
								<input type="text" :value="current.fat" name="fat" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.carbs2") }}</label>
								<input type="text" :value="current.carb" name="carb" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.protein") }}</label>
								<input type="text" :value="current.prot" name="prot" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.fiber") }}</label>
								<input type="text" :value="current.fib" name="fib" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
						</fieldset>
						<fieldset :disabled="!editMode">
							<div>
								<label>{{ $t("food.fatsat") }}</label>
								<input type="text" :value="current.fatsat" name="fatsat" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.fato3") }}</label>
								<input type="text" :value="current.fato3" name="fato3" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.fato6") }}</label>
								<input type="text" :value="current.fato6" name="fato6" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.sugar") }}</label>
								<input type="text" :value="current.sug" name="sug" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
							</div>
							<div>
								<label>{{ $t("food.salt") }}</label>
								<input type="text" :value="current.salt" name="salt" @change="onInput" />
								<span class="unit">{{ $t("unit.g") }}</span>
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
	border: none;
}

#details section.tags {
	padding: 0 3em 0.5em 0.25em;
	position: relative;
	display: block;
}

#details section.tags button.icon {
	position: absolute;
	right: 0.5em;
	bottom: 0.5em;
}

#details .nutrients form > div:not(:first-of-type) {
	margin-top: 3em;
}

@media only screen and (min-width: 400px) {
	#details .nutrients form > div {
		display: flex;
		justify-content: space-between;
	}

	#details .nutrients fieldset {
		flex-basis: 45%;
	}
}
</style>

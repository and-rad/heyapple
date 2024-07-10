<script setup>
import Main from "../components/Main.vue";
import Checkbox from "../components/Checkbox.vue";
import PasswordField from "../../../login/src/components/Password.vue";
import { ref, computed, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const csrf = inject("csrfToken");
const log = inject("log");
const prefs = inject("prefs");

/**
 * We don't work on the prefs object directly, but cache the
 * macro targets instead. This is a UX improvement because
 * it allows us to easily reset the values in the table without
 * reloading the page.
 */
const macros = ref([...prefs.value.macros]);

/**
 * If true, the table for entering macronutrient targets displays
 * all seven weekdays, regardless of whether they all contain
 * the same values for each nutrient.
 */
const forceDisplayFullMacroWeek = ref(false);

/**
 * Shows macro target values as percentage of the total kcal
 * goal. This is only an aid for editing values and won't be
 * saved between visits to the page.
 */
const displayMacroPercentage = ref(false);

/**
 * The index of the page section that is currently in view.
 * A section is considered in view if it is the lowest section
 * that starts above the vertical centerline of the viewport.
 * Used for styling navigation entries.
 */
const selectedSection = ref(0);

/**
 * The unit that macronutrient targets are displayed in. Can be
 * grams or percentage, depending on the display settings.
 */
const macroUnit = computed(() => (displayMacroPercentage.value ? "%" : t("unit.g")));

/**
 * True if all macronutrient targets for all days of the week
 * are the same.
 */
const allMacrosEqual = computed(() => {
	let arr = macros.value;
	return arr.every((elem) => JSON.stringify(elem) == JSON.stringify(arr[0]));
});

/**
 * True if macro targets for individual days are different or
 * all days of the table are forced to be visible.
 */
const shouldShowFullMacroTable = computed(() => !allMacrosEqual.value || forceDisplayFullMacroWeek.value);

/**
 * Returns the user's macronutrient targets, fomratted for display
 * in the macro target settings table.
 */
const formattedMacros = computed(() => {
	let result = [];

	// Prepare additional display elements
	if (shouldShowFullMacroTable.value) {
		result = macros.value.map((elem, idx) => ({ ...elem, l10nKey: `day.cal${idx + 1}`, weekend: idx > 4 }));
	} else {
		result = macros.value.slice(0, 1).map((elem) => ({ ...elem, l10nKey: "day.all", weekend: false }));
	}

	// Convert to percentage of kcal if necessary
	if (displayMacroPercentage.value) {
		result.forEach((elem) => {
			elem.fat *= 900.0 / elem.kcal;
			elem.carb *= 400.0 / elem.kcal;
			elem.prot *= 400.0 / elem.kcal;
		});
	}

	// Format numbers to fit the input fields
	const numFrac = displayMacroPercentage.value ? 1 : 0;
	result.forEach((elem) => {
		elem.fat = parseFloat(elem.fat).toFixed(numFrac);
		elem.carb = parseFloat(elem.carb).toFixed(numFrac);
		elem.prot = parseFloat(elem.prot).toFixed(numFrac);
	});

	return result;
});

function onNavItem(id) {
	let target = document.getElementById(id);
	target.scrollIntoView({ behavior: "smooth", block: "start" });
}

function onScroll() {
	const sections = document.querySelectorAll(".content > section");
	const center = window.innerHeight * 0.5;
	for (let i = sections.length - 1; i >= 0; --i) {
		if (sections[i].getBoundingClientRect().top < center) {
			selectedSection.value = i;
			return;
		}
	}
}

function onSaveEmail(evt) {
	evt.target.disabled = true;
	let form = evt.target.form;

	fetch("/auth/email", {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form)),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savemail.err" + response.status);
			}
			log.msg(t("savemail.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				evt.target.disabled = false;
			}, 500);
		});
}

function onRollUsername(evt) {
	evt.target.disabled = true;

	fetch("/api/v1/name", {
		method: "PUT",
		headers: { "X-CSRF-Token": csrf },
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savename.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			prefs.value.account.name = data.name;
			log.msg(t("savename.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				evt.target.disabled = false;
			}, 500);
		});
}

function onChangePassword(evt) {
	evt.target.disabled = true;
	let form = evt.target.form;

	fetch("/auth/pass", {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form)),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savepass.err" + response.status);
			}
			log.msg(t("savepass.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				evt.target.disabled = false;
			}, 500);
		});
}

/**
 * Event listener for the checkbox that controls whether to
 * display the full week in the table for macro targets. When
 * the checkbox is unchecked, all macros are set to the values
 * in the first table row.
 */
function onCheckedMacroDaily(evt) {
	forceDisplayFullMacroWeek.value = evt.target.checked;

	if (evt.target.checked) {
		macros.value = [...prefs.value.macros];
		return;
	}

	let arr = macros.value;
	for (let i = 0; i < arr.length; ++i) {
		arr[i] = arr[0];
	}
}

/**
 * Event listener for the checkbox that controls whether the
 * macronutrient targets are displayed as percentages of the
 * total daily calories.
 */
function onCheckedMacroPercentage(evt) {
	displayMacroPercentage.value = evt.target.checked;
}

function onChangeMacros(evt) {
	evt.target.disabled = true;
	let form = new FormData(evt.target.form);

	for (let macro of ["kcal", "fat", "carb", "prot"]) {
		let arr = form.getAll(macro);
		if (arr.length == 1) {
			arr = Array(7).fill(parseInt(arr[0]));
		}
		for (let i = 0; i < arr.length; ++i) {
			macros.value[i][macro] = parseInt(arr[i]);
		}
	}

	let data = { ...prefs.value };
	data.macros = macros.value;

	fetch("/api/v1/prefs", {
		method: "PUT",
		headers: {
			"Content-Type": "application/json",
			"X-CSRF-Token": csrf,
		},
		body: JSON.stringify(data),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savemacro.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			prefs.value = data;
			macros.value = [...prefs.value.macros];
			log.msg(t("savemacro.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				evt.target.disabled = false;
			}, 500);
		});
}

function onDeleteUser(evt) {
	evt.target.disabled = true;

	fetch("/api/v1/user", {
		method: "DELETE",
		headers: { "X-CSRF-Token": csrf },
	})
		.then((response) => {
			if (!response.ok) {
				throw t("deluser.err" + response.status);
			}
			log.msg(t("deluser.ok"));
			setTimeout(() => {
				window.location = "/";
			}, 3000);
		})
		.catch((err) => log.err(err))
		.finally(() => {
			setTimeout(function () {
				evt.target.disabled = false;
			}, 500);
		});
}
</script>

<template>
	<Main class="settings no-dt" @scroll="onScroll">
		<template #filter>
			<nav>
				<ul>
					<li>
						<a @click="onNavItem('head-macro')" :class="selectedSection == 0 ? 'selected' : ''">
							{{ t("nav.targets") }}
						</a>
					</li>
					<!--<li>
						<a @click="onNavItem('head-body')" :class="selectedSection == 1 ? 'selected' : ''">
							{{ t("nav.body") }}
						</a>
					</li>-->
					<li>
						<a @click="onNavItem('head-account')" :class="selectedSection == 1 ? 'selected' : ''">
							{{ t("nav.account") }}
						</a>
					</li>
					<li>
						<a @click="onNavItem('head-danger')" :class="selectedSection == 2 ? 'selected' : ''">
							{{ t("nav.danger") }}
						</a>
					</li>
				</ul>
			</nav>
		</template>

		<template #main>
			<section>
				<h2 id="head-macro">{{ t("nav.targets") }}</h2>
				<p v-html="t('profile.macrohint')"></p>

				<div class="checkbox-group">
					<Checkbox :checked="shouldShowFullMacroTable" @click="onCheckedMacroDaily">
						<span>{{ t("profile.macroeachday") }}</span>
					</Checkbox>
					<Checkbox :checked="displayMacroPercentage" @click="onCheckedMacroPercentage">
						<span>{{ t("profile.macropercent") }}</span>
					</Checkbox>
				</div>

				<form>
					<table>
						<thead>
							<tr>
								<th class="name"></th>
								<th class="num">{{ t("food.energy") }}</th>
								<th class="num">{{ t("food.fat") }}</th>
								<th class="num">{{ t("food.carbs2") }}</th>
								<th class="num">{{ t("food.protein") }}</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="macro in formattedMacros">
								<td class="name" :class="macro.weekend ? 'wknd' : ''">
									<div>{{ t(macro.l10nKey) }}</div>
								</td>
								<td class="num">
									<input type="number" name="kcal" :value="macro.kcal" />
									<span class="unit">&nbsp;{{ t("unit.cal") }}</span>
								</td>
								<td class="num">
									<input type="number" name="fat" :value="macro.fat" />
									<span class="unit">&nbsp;{{ macroUnit }}</span>
								</td>
								<td class="num">
									<input type="number" name="carb" :value="macro.carb" />
									<span class="unit">&nbsp;{{ macroUnit }}</span>
								</td>
								<td class="num">
									<input type="number" name="prot" :value="macro.prot" />
									<span class="unit">&nbsp;{{ macroUnit }}</span>
								</td>
							</tr>
						</tbody>
					</table>
					<button type="button" @click="onChangeMacros" class="async">
						{{ t("btn.changemacro") }}
					</button>
				</form>
			</section>

			<!--<section>
				<h2 id="head-body">{{ t("nav.body") }}</h2>
			</section>-->

			<section>
				<h2 id="head-account">{{ t("nav.account") }}</h2>
				<form @submit.prevent>
					<label>{{ t("profile.email") }}</label>
					<input type="email" name="email" :value="prefs.account.email" />
					<p v-html="t('profile.emailhint')"></p>
					<button type="button" @click="onSaveEmail" class="async">
						{{ t("btn.changemail") }}
					</button>
				</form>
				<form @submit.prevent>
					<label>
						{{ t("profile.name") }} <a href="#">{{ t("profile.namelink") }}</a>
					</label>
					<input readonly type="text" name="name" :value="prefs.account.name" />
					<p v-html="t('profile.namehint')"></p>
					<button type="button" @click="onRollUsername" class="async">
						{{ t("btn.reroll") }}
					</button>
				</form>
				<form @submit.prevent>
					<label>{{ t("profile.passold") }}</label>
					<PasswordField ref="passField" name="passold" />
					<label>{{ t("profile.passnew") }}</label>
					<PasswordField ref="passField" name="passnew" withBar="true" />
					<p v-html="t('profile.passhint')"></p>
					<button type="button" @click="onChangePassword" class="async">
						{{ t("btn.changepass") }}
					</button>
				</form>
			</section>

			<section class="danger">
				<h2 id="head-danger">{{ t("nav.danger") }}</h2>
				<form @submit.prevent>
					<p v-html="t('profile.deletehint')"></p>
					<button type="button" @click="onDeleteUser" class="async">
						{{ t("btn.deleteuser") }}
					</button>
				</form>
			</section>
		</template>
	</Main>
</template>

<style>
#filter nav ul {
	list-style: none;
	white-space: nowrap;
	padding: 0;
	margin: 1em 0;
}

#filter nav a {
	display: block;
	padding: 0.5em 1em;
	color: var(--color-text);
	border-left: 3px solid transparent;
	transition: color var(--transition-style);
	cursor: pointer;
}

#filter nav a:hover,
#filter nav a:active,
#filter nav a:focus {
	color: var(--color-primary);
}

#filter nav a.selected {
	border-color: var(--color-secondary);
}

main.settings #main .controls {
	border-bottom: var(--border-light);
}

main.settings .content {
	padding: 0 1em 1em;
}

main.settings .content h2 {
	margin: 2em 0 1em;
}

main.settings .content h2:first-child {
	margin-top: 1rem;
}

main.settings .content section {
	max-width: 480px;
	margin: auto;
}

main.settings .content section:last-child {
	margin-bottom: 50vh;
}

main.settings .content section.danger {
	background-color: var(--color-warn-light);
	border: 1px solid var(--color-warn);
	border-radius: 8px;
	padding: 1em;
}

main.settings .content section.danger form:last-child {
	margin-bottom: 0;
}

main.settings .content section.danger h2 {
	color: var(--color-secondary);
}

main.settings .content section.danger button {
	background-color: var(--color-secondary);
}

main.settings .content form {
	margin: 2em 0 4em;
}

main.settings .content label {
	white-space: nowrap;
	text-overflow: ellipsis;
	overflow-x: hidden;
	color: var(--color-text-light);
	margin-top: 1em;
	display: block;
}

main.settings .content .checkbox-group {
	margin: 2em 0;
}

main.settings .content label.checkbox {
	display: flex;
	margin: 1em 0;
}

main.settings .content label.checkbox > input + div {
	margin-right: 0.5em;
}

main.settings .content label.checkbox > input:checked ~ * {
	color: var(--color-text);
}

main.settings .content form > button {
	margin: 1em 0 1em;
}

main.settings label > a {
	float: right;
}

main.settings #main .content td.name {
	color: var(--color-text-light);
	width: 3em;
	vertical-align: middle;
	padding: 0;
}

main.settings .content td.name div {
	border: 1px solid var(--color-primary);
	background: var(--color-primary-lighter);
	border-radius: 1.25em;
	color: var(--color-primary-dark);
	padding: 0.5em;
	min-height: 2.5em;
	min-width: 4em;
	text-align: center;
}

main.settings .content td.name.wknd div {
	background: var(--color-primary-light);
}

main.settings #main .content table td.num {
	width: unset;
	max-width: unset;
}

main.settings .content table td.num input {
	text-align: right;
	border-radius: 0;
	border: none;
	border-bottom: var(--border);
	display: inline-block;
	width: 3em;
	padding: 0;
}

main.settings .content table span.unit {
	border-bottom: var(--border);
	display: inline-block;
	min-width: 1.1em;
}

main.settings .content table td.num input:focus + span.unit {
	border-color: var(--color-primary);
}

main.settings input + p,
main.settings span.password + p {
	margin-top: 1em;
}

@media only screen and (min-width: 500px) {
	main.settings .content form > button {
		margin-left: auto;
		width: 240px;
		min-width: 240px;
	}
}

@media only screen and (min-width: 800px) {
	main.settings #main .controls {
		display: none;
	}
}

@media only screen and (min-width: 1280px) {
	main.settings #main {
		right: 0;
	}

	main.settings .content {
		padding-right: calc(300px + 1em);
	}
}
</style>

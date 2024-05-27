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
 * If true, the table for entering macronutrient targets displays
 * all seven weekdays, regardless of whether they all contain
 * the same values for each nutrient.
 */
const forceDisplayFullMacroWeek = ref(false);

/**
 * True if all macronutrient targets for all days of the week
 * are the same.
 */
const allMacrosEqual = computed(() => {
	let arr = prefs.value.macros;
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
	if (shouldShowFullMacroTable.value) {
		return prefs.value.macros.map((elem, idx) => ({ ...elem, l10nKey: `day.cal${idx + 1}`, weekend: idx > 4 }));
	}
	return prefs.value.macros.slice(0, 1).map((elem) => ({ ...elem, l10nKey: "day.all", weekend: false }));
});

function onNavItem(id) {
	let target = document.getElementById(id);
	target.scrollIntoView({ behavior: "smooth", block: "start" });
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

function onCheckedMacroDaily(evt) {
	forceDisplayFullMacroWeek.value = evt.target.checked;
	if (!evt.target.checked) {
		let arr = prefs.value.macros;
		for (let i = 0; i < arr.length; ++i) {
			arr[i] = arr[0];
		}
	}
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
			prefs.value.macros[i][macro] = parseInt(arr[i]);
		}
	}

	fetch("/api/v1/prefs", {
		method: "PUT",
		headers: {
			"Content-Type": "application/json",
			"X-CSRF-Token": csrf,
		},
		body: JSON.stringify(prefs.value),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savemacro.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			prefs.value = data;
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
	<Main class="settings no-dt">
		<template #filter>
			<nav>
				<ul>
					<li>
						<a @click="onNavItem('head-account')"> {{ t("nav.account") }} </a>
					</li>
					<!--<li>
						<a @click="onNavItem('head-body')"> {{ t("nav.body") }} </a>
					</li>-->
					<li>
						<a @click="onNavItem('head-macro')"> {{ t("nav.targets") }} </a>
					</li>
					<li>
						<a @click="onNavItem('head-danger')"> {{ t("nav.danger") }} </a>
					</li>
				</ul>
			</nav>
		</template>

		<template #main>
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

			<!--<section>
				<h2 id="head-body">{{ t("nav.body") }}</h2>
			</section>-->

			<section>
				<h2 id="head-macro">{{ t("nav.targets") }}</h2>
				<p v-html="t('profile.macrohint')"></p>

				<div class="checkbox-group">
					<Checkbox :checked="shouldShowFullMacroTable" @click="onCheckedMacroDaily">
						<span>{{ t("profile.macroeachday") }}</span>
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
									<span class="unit">&nbsp;{{ t("unit.g") }}</span>
								</td>
								<td class="num">
									<input type="number" name="carb" :value="macro.carb" />
									<span class="unit">&nbsp;{{ t("unit.g") }}</span>
								</td>
								<td class="num">
									<input type="number" name="prot" :value="macro.prot" />
									<span class="unit">&nbsp;{{ t("unit.g") }}</span>
								</td>
							</tr>
						</tbody>
					</table>
					<button type="button" @click="onChangeMacros" class="async">
						{{ t("btn.changemacro") }}
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
